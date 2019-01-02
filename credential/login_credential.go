package credential

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	Username            = "username"
	Password            = "password"
	LoginCredentialType = "LoginCredential"
)

type LoginCredential struct {
	Username string
	Password string
}

func (c *LoginCredential) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		Username: c.Username,
		Password: c.Password,
		TypeKey:  LoginCredentialType,
	}, nil
}

func (c LoginCredential) RequireTransportSecurity() bool {
	return RequireTransportSecurity
}

func (c *LoginCredential) Parse(md metadata.MD) bool {
	if value, ok := md[Username]; ok {
		c.Username = value[0]
	} else {
		return false
	}
	if value, ok := md[Password]; ok {
		c.Password = value[0]
	} else {
		return false
	}

	return true
}

func NewLoginCredential(username, password string) *LoginCredential {
	return &LoginCredential{
		Username: username,
		Password: password,
	}
}

func ParseLoginCredential(md metadata.MD) (*LoginCredential, bool) {
	obj := new(LoginCredential)
	if !obj.Parse(md) {
		return nil, false
	}
	return obj, true
}
