// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/actor/sync.proto

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

// 推送的请求
type SyncPushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain       string            `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`                                                                                                 // 域的UUID
	Device       *DeviceEntity     `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`                                                                                                 // 设备实体
	UpProperty   map[string]string `protobuf:"bytes,3,rep,name=upProperty,proto3" json:"upProperty,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 上行属性（提交给服务的属性）
	DownProperty []string          `protobuf:"bytes,4,rep,name=downProperty,proto3" json:"downProperty,omitempty"`                                                                                     // 下行属性（希望服务在回复时包含的属性）
	Task         []string          `protobuf:"bytes,5,rep,name=task,proto3" json:"task,omitempty"`                                                                                                     // 完成的任务
}

func (x *SyncPushRequest) Reset() {
	*x = SyncPushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPushRequest) ProtoMessage() {}

func (x *SyncPushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPushRequest.ProtoReflect.Descriptor instead.
func (*SyncPushRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_sync_proto_rawDescGZIP(), []int{0}
}

func (x *SyncPushRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SyncPushRequest) GetDevice() *DeviceEntity {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *SyncPushRequest) GetUpProperty() map[string]string {
	if x != nil {
		return x.UpProperty
	}
	return nil
}

func (x *SyncPushRequest) GetDownProperty() []string {
	if x != nil {
		return x.DownProperty
	}
	return nil
}

func (x *SyncPushRequest) GetTask() []string {
	if x != nil {
		return x.Task
	}
	return nil
}

// 推送的回复
type SyncPushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *Status           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                                                                                             // 状态
	Access   int32             `protobuf:"varint,2,opt,name=access,proto3" json:"access,omitempty"`                                                                                            // 访问权限 (0:未处理，1:接受，2：拒绝)
	Alias    string            `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`                                                                                               // 别名
	Property map[string]string `protobuf:"bytes,4,rep,name=property,proto3" json:"property,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 属性
	Task     map[string]string `protobuf:"bytes,5,rep,name=task,proto3" json:"task,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`         // 执行的任务
}

func (x *SyncPushResponse) Reset() {
	*x = SyncPushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPushResponse) ProtoMessage() {}

func (x *SyncPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPushResponse.ProtoReflect.Descriptor instead.
func (*SyncPushResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_sync_proto_rawDescGZIP(), []int{1}
}

func (x *SyncPushResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SyncPushResponse) GetAccess() int32 {
	if x != nil {
		return x.Access
	}
	return 0
}

func (x *SyncPushResponse) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *SyncPushResponse) GetProperty() map[string]string {
	if x != nil {
		return x.Property
	}
	return nil
}

func (x *SyncPushResponse) GetTask() map[string]string {
	if x != nil {
		return x.Task
	}
	return nil
}

// 拉取的请求
type SyncPullRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain       string   `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`             // 域的名称
	DownProperty []string `protobuf:"bytes,2,rep,name=downProperty,proto3" json:"downProperty,omitempty"` // 下行属性（希望服务在回复时包含的属性）
}

func (x *SyncPullRequest) Reset() {
	*x = SyncPullRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_sync_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPullRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPullRequest) ProtoMessage() {}

func (x *SyncPullRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_sync_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPullRequest.ProtoReflect.Descriptor instead.
func (*SyncPullRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_sync_proto_rawDescGZIP(), []int{2}
}

func (x *SyncPullRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SyncPullRequest) GetDownProperty() []string {
	if x != nil {
		return x.DownProperty
	}
	return nil
}

// 拉取的回复
type SyncPullResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   *Status           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`                                                                                             // 状态
	Device   []*DeviceEntity   `protobuf:"bytes,2,rep,name=device,proto3" json:"device,omitempty"`                                                                                             // 设备列表
	Alias    map[string]string `protobuf:"bytes,3,rep,name=alias,proto3" json:"alias,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`       // 设备别名
	Property map[string]string `protobuf:"bytes,4,rep,name=property,proto3" json:"property,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 属性
	Healthy  map[string]int32  `protobuf:"bytes,5,rep,name=healthy,proto3" json:"healthy,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`  // 健康值
}

func (x *SyncPullResponse) Reset() {
	*x = SyncPullResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_sync_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPullResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPullResponse) ProtoMessage() {}

func (x *SyncPullResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_sync_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPullResponse.ProtoReflect.Descriptor instead.
func (*SyncPullResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_sync_proto_rawDescGZIP(), []int{3}
}

func (x *SyncPullResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SyncPullResponse) GetDevice() []*DeviceEntity {
	if x != nil {
		return x.Device
	}
	return nil
}

func (x *SyncPullResponse) GetAlias() map[string]string {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *SyncPullResponse) GetProperty() map[string]string {
	if x != nil {
		return x.Property
	}
	return nil
}

func (x *SyncPullResponse) GetHealthy() map[string]int32 {
	if x != nil {
		return x.Healthy
	}
	return nil
}

// 上传的请求
type SyncUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"` // 域的UUID
	Device string `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"` // 设备的序列号
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`     // 数据名
	Data   string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`     // 数据包(base64编码)
}

func (x *SyncUploadRequest) Reset() {
	*x = SyncUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_sync_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncUploadRequest) ProtoMessage() {}

func (x *SyncUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_sync_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncUploadRequest.ProtoReflect.Descriptor instead.
func (*SyncUploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_sync_proto_rawDescGZIP(), []int{4}
}

func (x *SyncUploadRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SyncUploadRequest) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *SyncUploadRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SyncUploadRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_proto_actor_sync_proto protoreflect.FileDescriptor

var file_proto_actor_sync_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x1a,
	0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x02, 0x0a, 0x0f, 0x53, 0x79,
	0x6e, 0x63, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55,
	0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x75, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x6f,
	0x77, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x1a, 0x3d, 0x0a, 0x0f, 0x55, 0x70, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xd7, 0x02, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x35,
	0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x3b, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4d, 0x0a, 0x0f, 0x53,
	0x79, 0x6e, 0x63, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x6f,
	0x77, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x22, 0xd6, 0x03, 0x0a, 0x10, 0x53,
	0x79, 0x6e, 0x63, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50,
	0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x41, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x75, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x75,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x1a, 0x38, 0x0a, 0x0a, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3b, 0x0a, 0x0d, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x6b, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0xb8, 0x01, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x75, 0x73,
	0x68, 0x12, 0x16, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x75,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04, 0x50, 0x75, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x50, 0x75, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x42, 0x6c, 0x61, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x3b, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_actor_sync_proto_rawDescOnce sync.Once
	file_proto_actor_sync_proto_rawDescData = file_proto_actor_sync_proto_rawDesc
)

func file_proto_actor_sync_proto_rawDescGZIP() []byte {
	file_proto_actor_sync_proto_rawDescOnce.Do(func() {
		file_proto_actor_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_actor_sync_proto_rawDescData)
	})
	return file_proto_actor_sync_proto_rawDescData
}

var file_proto_actor_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_actor_sync_proto_goTypes = []interface{}{
	(*SyncPushRequest)(nil),   // 0: actor.SyncPushRequest
	(*SyncPushResponse)(nil),  // 1: actor.SyncPushResponse
	(*SyncPullRequest)(nil),   // 2: actor.SyncPullRequest
	(*SyncPullResponse)(nil),  // 3: actor.SyncPullResponse
	(*SyncUploadRequest)(nil), // 4: actor.SyncUploadRequest
	nil,                       // 5: actor.SyncPushRequest.UpPropertyEntry
	nil,                       // 6: actor.SyncPushResponse.PropertyEntry
	nil,                       // 7: actor.SyncPushResponse.TaskEntry
	nil,                       // 8: actor.SyncPullResponse.AliasEntry
	nil,                       // 9: actor.SyncPullResponse.PropertyEntry
	nil,                       // 10: actor.SyncPullResponse.HealthyEntry
	(*DeviceEntity)(nil),      // 11: actor.DeviceEntity
	(*Status)(nil),            // 12: actor.Status
	(*BlankResponse)(nil),     // 13: actor.BlankResponse
}
var file_proto_actor_sync_proto_depIdxs = []int32{
	11, // 0: actor.SyncPushRequest.device:type_name -> actor.DeviceEntity
	5,  // 1: actor.SyncPushRequest.upProperty:type_name -> actor.SyncPushRequest.UpPropertyEntry
	12, // 2: actor.SyncPushResponse.status:type_name -> actor.Status
	6,  // 3: actor.SyncPushResponse.property:type_name -> actor.SyncPushResponse.PropertyEntry
	7,  // 4: actor.SyncPushResponse.task:type_name -> actor.SyncPushResponse.TaskEntry
	12, // 5: actor.SyncPullResponse.status:type_name -> actor.Status
	11, // 6: actor.SyncPullResponse.device:type_name -> actor.DeviceEntity
	8,  // 7: actor.SyncPullResponse.alias:type_name -> actor.SyncPullResponse.AliasEntry
	9,  // 8: actor.SyncPullResponse.property:type_name -> actor.SyncPullResponse.PropertyEntry
	10, // 9: actor.SyncPullResponse.healthy:type_name -> actor.SyncPullResponse.HealthyEntry
	0,  // 10: actor.Sync.Push:input_type -> actor.SyncPushRequest
	2,  // 11: actor.Sync.Pull:input_type -> actor.SyncPullRequest
	4,  // 12: actor.Sync.Upload:input_type -> actor.SyncUploadRequest
	1,  // 13: actor.Sync.Push:output_type -> actor.SyncPushResponse
	3,  // 14: actor.Sync.Pull:output_type -> actor.SyncPullResponse
	13, // 15: actor.Sync.Upload:output_type -> actor.BlankResponse
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_actor_sync_proto_init() }
func file_proto_actor_sync_proto_init() {
	if File_proto_actor_sync_proto != nil {
		return
	}
	file_proto_actor_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_actor_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPushRequest); i {
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
		file_proto_actor_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPushResponse); i {
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
		file_proto_actor_sync_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPullRequest); i {
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
		file_proto_actor_sync_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPullResponse); i {
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
		file_proto_actor_sync_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncUploadRequest); i {
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
			RawDescriptor: file_proto_actor_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_actor_sync_proto_goTypes,
		DependencyIndexes: file_proto_actor_sync_proto_depIdxs,
		MessageInfos:      file_proto_actor_sync_proto_msgTypes,
	}.Build()
	File_proto_actor_sync_proto = out.File
	file_proto_actor_sync_proto_rawDesc = nil
	file_proto_actor_sync_proto_goTypes = nil
	file_proto_actor_sync_proto_depIdxs = nil
}
