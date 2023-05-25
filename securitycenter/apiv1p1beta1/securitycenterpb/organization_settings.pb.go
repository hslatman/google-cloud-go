// Copyright 2021 Google LLC
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
// source: google/cloud/securitycenter/v1p1beta1/organization_settings.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The mode of inclusion when running Asset Discovery.
// Asset discovery can be limited by explicitly identifying projects to be
// included or excluded. If INCLUDE_ONLY is set, then only those projects
// within the organization and their children are discovered during asset
// discovery. If EXCLUDE is set, then projects that don't match those
// projects are discovered during asset discovery. If neither are set, then
// all projects within the organization are discovered during asset
// discovery.
type OrganizationSettings_AssetDiscoveryConfig_InclusionMode int32

const (
	// Unspecified. Setting the mode with this value will disable
	// inclusion/exclusion filtering for Asset Discovery.
	OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 0
	// Asset Discovery will capture only the resources within the projects
	// specified. All other resources will be ignored.
	OrganizationSettings_AssetDiscoveryConfig_INCLUDE_ONLY OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 1
	// Asset Discovery will ignore all resources under the projects specified.
	// All other resources will be retrieved.
	OrganizationSettings_AssetDiscoveryConfig_EXCLUDE OrganizationSettings_AssetDiscoveryConfig_InclusionMode = 2
)

// Enum value maps for OrganizationSettings_AssetDiscoveryConfig_InclusionMode.
var (
	OrganizationSettings_AssetDiscoveryConfig_InclusionMode_name = map[int32]string{
		0: "INCLUSION_MODE_UNSPECIFIED",
		1: "INCLUDE_ONLY",
		2: "EXCLUDE",
	}
	OrganizationSettings_AssetDiscoveryConfig_InclusionMode_value = map[string]int32{
		"INCLUSION_MODE_UNSPECIFIED": 0,
		"INCLUDE_ONLY":               1,
		"EXCLUDE":                    2,
	}
)

func (x OrganizationSettings_AssetDiscoveryConfig_InclusionMode) Enum() *OrganizationSettings_AssetDiscoveryConfig_InclusionMode {
	p := new(OrganizationSettings_AssetDiscoveryConfig_InclusionMode)
	*p = x
	return p
}

func (x OrganizationSettings_AssetDiscoveryConfig_InclusionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrganizationSettings_AssetDiscoveryConfig_InclusionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_enumTypes[0].Descriptor()
}

func (OrganizationSettings_AssetDiscoveryConfig_InclusionMode) Type() protoreflect.EnumType {
	return &file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_enumTypes[0]
}

func (x OrganizationSettings_AssetDiscoveryConfig_InclusionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrganizationSettings_AssetDiscoveryConfig_InclusionMode.Descriptor instead.
func (OrganizationSettings_AssetDiscoveryConfig_InclusionMode) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescGZIP(), []int{0, 0, 0}
}

// User specified settings that are attached to the Security Command
// Center organization.
type OrganizationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relative resource name of the settings. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/{organization_id}/organizationSettings".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A flag that indicates if Asset Discovery should be enabled. If the flag is
	// set to `true`, then discovery of assets will occur. If it is set to `false,
	// all historical assets will remain, but discovery of future assets will not
	// occur.
	EnableAssetDiscovery bool `protobuf:"varint,2,opt,name=enable_asset_discovery,json=enableAssetDiscovery,proto3" json:"enable_asset_discovery,omitempty"`
	// The configuration used for Asset Discovery runs.
	AssetDiscoveryConfig *OrganizationSettings_AssetDiscoveryConfig `protobuf:"bytes,3,opt,name=asset_discovery_config,json=assetDiscoveryConfig,proto3" json:"asset_discovery_config,omitempty"`
}

func (x *OrganizationSettings) Reset() {
	*x = OrganizationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationSettings) ProtoMessage() {}

func (x *OrganizationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationSettings.ProtoReflect.Descriptor instead.
func (*OrganizationSettings) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescGZIP(), []int{0}
}

func (x *OrganizationSettings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrganizationSettings) GetEnableAssetDiscovery() bool {
	if x != nil {
		return x.EnableAssetDiscovery
	}
	return false
}

func (x *OrganizationSettings) GetAssetDiscoveryConfig() *OrganizationSettings_AssetDiscoveryConfig {
	if x != nil {
		return x.AssetDiscoveryConfig
	}
	return nil
}

// The configuration used for Asset Discovery runs.
type OrganizationSettings_AssetDiscoveryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The project ids to use for filtering asset discovery.
	ProjectIds []string `protobuf:"bytes,1,rep,name=project_ids,json=projectIds,proto3" json:"project_ids,omitempty"`
	// The mode to use for filtering asset discovery.
	InclusionMode OrganizationSettings_AssetDiscoveryConfig_InclusionMode `protobuf:"varint,2,opt,name=inclusion_mode,json=inclusionMode,proto3,enum=google.cloud.securitycenter.v1p1beta1.OrganizationSettings_AssetDiscoveryConfig_InclusionMode" json:"inclusion_mode,omitempty"`
	// The folder ids to use for filtering asset discovery.
	// It consists of only digits, e.g., 756619654966.
	FolderIds []string `protobuf:"bytes,3,rep,name=folder_ids,json=folderIds,proto3" json:"folder_ids,omitempty"`
}

func (x *OrganizationSettings_AssetDiscoveryConfig) Reset() {
	*x = OrganizationSettings_AssetDiscoveryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrganizationSettings_AssetDiscoveryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrganizationSettings_AssetDiscoveryConfig) ProtoMessage() {}

func (x *OrganizationSettings_AssetDiscoveryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrganizationSettings_AssetDiscoveryConfig.ProtoReflect.Descriptor instead.
func (*OrganizationSettings_AssetDiscoveryConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *OrganizationSettings_AssetDiscoveryConfig) GetProjectIds() []string {
	if x != nil {
		return x.ProjectIds
	}
	return nil
}

func (x *OrganizationSettings_AssetDiscoveryConfig) GetInclusionMode() OrganizationSettings_AssetDiscoveryConfig_InclusionMode {
	if x != nil {
		return x.InclusionMode
	}
	return OrganizationSettings_AssetDiscoveryConfig_INCLUSION_MODE_UNSPECIFIED
}

func (x *OrganizationSettings_AssetDiscoveryConfig) GetFolderIds() []string {
	if x != nil {
		return x.FolderIds
	}
	return nil
}

var File_google_cloud_securitycenter_v1p1beta1_organization_settings_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x05, 0x0a, 0x14, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x86, 0x01, 0x0a, 0x16, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0xae, 0x02, 0x0a, 0x14, 0x41, 0x73, 0x73, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x85, 0x01, 0x0a, 0x0e,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x5e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0d, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x22, 0x4e, 0x0a, 0x0d, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x5f, 0x4f,
	0x4e, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x44, 0x45,
	0x10, 0x02, 0x3a, 0x6a, 0xea, 0x41, 0x67, 0x0a, 0x32, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x31, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0xfb,
	0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x50, 0x01, 0x5a, 0x51,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62,
	0x3b, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x62, 0xaa, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e,
	0x56, 0x31, 0x50, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xea, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescData = file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDesc
)

func file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDescData
}

var file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_goTypes = []interface{}{
	(OrganizationSettings_AssetDiscoveryConfig_InclusionMode)(0), // 0: google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig.InclusionMode
	(*OrganizationSettings)(nil),                                 // 1: google.cloud.securitycenter.v1p1beta1.OrganizationSettings
	(*OrganizationSettings_AssetDiscoveryConfig)(nil),            // 2: google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig
}
var file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_depIdxs = []int32{
	2, // 0: google.cloud.securitycenter.v1p1beta1.OrganizationSettings.asset_discovery_config:type_name -> google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig
	0, // 1: google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig.inclusion_mode:type_name -> google.cloud.securitycenter.v1p1beta1.OrganizationSettings.AssetDiscoveryConfig.InclusionMode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_init() }
func file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_init() {
	if File_google_cloud_securitycenter_v1p1beta1_organization_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationSettings); i {
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
		file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrganizationSettings_AssetDiscoveryConfig); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_depIdxs,
		EnumInfos:         file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_enumTypes,
		MessageInfos:      file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1p1beta1_organization_settings_proto = out.File
	file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_goTypes = nil
	file_google_cloud_securitycenter_v1p1beta1_organization_settings_proto_depIdxs = nil
}
