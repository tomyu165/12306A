// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: order.proto

package orderRPCpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Money          string `protobuf:"bytes,2,opt,name=Money,proto3" json:"Money,omitempty"`
	AffairID       string `protobuf:"bytes,3,opt,name=AffairID,proto3" json:"AffairID,omitempty"`
	ExpireDuration int32  `protobuf:"varint,4,opt,name=ExpireDuration,proto3" json:"ExpireDuration,omitempty"`
	OrderOutsideID string `protobuf:"bytes,5,opt,name=OrderOutsideID,proto3" json:"OrderOutsideID,omitempty"`
}

func (x *CreateInfo) Reset() {
	*x = CreateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInfo) ProtoMessage() {}

func (x *CreateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInfo.ProtoReflect.Descriptor instead.
func (*CreateInfo) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateInfo) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateInfo) GetMoney() string {
	if x != nil {
		return x.Money
	}
	return ""
}

func (x *CreateInfo) GetAffairID() string {
	if x != nil {
		return x.AffairID
	}
	return ""
}

func (x *CreateInfo) GetExpireDuration() int32 {
	if x != nil {
		return x.ExpireDuration
	}
	return 0
}

func (x *CreateInfo) GetOrderOutsideID() string {
	if x != nil {
		return x.OrderOutsideID
	}
	return ""
}

type Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Money          string `protobuf:"bytes,2,opt,name=Money,proto3" json:"Money,omitempty"`
	AffairID       string `protobuf:"bytes,3,opt,name=AffairID,proto3" json:"AffairID,omitempty"`
	ExpireDuration int32  `protobuf:"varint,4,opt,name=ExpireDuration,proto3" json:"ExpireDuration,omitempty"`
	OrderOutsideID string `protobuf:"bytes,5,opt,name=OrderOutsideID,proto3" json:"OrderOutsideID,omitempty"`
	State          int32  `protobuf:"varint,6,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *Info) Reset() {
	*x = Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Info) ProtoMessage() {}

func (x *Info) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Info.ProtoReflect.Descriptor instead.
func (*Info) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *Info) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Info) GetMoney() string {
	if x != nil {
		return x.Money
	}
	return ""
}

func (x *Info) GetAffairID() string {
	if x != nil {
		return x.AffairID
	}
	return ""
}

func (x *Info) GetExpireDuration() int32 {
	if x != nil {
		return x.ExpireDuration
	}
	return 0
}

func (x *Info) GetOrderOutsideID() string {
	if x != nil {
		return x.OrderOutsideID
	}
	return ""
}

func (x *Info) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type SearchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *SearchInfo) Reset() {
	*x = SearchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchInfo) ProtoMessage() {}

func (x *SearchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchInfo.ProtoReflect.Descriptor instead.
func (*SearchInfo) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *SearchInfo) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *Error) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateStateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutsideID string `protobuf:"bytes,1,opt,name=OutsideID,proto3" json:"OutsideID,omitempty"`
	State     int32  `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *UpdateStateInfo) Reset() {
	*x = UpdateStateInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStateInfo) ProtoMessage() {}

func (x *UpdateStateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStateInfo.ProtoReflect.Descriptor instead.
func (*UpdateStateInfo) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateStateInfo) GetOutsideID() string {
	if x != nil {
		return x.OutsideID
	}
	return ""
}

func (x *UpdateStateInfo) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type UpdateStateWithRInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutsideID  string `protobuf:"bytes,1,opt,name=OutsideID,proto3" json:"OutsideID,omitempty"`
	State      int32  `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	ROutsideID string `protobuf:"bytes,3,opt,name=ROutsideID,proto3" json:"ROutsideID,omitempty"`
}

func (x *UpdateStateWithRInfo) Reset() {
	*x = UpdateStateWithRInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStateWithRInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStateWithRInfo) ProtoMessage() {}

func (x *UpdateStateWithRInfo) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStateWithRInfo.ProtoReflect.Descriptor instead.
func (*UpdateStateWithRInfo) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateStateWithRInfo) GetOutsideID() string {
	if x != nil {
		return x.OutsideID
	}
	return ""
}

func (x *UpdateStateWithRInfo) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *UpdateStateWithRInfo) GetROutsideID() string {
	if x != nil {
		return x.ROutsideID
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x66, 0x66, 0x61, 0x69, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x66, 0x66, 0x61, 0x69, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a,
	0x0e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x75,
	0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44, 0x22, 0xb6, 0x01,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x66, 0x66, 0x61, 0x69, 0x72, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x66, 0x66, 0x61, 0x69, 0x72, 0x49, 0x44,
	0x12, 0x26, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x24, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x21, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x45, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x6a, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65,
	0x49, 0x44, 0x32, 0xec, 0x01, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x50, 0x43, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x50, 0x43, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_order_proto_goTypes = []interface{}{
	(*CreateInfo)(nil),           // 0: proto.CreateInfo
	(*Info)(nil),                 // 1: proto.Info
	(*SearchInfo)(nil),           // 2: proto.SearchInfo
	(*Error)(nil),                // 3: proto.Error
	(*UpdateStateInfo)(nil),      // 4: proto.UpdateStateInfo
	(*UpdateStateWithRInfo)(nil), // 5: proto.UpdateStateWithRInfo
}
var file_order_proto_depIdxs = []int32{
	0, // 0: proto.OrderRPCService.Create:input_type -> proto.CreateInfo
	2, // 1: proto.OrderRPCService.Read:input_type -> proto.SearchInfo
	4, // 2: proto.OrderRPCService.UpdateState:input_type -> proto.UpdateStateInfo
	5, // 3: proto.OrderRPCService.UpdateStateWithRelativeOrder:input_type -> proto.UpdateStateWithRInfo
	3, // 4: proto.OrderRPCService.Create:output_type -> proto.Error
	1, // 5: proto.OrderRPCService.Read:output_type -> proto.Info
	3, // 6: proto.OrderRPCService.UpdateState:output_type -> proto.Error
	3, // 7: proto.OrderRPCService.UpdateStateWithRelativeOrder:output_type -> proto.Error
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Info); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStateInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStateWithRInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderRPCServiceClient is the client API for OrderRPCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderRPCServiceClient interface {
	Create(ctx context.Context, in *CreateInfo, opts ...grpc.CallOption) (*Error, error)
	Read(ctx context.Context, in *SearchInfo, opts ...grpc.CallOption) (*Info, error)
	UpdateState(ctx context.Context, in *UpdateStateInfo, opts ...grpc.CallOption) (*Error, error)
	UpdateStateWithRelativeOrder(ctx context.Context, in *UpdateStateWithRInfo, opts ...grpc.CallOption) (*Error, error)
}

type orderRPCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderRPCServiceClient(cc grpc.ClientConnInterface) OrderRPCServiceClient {
	return &orderRPCServiceClient{cc}
}

func (c *orderRPCServiceClient) Create(ctx context.Context, in *CreateInfo, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/proto.OrderRPCService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderRPCServiceClient) Read(ctx context.Context, in *SearchInfo, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := c.cc.Invoke(ctx, "/proto.OrderRPCService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderRPCServiceClient) UpdateState(ctx context.Context, in *UpdateStateInfo, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/proto.OrderRPCService/UpdateState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderRPCServiceClient) UpdateStateWithRelativeOrder(ctx context.Context, in *UpdateStateWithRInfo, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/proto.OrderRPCService/UpdateStateWithRelativeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderRPCServiceServer is the server API for OrderRPCService service.
type OrderRPCServiceServer interface {
	Create(context.Context, *CreateInfo) (*Error, error)
	Read(context.Context, *SearchInfo) (*Info, error)
	UpdateState(context.Context, *UpdateStateInfo) (*Error, error)
	UpdateStateWithRelativeOrder(context.Context, *UpdateStateWithRInfo) (*Error, error)
}

// UnimplementedOrderRPCServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderRPCServiceServer struct {
}

func (*UnimplementedOrderRPCServiceServer) Create(context.Context, *CreateInfo) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedOrderRPCServiceServer) Read(context.Context, *SearchInfo) (*Info, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedOrderRPCServiceServer) UpdateState(context.Context, *UpdateStateInfo) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateState not implemented")
}
func (*UnimplementedOrderRPCServiceServer) UpdateStateWithRelativeOrder(context.Context, *UpdateStateWithRInfo) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStateWithRelativeOrder not implemented")
}

func RegisterOrderRPCServiceServer(s *grpc.Server, srv OrderRPCServiceServer) {
	s.RegisterService(&_OrderRPCService_serviceDesc, srv)
}

func _OrderRPCService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderRPCServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrderRPCService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderRPCServiceServer).Create(ctx, req.(*CreateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderRPCService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderRPCServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrderRPCService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderRPCServiceServer).Read(ctx, req.(*SearchInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderRPCService_UpdateState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStateInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderRPCServiceServer).UpdateState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrderRPCService/UpdateState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderRPCServiceServer).UpdateState(ctx, req.(*UpdateStateInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderRPCService_UpdateStateWithRelativeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStateWithRInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderRPCServiceServer).UpdateStateWithRelativeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OrderRPCService/UpdateStateWithRelativeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderRPCServiceServer).UpdateStateWithRelativeOrder(ctx, req.(*UpdateStateWithRInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderRPCService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OrderRPCService",
	HandlerType: (*OrderRPCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderRPCService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _OrderRPCService_Read_Handler,
		},
		{
			MethodName: "UpdateState",
			Handler:    _OrderRPCService_UpdateState_Handler,
		},
		{
			MethodName: "UpdateStateWithRelativeOrder",
			Handler:    _OrderRPCService_UpdateStateWithRelativeOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
