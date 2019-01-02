package credential

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	TokenKey            = "token"
	TokenCredentialType = "TokenCredential"
)

type TokenCredential struct {
	Token string
}

func (t *TokenCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		TokenKey: t.Token,
		TypeKey:  TokenCredentialType,
	}, nil
}

func (t TokenCredential) RequireTransportSecurity() bool {
	return RequireTransportSecurity
}

func (t *TokenCredential) Parse(md metadata.MD) bool {
	if value, ok := md[TokenKey]; ok {
		t.Token = value[0]
		return true
	}

	return false
}

func NewTokenCredential(token string) *TokenCredential {
	return &TokenCredential{
		Token: token,
	}
}

func ParseTokenCredential(md metadata.MD) (*TokenCredential, bool) {
	obj := new(TokenCredential)
	if !obj.Parse(md) {
		return nil, false
	}
	return obj, true
}
