// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/tx.proto

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

type MsgCreateOrder struct {
	Creator   string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Market    string   `protobuf:"bytes,2,opt,name=market,proto3" json:"market,omitempty"`
	Side      bool     `protobuf:"varint,3,opt,name=side,proto3" json:"side,omitempty"`
	OrderType string   `protobuf:"bytes,4,opt,name=orderType,proto3" json:"orderType,omitempty"`
	Price     string   `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	Quantity  string   `protobuf:"bytes,6,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Flags     []string `protobuf:"bytes,7,rep,name=flags,proto3" json:"flags,omitempty"`
}

func (m *MsgCreateOrder) Reset()         { *m = MsgCreateOrder{} }
func (m *MsgCreateOrder) String() string { return proto.CompactTextString(m) }
func (*MsgCreateOrder) ProtoMessage()    {}
func (*MsgCreateOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_463701e671e5a5e0, []int{0}
}
func (m *MsgCreateOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateOrder.Merge(m, src)
}
func (m *MsgCreateOrder) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateOrder proto.InternalMessageInfo

func (m *MsgCreateOrder) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateOrder) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *MsgCreateOrder) GetSide() bool {
	if m != nil {
		return m.Side
	}
	return false
}

func (m *MsgCreateOrder) GetOrderType() string {
	if m != nil {
		return m.OrderType
	}
	return ""
}

func (m *MsgCreateOrder) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *MsgCreateOrder) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *MsgCreateOrder) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

type MsgCreateOrderResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateOrderResponse) Reset()         { *m = MsgCreateOrderResponse{} }
func (m *MsgCreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateOrderResponse) ProtoMessage()    {}
func (*MsgCreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_463701e671e5a5e0, []int{1}
}
func (m *MsgCreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateOrderResponse.Merge(m, src)
}
func (m *MsgCreateOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateOrderResponse proto.InternalMessageInfo

func (m *MsgCreateOrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MsgCancelOrder struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Market  string `protobuf:"bytes,2,opt,name=market,proto3" json:"market,omitempty"`
	Side    bool   `protobuf:"varint,3,opt,name=side,proto3" json:"side,omitempty"`
	Price   string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Id      string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCancelOrder) Reset()         { *m = MsgCancelOrder{} }
func (m *MsgCancelOrder) String() string { return proto.CompactTextString(m) }
func (*MsgCancelOrder) ProtoMessage()    {}
func (*MsgCancelOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_463701e671e5a5e0, []int{2}
}
func (m *MsgCancelOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelOrder.Merge(m, src)
}
func (m *MsgCancelOrder) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelOrder proto.InternalMessageInfo

func (m *MsgCancelOrder) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCancelOrder) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *MsgCancelOrder) GetSide() bool {
	if m != nil {
		return m.Side
	}
	return false
}

func (m *MsgCancelOrder) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *MsgCancelOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MsgCancelOrderResponse struct {
}

func (m *MsgCancelOrderResponse) Reset()         { *m = MsgCancelOrderResponse{} }
func (m *MsgCancelOrderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCancelOrderResponse) ProtoMessage()    {}
func (*MsgCancelOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_463701e671e5a5e0, []int{3}
}
func (m *MsgCancelOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelOrderResponse.Merge(m, src)
}
func (m *MsgCancelOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelOrderResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateOrder)(nil), "soupyfinance.noodle.dex.MsgCreateOrder")
	proto.RegisterType((*MsgCreateOrderResponse)(nil), "soupyfinance.noodle.dex.MsgCreateOrderResponse")
	proto.RegisterType((*MsgCancelOrder)(nil), "soupyfinance.noodle.dex.MsgCancelOrder")
	proto.RegisterType((*MsgCancelOrderResponse)(nil), "soupyfinance.noodle.dex.MsgCancelOrderResponse")
}

func init() { proto.RegisterFile("dex/tx.proto", fileDescriptor_463701e671e5a5e0) }

var fileDescriptor_463701e671e5a5e0 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x3b, 0xfd, 0xee, 0xfd, 0xff, 0xe9, 0x62, 0x90, 0x3a, 0x14, 0x09, 0xa5, 0x1b, 0x0b,
	0x62, 0x02, 0xfa, 0x06, 0x8a, 0xcb, 0x22, 0x14, 0x57, 0xee, 0xd2, 0xcc, 0x6d, 0x1c, 0x6c, 0x33,
	0x71, 0x66, 0x0a, 0x09, 0xf8, 0x10, 0x3e, 0x90, 0x0f, 0xe0, 0xb2, 0x3b, 0x5d, 0x4a, 0xfb, 0x22,
	0x92, 0x49, 0x9b, 0xb6, 0x88, 0x1f, 0x0b, 0x77, 0x39, 0xb9, 0x77, 0xce, 0xb9, 0xbf, 0x99, 0x0b,
	0xff, 0x39, 0x26, 0x9e, 0x49, 0xdc, 0x58, 0x49, 0x23, 0xe9, 0xa1, 0x96, 0xf3, 0x38, 0x9d, 0x88,
	0xc8, 0x8f, 0x02, 0x74, 0x23, 0x29, 0xf9, 0x14, 0x5d, 0x8e, 0x49, 0xff, 0x99, 0x40, 0x7b, 0xa8,
	0xc3, 0x4b, 0x85, 0xbe, 0xc1, 0x6b, 0xc5, 0x51, 0x51, 0x06, 0x8d, 0x20, 0x93, 0x52, 0x31, 0xd2,
	0x23, 0x83, 0xd6, 0x68, 0x23, 0x69, 0x07, 0xea, 0x33, 0x5f, 0xdd, 0xa3, 0x61, 0x65, 0x5b, 0x58,
	0x2b, 0x4a, 0xa1, 0xaa, 0x05, 0x47, 0x56, 0xe9, 0x91, 0x41, 0x73, 0x64, 0xbf, 0xe9, 0x11, 0xb4,
	0x64, 0x66, 0x77, 0x93, 0xc6, 0xc8, 0xaa, 0xb6, 0x7d, 0xfb, 0x83, 0x1e, 0x40, 0x2d, 0x56, 0x22,
	0x40, 0x56, 0xb3, 0x95, 0x5c, 0xd0, 0x2e, 0x34, 0x1f, 0xe6, 0x7e, 0x64, 0x84, 0x49, 0x59, 0xdd,
	0x16, 0x0a, 0x9d, 0x9d, 0x98, 0x4c, 0xfd, 0x50, 0xb3, 0x46, 0xaf, 0x92, 0x9d, 0xb0, 0xa2, 0x3f,
	0x80, 0xce, 0xfe, 0xf4, 0x23, 0xd4, 0xb1, 0x8c, 0x34, 0xd2, 0x36, 0x94, 0x05, 0x5f, 0x03, 0x94,
	0x05, 0xef, 0x3f, 0xe6, 0x9c, 0x19, 0xff, 0xf4, 0x2f, 0x39, 0x0b, 0x92, 0xea, 0x2e, 0x49, 0x9e,
	0x5e, 0x2b, 0xd2, 0x59, 0x3e, 0xe7, 0x36, 0x7d, 0x33, 0xe7, 0xd9, 0x2b, 0x81, 0xca, 0x50, 0x87,
	0x34, 0x84, 0x7f, 0xbb, 0x8f, 0x70, 0xec, 0x7e, 0xf1, 0x62, 0xee, 0x3e, 0x6f, 0xd7, 0xfb, 0x65,
	0x63, 0x71, 0x31, 0x59, 0xd0, 0xce, 0x2d, 0x7c, 0x1f, 0xb4, 0x6d, 0xfc, 0x21, 0xe8, 0x33, 0xd9,
	0xc5, 0xd5, 0xcb, 0xd2, 0x21, 0x8b, 0xa5, 0x43, 0xde, 0x97, 0x0e, 0x79, 0x5a, 0x39, 0xa5, 0xc5,
	0xca, 0x29, 0xbd, 0xad, 0x9c, 0xd2, 0xed, 0x49, 0x28, 0xcc, 0xdd, 0x7c, 0xec, 0x06, 0x72, 0xe6,
	0x59, 0xd3, 0xd3, 0xb5, 0xab, 0x97, 0xbb, 0x7a, 0x89, 0x67, 0xb7, 0x37, 0x8d, 0x51, 0x8f, 0xeb,
	0x76, 0x83, 0xcf, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xde, 0xff, 0x20, 0xf7, 0xd1, 0x02, 0x00,
	0x00,
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
	CreateOrder(ctx context.Context, in *MsgCreateOrder, opts ...grpc.CallOption) (*MsgCreateOrderResponse, error)
	CancelOrder(ctx context.Context, in *MsgCancelOrder, opts ...grpc.CallOption) (*MsgCancelOrderResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateOrder(ctx context.Context, in *MsgCreateOrder, opts ...grpc.CallOption) (*MsgCreateOrderResponse, error) {
	out := new(MsgCreateOrderResponse)
	err := c.cc.Invoke(ctx, "/soupyfinance.noodle.dex.Msg/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelOrder(ctx context.Context, in *MsgCancelOrder, opts ...grpc.CallOption) (*MsgCancelOrderResponse, error) {
	out := new(MsgCancelOrderResponse)
	err := c.cc.Invoke(ctx, "/soupyfinance.noodle.dex.Msg/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateOrder(context.Context, *MsgCreateOrder) (*MsgCreateOrderResponse, error)
	CancelOrder(context.Context, *MsgCancelOrder) (*MsgCancelOrderResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateOrder(ctx context.Context, req *MsgCreateOrder) (*MsgCreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (*UnimplementedMsgServer) CancelOrder(ctx context.Context, req *MsgCancelOrder) (*MsgCancelOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soupyfinance.noodle.dex.Msg/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateOrder(ctx, req.(*MsgCreateOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/soupyfinance.noodle.dex.Msg/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelOrder(ctx, req.(*MsgCancelOrder))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "soupyfinance.noodle.dex.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Msg_CreateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Msg_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dex/tx.proto",
}

func (m *MsgCreateOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Flags) > 0 {
		for iNdEx := len(m.Flags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Flags[iNdEx])
			copy(dAtA[i:], m.Flags[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Flags[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Quantity) > 0 {
		i -= len(m.Quantity)
		copy(dAtA[i:], m.Quantity)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Quantity)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OrderType) > 0 {
		i -= len(m.OrderType)
		copy(dAtA[i:], m.OrderType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OrderType)))
		i--
		dAtA[i] = 0x22
	}
	if m.Side {
		i--
		if m.Side {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Market) > 0 {
		i -= len(m.Market)
		copy(dAtA[i:], m.Market)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Market)))
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

func (m *MsgCreateOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateOrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCancelOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x22
	}
	if m.Side {
		i--
		if m.Side {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Market) > 0 {
		i -= len(m.Market)
		copy(dAtA[i:], m.Market)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Market)))
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

func (m *MsgCancelOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelOrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Market)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Side {
		n += 2
	}
	l = len(m.OrderType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Quantity)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Flags) > 0 {
		for _, s := range m.Flags {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateOrderResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCancelOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Market)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Side {
		n += 2
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCancelOrderResponse) Size() (n int) {
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
func (m *MsgCreateOrder) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateOrder: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
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
			m.Market = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Side", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Side = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderType", wireType)
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
			m.OrderType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
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
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
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
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
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
			m.Flags = append(m.Flags, string(dAtA[iNdEx:postIndex]))
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
func (m *MsgCreateOrderResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCancelOrder) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCancelOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelOrder: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Market", wireType)
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
			m.Market = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Side", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Side = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
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
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCancelOrderResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCancelOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
