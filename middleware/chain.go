package grpc_middleware

import (
	"context"

	"google.golang.org/grpc"
)

// ChainUnaryServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three, and three
// will see context changes of one and two.
func ChainUnaryServer(interceptors ... grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	n := len(interceptors)

	// interceptors length more than 1.
	if n > 1 {
		// chain callback function.
		var chain func(idx int, ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)

		chain = func(idx int, ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			if idx == len(interceptors) {
				return handler(ctx, req)
			}
			// return next interceptor.
			return interceptors[idx](ctx, req, info, func(ctx context.Context, req interface{}) (interface{}, error) {
				// return next chain.
				return chain(idx+1, ctx, req, info, handler)
			})
		}

		return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			return chain(0, ctx, req, info, handler)
		}
	} else if n == 1 {
		// interceptors length is 1.
		return interceptors[0]
	}
	// interceptors length is 0.
	// Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}

// ChainStreamServer creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryServer(one, two, three) will execute one before two before three.
// If you want to pass context between interceptors, use WrapServerStream.
func ChainStreamServer(interceptors ...grpc.StreamServerInterceptor) grpc.StreamServerInterceptor {
	n := len(interceptors)

	// interceptors length more than 1.
	if n > 1 {
		// chain callback function.
		var chain func(idx int, srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error

		chain = func(idx int, srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			if idx == len(interceptors) {
				return handler(srv, ss)
			}
			// return next interceptor.
			return interceptors[0](srv, ss, info, func(srv interface{}, stream grpc.ServerStream) error {
				// return next chain.
				return chain(idx+1, srv, stream, info, handler)
			})
		}

		return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			return chain(0, srv, ss, info, handler)
		}
	} else if n == 1 {
		// interceptors length is 1.
		return interceptors[0]
	}
	// interceptors length is 0.
	// Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, ss)
	}
}

// ChainUnaryClient creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainUnaryClient(one, two, three) will execute one before two before three.
func ChainUnaryClient(interceptors ...grpc.UnaryClientInterceptor) grpc.UnaryClientInterceptor {
	n := len(interceptors)

	// interceptors length more than 1.
	if n > 1 {
		// chain callback function.
		var chain func(idx int, ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error

		chain = func(idx int, ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			if idx == len(interceptors) {
				return invoker(ctx, method, req, reply, cc, opts...)
			}
			// return next interceptor.
			return interceptors[idx](ctx, method, req, reply, cc, func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
				// return next chain.
				return chain(idx+1, ctx, method, req, reply, cc, invoker, opts...)
			}, opts...)
		}

		return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			return chain(0, ctx, method, req, reply, cc, invoker, opts...)
		}
	} else if n == 1 {
		// interceptors length is 1.
		return interceptors[0]
	}
	// interceptors length is 0.
	// Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// ChainStreamClient creates a single interceptor out of a chain of many interceptors.
//
// Execution is done in left-to-right order, including passing of context.
// For example ChainStreamClient(one, two, three) will execute one before two before three.
func ChainStreamClient(interceptors ...grpc.StreamClientInterceptor) grpc.StreamClientInterceptor {
	n := len(interceptors)

	// interceptors length more than 1.
	if n > 1 {
		// chain callback function.
		var chain func(idx int, ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error)

		chain = func(idx int, ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			if idx == len(interceptors) {
				return streamer(ctx, desc, cc, method, opts...)
			}
			// return next interceptor.
			return interceptors[idx](ctx, desc, cc, method, func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
				// return next chain.
				return chain(idx+1, ctx, desc, cc, method, streamer, opts...)
			}, opts...)
		}

		return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
			return chain(0, ctx, desc, cc, method, streamer, opts...)
		}
	} else if n == 1 {
		// interceptors length is 1.
		return interceptors[0]
	}
	// interceptors length is 0.
	// Dummy interceptor maintained for backward compatibility to avoid returning nil.
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(ctx, desc, cc, method, opts...)
	}
}

// Chain creates a single interceptor out of a chain of many interceptors.
//
// WithUnaryServerChain is a grpc.Server config option that accepts multiple unary interceptors.
// Basically syntactic sugar.
func WithUnaryServerChain(interceptors ...grpc.UnaryServerInterceptor) grpc.ServerOption {
	return grpc.UnaryInterceptor(ChainUnaryServer(interceptors...))
}

// WithStreamServerChain is a grpc.Server config option that accepts multiple stream interceptors.
// Basically syntactic sugar.
func WithStreamServerChain(interceptors ...grpc.StreamServerInterceptor) grpc.ServerOption {
	return grpc.StreamInterceptor(ChainStreamServer(interceptors...))
}

// Chain creates a single interceptor out of a chain of many interceptors.
//
// WithUnaryClientChain is a grpc.DialOption dial option that accepts multiple unary interceptors.
// Basically syntactic sugar.
func WithUnaryClientChain(interceptors ...grpc.UnaryClientInterceptor) grpc.DialOption {
	return grpc.WithUnaryInterceptor(ChainUnaryClient(interceptors...))
}

// Chain creates a single interceptor out of a chain of many interceptors.
//
// WithStreamInterceptor is a grpc.DialOption dial option that accepts multiple unary interceptors.
// Basically syntactic sugar.
func WithStreamClientChain(interceptors ...grpc.StreamClientInterceptor) grpc.DialOption {
	return grpc.WithStreamInterceptor(ChainStreamClient(interceptors...))
}
