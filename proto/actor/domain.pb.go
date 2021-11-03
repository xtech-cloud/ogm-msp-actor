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
		mi := &file_proto_actor_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainDeleteRequest) ProtoMessage() {}

func (x *DomainDeleteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DomainDeleteRequest.ProtoReflect.Descriptor instead.
func (*DomainDeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{1}
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
		mi := &file_proto_actor_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainListResponse) ProtoMessage() {}

func (x *DomainListResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DomainListResponse.ProtoReflect.Descriptor instead.
func (*DomainListResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{2}
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

// 执行的请求
type DomainExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`           // 域的uuid
	Command   string   `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`     // 指令
	Device    []string `protobuf:"bytes,3,rep,name=device,proto3" json:"device,omitempty"`       // 设备ID
	Parameter string   `protobuf:"bytes,4,opt,name=parameter,proto3" json:"parameter,omitempty"` // 参数(base64编码的json格式)
}

func (x *DomainExecuteRequest) Reset() {
	*x = DomainExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainExecuteRequest) ProtoMessage() {}

func (x *DomainExecuteRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DomainExecuteRequest.ProtoReflect.Descriptor instead.
func (*DomainExecuteRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{3}
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

// 获取设备的请求
type DomainFetchDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string            `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`                                                                                             // 域的uuid
	Offset int64             `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`                                                                                        // 偏移值
	Count  int64             `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`                                                                                          // 数量
	Filter map[string]string `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 过滤器
}

func (x *DomainFetchDeviceRequest) Reset() {
	*x = DomainFetchDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainFetchDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainFetchDeviceRequest) ProtoMessage() {}

func (x *DomainFetchDeviceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DomainFetchDeviceRequest.ProtoReflect.Descriptor instead.
func (*DomainFetchDeviceRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{4}
}

func (x *DomainFetchDeviceRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DomainFetchDeviceRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DomainFetchDeviceRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DomainFetchDeviceRequest) GetFilter() map[string]string {
	if x != nil {
		return x.Filter
	}
	return nil
}

// 获取设备的回复
type DomainFetchDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                                                                                          // 状态
	Total  int64             `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`                                                                                           // 总数
	Device []*DeviceEntity   `protobuf:"bytes,3,rep,name=device,proto3" json:"device,omitempty"`                                                                                          // 设备实体列表
	Access map[string]int32  `protobuf:"bytes,4,rep,name=access,proto3" json:"access,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"` // 设备访问权限
	Alias  map[string]string `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`    // 设备别名
}

func (x *DomainFetchDeviceResponse) Reset() {
	*x = DomainFetchDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainFetchDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainFetchDeviceResponse) ProtoMessage() {}

func (x *DomainFetchDeviceResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DomainFetchDeviceResponse.ProtoReflect.Descriptor instead.
func (*DomainFetchDeviceResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{5}
}

func (x *DomainFetchDeviceResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DomainFetchDeviceResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DomainFetchDeviceResponse) GetDevice() []*DeviceEntity {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *DomainFetchDeviceResponse) GetAccess() map[string]int32 {
	if x != nil {
		return x.Access
	}
	return nil
}

func (x *DomainFetchDeviceResponse) GetAlias() map[string]string {
	if x != nil {
		return x.Alias
	}
	return nil
}

// 编辑设备的请求
type DomainEditDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`      // 域的uuid
	Device string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`  // 设备的uuid
	Access int32  `protobuf:"varint,3,opt,name=access,proto3" json:"access,omitempty"` // 访问权限(1:允许，2：拒绝)
	Alias  string `protobuf:"bytes,4,opt,name=alias,proto3" json:"alias,omitempty"`    // 别名
}

func (x *DomainEditDeviceRequest) Reset() {
	*x = DomainEditDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_domain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainEditDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainEditDeviceRequest) ProtoMessage() {}

func (x *DomainEditDeviceRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DomainEditDeviceRequest.ProtoReflect.Descriptor instead.
func (*DomainEditDeviceRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_domain_proto_rawDescGZIP(), []int{6}
}

func (x *DomainEditDeviceRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DomainEditDeviceRequest) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *DomainEditDeviceRequest) GetAccess() int32 {
	if x != nil {
		return x.Access
	}
	return 0
}

func (x *DomainEditDeviceRequest) GetAlias() string {
	if x != nil {
		return x.Alias
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
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x7e, 0x0a, 0x12, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x7a, 0x0a, 0x14, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x22, 0xdc, 0x01,
	0x0a, 0x18, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x83, 0x03, 0x0a,
	0x19, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x73, 0x0a, 0x17, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x64, 0x69, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x32, 0x97, 0x03, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x42,
	0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0a, 0x45,
	0x64, 0x69, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x64, 0x69, 0x74, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x3b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_actor_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_actor_domain_proto_goTypes = []interface{}{
	(*DomainCreateRequest)(nil),       // 0: actor.DomainCreateRequest
	(*DomainDeleteRequest)(nil),       // 1: actor.DomainDeleteRequest
	(*DomainListResponse)(nil),        // 2: actor.DomainListResponse
	(*DomainExecuteRequest)(nil),      // 3: actor.DomainExecuteRequest
	(*DomainFetchDeviceRequest)(nil),  // 4: actor.DomainFetchDeviceRequest
	(*DomainFetchDeviceResponse)(nil), // 5: actor.DomainFetchDeviceResponse
	(*DomainEditDeviceRequest)(nil),   // 6: actor.DomainEditDeviceRequest
	nil,                               // 7: actor.DomainFetchDeviceRequest.FilterEntry
	nil,                               // 8: actor.DomainFetchDeviceResponse.AccessEntry
	nil,                               // 9: actor.DomainFetchDeviceResponse.AliasEntry
	(*Status)(nil),                    // 10: actor.Status
	(*DomainEntity)(nil),              // 11: actor.DomainEntity
	(*DeviceEntity)(nil),              // 12: actor.DeviceEntity
	(*ListRequest)(nil),               // 13: actor.ListRequest
	(*BlankResponse)(nil),             // 14: actor.BlankResponse
}
var file_proto_actor_domain_proto_depIdxs = []int32{
	10, // 0: actor.DomainListResponse.status:type_name -> actor.Status
	11, // 1: actor.DomainListResponse.domain:type_name -> actor.DomainEntity
	7,  // 2: actor.DomainFetchDeviceRequest.filter:type_name -> actor.DomainFetchDeviceRequest.FilterEntry
	10, // 3: actor.DomainFetchDeviceResponse.status:type_name -> actor.Status
	12, // 4: actor.DomainFetchDeviceResponse.device:type_name -> actor.DeviceEntity
	8,  // 5: actor.DomainFetchDeviceResponse.access:type_name -> actor.DomainFetchDeviceResponse.AccessEntry
	9,  // 6: actor.DomainFetchDeviceResponse.alias:type_name -> actor.DomainFetchDeviceResponse.AliasEntry
	0,  // 7: actor.Domain.Create:input_type -> actor.DomainCreateRequest
	1,  // 8: actor.Domain.Delete:input_type -> actor.DomainDeleteRequest
	13, // 9: actor.Domain.List:input_type -> actor.ListRequest
	3,  // 10: actor.Domain.Execute:input_type -> actor.DomainExecuteRequest
	4,  // 11: actor.Domain.FetchDevice:input_type -> actor.DomainFetchDeviceRequest
	6,  // 12: actor.Domain.EditDevice:input_type -> actor.DomainEditDeviceRequest
	14, // 13: actor.Domain.Create:output_type -> actor.BlankResponse
	14, // 14: actor.Domain.Delete:output_type -> actor.BlankResponse
	2,  // 15: actor.Domain.List:output_type -> actor.DomainListResponse
	14, // 16: actor.Domain.Execute:output_type -> actor.BlankResponse
	5,  // 17: actor.Domain.FetchDevice:output_type -> actor.DomainFetchDeviceResponse
	14, // 18: actor.Domain.EditDevice:output_type -> actor.BlankResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
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
		file_proto_actor_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_actor_domain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_actor_domain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainFetchDeviceRequest); i {
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
			switch v := v.(*DomainFetchDeviceResponse); i {
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
			switch v := v.(*DomainEditDeviceRequest); i {
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
			NumMessages:   10,
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
