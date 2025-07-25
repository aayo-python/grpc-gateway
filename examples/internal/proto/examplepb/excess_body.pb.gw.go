// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: examples/internal/proto/examplepb/excess_body.proto

/*
Package examplepb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package examplepb

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_ExcessBodyService_NoBodyRpc_0(ctx context.Context, marshaler runtime.Marshaler, client ExcessBodyServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	if req.Body != nil {
		_, _ = io.Copy(io.Discard, req.Body)
	}
	msg, err := client.NoBodyRpc(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_ExcessBodyService_NoBodyRpc_0(ctx context.Context, marshaler runtime.Marshaler, server ExcessBodyServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	msg, err := server.NoBodyRpc(ctx, &protoReq)
	return msg, metadata, err
}

func request_ExcessBodyService_NoBodyServerStream_0(ctx context.Context, marshaler runtime.Marshaler, client ExcessBodyServiceClient, req *http.Request, pathParams map[string]string) (ExcessBodyService_NoBodyServerStreamClient, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	if req.Body != nil {
		_, _ = io.Copy(io.Discard, req.Body)
	}
	stream, err := client.NoBodyServerStream(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

func request_ExcessBodyService_WithBodyRpc_0(ctx context.Context, marshaler runtime.Marshaler, client ExcessBodyServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if req.Body != nil {
		_, _ = io.Copy(io.Discard, req.Body)
	}
	msg, err := client.WithBodyRpc(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_ExcessBodyService_WithBodyRpc_0(ctx context.Context, marshaler runtime.Marshaler, server ExcessBodyServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.WithBodyRpc(ctx, &protoReq)
	return msg, metadata, err
}

func request_ExcessBodyService_WithBodyServerStream_0(ctx context.Context, marshaler runtime.Marshaler, client ExcessBodyServiceClient, req *http.Request, pathParams map[string]string) (ExcessBodyService_WithBodyServerStreamClient, runtime.ServerMetadata, error) {
	var (
		protoReq emptypb.Empty
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if req.Body != nil {
		_, _ = io.Copy(io.Discard, req.Body)
	}
	stream, err := client.WithBodyServerStream(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

// RegisterExcessBodyServiceHandlerServer registers the http handlers for service ExcessBodyService to "mux".
// UnaryRPC     :call ExcessBodyServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterExcessBodyServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterExcessBodyServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ExcessBodyServiceServer) error {
	mux.Handle(http.MethodPost, pattern_ExcessBodyService_NoBodyRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/NoBodyRpc", runtime.WithHTTPPathPattern("/rpc/excess-body/rpc"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ExcessBodyService_NoBodyRpc_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_ExcessBodyService_NoBodyRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	mux.Handle(http.MethodPost, pattern_ExcessBodyService_NoBodyServerStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})
	mux.Handle(http.MethodPost, pattern_ExcessBodyService_WithBodyRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/WithBodyRpc", runtime.WithHTTPPathPattern("/rpc/excess-body/rpc/with-body"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ExcessBodyService_WithBodyRpc_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_ExcessBodyService_WithBodyRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	mux.Handle(http.MethodPost, pattern_ExcessBodyService_WithBodyServerStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterExcessBodyServiceHandlerFromEndpoint is same as RegisterExcessBodyServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterExcessBodyServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterExcessBodyServiceHandler(ctx, mux, conn)
}

// RegisterExcessBodyServiceHandler registers the http handlers for service ExcessBodyService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterExcessBodyServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterExcessBodyServiceHandlerClient(ctx, mux, NewExcessBodyServiceClient(conn))
}

// RegisterExcessBodyServiceHandlerClient registers the http handlers for service ExcessBodyService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ExcessBodyServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ExcessBodyServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ExcessBodyServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterExcessBodyServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ExcessBodyServiceClient) error {
	mux.Handle(http.MethodPost, pattern_ExcessBodyService_NoBodyRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/NoBodyRpc", runtime.WithHTTPPathPattern("/rpc/excess-body/rpc"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ExcessBodyService_NoBodyRpc_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_ExcessBodyService_NoBodyRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_ExcessBodyService_NoBodyServerStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/NoBodyServerStream", runtime.WithHTTPPathPattern("/rpc/excess-body/stream"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ExcessBodyService_NoBodyServerStream_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_ExcessBodyService_NoBodyServerStream_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_ExcessBodyService_WithBodyRpc_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/WithBodyRpc", runtime.WithHTTPPathPattern("/rpc/excess-body/rpc/with-body"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ExcessBodyService_WithBodyRpc_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_ExcessBodyService_WithBodyRpc_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_ExcessBodyService_WithBodyServerStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/grpc.gateway.examples.internal.proto.examplepb.ExcessBodyService/WithBodyServerStream", runtime.WithHTTPPathPattern("/rpc/excess-body/stream/with-body"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ExcessBodyService_WithBodyServerStream_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_ExcessBodyService_WithBodyServerStream_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_ExcessBodyService_NoBodyRpc_0            = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 0}, []string{"rpc", "excess-body"}, ""))
	pattern_ExcessBodyService_NoBodyServerStream_0   = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"rpc", "excess-body", "stream"}, ""))
	pattern_ExcessBodyService_WithBodyRpc_0          = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 0, 2, 2}, []string{"rpc", "excess-body", "with-body"}, ""))
	pattern_ExcessBodyService_WithBodyServerStream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"rpc", "excess-body", "stream", "with-body"}, ""))
)

var (
	forward_ExcessBodyService_NoBodyRpc_0            = runtime.ForwardResponseMessage
	forward_ExcessBodyService_NoBodyServerStream_0   = runtime.ForwardResponseStream
	forward_ExcessBodyService_WithBodyRpc_0          = runtime.ForwardResponseMessage
	forward_ExcessBodyService_WithBodyServerStream_0 = runtime.ForwardResponseStream
)
