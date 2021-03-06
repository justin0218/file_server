// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/proto.proto

package proto

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

type R struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *R) Reset()         { *m = R{} }
func (m *R) String() string { return proto.CompactTextString(m) }
func (*R) ProtoMessage()    {}
func (*R) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef32c37ea206d67b, []int{0}
}

func (m *R) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_R.Unmarshal(m, b)
}
func (m *R) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_R.Marshal(b, m, deterministic)
}
func (m *R) XXX_Merge(src proto.Message) {
	xxx_messageInfo_R.Merge(m, src)
}
func (m *R) XXX_Size() int {
	return xxx_messageInfo_R.Size(m)
}
func (m *R) XXX_DiscardUnknown() {
	xxx_messageInfo_R.DiscardUnknown(m)
}

var xxx_messageInfo_R proto.InternalMessageInfo

func (m *R) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *R) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UploadLocalReq struct {
	FileBytes            []byte   `protobuf:"bytes,1,opt,name=file_bytes,json=fileBytes,proto3" json:"file_bytes,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Format               string   `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadLocalReq) Reset()         { *m = UploadLocalReq{} }
func (m *UploadLocalReq) String() string { return proto.CompactTextString(m) }
func (*UploadLocalReq) ProtoMessage()    {}
func (*UploadLocalReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef32c37ea206d67b, []int{1}
}

func (m *UploadLocalReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadLocalReq.Unmarshal(m, b)
}
func (m *UploadLocalReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadLocalReq.Marshal(b, m, deterministic)
}
func (m *UploadLocalReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadLocalReq.Merge(m, src)
}
func (m *UploadLocalReq) XXX_Size() int {
	return xxx_messageInfo_UploadLocalReq.Size(m)
}
func (m *UploadLocalReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadLocalReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadLocalReq proto.InternalMessageInfo

func (m *UploadLocalReq) GetFileBytes() []byte {
	if m != nil {
		return m.FileBytes
	}
	return nil
}

func (m *UploadLocalReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadLocalReq) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

type UploadLocalRes struct {
	Res                  *R       `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Format               string   `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	Size                 int64    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Url                  string   `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadLocalRes) Reset()         { *m = UploadLocalRes{} }
func (m *UploadLocalRes) String() string { return proto.CompactTextString(m) }
func (*UploadLocalRes) ProtoMessage()    {}
func (*UploadLocalRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef32c37ea206d67b, []int{2}
}

func (m *UploadLocalRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadLocalRes.Unmarshal(m, b)
}
func (m *UploadLocalRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadLocalRes.Marshal(b, m, deterministic)
}
func (m *UploadLocalRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadLocalRes.Merge(m, src)
}
func (m *UploadLocalRes) XXX_Size() int {
	return xxx_messageInfo_UploadLocalRes.Size(m)
}
func (m *UploadLocalRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadLocalRes.DiscardUnknown(m)
}

var xxx_messageInfo_UploadLocalRes proto.InternalMessageInfo

func (m *UploadLocalRes) GetRes() *R {
	if m != nil {
		return m.Res
	}
	return nil
}

func (m *UploadLocalRes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadLocalRes) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *UploadLocalRes) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *UploadLocalRes) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*R)(nil), "r")
	proto.RegisterType((*UploadLocalReq)(nil), "upload_local_req")
	proto.RegisterType((*UploadLocalRes)(nil), "upload_local_res")
}

func init() { proto.RegisterFile("api/proto/proto.proto", fileDescriptor_ef32c37ea206d67b) }

var fileDescriptor_ef32c37ea206d67b = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x90, 0x7a, 0x60, 0x52, 0x49, 0x93, 0x8b, 0xb1, 0x48, 0x48,
	0x88, 0x8b, 0x25, 0x39, 0x3f, 0x25, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x39, 0x08, 0xcc, 0x16,
	0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95,
	0x62, 0xb9, 0x04, 0x4a, 0x0b, 0x72, 0xf2, 0x13, 0x53, 0xe2, 0x73, 0xf2, 0x93, 0x13, 0x73, 0xe2,
	0x8b, 0x52, 0x0b, 0x85, 0x64, 0xb9, 0xb8, 0xd2, 0x32, 0x73, 0x52, 0xe3, 0x93, 0x2a, 0x4b, 0x52,
	0x8b, 0xc1, 0xfa, 0x79, 0x82, 0x38, 0x41, 0x22, 0x4e, 0x20, 0x01, 0x90, 0xc1, 0x79, 0x89, 0xb9,
	0xa9, 0x50, 0x53, 0xc0, 0x6c, 0x21, 0x31, 0x2e, 0xb6, 0xb4, 0xfc, 0xa2, 0xdc, 0xc4, 0x12, 0x09,
	0x66, 0xb0, 0x28, 0x94, 0xa7, 0x54, 0x85, 0x61, 0x7c, 0xb1, 0x90, 0x08, 0x17, 0x73, 0x11, 0xd4,
	0x5c, 0x6e, 0x23, 0x26, 0xbd, 0xa2, 0x20, 0x10, 0x97, 0x14, 0x53, 0x41, 0x6a, 0x8b, 0x33, 0xab,
	0x52, 0x25, 0x58, 0x20, 0x5e, 0x03, 0xb1, 0x41, 0x5e, 0x2b, 0x2d, 0xca, 0x91, 0x60, 0x85, 0x78,
	0xad, 0xb4, 0x28, 0xc7, 0xc8, 0x8e, 0x8b, 0x05, 0xe4, 0x68, 0x21, 0x33, 0x2e, 0x1e, 0x64, 0x37,
	0x08, 0x09, 0xea, 0xa1, 0xfb, 0x58, 0x0a, 0x43, 0xa8, 0x58, 0x89, 0xc1, 0x89, 0x3d, 0x8a, 0x15,
	0x1c, 0x9c, 0x49, 0x6c, 0x60, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x48, 0xb9, 0x83, 0x83,
	0x6e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileClient interface {
	UploadLocal(ctx context.Context, in *UploadLocalReq, opts ...grpc.CallOption) (*UploadLocalRes, error)
}

type fileClient struct {
	cc *grpc.ClientConn
}

func NewFileClient(cc *grpc.ClientConn) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) UploadLocal(ctx context.Context, in *UploadLocalReq, opts ...grpc.CallOption) (*UploadLocalRes, error) {
	out := new(UploadLocalRes)
	err := c.cc.Invoke(ctx, "/file/upload_local", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
type FileServer interface {
	UploadLocal(context.Context, *UploadLocalReq) (*UploadLocalRes, error)
}

// UnimplementedFileServer can be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (*UnimplementedFileServer) UploadLocal(ctx context.Context, req *UploadLocalReq) (*UploadLocalRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadLocal not implemented")
}

func RegisterFileServer(s *grpc.Server, srv FileServer) {
	s.RegisterService(&_File_serviceDesc, srv)
}

func _File_UploadLocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadLocalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).UploadLocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file/UploadLocal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).UploadLocal(ctx, req.(*UploadLocalReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _File_serviceDesc = grpc.ServiceDesc{
	ServiceName: "file",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "upload_local",
			Handler:    _File_UploadLocal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/proto.proto",
}
