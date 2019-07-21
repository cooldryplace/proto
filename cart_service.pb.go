// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cart_service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// LineItem represents an SKU with quantity.
type LineItem struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity             int64    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LineItem) Reset()         { *m = LineItem{} }
func (m *LineItem) String() string { return proto.CompactTextString(m) }
func (*LineItem) ProtoMessage()    {}
func (*LineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{0}
}

func (m *LineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineItem.Unmarshal(m, b)
}
func (m *LineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineItem.Marshal(b, m, deterministic)
}
func (m *LineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineItem.Merge(m, src)
}
func (m *LineItem) XXX_Size() int {
	return xxx_messageInfo_LineItem.Size(m)
}
func (m *LineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LineItem.DiscardUnknown(m)
}

var xxx_messageInfo_LineItem proto.InternalMessageInfo

func (m *LineItem) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *LineItem) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

// Cart holds selected LineItems.
type Cart struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64                `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Items                []*LineItem          `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Cart) Reset()         { *m = Cart{} }
func (m *Cart) String() string { return proto.CompactTextString(m) }
func (*Cart) ProtoMessage()    {}
func (*Cart) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{1}
}

func (m *Cart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cart.Unmarshal(m, b)
}
func (m *Cart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cart.Marshal(b, m, deterministic)
}
func (m *Cart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cart.Merge(m, src)
}
func (m *Cart) XXX_Size() int {
	return xxx_messageInfo_Cart.Size(m)
}
func (m *Cart) XXX_DiscardUnknown() {
	xxx_messageInfo_Cart.DiscardUnknown(m)
}

var xxx_messageInfo_Cart proto.InternalMessageInfo

func (m *Cart) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cart) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Cart) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Cart) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Cart) GetItems() []*LineItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// CartCreateRequest is used to create new Cart for User ID.
type CartCreateRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartCreateRequest) Reset()         { *m = CartCreateRequest{} }
func (m *CartCreateRequest) String() string { return proto.CompactTextString(m) }
func (*CartCreateRequest) ProtoMessage()    {}
func (*CartCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{2}
}

func (m *CartCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartCreateRequest.Unmarshal(m, b)
}
func (m *CartCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartCreateRequest.Marshal(b, m, deterministic)
}
func (m *CartCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartCreateRequest.Merge(m, src)
}
func (m *CartCreateRequest) XXX_Size() int {
	return xxx_messageInfo_CartCreateRequest.Size(m)
}
func (m *CartCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CartCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CartCreateRequest proto.InternalMessageInfo

func (m *CartCreateRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// CartRequest is used to fetch data about Cart state.
type CartRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartRequest) Reset()         { *m = CartRequest{} }
func (m *CartRequest) String() string { return proto.CompactTextString(m) }
func (*CartRequest) ProtoMessage()    {}
func (*CartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{3}
}

func (m *CartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartRequest.Unmarshal(m, b)
}
func (m *CartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartRequest.Marshal(b, m, deterministic)
}
func (m *CartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartRequest.Merge(m, src)
}
func (m *CartRequest) XXX_Size() int {
	return xxx_messageInfo_CartRequest.Size(m)
}
func (m *CartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CartRequest proto.InternalMessageInfo

func (m *CartRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// CartResponse is used to return Cart details.
type CartResponse struct {
	Cart                 *Cart    `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartResponse) Reset()         { *m = CartResponse{} }
func (m *CartResponse) String() string { return proto.CompactTextString(m) }
func (*CartResponse) ProtoMessage()    {}
func (*CartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{4}
}

func (m *CartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartResponse.Unmarshal(m, b)
}
func (m *CartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartResponse.Marshal(b, m, deterministic)
}
func (m *CartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartResponse.Merge(m, src)
}
func (m *CartResponse) XXX_Size() int {
	return xxx_messageInfo_CartResponse.Size(m)
}
func (m *CartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CartResponse proto.InternalMessageInfo

func (m *CartResponse) GetCart() *Cart {
	if m != nil {
		return m.Cart
	}
	return nil
}

// CartDeleteRequest is used to identify Cart that should be removed.
type CartDeleteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CartDeleteRequest) Reset()         { *m = CartDeleteRequest{} }
func (m *CartDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*CartDeleteRequest) ProtoMessage()    {}
func (*CartDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{5}
}

func (m *CartDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CartDeleteRequest.Unmarshal(m, b)
}
func (m *CartDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CartDeleteRequest.Marshal(b, m, deterministic)
}
func (m *CartDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CartDeleteRequest.Merge(m, src)
}
func (m *CartDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_CartDeleteRequest.Size(m)
}
func (m *CartDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CartDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CartDeleteRequest proto.InternalMessageInfo

func (m *CartDeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// AddProductRequest provides a way to specify what SKU should be added to a Cart.
type AddProductRequest struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity             int64    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CartId               int64    `protobuf:"varint,3,opt,name=cartId,proto3" json:"cartId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProductRequest) Reset()         { *m = AddProductRequest{} }
func (m *AddProductRequest) String() string { return proto.CompactTextString(m) }
func (*AddProductRequest) ProtoMessage()    {}
func (*AddProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{6}
}

func (m *AddProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProductRequest.Unmarshal(m, b)
}
func (m *AddProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProductRequest.Marshal(b, m, deterministic)
}
func (m *AddProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProductRequest.Merge(m, src)
}
func (m *AddProductRequest) XXX_Size() int {
	return xxx_messageInfo_AddProductRequest.Size(m)
}
func (m *AddProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddProductRequest proto.InternalMessageInfo

func (m *AddProductRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *AddProductRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *AddProductRequest) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

// DelProductRequest provides data to identify Cart that needs to be deleted.
type DelProductRequest struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=productId,proto3" json:"productId,omitempty"`
	CartId               int64    `protobuf:"varint,2,opt,name=cartId,proto3" json:"cartId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelProductRequest) Reset()         { *m = DelProductRequest{} }
func (m *DelProductRequest) String() string { return proto.CompactTextString(m) }
func (*DelProductRequest) ProtoMessage()    {}
func (*DelProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{7}
}

func (m *DelProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelProductRequest.Unmarshal(m, b)
}
func (m *DelProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelProductRequest.Marshal(b, m, deterministic)
}
func (m *DelProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelProductRequest.Merge(m, src)
}
func (m *DelProductRequest) XXX_Size() int {
	return xxx_messageInfo_DelProductRequest.Size(m)
}
func (m *DelProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DelProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DelProductRequest proto.InternalMessageInfo

func (m *DelProductRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *DelProductRequest) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

// EmptyCartRequest is used to remove all LineItems from a Cart.
type EmptyCartRequest struct {
	CartId               int64    `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyCartRequest) Reset()         { *m = EmptyCartRequest{} }
func (m *EmptyCartRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyCartRequest) ProtoMessage()    {}
func (*EmptyCartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9a99120c5507bc1, []int{8}
}

func (m *EmptyCartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyCartRequest.Unmarshal(m, b)
}
func (m *EmptyCartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyCartRequest.Marshal(b, m, deterministic)
}
func (m *EmptyCartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyCartRequest.Merge(m, src)
}
func (m *EmptyCartRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyCartRequest.Size(m)
}
func (m *EmptyCartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyCartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyCartRequest proto.InternalMessageInfo

func (m *EmptyCartRequest) GetCartId() int64 {
	if m != nil {
		return m.CartId
	}
	return 0
}

func init() {
	proto.RegisterType((*LineItem)(nil), "cooldryplace.protobuf.LineItem")
	proto.RegisterType((*Cart)(nil), "cooldryplace.protobuf.Cart")
	proto.RegisterType((*CartCreateRequest)(nil), "cooldryplace.protobuf.CartCreateRequest")
	proto.RegisterType((*CartRequest)(nil), "cooldryplace.protobuf.CartRequest")
	proto.RegisterType((*CartResponse)(nil), "cooldryplace.protobuf.CartResponse")
	proto.RegisterType((*CartDeleteRequest)(nil), "cooldryplace.protobuf.CartDeleteRequest")
	proto.RegisterType((*AddProductRequest)(nil), "cooldryplace.protobuf.AddProductRequest")
	proto.RegisterType((*DelProductRequest)(nil), "cooldryplace.protobuf.DelProductRequest")
	proto.RegisterType((*EmptyCartRequest)(nil), "cooldryplace.protobuf.EmptyCartRequest")
}

func init() { proto.RegisterFile("cart_service.proto", fileDescriptor_c9a99120c5507bc1) }

var fileDescriptor_c9a99120c5507bc1 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0xf4, 0x63, 0xdd, 0xde, 0x8a, 0xd8, 0x01, 0x97, 0x92, 0x75, 0x69, 0x99, 0x7d, 0xb0,
	0x28, 0xa4, 0x50, 0x11, 0x7c, 0x93, 0x75, 0x2b, 0x12, 0x10, 0x59, 0xc2, 0x3e, 0x09, 0x22, 0xd9,
	0xcc, 0xb5, 0x06, 0x92, 0x4e, 0x36, 0xb9, 0x11, 0xfa, 0xea, 0xbf, 0xf4, 0xdf, 0xc8, 0xcc, 0x24,
	0x4d, 0xda, 0xda, 0x86, 0xee, 0x53, 0x98, 0xb9, 0xe7, 0x9c, 0xb9, 0xf7, 0xdc, 0x43, 0x80, 0x05,
	0x7e, 0x4a, 0x3f, 0x32, 0x4c, 0x7f, 0x87, 0x01, 0x3a, 0x49, 0x2a, 0x49, 0xb2, 0x17, 0x81, 0x94,
	0x91, 0x48, 0x57, 0x49, 0xe4, 0x97, 0x77, 0xf7, 0xf9, 0x4f, 0xfb, 0x62, 0x21, 0xe5, 0x22, 0xc2,
	0x69, 0x79, 0x31, 0xc5, 0x38, 0xa1, 0x95, 0xa9, 0xdb, 0xa3, 0xed, 0x22, 0x85, 0x31, 0x66, 0xe4,
	0xc7, 0x89, 0x01, 0xf0, 0x39, 0x9c, 0x7d, 0x09, 0x97, 0xe8, 0x12, 0xc6, 0xec, 0x25, 0xf4, 0x92,
	0x54, 0x8a, 0x3c, 0x20, 0x57, 0x0c, 0xad, 0xb1, 0x35, 0x69, 0x7b, 0xd5, 0x05, 0xb3, 0xe1, 0xec,
	0x21, 0xf7, 0x97, 0x14, 0xd2, 0x6a, 0xd8, 0xd2, 0xc5, 0xf5, 0x99, 0xff, 0xb5, 0xa0, 0x73, 0xe3,
	0xa7, 0xc4, 0x9e, 0x41, 0x2b, 0x2c, 0xb9, 0xad, 0x50, 0xb0, 0x73, 0x38, 0xcd, 0x33, 0x4c, 0x5d,
	0x51, 0x50, 0x8a, 0x13, 0x7b, 0x0f, 0xbd, 0x20, 0x45, 0x9f, 0x50, 0x5c, 0xd3, 0xb0, 0x3d, 0xb6,
	0x26, 0xfd, 0x99, 0xed, 0x98, 0x5e, 0xd7, 0x93, 0x39, 0x77, 0x65, 0xaf, 0x5e, 0x05, 0x56, 0xcc,
	0x3c, 0x11, 0x05, 0xb3, 0xd3, 0xcc, 0x5c, 0x83, 0xd9, 0x3b, 0xe8, 0x86, 0x84, 0x71, 0x36, 0xec,
	0x8e, 0xdb, 0x93, 0xfe, 0x6c, 0xe4, 0xfc, 0xd7, 0x4f, 0xa7, 0xb4, 0xc3, 0x33, 0x68, 0xfe, 0x06,
	0x06, 0x6a, 0xb4, 0x1b, 0xdd, 0x81, 0x87, 0x0f, 0x39, 0x66, 0x54, 0x9b, 0xcb, 0xaa, 0xcf, 0xc5,
	0x2f, 0xa1, 0xaf, 0xc0, 0x25, 0x6c, 0xcb, 0x0e, 0xfe, 0x01, 0x9e, 0x9a, 0x72, 0x96, 0xc8, 0x65,
	0x86, 0x6c, 0x0a, 0x1d, 0xb5, 0x68, 0x8d, 0xe8, 0xcf, 0x2e, 0xf6, 0x74, 0xa4, 0x29, 0x1a, 0xc8,
	0xaf, 0x4c, 0x33, 0x73, 0x8c, 0xb0, 0x6a, 0x66, 0xfb, 0x15, 0x84, 0xc1, 0xb5, 0x10, 0xb7, 0x66,
	0x73, 0x25, 0xe8, 0xd1, 0xcb, 0x55, 0xb3, 0xaa, 0xb7, 0x5d, 0xa1, 0x17, 0xd5, 0xf6, 0x8a, 0x13,
	0x77, 0x61, 0x30, 0xc7, 0xe8, 0xa8, 0x67, 0x2a, 0xa9, 0xd6, 0x86, 0xd4, 0x6b, 0x78, 0xfe, 0x49,
	0xa5, 0xb6, 0xee, 0x5d, 0x85, 0xb5, 0xea, 0xd8, 0xd9, 0x9f, 0x0e, 0x74, 0x15, 0x2e, 0x63, 0xdf,
	0x01, 0xcc, 0x56, 0x74, 0xf4, 0x26, 0x07, 0xdc, 0xdb, 0x58, 0x9e, 0x7d, 0x75, 0xc8, 0xe7, 0x62,
	0x35, 0xfc, 0x84, 0xdd, 0xc1, 0x93, 0xcf, 0x48, 0x5a, 0x9b, 0x1f, 0x64, 0x1c, 0xa5, 0xfa, 0x15,
	0x7a, 0xeb, 0x51, 0xd9, 0xab, 0x3d, 0x9c, 0x6d, 0x33, 0xec, 0xf3, 0x9d, 0x88, 0x6b, 0x08, 0x3f,
	0x61, 0xb7, 0x00, 0x26, 0x0d, 0x8d, 0x26, 0x6c, 0x84, 0xe6, 0xb0, 0x62, 0x15, 0x9f, 0xbd, 0x8a,
	0x3b, 0x09, 0x6b, 0xec, 0xb1, 0x49, 0x71, 0x27, 0x4c, 0xfb, 0x15, 0x3f, 0x8e, 0xbe, 0x5d, 0x2e,
	0x42, 0xfa, 0x95, 0xdf, 0x3b, 0x81, 0x8c, 0xa7, 0x75, 0xbd, 0xe2, 0x57, 0x77, 0xaa, 0x3f, 0x6f,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x4c, 0x88, 0xbf, 0x49, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CartsClient is the client API for Carts service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CartsClient interface {
	// CreateCart will create new Cart for User ID.
	CreateCart(ctx context.Context, in *CartCreateRequest, opts ...grpc.CallOption) (*CartResponse, error)
	// GetCart returns current state of a Cart.
	GetCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponse, error)
	// EmptyCart will remove all LineItems from the Cart.
	EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// DeleteCart with provided Cart ID.
	DeleteCart(ctx context.Context, in *CartDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// AddProduct to existing Cart.
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// DelProduct from existing Cart.
	DelProduct(ctx context.Context, in *DelProductRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type cartsClient struct {
	cc *grpc.ClientConn
}

func NewCartsClient(cc *grpc.ClientConn) CartsClient {
	return &cartsClient{cc}
}

func (c *cartsClient) CreateCart(ctx context.Context, in *CartCreateRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, "/cooldryplace.protobuf.Carts/CreateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) GetCart(ctx context.Context, in *CartRequest, opts ...grpc.CallOption) (*CartResponse, error) {
	out := new(CartResponse)
	err := c.cc.Invoke(ctx, "/cooldryplace.protobuf.Carts/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) EmptyCart(ctx context.Context, in *EmptyCartRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cooldryplace.protobuf.Carts/EmptyCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) DeleteCart(ctx context.Context, in *CartDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cooldryplace.protobuf.Carts/DeleteCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cooldryplace.protobuf.Carts/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartsClient) DelProduct(ctx context.Context, in *DelProductRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cooldryplace.protobuf.Carts/DelProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartsServer is the server API for Carts service.
type CartsServer interface {
	// CreateCart will create new Cart for User ID.
	CreateCart(context.Context, *CartCreateRequest) (*CartResponse, error)
	// GetCart returns current state of a Cart.
	GetCart(context.Context, *CartRequest) (*CartResponse, error)
	// EmptyCart will remove all LineItems from the Cart.
	EmptyCart(context.Context, *EmptyCartRequest) (*empty.Empty, error)
	// DeleteCart with provided Cart ID.
	DeleteCart(context.Context, *CartDeleteRequest) (*empty.Empty, error)
	// AddProduct to existing Cart.
	AddProduct(context.Context, *AddProductRequest) (*empty.Empty, error)
	// DelProduct from existing Cart.
	DelProduct(context.Context, *DelProductRequest) (*empty.Empty, error)
}

// UnimplementedCartsServer can be embedded to have forward compatible implementations.
type UnimplementedCartsServer struct {
}

func (*UnimplementedCartsServer) CreateCart(ctx context.Context, req *CartCreateRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (*UnimplementedCartsServer) GetCart(ctx context.Context, req *CartRequest) (*CartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (*UnimplementedCartsServer) EmptyCart(ctx context.Context, req *EmptyCartRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyCart not implemented")
}
func (*UnimplementedCartsServer) DeleteCart(ctx context.Context, req *CartDeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (*UnimplementedCartsServer) AddProduct(ctx context.Context, req *AddProductRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (*UnimplementedCartsServer) DelProduct(ctx context.Context, req *DelProductRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelProduct not implemented")
}

func RegisterCartsServer(s *grpc.Server, srv CartsServer) {
	s.RegisterService(&_Carts_serviceDesc, srv)
}

func _Carts_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cooldryplace.protobuf.Carts/CreateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).CreateCart(ctx, req.(*CartCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cooldryplace.protobuf.Carts/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).GetCart(ctx, req.(*CartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_EmptyCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).EmptyCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cooldryplace.protobuf.Carts/EmptyCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).EmptyCart(ctx, req.(*EmptyCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cooldryplace.protobuf.Carts/DeleteCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).DeleteCart(ctx, req.(*CartDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cooldryplace.protobuf.Carts/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carts_DelProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartsServer).DelProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cooldryplace.protobuf.Carts/DelProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartsServer).DelProduct(ctx, req.(*DelProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Carts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cooldryplace.protobuf.Carts",
	HandlerType: (*CartsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCart",
			Handler:    _Carts_CreateCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _Carts_GetCart_Handler,
		},
		{
			MethodName: "EmptyCart",
			Handler:    _Carts_EmptyCart_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _Carts_DeleteCart_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _Carts_AddProduct_Handler,
		},
		{
			MethodName: "DelProduct",
			Handler:    _Carts_DelProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart_service.proto",
}
