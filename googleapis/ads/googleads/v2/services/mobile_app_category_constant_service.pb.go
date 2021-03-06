// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/mobile_app_category_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Request message for
// [MobileAppCategoryConstantService.GetMobileAppCategoryConstant][google.ads.googleads.v2.services.MobileAppCategoryConstantService.GetMobileAppCategoryConstant].
type GetMobileAppCategoryConstantRequest struct {
	// Resource name of the mobile app category constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMobileAppCategoryConstantRequest) Reset()         { *m = GetMobileAppCategoryConstantRequest{} }
func (m *GetMobileAppCategoryConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetMobileAppCategoryConstantRequest) ProtoMessage()    {}
func (*GetMobileAppCategoryConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d705bc81423cc8aa, []int{0}
}

func (m *GetMobileAppCategoryConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMobileAppCategoryConstantRequest.Unmarshal(m, b)
}
func (m *GetMobileAppCategoryConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMobileAppCategoryConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetMobileAppCategoryConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMobileAppCategoryConstantRequest.Merge(m, src)
}
func (m *GetMobileAppCategoryConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetMobileAppCategoryConstantRequest.Size(m)
}
func (m *GetMobileAppCategoryConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMobileAppCategoryConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMobileAppCategoryConstantRequest proto.InternalMessageInfo

func (m *GetMobileAppCategoryConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMobileAppCategoryConstantRequest)(nil), "google.ads.googleads.v2.services.GetMobileAppCategoryConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/mobile_app_category_constant_service.proto", fileDescriptor_d705bc81423cc8aa)
}

var fileDescriptor_d705bc81423cc8aa = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x4b, 0xf3, 0x40,
	0x1c, 0xc7, 0x49, 0x1e, 0x78, 0xe0, 0x09, 0x8f, 0x4b, 0x16, 0x4b, 0xec, 0x10, 0xda, 0x0a, 0xe2,
	0x70, 0x27, 0x71, 0x91, 0x53, 0x87, 0xb4, 0x4a, 0x45, 0x51, 0x4a, 0x85, 0x0e, 0x12, 0x08, 0xd7,
	0xe4, 0x08, 0x81, 0xe4, 0x2e, 0xe6, 0xae, 0x05, 0x11, 0x17, 0x67, 0x37, 0x5f, 0x80, 0xe0, 0xe8,
	0x4b, 0xe9, 0xea, 0xec, 0xe6, 0xe4, 0xab, 0x90, 0x34, 0x77, 0xa9, 0x0e, 0x31, 0x6e, 0x5f, 0xee,
	0xbe, 0xf7, 0xf9, 0xfe, 0xfe, 0x9c, 0x71, 0x16, 0x31, 0x16, 0x25, 0x04, 0xe2, 0x90, 0xc3, 0x52,
	0x16, 0x6a, 0xee, 0x40, 0x4e, 0xf2, 0x79, 0x1c, 0x10, 0x0e, 0x53, 0x36, 0x8d, 0x13, 0xe2, 0xe3,
	0x2c, 0xf3, 0x03, 0x2c, 0x48, 0xc4, 0xf2, 0x1b, 0x3f, 0x60, 0x94, 0x0b, 0x4c, 0x85, 0x2f, 0x5d,
	0x20, 0xcb, 0x99, 0x60, 0xa6, 0x5d, 0x12, 0x00, 0x0e, 0x39, 0xa8, 0x60, 0x60, 0xee, 0x00, 0x05,
	0xb3, 0x8e, 0xea, 0xe2, 0x72, 0xc2, 0xd9, 0x2c, 0x6f, 0xca, 0x2b, 0x73, 0xac, 0xb6, 0xa2, 0x64,
	0x31, 0xc4, 0x94, 0x32, 0x81, 0x45, 0xcc, 0x28, 0x97, 0xb7, 0xeb, 0x5f, 0x6e, 0x83, 0x24, 0x26,
	0xea, 0x59, 0xe7, 0xd4, 0xe8, 0x0e, 0x89, 0x38, 0x5f, 0xf2, 0xdd, 0x2c, 0x1b, 0x48, 0xfa, 0x40,
	0xc2, 0xc7, 0xe4, 0x7a, 0x46, 0xb8, 0x30, 0xbb, 0xc6, 0x9a, 0xaa, 0xc6, 0xa7, 0x38, 0x25, 0x2d,
	0xcd, 0xd6, 0xb6, 0xfe, 0x8d, 0xff, 0xab, 0xc3, 0x0b, 0x9c, 0x12, 0xe7, 0x49, 0x37, 0xec, 0x5a,
	0xd2, 0x65, 0xd9, 0xae, 0xf9, 0xa6, 0x19, 0xed, 0x9f, 0x12, 0xcd, 0x63, 0xd0, 0x34, 0x31, 0xf0,
	0x8b, 0x8a, 0xad, 0x83, 0x5a, 0x4c, 0x35, 0x56, 0x50, 0x0b, 0xe9, 0xec, 0xdd, 0xbf, 0xbe, 0x3f,
	0xea, 0x8e, 0xb9, 0x53, 0xec, 0xe1, 0xf6, 0x5b, 0xeb, 0x87, 0x69, 0xdd, 0x2b, 0x0e, 0xb7, 0xef,
	0xac, 0x8d, 0x85, 0xdb, 0x5a, 0xc5, 0x49, 0x95, 0xc5, 0x1c, 0x04, 0x2c, 0xed, 0x3f, 0xe8, 0x46,
	0x2f, 0x60, 0x69, 0x63, 0x87, 0xfd, 0xcd, 0xa6, 0x39, 0x8e, 0x8a, 0xed, 0x8d, 0xb4, 0xab, 0x13,
	0x89, 0x8a, 0x58, 0x82, 0x69, 0x04, 0x58, 0x1e, 0xc1, 0x88, 0xd0, 0xe5, 0x6e, 0xe1, 0x2a, 0xbc,
	0xfe, 0x2b, 0xef, 0x2b, 0xf1, 0xac, 0xff, 0x19, 0xba, 0xee, 0x8b, 0x6e, 0x0f, 0x4b, 0xa0, 0x1b,
	0x72, 0x50, 0xca, 0x42, 0x4d, 0x1c, 0x20, 0x83, 0xf9, 0x42, 0x59, 0x3c, 0x37, 0xe4, 0x5e, 0x65,
	0xf1, 0x26, 0x8e, 0xa7, 0x2c, 0x1f, 0x7a, 0xaf, 0x3c, 0x47, 0xc8, 0x0d, 0x39, 0x42, 0x95, 0x09,
	0xa1, 0x89, 0x83, 0x90, 0xb2, 0x4d, 0xff, 0x2e, 0xeb, 0xdc, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x28, 0x74, 0xb0, 0x71, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MobileAppCategoryConstantServiceClient is the client API for MobileAppCategoryConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MobileAppCategoryConstantServiceClient interface {
	// Returns the requested mobile app category constant.
	GetMobileAppCategoryConstant(ctx context.Context, in *GetMobileAppCategoryConstantRequest, opts ...grpc.CallOption) (*resources.MobileAppCategoryConstant, error)
}

type mobileAppCategoryConstantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileAppCategoryConstantServiceClient(cc grpc.ClientConnInterface) MobileAppCategoryConstantServiceClient {
	return &mobileAppCategoryConstantServiceClient{cc}
}

func (c *mobileAppCategoryConstantServiceClient) GetMobileAppCategoryConstant(ctx context.Context, in *GetMobileAppCategoryConstantRequest, opts ...grpc.CallOption) (*resources.MobileAppCategoryConstant, error) {
	out := new(resources.MobileAppCategoryConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.MobileAppCategoryConstantService/GetMobileAppCategoryConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileAppCategoryConstantServiceServer is the server API for MobileAppCategoryConstantService service.
type MobileAppCategoryConstantServiceServer interface {
	// Returns the requested mobile app category constant.
	GetMobileAppCategoryConstant(context.Context, *GetMobileAppCategoryConstantRequest) (*resources.MobileAppCategoryConstant, error)
}

// UnimplementedMobileAppCategoryConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMobileAppCategoryConstantServiceServer struct {
}

func (*UnimplementedMobileAppCategoryConstantServiceServer) GetMobileAppCategoryConstant(ctx context.Context, req *GetMobileAppCategoryConstantRequest) (*resources.MobileAppCategoryConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMobileAppCategoryConstant not implemented")
}

func RegisterMobileAppCategoryConstantServiceServer(s *grpc.Server, srv MobileAppCategoryConstantServiceServer) {
	s.RegisterService(&_MobileAppCategoryConstantService_serviceDesc, srv)
}

func _MobileAppCategoryConstantService_GetMobileAppCategoryConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMobileAppCategoryConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileAppCategoryConstantServiceServer).GetMobileAppCategoryConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.MobileAppCategoryConstantService/GetMobileAppCategoryConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileAppCategoryConstantServiceServer).GetMobileAppCategoryConstant(ctx, req.(*GetMobileAppCategoryConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MobileAppCategoryConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.MobileAppCategoryConstantService",
	HandlerType: (*MobileAppCategoryConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMobileAppCategoryConstant",
			Handler:    _MobileAppCategoryConstantService_GetMobileAppCategoryConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/mobile_app_category_constant_service.proto",
}
