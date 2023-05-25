// Copyright 2020 Google LLC
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
// source: google/cloud/recommendationengine/v1beta1/common.proto

package recommendationenginepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FeatureMap represents extra features that customers want to include in the
// recommendation model for catalogs/user events as categorical/numerical
// features.
type FeatureMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Categorical features that can take on one of a limited number of possible
	// values. Some examples would be the brand/maker of a product, or country of
	// a customer.
	//
	// Feature names and values must be UTF-8 encoded strings.
	//
	// For example: `{ "colors": {"value": ["yellow", "green"]},
	//                 "sizes": {"value":["S", "M"]}`
	CategoricalFeatures map[string]*FeatureMap_StringList `protobuf:"bytes,1,rep,name=categorical_features,json=categoricalFeatures,proto3" json:"categorical_features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Numerical features. Some examples would be the height/weight of a product,
	// or age of a customer.
	//
	// Feature names must be UTF-8 encoded strings.
	//
	// For example: `{ "lengths_cm": {"value":[2.3, 15.4]},
	//                 "heights_cm": {"value":[8.1, 6.4]} }`
	NumericalFeatures map[string]*FeatureMap_FloatList `protobuf:"bytes,2,rep,name=numerical_features,json=numericalFeatures,proto3" json:"numerical_features,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FeatureMap) Reset() {
	*x = FeatureMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureMap) ProtoMessage() {}

func (x *FeatureMap) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureMap.ProtoReflect.Descriptor instead.
func (*FeatureMap) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureMap) GetCategoricalFeatures() map[string]*FeatureMap_StringList {
	if x != nil {
		return x.CategoricalFeatures
	}
	return nil
}

func (x *FeatureMap) GetNumericalFeatures() map[string]*FeatureMap_FloatList {
	if x != nil {
		return x.NumericalFeatures
	}
	return nil
}

// A list of string features.
type FeatureMap_StringList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String feature value with a length limit of 128 bytes.
	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *FeatureMap_StringList) Reset() {
	*x = FeatureMap_StringList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureMap_StringList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureMap_StringList) ProtoMessage() {}

func (x *FeatureMap_StringList) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureMap_StringList.ProtoReflect.Descriptor instead.
func (*FeatureMap_StringList) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FeatureMap_StringList) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

// A list of float features.
type FeatureMap_FloatList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Float feature value.
	Value []float32 `protobuf:"fixed32,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *FeatureMap_FloatList) Reset() {
	*x = FeatureMap_FloatList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureMap_FloatList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureMap_FloatList) ProtoMessage() {}

func (x *FeatureMap_FloatList) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureMap_FloatList.ProtoReflect.Descriptor instead.
func (*FeatureMap_FloatList) Descriptor() ([]byte, []int) {
	return file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescGZIP(), []int{0, 1}
}

func (x *FeatureMap_FloatList) GetValue() []float32 {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_google_cloud_recommendationengine_v1beta1_common_proto protoreflect.FileDescriptor

var file_google_cloud_recommendationengine_v1beta1_common_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x22, 0xe7, 0x04, 0x0a, 0x0a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d,
	0x61, 0x70, 0x12, 0x81, 0x01, 0x0a, 0x14, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x63, 0x61, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x13, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x7b, 0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x61, 0x6c, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x61, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x11, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x1a, 0x22, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x21, 0x0a, 0x09, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x88, 0x01, 0x0a, 0x18, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x56, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x85, 0x01, 0x0a, 0x16, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x61, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x55, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x61, 0x70, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xa3, 0x02,
	0x0a, 0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50,
	0x01, 0x5a, 0x61, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x05, 0x52, 0x45, 0x43, 0x41, 0x49, 0xaa, 0x02, 0x29, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x29, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x2c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescOnce sync.Once
	file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescData = file_google_cloud_recommendationengine_v1beta1_common_proto_rawDesc
)

func file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescGZIP() []byte {
	file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescOnce.Do(func() {
		file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescData)
	})
	return file_google_cloud_recommendationengine_v1beta1_common_proto_rawDescData
}

var file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_recommendationengine_v1beta1_common_proto_goTypes = []interface{}{
	(*FeatureMap)(nil),            // 0: google.cloud.recommendationengine.v1beta1.FeatureMap
	(*FeatureMap_StringList)(nil), // 1: google.cloud.recommendationengine.v1beta1.FeatureMap.StringList
	(*FeatureMap_FloatList)(nil),  // 2: google.cloud.recommendationengine.v1beta1.FeatureMap.FloatList
	nil,                           // 3: google.cloud.recommendationengine.v1beta1.FeatureMap.CategoricalFeaturesEntry
	nil,                           // 4: google.cloud.recommendationengine.v1beta1.FeatureMap.NumericalFeaturesEntry
}
var file_google_cloud_recommendationengine_v1beta1_common_proto_depIdxs = []int32{
	3, // 0: google.cloud.recommendationengine.v1beta1.FeatureMap.categorical_features:type_name -> google.cloud.recommendationengine.v1beta1.FeatureMap.CategoricalFeaturesEntry
	4, // 1: google.cloud.recommendationengine.v1beta1.FeatureMap.numerical_features:type_name -> google.cloud.recommendationengine.v1beta1.FeatureMap.NumericalFeaturesEntry
	1, // 2: google.cloud.recommendationengine.v1beta1.FeatureMap.CategoricalFeaturesEntry.value:type_name -> google.cloud.recommendationengine.v1beta1.FeatureMap.StringList
	2, // 3: google.cloud.recommendationengine.v1beta1.FeatureMap.NumericalFeaturesEntry.value:type_name -> google.cloud.recommendationengine.v1beta1.FeatureMap.FloatList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_recommendationengine_v1beta1_common_proto_init() }
func file_google_cloud_recommendationengine_v1beta1_common_proto_init() {
	if File_google_cloud_recommendationengine_v1beta1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureMap); i {
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
		file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureMap_StringList); i {
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
		file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureMap_FloatList); i {
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
			RawDescriptor: file_google_cloud_recommendationengine_v1beta1_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_recommendationengine_v1beta1_common_proto_goTypes,
		DependencyIndexes: file_google_cloud_recommendationengine_v1beta1_common_proto_depIdxs,
		MessageInfos:      file_google_cloud_recommendationengine_v1beta1_common_proto_msgTypes,
	}.Build()
	File_google_cloud_recommendationengine_v1beta1_common_proto = out.File
	file_google_cloud_recommendationengine_v1beta1_common_proto_rawDesc = nil
	file_google_cloud_recommendationengine_v1beta1_common_proto_goTypes = nil
	file_google_cloud_recommendationengine_v1beta1_common_proto_depIdxs = nil
}
