// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: platform/industrial_data.proto

package platform

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

type GetControllerDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetControllerDataRequest) Reset() {
	*x = GetControllerDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_industrial_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetControllerDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetControllerDataRequest) ProtoMessage() {}

func (x *GetControllerDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_industrial_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetControllerDataRequest.ProtoReflect.Descriptor instead.
func (*GetControllerDataRequest) Descriptor() ([]byte, []int) {
	return file_platform_industrial_data_proto_rawDescGZIP(), []int{0}
}

type GetControllerDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ControllerMode      string `protobuf:"bytes,1,opt,name=controller_mode,json=controllerMode,proto3" json:"controller_mode,omitempty"`
	SwitchStatus        string `protobuf:"bytes,2,opt,name=switch_status,json=switchStatus,proto3" json:"switch_status,omitempty"`
	CommunicationStatus string `protobuf:"bytes,3,opt,name=communication_status,json=communicationStatus,proto3" json:"communication_status,omitempty"`
}

func (x *GetControllerDataReply) Reset() {
	*x = GetControllerDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_industrial_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetControllerDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetControllerDataReply) ProtoMessage() {}

func (x *GetControllerDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_platform_industrial_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetControllerDataReply.ProtoReflect.Descriptor instead.
func (*GetControllerDataReply) Descriptor() ([]byte, []int) {
	return file_platform_industrial_data_proto_rawDescGZIP(), []int{1}
}

func (x *GetControllerDataReply) GetControllerMode() string {
	if x != nil {
		return x.ControllerMode
	}
	return ""
}

func (x *GetControllerDataReply) GetSwitchStatus() string {
	if x != nil {
		return x.SwitchStatus
	}
	return ""
}

func (x *GetControllerDataReply) GetCommunicationStatus() string {
	if x != nil {
		return x.CommunicationStatus
	}
	return ""
}

type GetAllDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllDataRequest) Reset() {
	*x = GetAllDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_industrial_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDataRequest) ProtoMessage() {}

func (x *GetAllDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_industrial_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDataRequest.ProtoReflect.Descriptor instead.
func (*GetAllDataRequest) Descriptor() ([]byte, []int) {
	return file_platform_industrial_data_proto_rawDescGZIP(), []int{2}
}

type GetAllDataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllDataReply) Reset() {
	*x = GetAllDataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_industrial_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDataReply) ProtoMessage() {}

func (x *GetAllDataReply) ProtoReflect() protoreflect.Message {
	mi := &file_platform_industrial_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDataReply.ProtoReflect.Descriptor instead.
func (*GetAllDataReply) Descriptor() ([]byte, []int) {
	return file_platform_industrial_data_proto_rawDescGZIP(), []int{3}
}

var File_platform_industrial_data_proto protoreflect.FileDescriptor

var file_platform_industrial_data_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73,
	0x74, 0x72, 0x69, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xce,
	0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x6b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x63, 0x33, 0x68, 0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x69, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x4f,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x63, 0x33, 0x68,
	0x2f, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x61, 0x6c, 0x6c, 0x42,
	0x13, 0x5a, 0x11, 0x63, 0x33, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_industrial_data_proto_rawDescOnce sync.Once
	file_platform_industrial_data_proto_rawDescData = file_platform_industrial_data_proto_rawDesc
)

func file_platform_industrial_data_proto_rawDescGZIP() []byte {
	file_platform_industrial_data_proto_rawDescOnce.Do(func() {
		file_platform_industrial_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_industrial_data_proto_rawDescData)
	})
	return file_platform_industrial_data_proto_rawDescData
}

var file_platform_industrial_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_platform_industrial_data_proto_goTypes = []interface{}{
	(*GetControllerDataRequest)(nil), // 0: GetControllerDataRequest
	(*GetControllerDataReply)(nil),   // 1: GetControllerDataReply
	(*GetAllDataRequest)(nil),        // 2: GetAllDataRequest
	(*GetAllDataReply)(nil),          // 3: GetAllDataReply
}
var file_platform_industrial_data_proto_depIdxs = []int32{
	0, // 0: IndustrialData.GetControllerData:input_type -> GetControllerDataRequest
	2, // 1: IndustrialData.GetAllData:input_type -> GetAllDataRequest
	1, // 2: IndustrialData.GetControllerData:output_type -> GetControllerDataReply
	3, // 3: IndustrialData.GetAllData:output_type -> GetAllDataReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_platform_industrial_data_proto_init() }
func file_platform_industrial_data_proto_init() {
	if File_platform_industrial_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_industrial_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetControllerDataRequest); i {
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
		file_platform_industrial_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetControllerDataReply); i {
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
		file_platform_industrial_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDataRequest); i {
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
		file_platform_industrial_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDataReply); i {
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
			RawDescriptor: file_platform_industrial_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_industrial_data_proto_goTypes,
		DependencyIndexes: file_platform_industrial_data_proto_depIdxs,
		MessageInfos:      file_platform_industrial_data_proto_msgTypes,
	}.Build()
	File_platform_industrial_data_proto = out.File
	file_platform_industrial_data_proto_rawDesc = nil
	file_platform_industrial_data_proto_goTypes = nil
	file_platform_industrial_data_proto_depIdxs = nil
}
