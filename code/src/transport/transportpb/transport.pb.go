// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transport/transportpb/transport.proto

package transportpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	peerpb "github.com/nithinkatla00/RL-MultiLeaderProtocolSwitcher/peer/peerpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TransMsg struct {
	Msgs []peerpb.Message `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs"`
}

func (m *TransMsg) Reset()         { *m = TransMsg{} }
func (m *TransMsg) String() string { return proto.CompactTextString(m) }
func (*TransMsg) ProtoMessage()    {}
func (*TransMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7c1f8429b4b9ad, []int{0}
}
func (m *TransMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransMsg.Merge(m, src)
}
func (m *TransMsg) XXX_Size() int {
	return m.Size()
}
func (m *TransMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TransMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TransMsg proto.InternalMessageInfo

func (m *TransMsg) GetMsgs() []peerpb.Message {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// Empty is an empty message. It is identical to google.protobuf.Empty, but
// permits future modifications because it is custom.
type Empty struct {
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7c1f8429b4b9ad, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return m.Size()
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ClientPacket struct {
	Message   []byte `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ClientPacket) Reset()         { *m = ClientPacket{} }
func (m *ClientPacket) String() string { return proto.CompactTextString(m) }
func (*ClientPacket) ProtoMessage()    {}
func (*ClientPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7c1f8429b4b9ad, []int{2}
}
func (m *ClientPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientPacket.Merge(m, src)
}
func (m *ClientPacket) XXX_Size() int {
	return m.Size()
}
func (m *ClientPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientPacket.DiscardUnknown(m)
}

var xxx_messageInfo_ClientPacket proto.InternalMessageInfo

func (m *ClientPacket) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ClientPacket) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*TransMsg)(nil), "transportpb.TransMsg")
	proto.RegisterType((*Empty)(nil), "transportpb.Empty")
	proto.RegisterType((*ClientPacket)(nil), "transportpb.ClientPacket")
}

func init() {
	proto.RegisterFile("transport/transportpb/transport.proto", fileDescriptor_4b7c1f8429b4b9ad)
}

var fileDescriptor_4b7c1f8429b4b9ad = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x6e, 0x75, 0x73, 0x9a, 0x4d, 0x85, 0xa0, 0x52, 0x8b, 0xd4, 0x51, 0x10, 0xe6, 0xc1, 0x76,
	0x4c, 0x04, 0x6f, 0xe2, 0x86, 0x82, 0xb0, 0xc2, 0xdc, 0x76, 0xf2, 0x96, 0x76, 0x8f, 0x2c, 0xac,
	0x6d, 0x42, 0x9a, 0x29, 0xfb, 0x17, 0xfe, 0xac, 0x1d, 0x77, 0xf4, 0x24, 0xb2, 0xfd, 0x11, 0x69,
	0xeb, 0xdc, 0x06, 0x7a, 0x49, 0xbe, 0x7c, 0xdf, 0xf7, 0xc8, 0xf7, 0xde, 0x43, 0x17, 0x4a, 0x92,
	0x38, 0x11, 0x5c, 0x2a, 0xf7, 0x17, 0x09, 0x7f, 0x85, 0x1d, 0x21, 0xb9, 0xe2, 0xb8, 0xbc, 0x26,
	0x9a, 0x47, 0x94, 0x53, 0x9e, 0xf1, 0x6e, 0x8a, 0x72, 0x8b, 0x79, 0x22, 0x00, 0xa4, 0x9b, 0x1e,
	0xc2, 0xcf, 0xae, 0x9c, 0xb7, 0x6f, 0xd0, 0x6e, 0x3f, 0x2d, 0xf6, 0x12, 0x8a, 0x2f, 0x51, 0x21,
	0x4a, 0x68, 0x62, 0xe8, 0xd5, 0xed, 0x5a, 0xb9, 0x71, 0xe8, 0xe4, 0x6e, 0xc7, 0x83, 0x24, 0x21,
	0x14, 0x9a, 0x85, 0xe9, 0xe7, 0xb9, 0xd6, 0xcd, 0x2c, 0x76, 0x09, 0x15, 0x1f, 0x22, 0xa1, 0x26,
	0xf6, 0x23, 0xaa, 0xb4, 0x42, 0x06, 0xb1, 0xea, 0x90, 0x60, 0x04, 0x0a, 0x1b, 0xa8, 0x14, 0xe5,
	0x7e, 0x43, 0xaf, 0xea, 0xb5, 0x4a, 0x77, 0xf9, 0xc4, 0x67, 0x68, 0x2f, 0x61, 0x34, 0x26, 0x6a,
	0x2c, 0xc1, 0xd8, 0xca, 0xb4, 0x15, 0xd1, 0x78, 0x42, 0xfb, 0x1d, 0x00, 0xd9, 0x5f, 0x36, 0x82,
	0x6f, 0x51, 0xe9, 0xe7, 0x63, 0x7c, 0xec, 0xac, 0xf5, 0xe7, 0x2c, 0xe3, 0x9a, 0x78, 0x83, 0xce,
	0xe3, 0x68, 0x35, 0xbd, 0xf1, 0x8c, 0x0e, 0x5a, 0x3c, 0x8a, 0x48, 0x3c, 0xe8, 0x81, 0x7c, 0x65,
	0x01, 0xe0, 0x3b, 0x54, 0xbc, 0x17, 0x22, 0x9c, 0xe0, 0xd3, 0x8d, 0x92, 0xf5, 0xe0, 0xe6, 0xff,
	0x92, 0xad, 0x35, 0xe9, 0x74, 0x6e, 0xe9, 0xb3, 0xb9, 0xa5, 0x7f, 0xcd, 0x2d, 0xfd, 0x7d, 0x61,
	0x69, 0xb3, 0x85, 0xa5, 0x7d, 0x2c, 0x2c, 0xed, 0xc5, 0xa3, 0x4c, 0x0d, 0xc7, 0xbe, 0x13, 0xf0,
	0xc8, 0x8d, 0x99, 0x1a, 0xb2, 0x78, 0x44, 0x54, 0x48, 0xea, 0x75, 0xb7, 0xdb, 0xbe, 0xf2, 0xc6,
	0xa1, 0x62, 0x6d, 0x20, 0x03, 0x90, 0x9d, 0x74, 0xda, 0x01, 0x0f, 0x7b, 0x6f, 0x4c, 0x05, 0x43,
	0x90, 0xee, 0x9f, 0x9b, 0xf5, 0x77, 0xb2, 0xad, 0x5c, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x29,
	0x1a, 0xce, 0x76, 0xf9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeerTransportClient is the client API for PeerTransport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeerTransportClient interface {
	Message(ctx context.Context, opts ...grpc.CallOption) (PeerTransport_MessageClient, error)
}

type peerTransportClient struct {
	cc *grpc.ClientConn
}

func NewPeerTransportClient(cc *grpc.ClientConn) PeerTransportClient {
	return &peerTransportClient{cc}
}

func (c *peerTransportClient) Message(ctx context.Context, opts ...grpc.CallOption) (PeerTransport_MessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PeerTransport_serviceDesc.Streams[0], "/transportpb.PeerTransport/Message", opts...)
	if err != nil {
		return nil, err
	}
	x := &peerTransportMessageClient{stream}
	return x, nil
}

type PeerTransport_MessageClient interface {
	Send(*TransMsg) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type peerTransportMessageClient struct {
	grpc.ClientStream
}

func (x *peerTransportMessageClient) Send(m *TransMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *peerTransportMessageClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PeerTransportServer is the server API for PeerTransport service.
type PeerTransportServer interface {
	Message(PeerTransport_MessageServer) error
}

// UnimplementedPeerTransportServer can be embedded to have forward compatible implementations.
type UnimplementedPeerTransportServer struct {
}

func (*UnimplementedPeerTransportServer) Message(srv PeerTransport_MessageServer) error {
	return status.Errorf(codes.Unimplemented, "method Message not implemented")
}

func RegisterPeerTransportServer(s *grpc.Server, srv PeerTransportServer) {
	s.RegisterService(&_PeerTransport_serviceDesc, srv)
}

func _PeerTransport_Message_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PeerTransportServer).Message(&peerTransportMessageServer{stream})
}

type PeerTransport_MessageServer interface {
	SendAndClose(*Empty) error
	Recv() (*TransMsg, error)
	grpc.ServerStream
}

type peerTransportMessageServer struct {
	grpc.ServerStream
}

func (x *peerTransportMessageServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *peerTransportMessageServer) Recv() (*TransMsg, error) {
	m := new(TransMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PeerTransport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transportpb.PeerTransport",
	HandlerType: (*PeerTransportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Message",
			Handler:       _PeerTransport_Message_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "transport/transportpb/transport.proto",
}

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandServiceClient interface {
	Apply(ctx context.Context, in *ClientPacket, opts ...grpc.CallOption) (*ClientPacket, error)
}

type commandServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommandServiceClient(cc *grpc.ClientConn) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) Apply(ctx context.Context, in *ClientPacket, opts ...grpc.CallOption) (*ClientPacket, error) {
	out := new(ClientPacket)
	err := c.cc.Invoke(ctx, "/transportpb.CommandService/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServiceServer is the server API for CommandService service.
type CommandServiceServer interface {
	Apply(context.Context, *ClientPacket) (*ClientPacket, error)
}

// UnimplementedCommandServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (*UnimplementedCommandServiceServer) Apply(ctx context.Context, req *ClientPacket) (*ClientPacket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}

func RegisterCommandServiceServer(s *grpc.Server, srv CommandServiceServer) {
	s.RegisterService(&_CommandService_serviceDesc, srv)
}

func _CommandService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transportpb.CommandService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceServer).Apply(ctx, req.(*ClientPacket))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommandService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transportpb.CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _CommandService_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/transportpb/transport.proto",
}

func (m *TransMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTransport(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Empty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ClientPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintTransport(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTransport(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTransport(dAtA []byte, offset int, v uint64) int {
	offset -= sovTransport(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovTransport(uint64(l))
		}
	}
	return n
}

func (m *Empty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ClientPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTransport(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovTransport(uint64(l))
	}
	return n
}

func sovTransport(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTransport(x uint64) (n int) {
	return sovTransport(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TransMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTransport
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTransport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, peerpb.Message{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTransport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClientPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTransport
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransport
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = append(m.Message[:0], dAtA[iNdEx:postIndex]...)
			if m.Message == nil {
				m.Message = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTransport
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTransport
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTransport
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTransport(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTransport
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTransport(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTransport
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTransport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTransport
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTransport
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTransport
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTransport
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTransport        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTransport          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTransport = fmt.Errorf("proto: unexpected end of group")
)