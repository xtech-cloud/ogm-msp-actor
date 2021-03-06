// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/actor/domain.proto

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

// 创建的请求
type DomainCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 域的名称
}

func (x *DomainCreateRequest) Reset() {
	*x = DomainCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainCreateRequest) ProtoMessage() {}

func (x *DomainCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainCreateRequest.ProtoReflect.Descriptor instead.
func (*DomainCreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{0}
}

func (x *DomainCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 更新的请求
type DomainUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // 域的UUID
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // 域的名称
}

func (x *DomainUpdateRequest) Reset() {
	*x = DomainUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainUpdateRequest) ProtoMessage() {}

func (x *DomainUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainUpdateRequest.ProtoReflect.Descriptor instead.
func (*DomainUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{1}
}

func (x *DomainUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DomainUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 删除的请求
type DomainDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // 域的uuid
}

func (x *DomainDeleteRequest) Reset() {
	*x = DomainDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainDeleteRequest) ProtoMessage() {}

func (x *DomainDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainDeleteRequest.ProtoReflect.Descriptor instead.
func (*DomainDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{2}
}

func (x *DomainDeleteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 列举的回复
type DomainListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 总数
	Domain []*DomainEntity `protobuf:"bytes,3,rep,name=domain,proto3" json:"domain,omitempty"` //域实体
}

func (x *DomainListResponse) Reset() {
	*x = DomainListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainListResponse) ProtoMessage() {}

func (x *DomainListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainListResponse.ProtoReflect.Descriptor instead.
func (*DomainListResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{3}
}

func (x *DomainListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DomainListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DomainListResponse) GetDomain() []*DomainEntity {
	if x != nil {
		return x.Domain
	}
	return nil
}

// 获取的请求
type DomainGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // 域的uuid
}

func (x *DomainGetRequest) Reset() {
	*x = DomainGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainGetRequest) ProtoMessage() {}

func (x *DomainGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainGetRequest.ProtoReflect.Descriptor instead.
func (*DomainGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{4}
}

func (x *DomainGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取的回复
type DomainGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Domain *DomainEntity `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"` //域实体
}

func (x *DomainGetResponse) Reset() {
	*x = DomainGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainGetResponse) ProtoMessage() {}

func (x *DomainGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainGetResponse.ProtoReflect.Descriptor instead.
func (*DomainGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{5}
}

func (x *DomainGetResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DomainGetResponse) GetDomain() *DomainEntity {
	if x != nil {
		return x.Domain
	}
	return nil
}

// 查找的请求
type DomainFindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // 域的名称
}

func (x *DomainFindRequest) Reset() {
	*x = DomainFindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainFindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainFindRequest) ProtoMessage() {}

func (x *DomainFindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainFindRequest.ProtoReflect.Descriptor instead.
func (*DomainFindRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{6}
}

func (x *DomainFindRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 查找的回复
type DomainFindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Domain *DomainEntity `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"` //域实体
}

func (x *DomainFindResponse) Reset() {
	*x = DomainFindResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainFindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainFindResponse) ProtoMessage() {}

func (x *DomainFindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainFindResponse.ProtoReflect.Descriptor instead.
func (*DomainFindResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{7}
}

func (x *DomainFindResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DomainFindResponse) GetDomain() *DomainEntity {
	if x != nil {
		return x.Domain
	}
	return nil
}

// 搜索的请求
type DomainSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`      // 域的名称
}

func (x *DomainSearchRequest) Reset() {
	*x = DomainSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainSearchRequest) ProtoMessage() {}

func (x *DomainSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainSearchRequest.ProtoReflect.Descriptor instead.
func (*DomainSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{8}
}

func (x *DomainSearchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DomainSearchRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DomainSearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 搜索的回复
type DomainSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 总数
	Domain []*DomainEntity `protobuf:"bytes,3,rep,name=domain,proto3" json:"domain,omitempty"` //域实体
}

func (x *DomainSearchResponse) Reset() {
	*x = DomainSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainSearchResponse) ProtoMessage() {}

func (x *DomainSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainSearchResponse.ProtoReflect.Descriptor instead.
func (*DomainSearchResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{9}
}

func (x *DomainSearchResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DomainSearchResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DomainSearchResponse) GetDomain() []*DomainEntity {
	if x != nil {
		return x.Domain
	}
	return nil
}

// 执行的请求
type DomainExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`           // 域的uuid
	Command   string   `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`     // 指令
	Device    []string `protobuf:"bytes,3,rep,name=device,proto3" json:"device,omitempty"`       // 设备序列号
	Parameter string   `protobuf:"bytes,4,opt,name=parameter,proto3" json:"parameter,omitempty"` // 参数(base64编码的json格式)
}

func (x *DomainExecuteRequest) Reset() {
	*x = DomainExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainExecuteRequest) ProtoMessage() {}

func (x *DomainExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_domain_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainExecuteRequest.ProtoReflect.Descriptor instead.
func (*DomainExecuteRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{10}
}

func (x *DomainExecuteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DomainExecuteRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *DomainExecuteRequest) GetDevice() []string {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *DomainExecuteRequest) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

var File_proto_actor_domain_proto protoreflect.FileDescriptor

var file_proto_actor_domain_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x13, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3d, 0x0a, 0x13, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x7e, 0x0a, 0x12, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0x26, 0x0a, 0x10, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x11, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x12, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x22, 0x57, 0x0a, 0x13, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x80, 0x01,
	0x0a, 0x14, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0x7a, 0x0a, 0x14, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x32, 0xf8, 0x03, 0x0a,
	0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17,
	0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x3b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_actor_domain_proto_rawDescOnce sync.Once
	file_proto_actor_domain_proto_rawDescData = file_proto_actor_domain_proto_rawDesc
)

func file_proto_actor_domain_proto_rawDescGZIP() []byte {
	file_proto_actor_domain_proto_rawDescOnce.Do(func() {
		file_proto_actor_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_actor_domain_proto_rawDescData)
	})
	return file_proto_actor_domain_proto_rawDescData
}

var file_proto_actor_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_actor_domain_proto_goTypes = []interface{}{
	(*DomainCreateRequest)(nil),  // 0: actor.DomainCreateRequest
	(*DomainUpdateRequest)(nil),  // 1: actor.DomainUpdateRequest
	(*DomainDeleteRequest)(nil),  // 2: actor.DomainDeleteRequest
	(*DomainListResponse)(nil),   // 3: actor.DomainListResponse
	(*DomainGetRequest)(nil),     // 4: actor.DomainGetRequest
	(*DomainGetResponse)(nil),    // 5: actor.DomainGetResponse
	(*DomainFindRequest)(nil),    // 6: actor.DomainFindRequest
	(*DomainFindResponse)(nil),   // 7: actor.DomainFindResponse
	(*DomainSearchRequest)(nil),  // 8: actor.DomainSearchRequest
	(*DomainSearchResponse)(nil), // 9: actor.DomainSearchResponse
	(*DomainExecuteRequest)(nil), // 10: actor.DomainExecuteRequest
	(*Status)(nil),               // 11: actor.Status
	(*DomainEntity)(nil),         // 12: actor.DomainEntity
	(*ListRequest)(nil),          // 13: actor.ListRequest
	(*UuidResponse)(nil),         // 14: actor.UuidResponse
	(*BlankResponse)(nil),        // 15: actor.BlankResponse
}
var file_proto_actor_domain_proto_depIdxs = []int32{
	11, // 0: actor.DomainListResponse.status:type_name -> actor.Status
	12, // 1: actor.DomainListResponse.domain:type_name -> actor.DomainEntity
	11, // 2: actor.DomainGetResponse.status:type_name -> actor.Status
	12, // 3: actor.DomainGetResponse.domain:type_name -> actor.DomainEntity
	11, // 4: actor.DomainFindResponse.status:type_name -> actor.Status
	12, // 5: actor.DomainFindResponse.domain:type_name -> actor.DomainEntity
	11, // 6: actor.DomainSearchResponse.status:type_name -> actor.Status
	12, // 7: actor.DomainSearchResponse.domain:type_name -> actor.DomainEntity
	0,  // 8: actor.Domain.Create:input_type -> actor.DomainCreateRequest
	1,  // 9: actor.Domain.Update:input_type -> actor.DomainUpdateRequest
	2,  // 10: actor.Domain.Delete:input_type -> actor.DomainDeleteRequest
	13, // 11: actor.Domain.List:input_type -> actor.ListRequest
	4,  // 12: actor.Domain.Get:input_type -> actor.DomainGetRequest
	6,  // 13: actor.Domain.Find:input_type -> actor.DomainFindRequest
	8,  // 14: actor.Domain.Search:input_type -> actor.DomainSearchRequest
	10, // 15: actor.Domain.Execute:input_type -> actor.DomainExecuteRequest
	14, // 16: actor.Domain.Create:output_type -> actor.UuidResponse
	14, // 17: actor.Domain.Update:output_type -> actor.UuidResponse
	14, // 18: actor.Domain.Delete:output_type -> actor.UuidResponse
	3,  // 19: actor.Domain.List:output_type -> actor.DomainListResponse
	5,  // 20: actor.Domain.Get:output_type -> actor.DomainGetResponse
	7,  // 21: actor.Domain.Find:output_type -> actor.DomainFindResponse
	9,  // 22: actor.Domain.Search:output_type -> actor.DomainSearchResponse
	15, // 23: actor.Domain.Execute:output_type -> actor.BlankResponse
	16, // [16:24] is the sub-list for method output_type
	8,  // [8:16] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_actor_domain_proto_init() }
func file_proto_actor_domain_proto_init() {
	if File_proto_actor_domain_proto != nil {
		return
	}
	file_proto_actor_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_actor_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainCreateRequest); i {
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
		file_proto_actor_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainUpdateRequest); i {
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
		file_proto_actor_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainDeleteRequest); i {
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
		file_proto_actor_domain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainListResponse); i {
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
		file_proto_actor_domain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainGetRequest); i {
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
		file_proto_actor_domain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainGetResponse); i {
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
		file_proto_actor_domain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainFindRequest); i {
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
		file_proto_actor_domain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainFindResponse); i {
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
		file_proto_actor_domain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainSearchRequest); i {
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
		file_proto_actor_domain_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainSearchResponse); i {
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
		file_proto_actor_domain_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainExecuteRequest); i {
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
			RawDescriptor: file_proto_actor_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_actor_domain_proto_goTypes,
		DependencyIndexes: file_proto_actor_domain_proto_depIdxs,
		MessageInfos:      file_proto_actor_domain_proto_msgTypes,
	}.Build()
	File_proto_actor_domain_proto = out.File
	file_proto_actor_domain_proto_rawDesc = nil
	file_proto_actor_domain_proto_goTypes = nil
	file_proto_actor_domain_proto_depIdxs = nil
}
