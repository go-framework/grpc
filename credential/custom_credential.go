package credential

import (
	"context"
)

const (
	CustomCredentialType = "CustomCredential"
)

type CustomCredential map[string]string

func (c CustomCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	if _, ok := c[TypeKey]; !ok {
		c[TypeKey] = CustomCredentialType
	}

	return c, nil
}

func (c CustomCredential) RequireTransportSecurity() bool {
	return RequireTransportSecurity
}
