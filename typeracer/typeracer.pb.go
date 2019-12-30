// Code generated by protoc-gen-go. DO NOT EDIT.
// source: typeracer.proto

package typeracer

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

type PlayerMetrics struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	StrokesPerMinute     int32    `protobuf:"varint,2,opt,name=strokesPerMinute,proto3" json:"strokesPerMinute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerMetrics) Reset()         { *m = PlayerMetrics{} }
func (m *PlayerMetrics) String() string { return proto.CompactTextString(m) }
func (*PlayerMetrics) ProtoMessage()    {}
func (*PlayerMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf4c6564c4848df5, []int{0}
}

func (m *PlayerMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerMetrics.Unmarshal(m, b)
}
func (m *PlayerMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerMetrics.Marshal(b, m, deterministic)
}
func (m *PlayerMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerMetrics.Merge(m, src)
}
func (m *PlayerMetrics) XXX_Size() int {
	return xxx_messageInfo_PlayerMetrics.Size(m)
}
func (m *PlayerMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerMetrics proto.InternalMessageInfo

func (m *PlayerMetrics) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PlayerMetrics) GetStrokesPerMinute() int32 {
	if m != nil {
		return m.StrokesPerMinute
	}
	return 0
}

type Scoreboard struct {
	Reply                []*PlayerMetrics `protobuf:"bytes,1,rep,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Scoreboard) Reset()         { *m = Scoreboard{} }
func (m *Scoreboard) String() string { return proto.CompactTextString(m) }
func (*Scoreboard) ProtoMessage()    {}
func (*Scoreboard) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf4c6564c4848df5, []int{1}
}

func (m *Scoreboard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scoreboard.Unmarshal(m, b)
}
func (m *Scoreboard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scoreboard.Marshal(b, m, deterministic)
}
func (m *Scoreboard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scoreboard.Merge(m, src)
}
func (m *Scoreboard) XXX_Size() int {
	return xxx_messageInfo_Scoreboard.Size(m)
}
func (m *Scoreboard) XXX_DiscardUnknown() {
	xxx_messageInfo_Scoreboard.DiscardUnknown(m)
}

var xxx_messageInfo_Scoreboard proto.InternalMessageInfo

func (m *Scoreboard) GetReply() []*PlayerMetrics {
	if m != nil {
		return m.Reply
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayerMetrics)(nil), "typeracer.PlayerMetrics")
	proto.RegisterType((*Scoreboard)(nil), "typeracer.Scoreboard")
}

func init() { proto.RegisterFile("typeracer.proto", fileDescriptor_bf4c6564c4848df5) }

var fileDescriptor_bf4c6564c4848df5 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xa9, 0x2c, 0x48,
	0x2d, 0x4a, 0x4c, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0x85, 0x73, 0xf1, 0x06, 0xe4, 0x24, 0x56, 0xa6, 0x16, 0xf9, 0xa6, 0x96, 0x14, 0x65, 0x26, 0x17,
	0x0b, 0x49, 0x71, 0x71, 0x94, 0x16, 0xa7, 0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0x5a, 0x5c, 0x02, 0xc5, 0x25, 0x45, 0xf9, 0xd9, 0xa9, 0xc5,
	0x01, 0xa9, 0x45, 0xbe, 0x99, 0x79, 0xa5, 0x25, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41,
	0x18, 0xe2, 0x4a, 0x36, 0x5c, 0x5c, 0xc1, 0xc9, 0xf9, 0x45, 0xa9, 0x49, 0xf9, 0x89, 0x45, 0x29,
	0x42, 0x7a, 0x5c, 0xac, 0x45, 0xa9, 0x05, 0x39, 0x95, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46,
	0x12, 0x7a, 0x08, 0x27, 0xa1, 0x58, 0x1f, 0x04, 0x51, 0x66, 0x14, 0xce, 0xc5, 0x19, 0x52, 0x59,
	0x90, 0x1a, 0x04, 0x52, 0x21, 0xe4, 0xc5, 0x25, 0x18, 0x9c, 0x9a, 0x97, 0x82, 0xea, 0x4e, 0x9c,
	0x46, 0x48, 0x89, 0x22, 0xc9, 0x20, 0x9c, 0xa0, 0xc4, 0xa0, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06,
	0x0e, 0x01, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x82, 0xd1, 0x9f, 0x14, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TypeRacerClient is the client API for TypeRacer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TypeRacerClient interface {
	SendPlayerMetrics(ctx context.Context, opts ...grpc.CallOption) (TypeRacer_SendPlayerMetricsClient, error)
}

type typeRacerClient struct {
	cc *grpc.ClientConn
}

func NewTypeRacerClient(cc *grpc.ClientConn) TypeRacerClient {
	return &typeRacerClient{cc}
}

func (c *typeRacerClient) SendPlayerMetrics(ctx context.Context, opts ...grpc.CallOption) (TypeRacer_SendPlayerMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TypeRacer_serviceDesc.Streams[0], "/typeracer.TypeRacer/SendPlayerMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &typeRacerSendPlayerMetricsClient{stream}
	return x, nil
}

type TypeRacer_SendPlayerMetricsClient interface {
	Send(*PlayerMetrics) error
	Recv() (*Scoreboard, error)
	grpc.ClientStream
}

type typeRacerSendPlayerMetricsClient struct {
	grpc.ClientStream
}

func (x *typeRacerSendPlayerMetricsClient) Send(m *PlayerMetrics) error {
	return x.ClientStream.SendMsg(m)
}

func (x *typeRacerSendPlayerMetricsClient) Recv() (*Scoreboard, error) {
	m := new(Scoreboard)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TypeRacerServer is the server API for TypeRacer service.
type TypeRacerServer interface {
	SendPlayerMetrics(TypeRacer_SendPlayerMetricsServer) error
}

// UnimplementedTypeRacerServer can be embedded to have forward compatible implementations.
type UnimplementedTypeRacerServer struct {
}

func (*UnimplementedTypeRacerServer) SendPlayerMetrics(srv TypeRacer_SendPlayerMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendPlayerMetrics not implemented")
}

func RegisterTypeRacerServer(s *grpc.Server, srv TypeRacerServer) {
	s.RegisterService(&_TypeRacer_serviceDesc, srv)
}

func _TypeRacer_SendPlayerMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TypeRacerServer).SendPlayerMetrics(&typeRacerSendPlayerMetricsServer{stream})
}

type TypeRacer_SendPlayerMetricsServer interface {
	Send(*Scoreboard) error
	Recv() (*PlayerMetrics, error)
	grpc.ServerStream
}

type typeRacerSendPlayerMetricsServer struct {
	grpc.ServerStream
}

func (x *typeRacerSendPlayerMetricsServer) Send(m *Scoreboard) error {
	return x.ServerStream.SendMsg(m)
}

func (x *typeRacerSendPlayerMetricsServer) Recv() (*PlayerMetrics, error) {
	m := new(PlayerMetrics)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TypeRacer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "typeracer.TypeRacer",
	HandlerType: (*TypeRacerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendPlayerMetrics",
			Handler:       _TypeRacer_SendPlayerMetrics_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "typeracer.proto",
}