// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/actor/application.proto

package actor

import (
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

// 添加的请求
type ApplicationAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain   string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`      // 域的uuid
	Name     string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`         // 应用名
	Version  string `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`   // 应用版本
	Program  string `protobuf:"bytes,12,opt,name=program,proto3" json:"program,omitempty"`   // 可执行文件
	Location string `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"` // 安装路径
	Url      string `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`           // 下载地址
	Upgrade  int32  `protobuf:"varint,30,opt,name=upgrade,proto3" json:"upgrade,omitempty"`  // 更新策略
}

func (x *ApplicationAddRequest) Reset() {
	*x = ApplicationAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationAddRequest) ProtoMessage() {}

func (x *ApplicationAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationAddRequest.ProtoReflect.Descriptor instead.
func (*ApplicationAddRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_application_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationAddRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *ApplicationAddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationAddRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ApplicationAddRequest) GetProgram() string {
	if x != nil {
		return x.Program
	}
	return ""
}

func (x *ApplicationAddRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ApplicationAddRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ApplicationAddRequest) GetUpgrade() int32 {
	if x != nil {
		return x.Upgrade
	}
	return 0
}

// 更新的请求
type ApplicationUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`          // 应用的UUID
	Name     string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`         // 应用名
	Version  string `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`   // 应用版本
	Program  string `protobuf:"bytes,12,opt,name=program,proto3" json:"program,omitempty"`   // 可执行文件
	Location string `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"` // 安装路径
	Url      string `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`           // 下载地址
	Upgrade  int32  `protobuf:"varint,30,opt,name=upgrade,proto3" json:"upgrade,omitempty"`  // 更新策略
}

func (x *ApplicationUpdateRequest) Reset() {
	*x = ApplicationUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationUpdateRequest) ProtoMessage() {}

func (x *ApplicationUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationUpdateRequest.ProtoReflect.Descriptor instead.
func (*ApplicationUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_application_proto_rawDescGZIP(), []int{1}
}

func (x *ApplicationUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ApplicationUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApplicationUpdateRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ApplicationUpdateRequest) GetProgram() string {
	if x != nil {
		return x.Program
	}
	return ""
}

func (x *ApplicationUpdateRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ApplicationUpdateRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ApplicationUpdateRequest) GetUpgrade() int32 {
	if x != nil {
		return x.Upgrade
	}
	return 0
}

// 移除的请求
type ApplicationRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // 应用的uuid
}

func (x *ApplicationRemoveRequest) Reset() {
	*x = ApplicationRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationRemoveRequest) ProtoMessage() {}

func (x *ApplicationRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationRemoveRequest.ProtoReflect.Descriptor instead.
func (*ApplicationRemoveRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_application_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationRemoveRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取的请求
type ApplicationGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // 应用的uuid
}

func (x *ApplicationGetRequest) Reset() {
	*x = ApplicationGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationGetRequest) ProtoMessage() {}

func (x *ApplicationGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationGetRequest.ProtoReflect.Descriptor instead.
func (*ApplicationGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_application_proto_rawDescGZIP(), []int{3}
}

func (x *ApplicationGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取的回复
type ApplicationGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      *Status            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`           // 状态
	Application *ApplicationEntity `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"` // 应用实体
}

func (x *ApplicationGetResponse) Reset() {
	*x = ApplicationGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_application_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationGetResponse) ProtoMessage() {}

func (x *ApplicationGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_application_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationGetResponse.ProtoReflect.Descriptor instead.
func (*ApplicationGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_application_proto_rawDescGZIP(), []int{4}
}

func (x *ApplicationGetResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ApplicationGetResponse) GetApplication() *ApplicationEntity {
	if x != nil {
		return x.Application
	}
	return nil
}

// 列举的请求
type ApplicationListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`  // 域的uuid
}

func (x *ApplicationListRequest) Reset() {
	*x = ApplicationListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_application_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationListRequest) ProtoMessage() {}

func (x *ApplicationListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_application_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationListRequest.ProtoReflect.Descriptor instead.
func (*ApplicationListRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_application_proto_rawDescGZIP(), []int{5}
}

func (x *ApplicationListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ApplicationListRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ApplicationListRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

// 列举的回复
type ApplicationListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      *Status              `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`           // 状态
	Total       int64                `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`            // 总数
	Application []*ApplicationEntity `protobuf:"bytes,3,rep,name=application,proto3" json:"application,omitempty"` // 应用实体列表
}

func (x *ApplicationListResponse) Reset() {
	*x = ApplicationListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_application_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationListResponse) ProtoMessage() {}

func (x *ApplicationListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_application_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationListResponse.ProtoReflect.Descriptor instead.
func (*ApplicationListResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_application_proto_rawDescGZIP(), []int{6}
}

func (x *ApplicationListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ApplicationListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ApplicationListResponse) GetApplication() []*ApplicationEntity {
	if x != nil {
		return x.Application
	}
	return nil
}

var File_proto_actor_application_proto protoreflect.FileDescriptor

var file_proto_actor_application_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbf, 0x01, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x70, 0x67, 0x72,
	0x61, 0x64, 0x65, 0x22, 0x2e, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x15, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x7b, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5e, 0x0a,
	0x16, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x92, 0x01,
	0x0a, 0x17, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x32, 0xdc, 0x02, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1c, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1c, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x3b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_actor_application_proto_rawDescOnce sync.Once
	file_proto_actor_application_proto_rawDescData = file_proto_actor_application_proto_rawDesc
)

func file_proto_actor_application_proto_rawDescGZIP() []byte {
	file_proto_actor_application_proto_rawDescOnce.Do(func() {
		file_proto_actor_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_actor_application_proto_rawDescData)
	})
	return file_proto_actor_application_proto_rawDescData
}

var file_proto_actor_application_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_actor_application_proto_goTypes = []interface{}{
	(*ApplicationAddRequest)(nil),    // 0: actor.ApplicationAddRequest
	(*ApplicationUpdateRequest)(nil), // 1: actor.ApplicationUpdateRequest
	(*ApplicationRemoveRequest)(nil), // 2: actor.ApplicationRemoveRequest
	(*ApplicationGetRequest)(nil),    // 3: actor.ApplicationGetRequest
	(*ApplicationGetResponse)(nil),   // 4: actor.ApplicationGetResponse
	(*ApplicationListRequest)(nil),   // 5: actor.ApplicationListRequest
	(*ApplicationListResponse)(nil),  // 6: actor.ApplicationListResponse
	(*Status)(nil),                   // 7: actor.Status
	(*ApplicationEntity)(nil),        // 8: actor.ApplicationEntity
	(*UuidResponse)(nil),             // 9: actor.UuidResponse
}
var file_proto_actor_application_proto_depIdxs = []int32{
	7, // 0: actor.ApplicationGetResponse.status:type_name -> actor.Status
	8, // 1: actor.ApplicationGetResponse.application:type_name -> actor.ApplicationEntity
	7, // 2: actor.ApplicationListResponse.status:type_name -> actor.Status
	8, // 3: actor.ApplicationListResponse.application:type_name -> actor.ApplicationEntity
	0, // 4: actor.Application.Add:input_type -> actor.ApplicationAddRequest
	2, // 5: actor.Application.Remove:input_type -> actor.ApplicationRemoveRequest
	5, // 6: actor.Application.List:input_type -> actor.ApplicationListRequest
	3, // 7: actor.Application.Get:input_type -> actor.ApplicationGetRequest
	1, // 8: actor.Application.Update:input_type -> actor.ApplicationUpdateRequest
	9, // 9: actor.Application.Add:output_type -> actor.UuidResponse
	9, // 10: actor.Application.Remove:output_type -> actor.UuidResponse
	6, // 11: actor.Application.List:output_type -> actor.ApplicationListResponse
	4, // 12: actor.Application.Get:output_type -> actor.ApplicationGetResponse
	9, // 13: actor.Application.Update:output_type -> actor.UuidResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_actor_application_proto_init() }
func file_proto_actor_application_proto_init() {
	if File_proto_actor_application_proto != nil {
		return
	}
	file_proto_actor_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_actor_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationAddRequest); i {
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
		file_proto_actor_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationUpdateRequest); i {
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
		file_proto_actor_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationRemoveRequest); i {
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
		file_proto_actor_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationGetRequest); i {
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
		file_proto_actor_application_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationGetResponse); i {
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
		file_proto_actor_application_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationListRequest); i {
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
		file_proto_actor_application_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationListResponse); i {
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
			RawDescriptor: file_proto_actor_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_actor_application_proto_goTypes,
		DependencyIndexes: file_proto_actor_application_proto_depIdxs,
		MessageInfos:      file_proto_actor_application_proto_msgTypes,
	}.Build()
	File_proto_actor_application_proto = out.File
	file_proto_actor_application_proto_rawDesc = nil
	file_proto_actor_application_proto_goTypes = nil
	file_proto_actor_application_proto_depIdxs = nil
}
