// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/actor/device.proto

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

// 列举的回复
type DeviceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 总数
	Device []*DeviceEntity `protobuf:"bytes,3,rep,name=device,proto3" json:"device,omitempty"` // 设备实体
}

func (x *DeviceListResponse) Reset() {
	*x = DeviceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceListResponse) ProtoMessage() {}

func (x *DeviceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceListResponse.ProtoReflect.Descriptor instead.
func (*DeviceListResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_device_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DeviceListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DeviceListResponse) GetDevice() []*DeviceEntity {
	if x != nil {
		return x.Device
	}
	return nil
}

// 搜索的请求
type DeviceSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset          int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`                  // 偏移值
	Count           int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`                    // 数量
	SerialNumber    string `protobuf:"bytes,3,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`       // 序列号
	Name            string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                       // 设备名称
	OperatingSystem string `protobuf:"bytes,5,opt,name=operatingSystem,proto3" json:"operatingSystem,omitempty"` // 操作系统
	SystemVersion   string `protobuf:"bytes,6,opt,name=systemVersion,proto3" json:"systemVersion,omitempty"`     // 系统版本
	Shape           string `protobuf:"bytes,7,opt,name=shape,proto3" json:"shape,omitempty"`                     // 形态
}

func (x *DeviceSearchRequest) Reset() {
	*x = DeviceSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceSearchRequest) ProtoMessage() {}

func (x *DeviceSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceSearchRequest.ProtoReflect.Descriptor instead.
func (*DeviceSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_actor_device_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceSearchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DeviceSearchRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *DeviceSearchRequest) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *DeviceSearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeviceSearchRequest) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *DeviceSearchRequest) GetSystemVersion() string {
	if x != nil {
		return x.SystemVersion
	}
	return ""
}

func (x *DeviceSearchRequest) GetShape() string {
	if x != nil {
		return x.Shape
	}
	return ""
}

// 搜索的回复
type DeviceSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  int64           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 总数
	Device []*DeviceEntity `protobuf:"bytes,3,rep,name=device,proto3" json:"device,omitempty"` // 设备实体
}

func (x *DeviceSearchResponse) Reset() {
	*x = DeviceSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_actor_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceSearchResponse) ProtoMessage() {}

func (x *DeviceSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_actor_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceSearchResponse.ProtoReflect.Descriptor instead.
func (*DeviceSearchResponse) Descriptor() ([]byte, []int) {
	return file_proto_actor_device_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceSearchResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DeviceSearchResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DeviceSearchResponse) GetDevice() []*DeviceEntity {
	if x != nil {
		return x.Device
	}
	return nil
}

var File_proto_actor_device_proto protoreflect.FileDescriptor

var file_proto_actor_device_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x12, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0xe1, 0x01, 0x0a, 0x13,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65, 0x22,
	0x80, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x32, 0x86, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x1a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x3b, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_actor_device_proto_rawDescOnce sync.Once
	file_proto_actor_device_proto_rawDescData = file_proto_actor_device_proto_rawDesc
)

func file_proto_actor_device_proto_rawDescGZIP() []byte {
	file_proto_actor_device_proto_rawDescOnce.Do(func() {
		file_proto_actor_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_actor_device_proto_rawDescData)
	})
	return file_proto_actor_device_proto_rawDescData
}

var file_proto_actor_device_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_actor_device_proto_goTypes = []interface{}{
	(*DeviceListResponse)(nil),   // 0: actor.DeviceListResponse
	(*DeviceSearchRequest)(nil),  // 1: actor.DeviceSearchRequest
	(*DeviceSearchResponse)(nil), // 2: actor.DeviceSearchResponse
	(*Status)(nil),               // 3: actor.Status
	(*DeviceEntity)(nil),         // 4: actor.DeviceEntity
	(*ListRequest)(nil),          // 5: actor.ListRequest
}
var file_proto_actor_device_proto_depIdxs = []int32{
	3, // 0: actor.DeviceListResponse.status:type_name -> actor.Status
	4, // 1: actor.DeviceListResponse.device:type_name -> actor.DeviceEntity
	3, // 2: actor.DeviceSearchResponse.status:type_name -> actor.Status
	4, // 3: actor.DeviceSearchResponse.device:type_name -> actor.DeviceEntity
	5, // 4: actor.Device.List:input_type -> actor.ListRequest
	1, // 5: actor.Device.Search:input_type -> actor.DeviceSearchRequest
	0, // 6: actor.Device.List:output_type -> actor.DeviceListResponse
	2, // 7: actor.Device.Search:output_type -> actor.DeviceSearchResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_actor_device_proto_init() }
func file_proto_actor_device_proto_init() {
	if File_proto_actor_device_proto != nil {
		return
	}
	file_proto_actor_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_actor_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceListResponse); i {
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
		file_proto_actor_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceSearchRequest); i {
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
		file_proto_actor_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceSearchResponse); i {
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
			RawDescriptor: file_proto_actor_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_actor_device_proto_goTypes,
		DependencyIndexes: file_proto_actor_device_proto_depIdxs,
		MessageInfos:      file_proto_actor_device_proto_msgTypes,
	}.Build()
	File_proto_actor_device_proto = out.File
	file_proto_actor_device_proto_rawDesc = nil
	file_proto_actor_device_proto_goTypes = nil
	file_proto_actor_device_proto_depIdxs = nil
}
