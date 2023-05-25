// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: google/cloud/optimization/v1/async_model.proto

package optimizationpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Data formats for input and output files.
type DataFormat int32

const (
	// Default value.
	DataFormat_DATA_FORMAT_UNSPECIFIED DataFormat = 0
	// Input data in json format.
	DataFormat_JSON DataFormat = 1
	// Input data in string format.
	DataFormat_STRING DataFormat = 2
)

// Enum value maps for DataFormat.
var (
	DataFormat_name = map[int32]string{
		0: "DATA_FORMAT_UNSPECIFIED",
		1: "JSON",
		2: "STRING",
	}
	DataFormat_value = map[string]int32{
		"DATA_FORMAT_UNSPECIFIED": 0,
		"JSON":                    1,
		"STRING":                  2,
	}
)

func (x DataFormat) Enum() *DataFormat {
	p := new(DataFormat)
	*p = x
	return p
}

func (x DataFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_optimization_v1_async_model_proto_enumTypes[0].Descriptor()
}

func (DataFormat) Type() protoreflect.EnumType {
	return &file_google_cloud_optimization_v1_async_model_proto_enumTypes[0]
}

func (x DataFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataFormat.Descriptor instead.
func (DataFormat) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_optimization_v1_async_model_proto_rawDescGZIP(), []int{0}
}

// Possible states of the operation.
type AsyncModelMetadata_State int32

const (
	// The default value. This value is used if the state is omitted.
	AsyncModelMetadata_STATE_UNSPECIFIED AsyncModelMetadata_State = 0
	// Request is being processed.
	AsyncModelMetadata_RUNNING AsyncModelMetadata_State = 1
	// The operation completed successfully.
	AsyncModelMetadata_SUCCEEDED AsyncModelMetadata_State = 2
	// The operation was cancelled.
	AsyncModelMetadata_CANCELLED AsyncModelMetadata_State = 3
	// The operation has failed.
	AsyncModelMetadata_FAILED AsyncModelMetadata_State = 4
)

// Enum value maps for AsyncModelMetadata_State.
var (
	AsyncModelMetadata_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "RUNNING",
		2: "SUCCEEDED",
		3: "CANCELLED",
		4: "FAILED",
	}
	AsyncModelMetadata_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"RUNNING":           1,
		"SUCCEEDED":         2,
		"CANCELLED":         3,
		"FAILED":            4,
	}
)

func (x AsyncModelMetadata_State) Enum() *AsyncModelMetadata_State {
	p := new(AsyncModelMetadata_State)
	*p = x
	return p
}

func (x AsyncModelMetadata_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AsyncModelMetadata_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_optimization_v1_async_model_proto_enumTypes[1].Descriptor()
}

func (AsyncModelMetadata_State) Type() protoreflect.EnumType {
	return &file_google_cloud_optimization_v1_async_model_proto_enumTypes[1]
}

func (x AsyncModelMetadata_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AsyncModelMetadata_State.Descriptor instead.
func (AsyncModelMetadata_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_optimization_v1_async_model_proto_rawDescGZIP(), []int{4, 0}
}

// The desired input location information.
type InputConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location of the input model in cloud storage.
	// Required.
	//
	// Types that are assignable to Source:
	//	*InputConfig_GcsSource
	Source isInputConfig_Source `protobuf_oneof:"source"`
	// The input data format that used to store the model in Cloud Storage.
	DataFormat DataFormat `protobuf:"varint,2,opt,name=data_format,json=dataFormat,proto3,enum=google.cloud.optimization.v1.DataFormat" json:"data_format,omitempty"`
}

func (x *InputConfig) Reset() {
	*x = InputConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputConfig) ProtoMessage() {}

func (x *InputConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputConfig.ProtoReflect.Descriptor instead.
func (*InputConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_optimization_v1_async_model_proto_rawDescGZIP(), []int{0}
}

func (m *InputConfig) GetSource() isInputConfig_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *InputConfig) GetGcsSource() *GcsSource {
	if x, ok := x.GetSource().(*InputConfig_GcsSource); ok {
		return x.GcsSource
	}
	return nil
}

func (x *InputConfig) GetDataFormat() DataFormat {
	if x != nil {
		return x.DataFormat
	}
	return DataFormat_DATA_FORMAT_UNSPECIFIED
}

type isInputConfig_Source interface {
	isInputConfig_Source()
}

type InputConfig_GcsSource struct {
	// The Google Cloud Storage location to read the input from. This must be a
	// single file.
	GcsSource *GcsSource `protobuf:"bytes,1,opt,name=gcs_source,json=gcsSource,proto3,oneof"`
}

func (*InputConfig_GcsSource) isInputConfig_Source() {}

// The desired output location.
type OutputConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location of the output result in cloud storage.
	// Required.
	//
	// Types that are assignable to Destination:
	//	*OutputConfig_GcsDestination
	Destination isOutputConfig_Destination `protobuf_oneof:"destination"`
	// The output data format that used to store the results in Cloud Storage.
	DataFormat DataFormat `protobuf:"varint,2,opt,name=data_format,json=dataFormat,proto3,enum=google.cloud.optimization.v1.DataFormat" json:"data_format,omitempty"`
}

func (x *OutputConfig) Reset() {
	*x = OutputConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputConfig) ProtoMessage() {}

func (x *OutputConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputConfig.ProtoReflect.Descriptor instead.
func (*OutputConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_optimization_v1_async_model_proto_rawDescGZIP(), []int{1}
}

func (m *OutputConfig) GetDestination() isOutputConfig_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *OutputConfig) GetGcsDestination() *GcsDestination {
	if x, ok := x.GetDestination().(*OutputConfig_GcsDestination); ok {
		return x.GcsDestination
	}
	return nil
}

func (x *OutputConfig) GetDataFormat() DataFormat {
	if x != nil {
		return x.DataFormat
	}
	return DataFormat_DATA_FORMAT_UNSPECIFIED
}

type isOutputConfig_Destination interface {
	isOutputConfig_Destination()
}

type OutputConfig_GcsDestination struct {
	// The Google Cloud Storage location to write the output to.
	GcsDestination *GcsDestination `protobuf:"bytes,1,opt,name=gcs_destination,json=gcsDestination,proto3,oneof"`
}

func (*OutputConfig_GcsDestination) isOutputConfig_Destination() {}

// The Google Cloud Storage location where the input file will be read from.
type GcsSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. URI of the Google Cloud Storage location.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *GcsSource) Reset() {
	*x = GcsSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsSource) ProtoMessage() {}

func (x *GcsSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsSource.ProtoReflect.Descriptor instead.
func (*GcsSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_optimization_v1_async_model_proto_rawDescGZIP(), []int{2}
}

func (x *GcsSource) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

// The Google Cloud Storage location where the output file will be written to.
type GcsDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. URI of the Google Cloud Storage location.
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *GcsDestination) Reset() {
	*x = GcsDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsDestination) ProtoMessage() {}

func (x *GcsDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsDestination.ProtoReflect.Descriptor instead.
func (*GcsDestination) Descriptor() ([]byte, []int) {
	return file_google_cloud_optimization_v1_async_model_proto_rawDescGZIP(), []int{3}
}

func (x *GcsDestination) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

// The long running operation metadata for async model related methods.
type AsyncModelMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The state of the current operation.
	State AsyncModelMetadata_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.optimization.v1.AsyncModelMetadata_State" json:"state,omitempty"`
	// A message providing more details about the current state of the operation.
	// For example, the error message if the operation is failed.
	StateMessage string `protobuf:"bytes,2,opt,name=state_message,json=stateMessage,proto3" json:"state_message,omitempty"`
	// The creation time of the operation.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last update time of the operation.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *AsyncModelMetadata) Reset() {
	*x = AsyncModelMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncModelMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncModelMetadata) ProtoMessage() {}

func (x *AsyncModelMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_optimization_v1_async_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncModelMetadata.ProtoReflect.Descriptor instead.
func (*AsyncModelMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_optimization_v1_async_model_proto_rawDescGZIP(), []int{4}
}

func (x *AsyncModelMetadata) GetState() AsyncModelMetadata_State {
	if x != nil {
		return x.State
	}
	return AsyncModelMetadata_STATE_UNSPECIFIED
}

func (x *AsyncModelMetadata) GetStateMessage() string {
	if x != nil {
		return x.StateMessage
	}
	return ""
}

func (x *AsyncModelMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *AsyncModelMetadata) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_google_cloud_optimization_v1_async_model_proto protoreflect.FileDescriptor

var file_google_cloud_optimization_v1_async_model_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xac, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x48, 0x0a, 0x0a, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52,
	0x09, 0x67, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0xc1, 0x01, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x57, 0x0a, 0x0f, 0x67, 0x63, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x73, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x67, 0x63, 0x73, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x09, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x03, 0x75, 0x72, 0x69, 0x22, 0x27, 0x0a, 0x0e, 0x47, 0x63, 0x73, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x03, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x03, 0x75, 0x72, 0x69,
	0x22, 0xd8, 0x02, 0x0a, 0x12, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a,
	0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x3f, 0x0a, 0x0a, 0x44,
	0x61, 0x74, 0x61, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x41, 0x54,
	0x41, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x7b, 0x0a, 0x20,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x42, 0x0f, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x44, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x3b, 0x6f, 0x70, 0x74, 0x69, 0x6d,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_optimization_v1_async_model_proto_rawDescOnce sync.Once
	file_google_cloud_optimization_v1_async_model_proto_rawDescData = file_google_cloud_optimization_v1_async_model_proto_rawDesc
)

func file_google_cloud_optimization_v1_async_model_proto_rawDescGZIP() []byte {
	file_google_cloud_optimization_v1_async_model_proto_rawDescOnce.Do(func() {
		file_google_cloud_optimization_v1_async_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_optimization_v1_async_model_proto_rawDescData)
	})
	return file_google_cloud_optimization_v1_async_model_proto_rawDescData
}

var file_google_cloud_optimization_v1_async_model_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_optimization_v1_async_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_optimization_v1_async_model_proto_goTypes = []interface{}{
	(DataFormat)(0),               // 0: google.cloud.optimization.v1.DataFormat
	(AsyncModelMetadata_State)(0), // 1: google.cloud.optimization.v1.AsyncModelMetadata.State
	(*InputConfig)(nil),           // 2: google.cloud.optimization.v1.InputConfig
	(*OutputConfig)(nil),          // 3: google.cloud.optimization.v1.OutputConfig
	(*GcsSource)(nil),             // 4: google.cloud.optimization.v1.GcsSource
	(*GcsDestination)(nil),        // 5: google.cloud.optimization.v1.GcsDestination
	(*AsyncModelMetadata)(nil),    // 6: google.cloud.optimization.v1.AsyncModelMetadata
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_google_cloud_optimization_v1_async_model_proto_depIdxs = []int32{
	4, // 0: google.cloud.optimization.v1.InputConfig.gcs_source:type_name -> google.cloud.optimization.v1.GcsSource
	0, // 1: google.cloud.optimization.v1.InputConfig.data_format:type_name -> google.cloud.optimization.v1.DataFormat
	5, // 2: google.cloud.optimization.v1.OutputConfig.gcs_destination:type_name -> google.cloud.optimization.v1.GcsDestination
	0, // 3: google.cloud.optimization.v1.OutputConfig.data_format:type_name -> google.cloud.optimization.v1.DataFormat
	1, // 4: google.cloud.optimization.v1.AsyncModelMetadata.state:type_name -> google.cloud.optimization.v1.AsyncModelMetadata.State
	7, // 5: google.cloud.optimization.v1.AsyncModelMetadata.create_time:type_name -> google.protobuf.Timestamp
	7, // 6: google.cloud.optimization.v1.AsyncModelMetadata.update_time:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_optimization_v1_async_model_proto_init() }
func file_google_cloud_optimization_v1_async_model_proto_init() {
	if File_google_cloud_optimization_v1_async_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_optimization_v1_async_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputConfig); i {
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
		file_google_cloud_optimization_v1_async_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputConfig); i {
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
		file_google_cloud_optimization_v1_async_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsSource); i {
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
		file_google_cloud_optimization_v1_async_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsDestination); i {
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
		file_google_cloud_optimization_v1_async_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncModelMetadata); i {
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
	file_google_cloud_optimization_v1_async_model_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*InputConfig_GcsSource)(nil),
	}
	file_google_cloud_optimization_v1_async_model_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*OutputConfig_GcsDestination)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_optimization_v1_async_model_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_optimization_v1_async_model_proto_goTypes,
		DependencyIndexes: file_google_cloud_optimization_v1_async_model_proto_depIdxs,
		EnumInfos:         file_google_cloud_optimization_v1_async_model_proto_enumTypes,
		MessageInfos:      file_google_cloud_optimization_v1_async_model_proto_msgTypes,
	}.Build()
	File_google_cloud_optimization_v1_async_model_proto = out.File
	file_google_cloud_optimization_v1_async_model_proto_rawDesc = nil
	file_google_cloud_optimization_v1_async_model_proto_goTypes = nil
	file_google_cloud_optimization_v1_async_model_proto_depIdxs = nil
}
