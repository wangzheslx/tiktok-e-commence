// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v3.21.0
// source: service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 购物车商品
type CartItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     uint32                 `protobuf:"varint,1,opt,name=product_id,proto3" json:"product_id,omitempty"` // 商品ID
	Quantity      int32                  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`     // 商品数量
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	mi := &file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItem) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

// 新增购物车商品
type UpsertItemReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Item          *CartItem              `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertItemReq) Reset() {
	*x = UpsertItemReq{}
	mi := &file_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertItemReq) ProtoMessage() {}

func (x *UpsertItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertItemReq.ProtoReflect.Descriptor instead.
func (*UpsertItemReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpsertItemReq) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *UpsertItemReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpsertItemReq) GetItem() *CartItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpsertItemResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 操作是否成功
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertItemResp) Reset() {
	*x = UpsertItemResp{}
	mi := &file_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertItemResp) ProtoMessage() {}

func (x *UpsertItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertItemResp.ProtoReflect.Descriptor instead.
func (*UpsertItemResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpsertItemResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 清空购物车
type EmptyCartReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyCartReq) Reset() {
	*x = EmptyCartReq{}
	mi := &file_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCartReq) ProtoMessage() {}

func (x *EmptyCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCartReq.ProtoReflect.Descriptor instead.
func (*EmptyCartReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *EmptyCartReq) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *EmptyCartReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EmptyCartResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 操作是否成功
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmptyCartResp) Reset() {
	*x = EmptyCartResp{}
	mi := &file_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyCartResp) ProtoMessage() {}

func (x *EmptyCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyCartResp.ProtoReflect.Descriptor instead.
func (*EmptyCartResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *EmptyCartResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 获取购物车信息
type GetCartReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCartReq) Reset() {
	*x = GetCartReq{}
	mi := &file_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartReq) ProtoMessage() {}

func (x *GetCartReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartReq.ProtoReflect.Descriptor instead.
func (*GetCartReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartReq) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetCartReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCartResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Cart          *Cart                  `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCartResp) Reset() {
	*x = GetCartResp{}
	mi := &file_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResp) ProtoMessage() {}

func (x *GetCartResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResp.ProtoReflect.Descriptor instead.
func (*GetCartResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetCartResp) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

// 购物车信息
type Cart struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   // 用户ID
	Items         []*CartItem            `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"` // 购物车商品列表
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cart) Reset() {
	*x = Cart{}
	mi := &file_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *Cart) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Cart) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// 删除购物车商品
type RemoveCartItemReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ProductId     uint32                 `protobuf:"varint,3,opt,name=product_id,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveCartItemReq) Reset() {
	*x = RemoveCartItemReq{}
	mi := &file_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCartItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCartItemReq) ProtoMessage() {}

func (x *RemoveCartItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCartItemReq.ProtoReflect.Descriptor instead.
func (*RemoveCartItemReq) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveCartItemReq) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *RemoveCartItemReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RemoveCartItemReq) GetProductId() uint32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type RemoveCartItemResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // 操作是否成功
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveCartItemResp) Reset() {
	*x = RemoveCartItemResp{}
	mi := &file_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCartItemResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCartItemResp) ProtoMessage() {}

func (x *RemoveCartItemResp) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCartItemResp.ProtoReflect.Descriptor instead.
func (*RemoveCartItemResp) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveCartItemResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x22, 0x64, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x2a, 0x0a, 0x0e, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x38, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29,
	0x0a, 0x0d, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x36, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x25, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x22, 0x5d, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x5d, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x86, 0x03, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0a, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x13, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72,
	0x74, 0x12, 0x4e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x12, 0x17, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72,
	0x74, 0x12, 0x54, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x12, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x2a, 0x08, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x12, 0x75, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x69, 0x74, 0x65,
	0x6d, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x0d,
	0x5a, 0x0b, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_proto_goTypes = []any{
	(*CartItem)(nil),           // 0: api.cart.v1.CartItem
	(*UpsertItemReq)(nil),      // 1: api.cart.v1.UpsertItemReq
	(*UpsertItemResp)(nil),     // 2: api.cart.v1.UpsertItemResp
	(*EmptyCartReq)(nil),       // 3: api.cart.v1.EmptyCartReq
	(*EmptyCartResp)(nil),      // 4: api.cart.v1.EmptyCartResp
	(*GetCartReq)(nil),         // 5: api.cart.v1.GetCartReq
	(*GetCartResp)(nil),        // 6: api.cart.v1.GetCartResp
	(*Cart)(nil),               // 7: api.cart.v1.Cart
	(*RemoveCartItemReq)(nil),  // 8: api.cart.v1.RemoveCartItemReq
	(*RemoveCartItemResp)(nil), // 9: api.cart.v1.RemoveCartItemResp
}
var file_service_proto_depIdxs = []int32{
	0, // 0: api.cart.v1.UpsertItemReq.item:type_name -> api.cart.v1.CartItem
	7, // 1: api.cart.v1.GetCartResp.cart:type_name -> api.cart.v1.Cart
	0, // 2: api.cart.v1.Cart.items:type_name -> api.cart.v1.CartItem
	1, // 3: api.cart.v1.CartService.UpsertItem:input_type -> api.cart.v1.UpsertItemReq
	5, // 4: api.cart.v1.CartService.GetCart:input_type -> api.cart.v1.GetCartReq
	3, // 5: api.cart.v1.CartService.EmptyCart:input_type -> api.cart.v1.EmptyCartReq
	8, // 6: api.cart.v1.CartService.RemoveCartItem:input_type -> api.cart.v1.RemoveCartItemReq
	2, // 7: api.cart.v1.CartService.UpsertItem:output_type -> api.cart.v1.UpsertItemResp
	6, // 8: api.cart.v1.CartService.GetCart:output_type -> api.cart.v1.GetCartResp
	4, // 9: api.cart.v1.CartService.EmptyCart:output_type -> api.cart.v1.EmptyCartResp
	9, // 10: api.cart.v1.CartService.RemoveCartItem:output_type -> api.cart.v1.RemoveCartItemResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
