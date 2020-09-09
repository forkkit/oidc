package op

import (
	"context"
	"errors"
	"net/http"
	"time"

	"gopkg.in/square/go-jose.v2"

	"github.com/caos/oidc/pkg/oidc"
	"github.com/caos/oidc/pkg/rp"
	"github.com/caos/oidc/pkg/utils"
)

type Exchanger interface {
	Issuer() string
	Storage() Storage
	Decoder() utils.Decoder
	Signer() Signer
	Crypto() Crypto
	AuthMethodPostSupported() bool
}

type VerifyExchanger interface {
	Exchanger
	ClientJWTVerifier() rp.Verifier
}

func tokenHandler(exchanger Exchanger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("grant_type") {
		case string(oidc.GrantTypeCode):
			CodeExchange(w, r, exchanger)
			return
		case string(oidc.GrantTypeBearer):
			JWTExchange(w, r, exchanger)
			return
		case "excahnge":
			TokenExchange(w, r, exchanger)
		case "":
			RequestError(w, r, ErrInvalidRequest("grant_type missing"))
			return
		default:

		}
	}
}

func CodeExchange(w http.ResponseWriter, r *http.Request, exchanger Exchanger) {
	tokenReq, err := ParseAccessTokenRequest(r, exchanger.Decoder())
	if err != nil {
		RequestError(w, r, err)
	}
	if tokenReq.Code == "" {
		RequestError(w, r, ErrInvalidRequest("code missing"))
		return
	}
	authReq, client, err := ValidateAccessTokenRequest(r.Context(), tokenReq, exchanger)
	if err != nil {
		RequestError(w, r, err)
		return
	}
	resp, err := CreateTokenResponse(r.Context(), authReq, client, exchanger, true, tokenReq.Code)
	if err != nil {
		RequestError(w, r, err)
		return
	}
	utils.MarshalJSON(w, resp)
}

func ParseAccessTokenRequest(r *http.Request, decoder utils.Decoder) (*oidc.AccessTokenRequest, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, ErrInvalidRequest("error parsing form")
	}
	tokenReq := new(oidc.AccessTokenRequest)
	err = decoder.Decode(tokenReq, r.Form)
	if err != nil {
		return nil, ErrInvalidRequest("error decoding form")
	}
	clientID, clientSecret, ok := r.BasicAuth()
	if ok {
		tokenReq.ClientID = clientID
		tokenReq.ClientSecret = clientSecret

	}
	return tokenReq, nil
}

func ValidateAccessTokenRequest(ctx context.Context, tokenReq *oidc.AccessTokenRequest, exchanger Exchanger) (AuthRequest, Client, error) {
	authReq, client, err := AuthorizeClient(ctx, tokenReq, exchanger)
	if err != nil {
		return nil, nil, err
	}
	if client.GetID() != authReq.GetClientID() {
		return nil, nil, ErrInvalidRequest("invalid auth code")
	}
	if tokenReq.RedirectURI != authReq.GetRedirectURI() {
		return nil, nil, ErrInvalidRequest("redirect_uri does no correspond")
	}
	return authReq, client, nil
}

func AuthorizeClient(ctx context.Context, tokenReq *oidc.AccessTokenRequest, exchanger Exchanger) (AuthRequest, Client, error) {
	client, err := exchanger.Storage().GetClientByClientID(ctx, tokenReq.ClientID)
	if err != nil {
		return nil, nil, err
	}
	if client.AuthMethod() == AuthMethodNone {
		authReq, err := AuthorizeCodeChallenge(ctx, tokenReq, exchanger)
		return authReq, client, err
	}
	if client.AuthMethod() == AuthMethodPost && !exchanger.AuthMethodPostSupported() {
		return nil, nil, errors.New("basic not supported")
	}
	err = AuthorizeClientIDSecret(ctx, tokenReq.ClientID, tokenReq.ClientSecret, exchanger.Storage())
	if err != nil {
		return nil, nil, err
	}
	authReq, err := exchanger.Storage().AuthRequestByCode(ctx, tokenReq.Code)
	if err != nil {
		return nil, nil, ErrInvalidRequest("invalid code")
	}
	return authReq, client, nil
}

func AuthorizeClientIDSecret(ctx context.Context, clientID, clientSecret string, storage OPStorage) error {
	return storage.AuthorizeClientIDSecret(ctx, clientID, clientSecret)
}

func AuthorizeCodeChallenge(ctx context.Context, tokenReq *oidc.AccessTokenRequest, exchanger Exchanger) (AuthRequest, error) {
	if tokenReq.CodeVerifier == "" {
		return nil, ErrInvalidRequest("code_challenge required")
	}
	authReq, err := exchanger.Storage().AuthRequestByCode(ctx, tokenReq.Code)
	if err != nil {
		return nil, ErrInvalidRequest("invalid code")
	}
	if !oidc.VerifyCodeChallenge(authReq.GetCodeChallenge(), tokenReq.CodeVerifier) {
		return nil, ErrInvalidRequest("code_challenge invalid")
	}
	return authReq, nil
}

type ClientJWTVerifier struct {
	claims *oidc.JWTTokenRequest
	Storage
}

func (c ClientJWTVerifier) Storage() Storage {
	panic("implement me")
}

func (c ClientJWTVerifier) Issuer() string {
	panic("implement me")
}

func (c ClientJWTVerifier) ClientID() string {
	panic("implement me")
}

func (c ClientJWTVerifier) SupportedSignAlgs() []string {
	panic("implement me")
}

func (c ClientJWTVerifier) KeySet() oidc.KeySet {
	return c.claims
}

func (c ClientJWTVerifier) ACR() oidc.ACRVerifier {
	panic("implement me")
}

func (c ClientJWTVerifier) MaxAge() time.Duration {
	panic("implement me")
}

func (c ClientJWTVerifier) MaxAgeIAT() time.Duration {
	panic("implement me")
}

func (c ClientJWTVerifier) Offset() time.Duration {
	panic("implement me")
}

func JWTExchange(w http.ResponseWriter, r *http.Request, exchanger VerifyExchanger) {
	assertion, err := ParseJWTTokenRequest(r, exchanger.Decoder())
	if err != nil {
		RequestError(w, r, err)
	}
	claims := new(oidc.JWTTokenRequest)
	//var keyset oidc.KeySet
	verifier := new(ClientJWTVerifier)
	req, err := VerifyJWTAssertion(r.Context(), assertion, verifier)
	if err != nil {
		RequestError(w, r, err)
	}

	resp, err := CreateJWTTokenResponse(r.Context(), claims, exchanger)
	if err != nil {
		RequestError(w, r, err)
		return
	}
	utils.MarshalJSON(w, resp)
}

type JWTAssertionVerifier interface {
	Storage() Storage
	oidc.Verifier
}

func VerifyJWTAssertion(ctx context.Context, assertion string, v JWTAssertionVerifier) (*oidc.JWTTokenRequest, error) {
	claims := new(oidc.JWTTokenRequest)
	payload, err := oidc.ParseToken(assertion, claims)

	oidc.CheckAudience(claims.Audience, v)

	oidc.CheckExpiration(claims.ExpiresAt, v)

	oidc.CheckIssuedAt(claims.IssuedAt, v)

	if claims.Issuer != claims.Subject {

	}
	v.Storage().GetClientByClientID(ctx, claims.Issuer)

	keySet := &ClientAssertionKeySet{v.Storage(), claims.Issuer}

	oidc.CheckSignature(ctx, assertion, payload, claims, nil, keySet)
}

type ClientAssertionKeySet struct {
	Storage
	id string
}

func (c *ClientAssertionKeySet) VerifySignature(ctx context.Context, jws *jose.JSONWebSignature) (payload []byte, err error) {
	keyID := ""
	for _, sig := range jws.Signatures {
		keyID = sig.Header.KeyID
		break
	}
	keySet, err := c.Storage.GetKeysByServiceAccount(id)
	if err != nil {
		return nil, errors.New("error fetching keys")
	}
	payload, err, ok := rp.CheckKey(keyID, keySet.Keys, jws)
	if !ok {
		return nil, errors.New("invalid kid")
	}
	return payload, err
}

func ParseJWTTokenRequest(r *http.Request, decoder utils.Decoder) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", ErrInvalidRequest("error parsing form")
	}
	tokenReq := new(struct {
		Token string `schema:"assertion"`
	})
	err = decoder.Decode(tokenReq, r.Form)
	if err != nil {
		return "", ErrInvalidRequest("error decoding form")
	}
	//TODO: validations
	return tokenReq.Token, nil
}

func TokenExchange(w http.ResponseWriter, r *http.Request, exchanger Exchanger) {
	tokenRequest, err := ParseTokenExchangeRequest(w, r)
	if err != nil {
		RequestError(w, r, err)
		return
	}
	err = ValidateTokenExchangeRequest(tokenRequest, exchanger.Storage())
	if err != nil {
		RequestError(w, r, err)
		return
	}
}

func ParseTokenExchangeRequest(w http.ResponseWriter, r *http.Request) (oidc.TokenRequest, error) {
	return nil, errors.New("Unimplemented") //TODO: impl
}

func ValidateTokenExchangeRequest(tokenReq oidc.TokenRequest, storage Storage) error {
	return errors.New("Unimplemented") //TODO: impl
}
