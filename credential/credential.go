package credential

import "google.golang.org/grpc/metadata"

// global require transport security.
var RequireTransportSecurity = true

// credential type key.
const TypeKey = "credential_type"

// get the meta data type by TypeKey, if not have TypeKey then parse false.
func ParseType(md metadata.MD) (string, bool) {
	if kind, ok := md[TypeKey]; ok {
		return kind[0], true
	}
	return "", false
}
