// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bridge/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgObserveDeposit struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	ChainId   string `protobuf:"bytes,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	Depositor string `protobuf:"bytes,3,opt,name=depositor,proto3" json:"depositor,omitempty"`
	DepositId string `protobuf:"bytes,4,opt,name=depositId,proto3" json:"depositId,omitempty"`
	Quantity  string `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Asset     string `protobuf:"bytes,6,opt,name=asset,proto3" json:"asset,omitempty"`
}

func (m *MsgObserveDeposit) Reset()         { *m = MsgObserveDeposit{} }
func (m *MsgObserveDeposit) String() string { return proto.CompactTextString(m) }
func (*MsgObserveDeposit) ProtoMessage()    {}
func (*MsgObserveDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_41dc2e30e6049cae, []int{0}
}
func (m *MsgObserveDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgObserveDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgObserveDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgObserveDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgObserveDeposit.Merge(m, src)
}
func (m *MsgObserveDeposit) XXX_Size() int {
	return m.Size()
}
func (m *MsgObserveDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgObserveDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_MsgObserveDeposit proto.InternalMessageInfo

func (m *MsgObserveDeposit) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgObserveDeposit) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *MsgObserveDeposit) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *MsgObserveDeposit) GetDepositId() string {
	if m != nil {
		return m.DepositId
	}
	return ""
}

func (m *MsgObserveDeposit) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *MsgObserveDeposit) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

type MsgObserveDepositResponse struct {
}

func (m *MsgObserveDepositResponse) Reset()         { *m = MsgObserveDepositResponse{} }
func (m *MsgObserveDepositResponse) String() string { return proto.CompactTextString(m) }
func (*MsgObserveDepositResponse) ProtoMessage()    {}
func (*MsgObserveDepositResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41dc2e30e6049cae, []int{1}
}
func (m *MsgObserveDepositResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgObserveDepositResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgObserveDepositResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgObserveDepositResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgObserveDepositResponse.Merge(m, src)
}
func (m *MsgObserveDepositResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgObserveDepositResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgObserveDepositResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgObserveDepositResponse proto.InternalMessageInfo

type MsgWithdraw struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Asset    string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Quantity string `protobuf:"bytes,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ChainId  string `protobuf:"bytes,4,opt,name=chainId,proto3" json:"chainId,omitempty"`
}

func (m *MsgWithdraw) Reset()         { *m = MsgWithdraw{} }
func (m *MsgWithdraw) String() string { return proto.CompactTextString(m) }
func (*MsgWithdraw) ProtoMessage()    {}
func (*MsgWithdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_41dc2e30e6049cae, []int{2}
}
func (m *MsgWithdraw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdraw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdraw.Merge(m, src)
}
func (m *MsgWithdraw) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdraw.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdraw proto.InternalMessageInfo

func (m *MsgWithdraw) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgWithdraw) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

func (m *MsgWithdraw) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *MsgWithdraw) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

type MsgWithdrawResponse struct {
}

func (m *MsgWithdrawResponse) Reset()         { *m = MsgWithdrawResponse{} }
func (m *MsgWithdrawResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawResponse) ProtoMessage()    {}
func (*MsgWithdrawResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_41dc2e30e6049cae, []int{3}
}
func (m *MsgWithdrawResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawResponse.Merge(m, src)
}
func (m *MsgWithdrawResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgObserveDeposit)(nil), "soupyfinance.noodle.bridge.MsgObserveDeposit")
	proto.RegisterType((*MsgObserveDepositResponse)(nil), "soupyfinance.noodle.bridge.MsgObserveDepositResponse")
	proto.RegisterType((*MsgWithdraw)(nil), "soupyfinance.noodle.bridge.MsgWithdraw")
	proto.RegisterType((*MsgWithdrawResponse)(nil), "soupyfinance.noodle.bridge.MsgWithdrawResponse")
}

func init() { proto.RegisterFile("bridge/tx.proto", fileDescriptor_41dc2e30e6049cae) }

var fileDescriptor_41dc2e30e6049cae = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4b, 0x4e, 0xc3, 0x30,
	0x14, 0xac, 0xfb, 0xa3, 0x35, 0x12, 0x08, 0x03, 0x92, 0x09, 0x28, 0x42, 0xd9, 0xc0, 0xa6, 0x8e,
	0x04, 0xe2, 0x02, 0x88, 0x05, 0x5d, 0x54, 0x48, 0xdd, 0x20, 0xb1, 0x4b, 0x6a, 0x93, 0x5a, 0x82,
	0x38, 0xf8, 0xb9, 0xa5, 0xbd, 0x05, 0x97, 0xe1, 0x0e, 0x2c, 0xbb, 0x64, 0x89, 0x9a, 0x8b, 0x20,
	0x25, 0x75, 0x49, 0x5a, 0xf1, 0x5b, 0x4e, 0x66, 0x32, 0x6f, 0xe6, 0xf9, 0xe1, 0xed, 0x50, 0x4b,
	0x1e, 0x09, 0xdf, 0x4c, 0x58, 0xa2, 0x95, 0x51, 0xc4, 0x01, 0x35, 0x4a, 0xa6, 0xf7, 0x32, 0x0e,
	0xe2, 0x81, 0x60, 0xb1, 0x52, 0xfc, 0x41, 0xb0, 0x5c, 0xe4, 0xbd, 0x22, 0xbc, 0xd3, 0x83, 0xe8,
	0x26, 0x04, 0xa1, 0xc7, 0xe2, 0x4a, 0x24, 0x0a, 0xa4, 0x21, 0x14, 0x6f, 0x0c, 0xb4, 0x08, 0x8c,
	0xd2, 0x14, 0x1d, 0xa3, 0xd3, 0x76, 0xdf, 0xc2, 0x8c, 0x19, 0x06, 0x32, 0xee, 0x72, 0x5a, 0x5d,
	0x30, 0x39, 0x24, 0x47, 0xb8, 0xcd, 0xf3, 0xdf, 0x95, 0xa6, 0xb5, 0x8c, 0xfb, 0xfa, 0x50, 0x60,
	0xbb, 0x9c, 0xd6, 0x4b, 0x6c, 0x97, 0x13, 0x07, 0xb7, 0x9e, 0x46, 0x41, 0x6c, 0xa4, 0x99, 0xd2,
	0x46, 0x46, 0x2e, 0x31, 0xd9, 0xc3, 0x8d, 0x00, 0x40, 0x18, 0xda, 0xcc, 0x88, 0x1c, 0x78, 0x87,
	0xf8, 0x60, 0x2d, 0x76, 0x5f, 0x40, 0xa2, 0x62, 0x10, 0x1e, 0xe0, 0xcd, 0x1e, 0x44, 0xb7, 0xd2,
	0x0c, 0xb9, 0x0e, 0x9e, 0x7f, 0x68, 0xb3, 0xf4, 0xae, 0x16, 0xbc, 0x4b, 0x69, 0x6a, 0x2b, 0x69,
	0x0a, 0xfd, 0xeb, 0xa5, 0xfe, 0xde, 0x3e, 0xde, 0x2d, 0x0c, 0xb5, 0x59, 0xce, 0x52, 0x84, 0x6b,
	0x3d, 0x88, 0xc8, 0x18, 0x6f, 0xad, 0x2c, 0xb9, 0xc3, 0xbe, 0x7f, 0x17, 0xb6, 0x56, 0xce, 0xb9,
	0xf8, 0x97, 0xdc, 0xce, 0x27, 0x1c, 0xb7, 0x96, 0x8b, 0x38, 0xf9, 0xc5, 0xc2, 0x0a, 0x1d, 0xff,
	0x8f, 0x42, 0x3b, 0xe5, 0xf2, 0xfa, 0x6d, 0xee, 0xa2, 0xd9, 0xdc, 0x45, 0x1f, 0x73, 0x17, 0xbd,
	0xa4, 0x6e, 0x65, 0x96, 0xba, 0x95, 0xf7, 0xd4, 0xad, 0xdc, 0xb1, 0x48, 0x9a, 0xe1, 0x28, 0x64,
	0x03, 0xf5, 0xe8, 0x67, 0xa6, 0x9d, 0x85, 0xab, 0x9f, 0xbb, 0xfa, 0x13, 0xdf, 0xde, 0xeb, 0x34,
	0x11, 0x10, 0x36, 0xb3, 0x9b, 0x3d, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xf6, 0x62, 0xca,
	0xc6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	ObserveDeposit(ctx context.Context, in *MsgObserveDeposit, opts ...grpc.CallOption) (*MsgObserveDepositResponse, error)
	Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ObserveDeposit(ctx context.Context, in *MsgObserveDeposit, opts ...grpc.CallOption) (*MsgObserveDepositResponse, error) {
	out := new(MsgObserveDepositResponse)
	err := c.cc.Invoke(ctx, "/soupyfinance.noodle.bridge.Msg/ObserveDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Withdraw(ctx context.Context, in *MsgWithdraw, opts ...grpc.CallOption) (*MsgWithdrawResponse, error) {
	out := new(MsgWithdrawResponse)
	err := c.cc.Invoke(ctx, "/soupyfinance.noodle.bridge.Msg/Withdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	ObserveDeposit(context.Context, *MsgObserveDeposit) (*MsgObserveDepositResponse, error)
	Withdraw(context.Context, *MsgWithdraw) (*MsgWithdrawResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) ObserveDeposit(ctx context.Context, req *MsgObserveDeposit) (*MsgObserveDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObserveDeposit not implemented")
}
func (*UnimplementedMsgServer) Withdraw(ctx context.Context, req *MsgWithdraw) (*MsgWithdrawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Withdraw not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_ObserveDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgObserveDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ObserveDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soupyfinance.noodle.bridge.Msg/ObserveDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ObserveDeposit(ctx, req.(*MsgObserveDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Withdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdraw)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Withdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soupyfinance.noodle.bridge.Msg/Withdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Withdraw(ctx, req.(*MsgWithdraw))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "soupyfinance.noodle.bridge.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ObserveDeposit",
			Handler:    _Msg_ObserveDeposit_Handler,
		},
		{
			MethodName: "Withdraw",
			Handler:    _Msg_Withdraw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bridge/tx.proto",
}

func (m *MsgObserveDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgObserveDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgObserveDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Quantity) > 0 {
		i -= len(m.Quantity)
		copy(dAtA[i:], m.Quantity)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Quantity)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DepositId) > 0 {
		i -= len(m.DepositId)
		copy(dAtA[i:], m.DepositId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DepositId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgObserveDepositResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgObserveDepositResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgObserveDepositResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdraw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdraw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdraw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Quantity) > 0 {
		i -= len(m.Quantity)
		copy(dAtA[i:], m.Quantity)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Quantity)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgObserveDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DepositId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Quantity)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgObserveDepositResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdraw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Quantity)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgObserveDeposit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgObserveDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgObserveDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quantity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgObserveDepositResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgObserveDepositResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgObserveDepositResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdraw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdraw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdraw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quantity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgWithdrawResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgWithdrawResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
