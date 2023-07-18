// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: product_net/product_network.proto

package product_net

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

type SwitchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SwitchRequest) Reset() {
	*x = SwitchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchRequest) ProtoMessage() {}

func (x *SwitchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchRequest.ProtoReflect.Descriptor instead.
func (*SwitchRequest) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{0}
}

type SwitchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DC_402_AB          int32  `protobuf:"varint,1,opt,name=DC_402_AB,json=DC402AB,proto3" json:"DC_402_AB,omitempty"`
	DC_402Master       int32  `protobuf:"varint,2,opt,name=DC_402_master,json=DC402Master,proto3" json:"DC_402_master,omitempty"`
	DC_402Manual       int32  `protobuf:"varint,3,opt,name=DC_402_manual,json=DC402Manual,proto3" json:"DC_402_manual,omitempty"`
	Export_MAPD        string `protobuf:"bytes,4,opt,name=export_MAPD,json=exportMAPD,proto3" json:"export_MAPD,omitempty"`
	ResetButton        int32  `protobuf:"varint,5,opt,name=reset_button,json=resetButton,proto3" json:"reset_button,omitempty"`
	DesiccantButton    int32  `protobuf:"varint,6,opt,name=desiccant_button,json=desiccantButton,proto3" json:"desiccant_button,omitempty"`
	CommunicationAlarm int32  `protobuf:"varint,7,opt,name=communication_alarm,json=communicationAlarm,proto3" json:"communication_alarm,omitempty"`
}

func (x *SwitchReply) Reset() {
	*x = SwitchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwitchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwitchReply) ProtoMessage() {}

func (x *SwitchReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwitchReply.ProtoReflect.Descriptor instead.
func (*SwitchReply) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{1}
}

func (x *SwitchReply) GetDC_402_AB() int32 {
	if x != nil {
		return x.DC_402_AB
	}
	return 0
}

func (x *SwitchReply) GetDC_402Master() int32 {
	if x != nil {
		return x.DC_402Master
	}
	return 0
}

func (x *SwitchReply) GetDC_402Manual() int32 {
	if x != nil {
		return x.DC_402Manual
	}
	return 0
}

func (x *SwitchReply) GetExport_MAPD() string {
	if x != nil {
		return x.Export_MAPD
	}
	return ""
}

func (x *SwitchReply) GetResetButton() int32 {
	if x != nil {
		return x.ResetButton
	}
	return 0
}

func (x *SwitchReply) GetDesiccantButton() int32 {
	if x != nil {
		return x.DesiccantButton
	}
	return 0
}

func (x *SwitchReply) GetCommunicationAlarm() int32 {
	if x != nil {
		return x.CommunicationAlarm
	}
	return 0
}

type ControlVarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ControlVarRequest) Reset() {
	*x = ControlVarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlVarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlVarRequest) ProtoMessage() {}

func (x *ControlVarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlVarRequest.ProtoReflect.Descriptor instead.
func (*ControlVarRequest) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{2}
}

type ControlVarReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MAPDInputConcentration  string `protobuf:"bytes,1,opt,name=MAPD_input_concentration,json=MAPDInputConcentration,proto3" json:"MAPD_input_concentration,omitempty"`
	MAPDOutputConcentration string `protobuf:"bytes,2,opt,name=MAPD_output_concentration,json=MAPDOutputConcentration,proto3" json:"MAPD_output_concentration,omitempty"`
	C3InputConcentration    string `protobuf:"bytes,3,opt,name=C3_input_concentration,json=C3InputConcentration,proto3" json:"C3_input_concentration,omitempty"`
	C3OutputConcentration   string `protobuf:"bytes,4,opt,name=C3_output_concentration,json=C3OutputConcentration,proto3" json:"C3_output_concentration,omitempty"`
}

func (x *ControlVarReply) Reset() {
	*x = ControlVarReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlVarReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlVarReply) ProtoMessage() {}

func (x *ControlVarReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlVarReply.ProtoReflect.Descriptor instead.
func (*ControlVarReply) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{3}
}

func (x *ControlVarReply) GetMAPDInputConcentration() string {
	if x != nil {
		return x.MAPDInputConcentration
	}
	return ""
}

func (x *ControlVarReply) GetMAPDOutputConcentration() string {
	if x != nil {
		return x.MAPDOutputConcentration
	}
	return ""
}

func (x *ControlVarReply) GetC3InputConcentration() string {
	if x != nil {
		return x.C3InputConcentration
	}
	return ""
}

func (x *ControlVarReply) GetC3OutputConcentration() string {
	if x != nil {
		return x.C3OutputConcentration
	}
	return ""
}

type ControlledVarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ControlledVarRequest) Reset() {
	*x = ControlledVarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlledVarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlledVarRequest) ProtoMessage() {}

func (x *ControlledVarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlledVarRequest.ProtoReflect.Descriptor instead.
func (*ControlledVarRequest) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{4}
}

type ControlledVarReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpperH2Flow  string `protobuf:"bytes,1,opt,name=upper_h2_flow,json=upperH2Flow,proto3" json:"upper_h2_flow,omitempty"`
	LowerH2Flow  string `protobuf:"bytes,2,opt,name=lower_h2_flow,json=lowerH2Flow,proto3" json:"lower_h2_flow,omitempty"`
	TotalH2Flow  string `protobuf:"bytes,3,opt,name=total_h2_flow,json=totalH2Flow,proto3" json:"total_h2_flow,omitempty"`
	UpperHaRatio string `protobuf:"bytes,4,opt,name=upper_ha_ratio,json=upperHaRatio,proto3" json:"upper_ha_ratio,omitempty"`
}

func (x *ControlledVarReply) Reset() {
	*x = ControlledVarReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlledVarReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlledVarReply) ProtoMessage() {}

func (x *ControlledVarReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlledVarReply.ProtoReflect.Descriptor instead.
func (*ControlledVarReply) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{5}
}

func (x *ControlledVarReply) GetUpperH2Flow() string {
	if x != nil {
		return x.UpperH2Flow
	}
	return ""
}

func (x *ControlledVarReply) GetLowerH2Flow() string {
	if x != nil {
		return x.LowerH2Flow
	}
	return ""
}

func (x *ControlledVarReply) GetTotalH2Flow() string {
	if x != nil {
		return x.TotalH2Flow
	}
	return ""
}

func (x *ControlledVarReply) GetUpperHaRatio() string {
	if x != nil {
		return x.UpperHaRatio
	}
	return ""
}

type ConfoundingVarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfoundingVarRequest) Reset() {
	*x = ConfoundingVarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfoundingVarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfoundingVarRequest) ProtoMessage() {}

func (x *ConfoundingVarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfoundingVarRequest.ProtoReflect.Descriptor instead.
func (*ConfoundingVarRequest) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{6}
}

type ConfoundingVarReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ABInputTemperature  string `protobuf:"bytes,1,opt,name=AB_input_temperature,json=ABInputTemperature,proto3" json:"AB_input_temperature,omitempty"`
	ABOutputTemperature string `protobuf:"bytes,2,opt,name=AB_output_temperature,json=ABOutputTemperature,proto3" json:"AB_output_temperature,omitempty"`
	FA_409Pressure      string `protobuf:"bytes,3,opt,name=FA_409_pressure,json=FA409Pressure,proto3" json:"FA_409_pressure,omitempty"`
}

func (x *ConfoundingVarReply) Reset() {
	*x = ConfoundingVarReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfoundingVarReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfoundingVarReply) ProtoMessage() {}

func (x *ConfoundingVarReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfoundingVarReply.ProtoReflect.Descriptor instead.
func (*ConfoundingVarReply) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{7}
}

func (x *ConfoundingVarReply) GetABInputTemperature() string {
	if x != nil {
		return x.ABInputTemperature
	}
	return ""
}

func (x *ConfoundingVarReply) GetABOutputTemperature() string {
	if x != nil {
		return x.ABOutputTemperature
	}
	return ""
}

func (x *ConfoundingVarReply) GetFA_409Pressure() string {
	if x != nil {
		return x.FA_409Pressure
	}
	return ""
}

type CatalystRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CatalystRequest) Reset() {
	*x = CatalystRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatalystRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalystRequest) ProtoMessage() {}

func (x *CatalystRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalystRequest.ProtoReflect.Descriptor instead.
func (*CatalystRequest) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{8}
}

type CatalystReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ch3C          string `protobuf:"bytes,1,opt,name=ch3c,proto3" json:"ch3c,omitempty"`
	C3H4          string `protobuf:"bytes,2,opt,name=c3h4,proto3" json:"c3h4,omitempty"`
	C3H6          string `protobuf:"bytes,3,opt,name=c3h6,proto3" json:"c3h6,omitempty"`
	C3H8          string `protobuf:"bytes,4,opt,name=c3h8,proto3" json:"c3h8,omitempty"`
	C3ManualInput string `protobuf:"bytes,5,opt,name=c3_manual_input,json=c3ManualInput,proto3" json:"c3_manual_input,omitempty"`
}

func (x *CatalystReply) Reset() {
	*x = CatalystReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_net_product_network_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CatalystReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CatalystReply) ProtoMessage() {}

func (x *CatalystReply) ProtoReflect() protoreflect.Message {
	mi := &file_product_net_product_network_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CatalystReply.ProtoReflect.Descriptor instead.
func (*CatalystReply) Descriptor() ([]byte, []int) {
	return file_product_net_product_network_proto_rawDescGZIP(), []int{9}
}

func (x *CatalystReply) GetCh3C() string {
	if x != nil {
		return x.Ch3C
	}
	return ""
}

func (x *CatalystReply) GetC3H4() string {
	if x != nil {
		return x.C3H4
	}
	return ""
}

func (x *CatalystReply) GetC3H6() string {
	if x != nil {
		return x.C3H6
	}
	return ""
}

func (x *CatalystReply) GetC3H8() string {
	if x != nil {
		return x.C3H8
	}
	return ""
}

func (x *CatalystReply) GetC3ManualInput() string {
	if x != nil {
		return x.C3ManualInput
	}
	return ""
}

var File_product_net_product_network_proto protoreflect.FileDescriptor

var file_product_net_product_network_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x6e, 0x65, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x91, 0x02, 0x0a, 0x0b, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x09, 0x44, 0x43, 0x5f, 0x34, 0x30, 0x32, 0x5f, 0x41, 0x42,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x44, 0x43, 0x34, 0x30, 0x32, 0x41, 0x42, 0x12,
	0x22, 0x0a, 0x0d, 0x44, 0x43, 0x5f, 0x34, 0x30, 0x32, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x44, 0x43, 0x34, 0x30, 0x32, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x44, 0x43, 0x5f, 0x34, 0x30, 0x32, 0x5f, 0x6d, 0x61,
	0x6e, 0x75, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x44, 0x43, 0x34, 0x30,
	0x32, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x4d, 0x41, 0x50, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x4d, 0x41, 0x50, 0x44, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x5f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64,
	0x65, 0x73, 0x69, 0x63, 0x63, 0x61, 0x6e, 0x74, 0x5f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x69, 0x63, 0x63, 0x61, 0x6e, 0x74,
	0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xf5, 0x01, 0x0a,
	0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x56, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x38, 0x0a, 0x18, 0x4d, 0x41, 0x50, 0x44, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x16, 0x4d, 0x41, 0x50, 0x44, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e,
	0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x19, 0x4d, 0x41,
	0x50, 0x44, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x4d,
	0x41, 0x50, 0x44, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x16, 0x43, 0x33, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x43, 0x33, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17,
	0x43, 0x33, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x43,
	0x33, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x64, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa6, 0x01, 0x0a,
	0x12, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x56, 0x61, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x68, 0x32, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x48, 0x32, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x22, 0x0a, 0x0d, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x5f, 0x68, 0x32, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x48, 0x32, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x22, 0x0a, 0x0d, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x68, 0x32, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x48, 0x32, 0x46, 0x6c, 0x6f, 0x77, 0x12,
	0x24, 0x0a, 0x0e, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x5f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x70, 0x70, 0x65, 0x72, 0x48, 0x61,
	0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa3,
	0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x41, 0x42, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x41, 0x42, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x65, 0x6d,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x41, 0x42, 0x5f, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x41, 0x42, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x46, 0x41, 0x5f, 0x34, 0x30, 0x39, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x46, 0x41, 0x34, 0x30, 0x39, 0x50, 0x72, 0x65, 0x73,
	0x73, 0x75, 0x72, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x79, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x0d, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x79, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x68, 0x33,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x68, 0x33, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x33, 0x68, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x33, 0x68,
	0x34, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x33, 0x68, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x33, 0x68, 0x36, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x33, 0x68, 0x38, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x33, 0x68, 0x38, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x33, 0x5f,
	0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x33, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x32, 0xa0, 0x05, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x65, 0x74,
	0x12, 0x72, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6e, 0x65, 0x74, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x6e, 0x65, 0x74, 0x2e, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x63, 0x33, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x73, 0x77,
	0x69, 0x74, 0x63, 0x68, 0x12, 0x81, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x56, 0x61, 0x72, 0x73, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x56, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x63, 0x33, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x76, 0x61, 0x72, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x56, 0x61, 0x72, 0x73, 0x12, 0x25,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x64, 0x56, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x64, 0x56, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x26, 0x12, 0x24, 0x2f, 0x63, 0x33, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x64, 0x2d, 0x76, 0x61, 0x72, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72, 0x73, 0x12,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65,
	0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x63, 0x33, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x63, 0x61, 0x72, 0x73, 0x12, 0x76, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x79, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x2e, 0x43, 0x61,
	0x74, 0x61, 0x6c, 0x79, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x2e,
	0x43, 0x61, 0x74, 0x61, 0x6c, 0x79, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x25, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x63, 0x33, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x79, 0x73, 0x74, 0x42, 0x34, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x50, 0x01, 0x5a, 0x1f, 0x63, 0x33, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x3b, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_product_net_product_network_proto_rawDescOnce sync.Once
	file_product_net_product_network_proto_rawDescData = file_product_net_product_network_proto_rawDesc
)

func file_product_net_product_network_proto_rawDescGZIP() []byte {
	file_product_net_product_network_proto_rawDescOnce.Do(func() {
		file_product_net_product_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_net_product_network_proto_rawDescData)
	})
	return file_product_net_product_network_proto_rawDescData
}

var file_product_net_product_network_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_product_net_product_network_proto_goTypes = []interface{}{
	(*SwitchRequest)(nil),         // 0: api.product_net.SwitchRequest
	(*SwitchReply)(nil),           // 1: api.product_net.SwitchReply
	(*ControlVarRequest)(nil),     // 2: api.product_net.ControlVarRequest
	(*ControlVarReply)(nil),       // 3: api.product_net.ControlVarReply
	(*ControlledVarRequest)(nil),  // 4: api.product_net.ControlledVarRequest
	(*ControlledVarReply)(nil),    // 5: api.product_net.ControlledVarReply
	(*ConfoundingVarRequest)(nil), // 6: api.product_net.ConfoundingVarRequest
	(*ConfoundingVarReply)(nil),   // 7: api.product_net.ConfoundingVarReply
	(*CatalystRequest)(nil),       // 8: api.product_net.CatalystRequest
	(*CatalystReply)(nil),         // 9: api.product_net.CatalystReply
}
var file_product_net_product_network_proto_depIdxs = []int32{
	0, // 0: api.product_net.ProductNet.GetSwitchInfo:input_type -> api.product_net.SwitchRequest
	2, // 1: api.product_net.ProductNet.GetControlVars:input_type -> api.product_net.ControlVarRequest
	4, // 2: api.product_net.ProductNet.GetControlledVars:input_type -> api.product_net.ControlledVarRequest
	6, // 3: api.product_net.ProductNet.GetConfoundingVars:input_type -> api.product_net.ConfoundingVarRequest
	8, // 4: api.product_net.ProductNet.GetCatalyst:input_type -> api.product_net.CatalystRequest
	1, // 5: api.product_net.ProductNet.GetSwitchInfo:output_type -> api.product_net.SwitchReply
	3, // 6: api.product_net.ProductNet.GetControlVars:output_type -> api.product_net.ControlVarReply
	5, // 7: api.product_net.ProductNet.GetControlledVars:output_type -> api.product_net.ControlledVarReply
	7, // 8: api.product_net.ProductNet.GetConfoundingVars:output_type -> api.product_net.ConfoundingVarReply
	9, // 9: api.product_net.ProductNet.GetCatalyst:output_type -> api.product_net.CatalystReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_net_product_network_proto_init() }
func file_product_net_product_network_proto_init() {
	if File_product_net_product_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_net_product_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchRequest); i {
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
		file_product_net_product_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwitchReply); i {
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
		file_product_net_product_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlVarRequest); i {
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
		file_product_net_product_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlVarReply); i {
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
		file_product_net_product_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlledVarRequest); i {
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
		file_product_net_product_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlledVarReply); i {
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
		file_product_net_product_network_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfoundingVarRequest); i {
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
		file_product_net_product_network_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfoundingVarReply); i {
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
		file_product_net_product_network_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatalystRequest); i {
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
		file_product_net_product_network_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CatalystReply); i {
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
			RawDescriptor: file_product_net_product_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_net_product_network_proto_goTypes,
		DependencyIndexes: file_product_net_product_network_proto_depIdxs,
		MessageInfos:      file_product_net_product_network_proto_msgTypes,
	}.Build()
	File_product_net_product_network_proto = out.File
	file_product_net_product_network_proto_rawDesc = nil
	file_product_net_product_network_proto_goTypes = nil
	file_product_net_product_network_proto_depIdxs = nil
}