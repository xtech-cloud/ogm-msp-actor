// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/actor/shared.proto

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

// 状态
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 状态信息
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_actor_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// 空白的请求
type BlankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlankRequest) Reset() {
	*x = BlankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankRequest) ProtoMessage() {}

func (x *BlankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankRequest.ProtoReflect.Descriptor instead.
func (*BlankRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_shared_proto_rawDescGZIP(), []int{1}
}

// 空白的回复
type BlankResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 请求状态
}

func (x *BlankResponse) Reset() {
	*x = BlankResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankResponse) ProtoMessage() {}

func (x *BlankResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankResponse.ProtoReflect.Descriptor instead.
func (*BlankResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_shared_proto_rawDescGZIP(), []int{2}
}

func (x *BlankResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 列举的请求
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移
	Count  int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_shared_proto_rawDescGZIP(), []int{3}
}

func (x *ListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 域实体
type DomainEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // 域的名称
}

func (x *DomainEntity) Reset() {
	*x = DomainEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainEntity) ProtoMessage() {}

func (x *DomainEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainEntity.ProtoReflect.Descriptor instead.
func (*DomainEntity) Descriptor() ([]byte, []int) {
	return file_proto_actor_shared_proto_rawDescGZIP(), []int{4}
}

func (x *DomainEntity) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *DomainEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 设备实体
type DeviceEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerialNumber     string            `protobuf:"bytes,1,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`                                                                                 // 序列号
	Name             string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                                                                                 // 设备名称
	OperatingSystem  string            `protobuf:"bytes,10,opt,name=operatingSystem,proto3" json:"operatingSystem,omitempty"`                                                                          // 操作系统
	SystemVersion    string            `protobuf:"bytes,11,opt,name=systemVersion,proto3" json:"systemVersion,omitempty"`                                                                              // 系统版本
	Shape            string            `protobuf:"bytes,20,opt,name=shape,proto3" json:"shape,omitempty"`                                                                                              // 形态
	Battery          int32             `protobuf:"varint,21,opt,name=battery,proto3" json:"battery,omitempty"`                                                                                         // 电量
	Volume           int32             `protobuf:"varint,22,opt,name=volume,proto3" json:"volume,omitempty"`                                                                                           // 音量
	Brightness       int32             `protobuf:"varint,23,opt,name=brightness,proto3" json:"brightness,omitempty"`                                                                                   // 亮度
	Storage          string            `protobuf:"bytes,30,opt,name=storage,proto3" json:"storage,omitempty"`                                                                                          // 存储
	StorageBlocks    int64             `protobuf:"varint,31,opt,name=storageBlocks,proto3" json:"storageBlocks,omitempty"`                                                                             // 数据目录所在的存储的容量
	StorageAvailable int64             `protobuf:"varint,32,opt,name=storageAvailable,proto3" json:"storageAvailable,omitempty"`                                                                       // 数据目录所在的存储的剩余
	Network          string            `protobuf:"bytes,40,opt,name=network,proto3" json:"network,omitempty"`                                                                                          // 网络
	NetworkStrength  int32             `protobuf:"varint,41,opt,name=networkStrength,proto3" json:"networkStrength,omitempty"`                                                                         // 网络信号强度
	Program          map[string]string `protobuf:"bytes,101,rep,name=program,proto3" json:"program,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 程序信息<程序名，程序版本>
}

func (x *DeviceEntity) Reset() {
	*x = DeviceEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_shared_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceEntity) ProtoMessage() {}

func (x *DeviceEntity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_shared_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceEntity.ProtoReflect.Descriptor instead.
func (*DeviceEntity) Descriptor() ([]byte, []int) {
	return file_proto_actor_shared_proto_rawDescGZIP(), []int{5}
}

func (x *DeviceEntity) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *DeviceEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceEntity) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *DeviceEntity) GetSystemVersion() string {
	if x != nil {
		return x.SystemVersion
	}
	return ""
}

func (x *DeviceEntity) GetShape() string {
	if x != nil {
		return x.Shape
	}
	return ""
}

func (x *DeviceEntity) GetBattery() int32 {
	if x != nil {
		return x.Battery
	}
	return 0
}

func (x *DeviceEntity) GetVolume() int32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *DeviceEntity) GetBrightness() int32 {
	if x != nil {
		return x.Brightness
	}
	return 0
}

func (x *DeviceEntity) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *DeviceEntity) GetStorageBlocks() int64 {
	if x != nil {
		return x.StorageBlocks
	}
	return 0
}

func (x *DeviceEntity) GetStorageAvailable() int64 {
	if x != nil {
		return x.StorageAvailable
	}
	return 0
}

func (x *DeviceEntity) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *DeviceEntity) GetNetworkStrength() int32 {
	if x != nil {
		return x.NetworkStrength
	}
	return 0
}

func (x *DeviceEntity) GetProgram() map[string]string {
	if x != nil {
		return x.Program
	}
	return nil
}

var File_proto_actor_shared_proto protoreflect.FileDescriptor

var file_proto_actor_shared_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x22, 0x36, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x42, 0x6c, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x0d, 0x42, 0x6c, 0x61,
	0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x3b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x36,
	0x0a, 0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa6, 0x04, 0x0a, 0x0c, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x72, 0x69,
	0x67, 0x68, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x28, 0x0a,
	0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x29, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x53,
	0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x3a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x61, 0x6d, 0x18, 0x65, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x61, 0x6d, 0x1a, 0x3a, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x3b, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_actor_shared_proto_rawDescOnce sync.Once
	file_proto_actor_shared_proto_rawDescData = file_proto_actor_shared_proto_rawDesc
)

func file_proto_actor_shared_proto_rawDescGZIP() []byte {
	file_proto_actor_shared_proto_rawDescOnce.Do(func() {
		file_proto_actor_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_actor_shared_proto_rawDescData)
	})
	return file_proto_actor_shared_proto_rawDescData
}

var file_proto_actor_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_actor_shared_proto_goTypes = []interface{}{
	(*Status)(nil),        // 0: actor.Status
	(*BlankRequest)(nil),  // 1: actor.BlankRequest
	(*BlankResponse)(nil), // 2: actor.BlankResponse
	(*ListRequest)(nil),   // 3: actor.ListRequest
	(*DomainEntity)(nil),  // 4: actor.DomainEntity
	(*DeviceEntity)(nil),  // 5: actor.DeviceEntity
	nil,                   // 6: actor.DeviceEntity.ProgramEntry
}
var file_proto_actor_shared_proto_depIdxs = []int32{
	0, // 0: actor.BlankResponse.status:type_name -> actor.Status
	6, // 1: actor.DeviceEntity.program:type_name -> actor.DeviceEntity.ProgramEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_actor_shared_proto_init() }
func file_proto_actor_shared_proto_init() {
	if File_proto_actor_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_actor_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_proto_actor_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankRequest); i {
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
		file_proto_actor_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankResponse); i {
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
		file_proto_actor_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_proto_actor_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainEntity); i {
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
		file_proto_actor_shared_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceEntity); i {
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
			RawDescriptor: file_proto_actor_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_actor_shared_proto_goTypes,
		DependencyIndexes: file_proto_actor_shared_proto_depIdxs,
		MessageInfos:      file_proto_actor_shared_proto_msgTypes,
	}.Build()
	File_proto_actor_shared_proto = out.File
	file_proto_actor_shared_proto_rawDesc = nil
	file_proto_actor_shared_proto_goTypes = nil
	file_proto_actor_shared_proto_depIdxs = nil
}
