package op

import "github.com/caos/oidc/pkg/oidc"

func NeedsExistingSession(authRequest *oidc.AuthRequest) bool {
	if authRequest == nil {
		return true
	}
	return authRequest.IDTokenHint != "" //TODO: impl: https://openid.net/specs/openid-connect-core-1_0.html#rfc.section.3.1.2.2
}
