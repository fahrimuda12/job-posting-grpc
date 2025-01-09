// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: protos/company/company.proto

package companypb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CompanyService_GetCompany_FullMethodName    = "/protos.company.CompanyService/GetCompany"
	CompanyService_DetailCompany_FullMethodName = "/protos.company.CompanyService/DetailCompany"
	CompanyService_CreateCompany_FullMethodName = "/protos.company.CompanyService/CreateCompany"
	CompanyService_UpdateCompany_FullMethodName = "/protos.company.CompanyService/UpdateCompany"
	CompanyService_DeleteCompany_FullMethodName = "/protos.company.CompanyService/DeleteCompany"
)

// CompanyServiceClient is the client API for CompanyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompanyServiceClient interface {
	GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error)
	DetailCompany(ctx context.Context, in *DetailCompanyRequest, opts ...grpc.CallOption) (*DetailCompanyResponse, error)
	CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error)
	UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error)
	DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*DeleteCompanyResponse, error)
}

type companyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompanyServiceClient(cc grpc.ClientConnInterface) CompanyServiceClient {
	return &companyServiceClient{cc}
}

func (c *companyServiceClient) GetCompany(ctx context.Context, in *GetCompanyRequest, opts ...grpc.CallOption) (*GetCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_GetCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DetailCompany(ctx context.Context, in *DetailCompanyRequest, opts ...grpc.CallOption) (*DetailCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_DetailCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) CreateCompany(ctx context.Context, in *CreateCompanyRequest, opts ...grpc.CallOption) (*CreateCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_CreateCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) UpdateCompany(ctx context.Context, in *UpdateCompanyRequest, opts ...grpc.CallOption) (*UpdateCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_UpdateCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyServiceClient) DeleteCompany(ctx context.Context, in *DeleteCompanyRequest, opts ...grpc.CallOption) (*DeleteCompanyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCompanyResponse)
	err := c.cc.Invoke(ctx, CompanyService_DeleteCompany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompanyServiceServer is the server API for CompanyService service.
// All implementations should embed UnimplementedCompanyServiceServer
// for forward compatibility.
type CompanyServiceServer interface {
	GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error)
	DetailCompany(context.Context, *DetailCompanyRequest) (*DetailCompanyResponse, error)
	CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error)
	UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error)
	DeleteCompany(context.Context, *DeleteCompanyRequest) (*DeleteCompanyResponse, error)
}

// UnimplementedCompanyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCompanyServiceServer struct{}

func (UnimplementedCompanyServiceServer) GetCompany(context.Context, *GetCompanyRequest) (*GetCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompany not implemented")
}
func (UnimplementedCompanyServiceServer) DetailCompany(context.Context, *DetailCompanyRequest) (*DetailCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailCompany not implemented")
}
func (UnimplementedCompanyServiceServer) CreateCompany(context.Context, *CreateCompanyRequest) (*CreateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompany not implemented")
}
func (UnimplementedCompanyServiceServer) UpdateCompany(context.Context, *UpdateCompanyRequest) (*UpdateCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompany not implemented")
}
func (UnimplementedCompanyServiceServer) DeleteCompany(context.Context, *DeleteCompanyRequest) (*DeleteCompanyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompany not implemented")
}
func (UnimplementedCompanyServiceServer) testEmbeddedByValue() {}

// UnsafeCompanyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompanyServiceServer will
// result in compilation errors.
type UnsafeCompanyServiceServer interface {
	mustEmbedUnimplementedCompanyServiceServer()
}

func RegisterCompanyServiceServer(s grpc.ServiceRegistrar, srv CompanyServiceServer) {
	// If the following call pancis, it indicates UnimplementedCompanyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CompanyService_ServiceDesc, srv)
}

func _CompanyService_GetCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).GetCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_GetCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).GetCompany(ctx, req.(*GetCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DetailCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DetailCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_DetailCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DetailCompany(ctx, req.(*DetailCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_CreateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).CreateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_CreateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).CreateCompany(ctx, req.(*CreateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_UpdateCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_UpdateCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).UpdateCompany(ctx, req.(*UpdateCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompanyService_DeleteCompany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CompanyService_DeleteCompany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServiceServer).DeleteCompany(ctx, req.(*DeleteCompanyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompanyService_ServiceDesc is the grpc.ServiceDesc for CompanyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompanyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.company.CompanyService",
	HandlerType: (*CompanyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompany",
			Handler:    _CompanyService_GetCompany_Handler,
		},
		{
			MethodName: "DetailCompany",
			Handler:    _CompanyService_DetailCompany_Handler,
		},
		{
			MethodName: "CreateCompany",
			Handler:    _CompanyService_CreateCompany_Handler,
		},
		{
			MethodName: "UpdateCompany",
			Handler:    _CompanyService_UpdateCompany_Handler,
		},
		{
			MethodName: "DeleteCompany",
			Handler:    _CompanyService_DeleteCompany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/company/company.proto",
}
