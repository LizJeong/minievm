// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: minievm/evm/v1/query.proto

package evmv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Query_Code_FullMethodName                = "/minievm.evm.v1.Query/Code"
	Query_State_FullMethodName               = "/minievm.evm.v1.Query/State"
	Query_ContractAddrByDenom_FullMethodName = "/minievm.evm.v1.Query/ContractAddrByDenom"
	Query_Denom_FullMethodName               = "/minievm.evm.v1.Query/Denom"
	Query_Call_FullMethodName                = "/minievm.evm.v1.Query/Call"
	Query_Params_FullMethodName              = "/minievm.evm.v1.Query/Params"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Code gets the module info.
	Code(ctx context.Context, in *QueryCodeRequest, opts ...grpc.CallOption) (*QueryCodeResponse, error)
	// State gets the state bytes of the given address and key bytes.
	State(ctx context.Context, in *QueryStateRequest, opts ...grpc.CallOption) (*QueryStateResponse, error)
	ContractAddrByDenom(ctx context.Context, in *QueryContractAddrByDenomRequest, opts ...grpc.CallOption) (*QueryContractAddrByDenomResponse, error)
	Denom(ctx context.Context, in *QueryDenomRequest, opts ...grpc.CallOption) (*QueryDenomResponse, error)
	// Call execute entry function and return  the function result
	Call(ctx context.Context, in *QueryCallRequest, opts ...grpc.CallOption) (*QueryCallResponse, error)
	// Params queries all parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Code(ctx context.Context, in *QueryCodeRequest, opts ...grpc.CallOption) (*QueryCodeResponse, error) {
	out := new(QueryCodeResponse)
	err := c.cc.Invoke(ctx, Query_Code_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) State(ctx context.Context, in *QueryStateRequest, opts ...grpc.CallOption) (*QueryStateResponse, error) {
	out := new(QueryStateResponse)
	err := c.cc.Invoke(ctx, Query_State_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ContractAddrByDenom(ctx context.Context, in *QueryContractAddrByDenomRequest, opts ...grpc.CallOption) (*QueryContractAddrByDenomResponse, error) {
	out := new(QueryContractAddrByDenomResponse)
	err := c.cc.Invoke(ctx, Query_ContractAddrByDenom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Denom(ctx context.Context, in *QueryDenomRequest, opts ...grpc.CallOption) (*QueryDenomResponse, error) {
	out := new(QueryDenomResponse)
	err := c.cc.Invoke(ctx, Query_Denom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Call(ctx context.Context, in *QueryCallRequest, opts ...grpc.CallOption) (*QueryCallResponse, error) {
	out := new(QueryCallResponse)
	err := c.cc.Invoke(ctx, Query_Call_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Code gets the module info.
	Code(context.Context, *QueryCodeRequest) (*QueryCodeResponse, error)
	// State gets the state bytes of the given address and key bytes.
	State(context.Context, *QueryStateRequest) (*QueryStateResponse, error)
	ContractAddrByDenom(context.Context, *QueryContractAddrByDenomRequest) (*QueryContractAddrByDenomResponse, error)
	Denom(context.Context, *QueryDenomRequest) (*QueryDenomResponse, error)
	// Call execute entry function and return  the function result
	Call(context.Context, *QueryCallRequest) (*QueryCallResponse, error)
	// Params queries all parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Code(context.Context, *QueryCodeRequest) (*QueryCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Code not implemented")
}
func (UnimplementedQueryServer) State(context.Context, *QueryStateRequest) (*QueryStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (UnimplementedQueryServer) ContractAddrByDenom(context.Context, *QueryContractAddrByDenomRequest) (*QueryContractAddrByDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContractAddrByDenom not implemented")
}
func (UnimplementedQueryServer) Denom(context.Context, *QueryDenomRequest) (*QueryDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Denom not implemented")
}
func (UnimplementedQueryServer) Call(context.Context, *QueryCallRequest) (*QueryCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Code_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Code(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Code_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Code(ctx, req.(*QueryCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_State_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).State(ctx, req.(*QueryStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ContractAddrByDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryContractAddrByDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ContractAddrByDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ContractAddrByDenom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ContractAddrByDenom(ctx, req.(*QueryContractAddrByDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Denom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Denom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Denom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Denom(ctx, req.(*QueryDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Call_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Call(ctx, req.(*QueryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "minievm.evm.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Code",
			Handler:    _Query_Code_Handler,
		},
		{
			MethodName: "State",
			Handler:    _Query_State_Handler,
		},
		{
			MethodName: "ContractAddrByDenom",
			Handler:    _Query_ContractAddrByDenom_Handler,
		},
		{
			MethodName: "Denom",
			Handler:    _Query_Denom_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _Query_Call_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minievm/evm/v1/query.proto",
}
