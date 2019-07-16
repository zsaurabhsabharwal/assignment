// Code generated by protoc-gen-go. DO NOT EDIT.
// source: communications.proto

package communications

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type RestarauntInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rating               string   `protobuf:"bytes,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Cuisines             string   `protobuf:"bytes,3,opt,name=cuisines,proto3" json:"cuisines,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Time                 string   `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Cft                  string   `protobuf:"bytes,6,opt,name=cft,proto3" json:"cft,omitempty"`
	ID                   string   `protobuf:"bytes,7,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestarauntInfo) Reset()         { *m = RestarauntInfo{} }
func (m *RestarauntInfo) String() string { return proto.CompactTextString(m) }
func (*RestarauntInfo) ProtoMessage()    {}
func (*RestarauntInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3e14bb7aab8877, []int{0}
}

func (m *RestarauntInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestarauntInfo.Unmarshal(m, b)
}
func (m *RestarauntInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestarauntInfo.Marshal(b, m, deterministic)
}
func (m *RestarauntInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestarauntInfo.Merge(m, src)
}
func (m *RestarauntInfo) XXX_Size() int {
	return xxx_messageInfo_RestarauntInfo.Size(m)
}
func (m *RestarauntInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RestarauntInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RestarauntInfo proto.InternalMessageInfo

func (m *RestarauntInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RestarauntInfo) GetRating() string {
	if m != nil {
		return m.Rating
	}
	return ""
}

func (m *RestarauntInfo) GetCuisines() string {
	if m != nil {
		return m.Cuisines
	}
	return ""
}

func (m *RestarauntInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RestarauntInfo) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *RestarauntInfo) GetCft() string {
	if m != nil {
		return m.Cft
	}
	return ""
}

func (m *RestarauntInfo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type ListRestaraunt struct {
	List                 []*RestarauntInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListRestaraunt) Reset()         { *m = ListRestaraunt{} }
func (m *ListRestaraunt) String() string { return proto.CompactTextString(m) }
func (*ListRestaraunt) ProtoMessage()    {}
func (*ListRestaraunt) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3e14bb7aab8877, []int{1}
}

func (m *ListRestaraunt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRestaraunt.Unmarshal(m, b)
}
func (m *ListRestaraunt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRestaraunt.Marshal(b, m, deterministic)
}
func (m *ListRestaraunt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRestaraunt.Merge(m, src)
}
func (m *ListRestaraunt) XXX_Size() int {
	return xxx_messageInfo_ListRestaraunt.Size(m)
}
func (m *ListRestaraunt) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRestaraunt.DiscardUnknown(m)
}

var xxx_messageInfo_ListRestaraunt proto.InternalMessageInfo

func (m *ListRestaraunt) GetList() []*RestarauntInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type RestarauntID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestarauntID) Reset()         { *m = RestarauntID{} }
func (m *RestarauntID) String() string { return proto.CompactTextString(m) }
func (*RestarauntID) ProtoMessage()    {}
func (*RestarauntID) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3e14bb7aab8877, []int{2}
}

func (m *RestarauntID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestarauntID.Unmarshal(m, b)
}
func (m *RestarauntID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestarauntID.Marshal(b, m, deterministic)
}
func (m *RestarauntID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestarauntID.Merge(m, src)
}
func (m *RestarauntID) XXX_Size() int {
	return xxx_messageInfo_RestarauntID.Size(m)
}
func (m *RestarauntID) XXX_DiscardUnknown() {
	xxx_messageInfo_RestarauntID.DiscardUnknown(m)
}

var xxx_messageInfo_RestarauntID proto.InternalMessageInfo

func (m *RestarauntID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3e14bb7aab8877, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*RestarauntInfo)(nil), "communications.RestarauntInfo")
	proto.RegisterType((*ListRestaraunt)(nil), "communications.ListRestaraunt")
	proto.RegisterType((*RestarauntID)(nil), "communications.RestarauntID")
	proto.RegisterType((*Response)(nil), "communications.Response")
}

func init() { proto.RegisterFile("communications.proto", fileDescriptor_9a3e14bb7aab8877) }

var fileDescriptor_9a3e14bb7aab8877 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xed, 0x4a, 0xc3, 0x30,
	0x14, 0x5d, 0xbb, 0xb9, 0xcd, 0xab, 0x2b, 0xe3, 0x22, 0x12, 0x8a, 0x8c, 0x91, 0x5f, 0xfb, 0xb5,
	0x1f, 0xf3, 0x09, 0x84, 0x8a, 0x0c, 0x07, 0x4a, 0xdf, 0x20, 0x36, 0x99, 0x04, 0xd6, 0x64, 0xf4,
	0xde, 0xbe, 0x93, 0x6f, 0xe2, 0x6b, 0xc9, 0x62, 0x9d, 0x9d, 0x1f, 0xf3, 0x8f, 0xff, 0xee, 0x39,
	0x27, 0xb9, 0xe7, 0x9c, 0x10, 0xb8, 0x28, 0x7c, 0x59, 0xd6, 0xce, 0x16, 0x8a, 0xad, 0x77, 0x34,
	0xdf, 0x56, 0x9e, 0x3d, 0x26, 0x87, 0xac, 0x7c, 0x89, 0x20, 0xc9, 0x0d, 0xb1, 0xaa, 0x54, 0xed,
	0x78, 0xe9, 0xd6, 0x1e, 0x11, 0x7a, 0x4e, 0x95, 0x46, 0x44, 0xd3, 0x68, 0x76, 0x9a, 0x87, 0x19,
	0x2f, 0xa1, 0x5f, 0x29, 0xb6, 0xee, 0x59, 0xc4, 0x81, 0x6d, 0x10, 0xa6, 0x30, 0x2c, 0x6a, 0x4b,
	0xd6, 0x19, 0x12, 0xdd, 0xa0, 0xec, 0x31, 0x0a, 0x18, 0x28, 0xad, 0x2b, 0x43, 0x24, 0x7a, 0x41,
	0xfa, 0x80, 0x3b, 0x07, 0xb6, 0xa5, 0x11, 0x27, 0xef, 0x0e, 0xbb, 0x19, 0xc7, 0xd0, 0x2d, 0xd6,
	0x2c, 0xfa, 0x81, 0xda, 0x8d, 0x98, 0x40, 0xbc, 0xcc, 0xc4, 0x20, 0x10, 0xf1, 0x32, 0x93, 0x19,
	0x24, 0x2b, 0x4b, 0xfc, 0x99, 0x16, 0x17, 0xd0, 0xdb, 0x58, 0x62, 0x11, 0x4d, 0xbb, 0xb3, 0xb3,
	0xc5, 0x64, 0xfe, 0xa5, 0xf1, 0x61, 0xaf, 0x3c, 0x9c, 0x95, 0x13, 0x38, 0x6f, 0xf1, 0x59, 0xe3,
	0x12, 0xed, 0x5d, 0x24, 0x0c, 0x73, 0x43, 0x5b, 0xef, 0x28, 0xb4, 0x26, 0x56, 0x5c, 0x53, 0xa3,
	0x37, 0x68, 0xf1, 0x1a, 0x87, 0x6a, 0x46, 0x5b, 0xc6, 0x07, 0x18, 0xdd, 0x68, 0xdd, 0x0a, 0xf5,
	0x47, 0x8c, 0xf4, 0xea, 0x88, 0x9e, 0xc9, 0x0e, 0x3e, 0x42, 0x72, 0xab, 0x2d, 0xff, 0xe3, 0xc6,
	0x7b, 0x18, 0xdd, 0x99, 0xf6, 0x42, 0xf1, 0xc3, 0x85, 0xd0, 0x38, 0xfd, 0x66, 0x75, 0xf8, 0xe2,
	0xb2, 0x83, 0x2b, 0x18, 0x67, 0x66, 0x63, 0xd8, 0xb4, 0xf6, 0x1d, 0x0d, 0x90, 0xfe, 0xea, 0x26,
	0x3b, 0x4f, 0xfd, 0xf0, 0x2b, 0xaf, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x8a, 0x71, 0x0a,
	0xad, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddeditClient is the client API for Addedit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddeditClient interface {
	AddRestaraunt(ctx context.Context, in *RestarauntInfo, opts ...grpc.CallOption) (*RestarauntID, error)
	EditRestaraunt(ctx context.Context, in *RestarauntInfo, opts ...grpc.CallOption) (*RestarauntID, error)
	GetRestaraunt(ctx context.Context, in *Response, opts ...grpc.CallOption) (*ListRestaraunt, error)
	DeleteRestaraunt(ctx context.Context, in *RestarauntID, opts ...grpc.CallOption) (*Response, error)
}

type addeditClient struct {
	cc *grpc.ClientConn
}

func NewAddeditClient(cc *grpc.ClientConn) AddeditClient {
	return &addeditClient{cc}
}

func (c *addeditClient) AddRestaraunt(ctx context.Context, in *RestarauntInfo, opts ...grpc.CallOption) (*RestarauntID, error) {
	out := new(RestarauntID)
	err := c.cc.Invoke(ctx, "/communications.addedit/AddRestaraunt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addeditClient) EditRestaraunt(ctx context.Context, in *RestarauntInfo, opts ...grpc.CallOption) (*RestarauntID, error) {
	out := new(RestarauntID)
	err := c.cc.Invoke(ctx, "/communications.addedit/EditRestaraunt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addeditClient) GetRestaraunt(ctx context.Context, in *Response, opts ...grpc.CallOption) (*ListRestaraunt, error) {
	out := new(ListRestaraunt)
	err := c.cc.Invoke(ctx, "/communications.addedit/GetRestaraunt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addeditClient) DeleteRestaraunt(ctx context.Context, in *RestarauntID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/communications.addedit/DeleteRestaraunt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddeditServer is the server API for Addedit service.
type AddeditServer interface {
	AddRestaraunt(context.Context, *RestarauntInfo) (*RestarauntID, error)
	EditRestaraunt(context.Context, *RestarauntInfo) (*RestarauntID, error)
	GetRestaraunt(context.Context, *Response) (*ListRestaraunt, error)
	DeleteRestaraunt(context.Context, *RestarauntID) (*Response, error)
}

// UnimplementedAddeditServer can be embedded to have forward compatible implementations.
type UnimplementedAddeditServer struct {
}

func (*UnimplementedAddeditServer) AddRestaraunt(ctx context.Context, req *RestarauntInfo) (*RestarauntID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRestaraunt not implemented")
}
func (*UnimplementedAddeditServer) EditRestaraunt(ctx context.Context, req *RestarauntInfo) (*RestarauntID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditRestaraunt not implemented")
}
func (*UnimplementedAddeditServer) GetRestaraunt(ctx context.Context, req *Response) (*ListRestaraunt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaraunt not implemented")
}
func (*UnimplementedAddeditServer) DeleteRestaraunt(ctx context.Context, req *RestarauntID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRestaraunt not implemented")
}

func RegisterAddeditServer(s *grpc.Server, srv AddeditServer) {
	s.RegisterService(&_Addedit_serviceDesc, srv)
}

func _Addedit_AddRestaraunt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestarauntInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddeditServer).AddRestaraunt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.addedit/AddRestaraunt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddeditServer).AddRestaraunt(ctx, req.(*RestarauntInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addedit_EditRestaraunt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestarauntInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddeditServer).EditRestaraunt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.addedit/EditRestaraunt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddeditServer).EditRestaraunt(ctx, req.(*RestarauntInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addedit_GetRestaraunt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Response)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddeditServer).GetRestaraunt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.addedit/GetRestaraunt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddeditServer).GetRestaraunt(ctx, req.(*Response))
	}
	return interceptor(ctx, in, info, handler)
}

func _Addedit_DeleteRestaraunt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestarauntID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddeditServer).DeleteRestaraunt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communications.addedit/DeleteRestaraunt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddeditServer).DeleteRestaraunt(ctx, req.(*RestarauntID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Addedit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "communications.addedit",
	HandlerType: (*AddeditServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRestaraunt",
			Handler:    _Addedit_AddRestaraunt_Handler,
		},
		{
			MethodName: "EditRestaraunt",
			Handler:    _Addedit_EditRestaraunt_Handler,
		},
		{
			MethodName: "GetRestaraunt",
			Handler:    _Addedit_GetRestaraunt_Handler,
		},
		{
			MethodName: "DeleteRestaraunt",
			Handler:    _Addedit_DeleteRestaraunt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communications.proto",
}