package credential

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	AppKey            = "app_key"
	AppSecret         = "app_secret"
	APICredentialType = "APICredential"
)

type APICredential struct {
	AppKey    string
	AppSecret string
}

func (a *APICredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		AppKey:    a.AppKey,
		AppSecret: a.AppSecret,
		TypeKey:   APICredentialType,
	}, nil
}

func (a APICredential) RequireTransportSecurity() bool {
	return RequireTransportSecurity
}

func (a *APICredential) Parse(md metadata.MD) bool {
	if value, ok := md[AppKey]; ok {
		a.AppKey = value[0]
	} else {
		return false
	}
	if value, ok := md[AppSecret]; ok {
		a.AppSecret = value[0]
	} else {
		return false
	}

	return true
}

func NewAPICredential(key, secret string) *APICredential {
	return &APICredential{
		AppKey:    key,
		AppSecret: secret,
	}
}

func ParseAPICredential(md metadata.MD) (*APICredential, bool) {
	obj := new(APICredential)
	if !obj.Parse(md) {
		return nil, false
	}
	return obj, true
}
