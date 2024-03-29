// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: wearable/device/wearable_service.proto

package wearable

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

type BeatsPerMinutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BeatsPerMinutesRequest) Reset() {
	*x = BeatsPerMinutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_device_wearable_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeatsPerMinutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeatsPerMinutesRequest) ProtoMessage() {}

func (x *BeatsPerMinutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_device_wearable_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeatsPerMinutesRequest.ProtoReflect.Descriptor instead.
func (*BeatsPerMinutesRequest) Descriptor() ([]byte, []int) {
	return file_wearable_device_wearable_service_proto_rawDescGZIP(), []int{0}
}

func (x *BeatsPerMinutesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BeatsPerMinuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Minute uint32 `protobuf:"varint,2,opt,name=minute,proto3" json:"minute,omitempty"`
}

func (x *BeatsPerMinuteResponse) Reset() {
	*x = BeatsPerMinuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_device_wearable_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeatsPerMinuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeatsPerMinuteResponse) ProtoMessage() {}

func (x *BeatsPerMinuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_device_wearable_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeatsPerMinuteResponse.ProtoReflect.Descriptor instead.
func (*BeatsPerMinuteResponse) Descriptor() ([]byte, []int) {
	return file_wearable_device_wearable_service_proto_rawDescGZIP(), []int{1}
}

func (x *BeatsPerMinuteResponse) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *BeatsPerMinuteResponse) GetMinute() uint32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

type CalculateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Value  uint32 `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Minute uint32 `protobuf:"varint,3,opt,name=Minute,proto3" json:"Minute,omitempty"`
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_device_wearable_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_device_wearable_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_wearable_device_wearable_service_proto_rawDescGZIP(), []int{2}
}

func (x *CalculateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CalculateRequest) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CalculateRequest) GetMinute() uint32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float32 `protobuf:"fixed32,1,opt,name=Average,proto3" json:"Average,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_device_wearable_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_device_wearable_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_wearable_device_wearable_service_proto_rawDescGZIP(), []int{3}
}

func (x *CalculateResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

var File_wearable_device_wearable_service_proto protoreflect.FileDescriptor

var file_wearable_device_wearable_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x42, 0x65, 0x61,
	0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x16, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x22, 0x50, 0x0a, 0x10, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x22, 0x2d, 0x0a,
	0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x32, 0xe1, 0x01, 0x0a,
	0x0f, 0x57, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x66, 0x0a, 0x0e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x12, 0x27, 0x2e, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x77, 0x65,
	0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x65,
	0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x66, 0x0a, 0x17, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x18, 0x5a, 0x16, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wearable_device_wearable_service_proto_rawDescOnce sync.Once
	file_wearable_device_wearable_service_proto_rawDescData = file_wearable_device_wearable_service_proto_rawDesc
)

func file_wearable_device_wearable_service_proto_rawDescGZIP() []byte {
	file_wearable_device_wearable_service_proto_rawDescOnce.Do(func() {
		file_wearable_device_wearable_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_wearable_device_wearable_service_proto_rawDescData)
	})
	return file_wearable_device_wearable_service_proto_rawDescData
}

var file_wearable_device_wearable_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wearable_device_wearable_service_proto_goTypes = []interface{}{
	(*BeatsPerMinutesRequest)(nil), // 0: wearable.device.BeatsPerMinutesRequest
	(*BeatsPerMinuteResponse)(nil), // 1: wearable.device.BeatsPerMinuteResponse
	(*CalculateRequest)(nil),       // 2: wearable.device.CalculateRequest
	(*CalculateResponse)(nil),      // 3: wearable.device.CalculateResponse
}
var file_wearable_device_wearable_service_proto_depIdxs = []int32{
	0, // 0: wearable.device.WearableService.BeatsPerMinute:input_type -> wearable.device.BeatsPerMinutesRequest
	2, // 1: wearable.device.WearableService.CalculateBeatsPerMinute:input_type -> wearable.device.CalculateRequest
	1, // 2: wearable.device.WearableService.BeatsPerMinute:output_type -> wearable.device.BeatsPerMinuteResponse
	3, // 3: wearable.device.WearableService.CalculateBeatsPerMinute:output_type -> wearable.device.CalculateResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wearable_device_wearable_service_proto_init() }
func file_wearable_device_wearable_service_proto_init() {
	if File_wearable_device_wearable_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wearable_device_wearable_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeatsPerMinutesRequest); i {
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
		file_wearable_device_wearable_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeatsPerMinuteResponse); i {
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
		file_wearable_device_wearable_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRequest); i {
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
		file_wearable_device_wearable_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse); i {
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
			RawDescriptor: file_wearable_device_wearable_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wearable_device_wearable_service_proto_goTypes,
		DependencyIndexes: file_wearable_device_wearable_service_proto_depIdxs,
		MessageInfos:      file_wearable_device_wearable_service_proto_msgTypes,
	}.Build()
	File_wearable_device_wearable_service_proto = out.File
	file_wearable_device_wearable_service_proto_rawDesc = nil
	file_wearable_device_wearable_service_proto_goTypes = nil
	file_wearable_device_wearable_service_proto_depIdxs = nil
}
