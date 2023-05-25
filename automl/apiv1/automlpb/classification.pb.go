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
// source: google/cloud/automl/v1/classification.proto

package automlpb

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

// Type of the classification problem.
type ClassificationType int32

const (
	// An un-set value of this enum.
	ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED ClassificationType = 0
	// At most one label is allowed per example.
	ClassificationType_MULTICLASS ClassificationType = 1
	// Multiple labels are allowed for one example.
	ClassificationType_MULTILABEL ClassificationType = 2
)

// Enum value maps for ClassificationType.
var (
	ClassificationType_name = map[int32]string{
		0: "CLASSIFICATION_TYPE_UNSPECIFIED",
		1: "MULTICLASS",
		2: "MULTILABEL",
	}
	ClassificationType_value = map[string]int32{
		"CLASSIFICATION_TYPE_UNSPECIFIED": 0,
		"MULTICLASS":                      1,
		"MULTILABEL":                      2,
	}
)

func (x ClassificationType) Enum() *ClassificationType {
	p := new(ClassificationType)
	*p = x
	return p
}

func (x ClassificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClassificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_automl_v1_classification_proto_enumTypes[0].Descriptor()
}

func (ClassificationType) Type() protoreflect.EnumType {
	return &file_google_cloud_automl_v1_classification_proto_enumTypes[0]
}

func (x ClassificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClassificationType.Descriptor instead.
func (ClassificationType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_classification_proto_rawDescGZIP(), []int{0}
}

// Contains annotation details specific to classification.
type ClassificationAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. A confidence estimate between 0.0 and 1.0. A higher value
	// means greater confidence that the annotation is positive. If a user
	// approves an annotation as negative or positive, the score value remains
	// unchanged. If a user creates an annotation, the score is 0 for negative or
	// 1 for positive.
	Score float32 `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *ClassificationAnnotation) Reset() {
	*x = ClassificationAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationAnnotation) ProtoMessage() {}

func (x *ClassificationAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationAnnotation.ProtoReflect.Descriptor instead.
func (*ClassificationAnnotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_classification_proto_rawDescGZIP(), []int{0}
}

func (x *ClassificationAnnotation) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

// Model evaluation metrics for classification problems.
// Note: For Video Classification this metrics only describe quality of the
// Video Classification predictions of "segment_classification" type.
type ClassificationEvaluationMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The Area Under Precision-Recall Curve metric. Micro-averaged
	// for the overall evaluation.
	AuPrc float32 `protobuf:"fixed32,1,opt,name=au_prc,json=auPrc,proto3" json:"au_prc,omitempty"`
	// Output only. The Area Under Receiver Operating Characteristic curve metric.
	// Micro-averaged for the overall evaluation.
	AuRoc float32 `protobuf:"fixed32,6,opt,name=au_roc,json=auRoc,proto3" json:"au_roc,omitempty"`
	// Output only. The Log Loss metric.
	LogLoss float32 `protobuf:"fixed32,7,opt,name=log_loss,json=logLoss,proto3" json:"log_loss,omitempty"`
	// Output only. Metrics for each confidence_threshold in
	// 0.00,0.05,0.10,...,0.95,0.96,0.97,0.98,0.99 and
	// position_threshold = INT32_MAX_VALUE.
	// ROC and precision-recall curves, and other aggregated metrics are derived
	// from them. The confidence metrics entries may also be supplied for
	// additional values of position_threshold, but from these no aggregated
	// metrics are computed.
	ConfidenceMetricsEntry []*ClassificationEvaluationMetrics_ConfidenceMetricsEntry `protobuf:"bytes,3,rep,name=confidence_metrics_entry,json=confidenceMetricsEntry,proto3" json:"confidence_metrics_entry,omitempty"`
	// Output only. Confusion matrix of the evaluation.
	// Only set for MULTICLASS classification problems where number
	// of labels is no more than 10.
	// Only set for model level evaluation, not for evaluation per label.
	ConfusionMatrix *ClassificationEvaluationMetrics_ConfusionMatrix `protobuf:"bytes,4,opt,name=confusion_matrix,json=confusionMatrix,proto3" json:"confusion_matrix,omitempty"`
	// Output only. The annotation spec ids used for this evaluation.
	AnnotationSpecId []string `protobuf:"bytes,5,rep,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
}

func (x *ClassificationEvaluationMetrics) Reset() {
	*x = ClassificationEvaluationMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationEvaluationMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationEvaluationMetrics) ProtoMessage() {}

func (x *ClassificationEvaluationMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationEvaluationMetrics.ProtoReflect.Descriptor instead.
func (*ClassificationEvaluationMetrics) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_classification_proto_rawDescGZIP(), []int{1}
}

func (x *ClassificationEvaluationMetrics) GetAuPrc() float32 {
	if x != nil {
		return x.AuPrc
	}
	return 0
}

func (x *ClassificationEvaluationMetrics) GetAuRoc() float32 {
	if x != nil {
		return x.AuRoc
	}
	return 0
}

func (x *ClassificationEvaluationMetrics) GetLogLoss() float32 {
	if x != nil {
		return x.LogLoss
	}
	return 0
}

func (x *ClassificationEvaluationMetrics) GetConfidenceMetricsEntry() []*ClassificationEvaluationMetrics_ConfidenceMetricsEntry {
	if x != nil {
		return x.ConfidenceMetricsEntry
	}
	return nil
}

func (x *ClassificationEvaluationMetrics) GetConfusionMatrix() *ClassificationEvaluationMetrics_ConfusionMatrix {
	if x != nil {
		return x.ConfusionMatrix
	}
	return nil
}

func (x *ClassificationEvaluationMetrics) GetAnnotationSpecId() []string {
	if x != nil {
		return x.AnnotationSpecId
	}
	return nil
}

// Metrics for a single confidence threshold.
type ClassificationEvaluationMetrics_ConfidenceMetricsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Metrics are computed with an assumption that the model
	// never returns predictions with score lower than this value.
	ConfidenceThreshold float32 `protobuf:"fixed32,1,opt,name=confidence_threshold,json=confidenceThreshold,proto3" json:"confidence_threshold,omitempty"`
	// Output only. Metrics are computed with an assumption that the model
	// always returns at most this many predictions (ordered by their score,
	// descendingly), but they all still need to meet the confidence_threshold.
	PositionThreshold int32 `protobuf:"varint,14,opt,name=position_threshold,json=positionThreshold,proto3" json:"position_threshold,omitempty"`
	// Output only. Recall (True Positive Rate) for the given confidence
	// threshold.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. Precision for the given confidence threshold.
	Precision float32 `protobuf:"fixed32,3,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. False Positive Rate for the given confidence threshold.
	FalsePositiveRate float32 `protobuf:"fixed32,8,opt,name=false_positive_rate,json=falsePositiveRate,proto3" json:"false_positive_rate,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,4,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	// Output only. The Recall (True Positive Rate) when only considering the
	// label that has the highest prediction score and not below the confidence
	// threshold for each example.
	RecallAt1 float32 `protobuf:"fixed32,5,opt,name=recall_at1,json=recallAt1,proto3" json:"recall_at1,omitempty"`
	// Output only. The precision when only considering the label that has the
	// highest prediction score and not below the confidence threshold for each
	// example.
	PrecisionAt1 float32 `protobuf:"fixed32,6,opt,name=precision_at1,json=precisionAt1,proto3" json:"precision_at1,omitempty"`
	// Output only. The False Positive Rate when only considering the label that
	// has the highest prediction score and not below the confidence threshold
	// for each example.
	FalsePositiveRateAt1 float32 `protobuf:"fixed32,9,opt,name=false_positive_rate_at1,json=falsePositiveRateAt1,proto3" json:"false_positive_rate_at1,omitempty"`
	// Output only. The harmonic mean of [recall_at1][google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.recall_at1] and [precision_at1][google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry.precision_at1].
	F1ScoreAt1 float32 `protobuf:"fixed32,7,opt,name=f1_score_at1,json=f1ScoreAt1,proto3" json:"f1_score_at1,omitempty"`
	// Output only. The number of model created labels that match a ground truth
	// label.
	TruePositiveCount int64 `protobuf:"varint,10,opt,name=true_positive_count,json=truePositiveCount,proto3" json:"true_positive_count,omitempty"`
	// Output only. The number of model created labels that do not match a
	// ground truth label.
	FalsePositiveCount int64 `protobuf:"varint,11,opt,name=false_positive_count,json=falsePositiveCount,proto3" json:"false_positive_count,omitempty"`
	// Output only. The number of ground truth labels that are not matched
	// by a model created label.
	FalseNegativeCount int64 `protobuf:"varint,12,opt,name=false_negative_count,json=falseNegativeCount,proto3" json:"false_negative_count,omitempty"`
	// Output only. The number of labels that were not created by the model,
	// but if they would, they would not match a ground truth label.
	TrueNegativeCount int64 `protobuf:"varint,13,opt,name=true_negative_count,json=trueNegativeCount,proto3" json:"true_negative_count,omitempty"`
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) Reset() {
	*x = ClassificationEvaluationMetrics_ConfidenceMetricsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationEvaluationMetrics_ConfidenceMetricsEntry) ProtoMessage() {}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationEvaluationMetrics_ConfidenceMetricsEntry.ProtoReflect.Descriptor instead.
func (*ClassificationEvaluationMetrics_ConfidenceMetricsEntry) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_classification_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetConfidenceThreshold() float32 {
	if x != nil {
		return x.ConfidenceThreshold
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetPositionThreshold() int32 {
	if x != nil {
		return x.PositionThreshold
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetRecall() float32 {
	if x != nil {
		return x.Recall
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetPrecision() float32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetFalsePositiveRate() float32 {
	if x != nil {
		return x.FalsePositiveRate
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetF1Score() float32 {
	if x != nil {
		return x.F1Score
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetRecallAt1() float32 {
	if x != nil {
		return x.RecallAt1
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetPrecisionAt1() float32 {
	if x != nil {
		return x.PrecisionAt1
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetFalsePositiveRateAt1() float32 {
	if x != nil {
		return x.FalsePositiveRateAt1
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetF1ScoreAt1() float32 {
	if x != nil {
		return x.F1ScoreAt1
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetTruePositiveCount() int64 {
	if x != nil {
		return x.TruePositiveCount
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetFalsePositiveCount() int64 {
	if x != nil {
		return x.FalsePositiveCount
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetFalseNegativeCount() int64 {
	if x != nil {
		return x.FalseNegativeCount
	}
	return 0
}

func (x *ClassificationEvaluationMetrics_ConfidenceMetricsEntry) GetTrueNegativeCount() int64 {
	if x != nil {
		return x.TrueNegativeCount
	}
	return 0
}

// Confusion matrix of the model running the classification.
type ClassificationEvaluationMetrics_ConfusionMatrix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. IDs of the annotation specs used in the confusion matrix.
	// For Tables CLASSIFICATION
	// [prediction_type][google.cloud.automl.v1p1beta.TablesModelMetadata.prediction_type]
	// only list of [annotation_spec_display_name-s][] is populated.
	AnnotationSpecId []string `protobuf:"bytes,1,rep,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only. Display name of the annotation specs used in the confusion
	// matrix, as they were at the moment of the evaluation. For Tables
	// CLASSIFICATION
	// [prediction_type-s][google.cloud.automl.v1p1beta.TablesModelMetadata.prediction_type],
	// distinct values of the target column at the moment of the model
	// evaluation are populated here.
	DisplayName []string `protobuf:"bytes,3,rep,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Output only. Rows in the confusion matrix. The number of rows is equal to
	// the size of `annotation_spec_id`.
	// `row[i].example_count[j]` is the number of examples that have ground
	// truth of the `annotation_spec_id[i]` and are predicted as
	// `annotation_spec_id[j]` by the model being evaluated.
	Row []*ClassificationEvaluationMetrics_ConfusionMatrix_Row `protobuf:"bytes,2,rep,name=row,proto3" json:"row,omitempty"`
}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix) Reset() {
	*x = ClassificationEvaluationMetrics_ConfusionMatrix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationEvaluationMetrics_ConfusionMatrix) ProtoMessage() {}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationEvaluationMetrics_ConfusionMatrix.ProtoReflect.Descriptor instead.
func (*ClassificationEvaluationMetrics_ConfusionMatrix) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_classification_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix) GetAnnotationSpecId() []string {
	if x != nil {
		return x.AnnotationSpecId
	}
	return nil
}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix) GetDisplayName() []string {
	if x != nil {
		return x.DisplayName
	}
	return nil
}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix) GetRow() []*ClassificationEvaluationMetrics_ConfusionMatrix_Row {
	if x != nil {
		return x.Row
	}
	return nil
}

// Output only. A row in the confusion matrix.
type ClassificationEvaluationMetrics_ConfusionMatrix_Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Value of the specific cell in the confusion matrix.
	// The number of values each row has (i.e. the length of the row) is equal
	// to the length of the `annotation_spec_id` field or, if that one is not
	// populated, length of the [display_name][google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.display_name] field.
	ExampleCount []int32 `protobuf:"varint,1,rep,packed,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix_Row) Reset() {
	*x = ClassificationEvaluationMetrics_ConfusionMatrix_Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix_Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationEvaluationMetrics_ConfusionMatrix_Row) ProtoMessage() {}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix_Row) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_automl_v1_classification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationEvaluationMetrics_ConfusionMatrix_Row.ProtoReflect.Descriptor instead.
func (*ClassificationEvaluationMetrics_ConfusionMatrix_Row) Descriptor() ([]byte, []int) {
	return file_google_cloud_automl_v1_classification_proto_rawDescGZIP(), []int{1, 1, 0}
}

func (x *ClassificationEvaluationMetrics_ConfusionMatrix_Row) GetExampleCount() []int32 {
	if x != nil {
		return x.ExampleCount
	}
	return nil
}

var File_google_cloud_automl_v1_classification_proto protoreflect.FileDescriptor

var file_google_cloud_automl_v1_classification_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x18, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xe6, 0x09, 0x0a, 0x1f, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x75, 0x5f, 0x70, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x61, 0x75, 0x50,
	0x72, 0x63, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x75, 0x5f, 0x72, 0x6f, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x61, 0x75, 0x52, 0x6f, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x6c, 0x6f, 0x67,
	0x4c, 0x6f, 0x73, 0x73, 0x12, 0x88, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x72, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x72,
	0x69, 0x78, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74,
	0x72, 0x69, 0x78, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x49,
	0x64, 0x1a, 0xdc, 0x04, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x14,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x13, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12,
	0x2d, 0x0a, 0x12, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x11, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x31, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x31, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x74, 0x31, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x09, 0x72, 0x65, 0x63, 0x61, 0x6c, 0x6c, 0x41, 0x74, 0x31, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x74, 0x31, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x41, 0x74, 0x31, 0x12, 0x35, 0x0a, 0x17, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x31, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x14, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x61, 0x74, 0x65, 0x41, 0x74, 0x31, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x31,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x66, 0x31, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x41, 0x74, 0x31, 0x12, 0x2e, 0x0a, 0x13,
	0x74, 0x72, 0x75, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x72, 0x75, 0x65, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14,
	0x66, 0x61, 0x6c, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x66, 0x61, 0x6c, 0x73,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30,
	0x0a, 0x14, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x2e, 0x0a, 0x13, 0x74, 0x72, 0x75, 0x65, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x74,
	0x72, 0x75, 0x65, 0x4e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0xed, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x5d, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x74, 0x72, 0x69, 0x78, 0x2e, 0x52, 0x6f, 0x77, 0x52,
	0x03, 0x72, 0x6f, 0x77, 0x1a, 0x2a, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0c, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x2a, 0x59, 0x0a, 0x12, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49,
	0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4d,
	0x55, 0x4c, 0x54, 0x49, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x02, 0x42, 0xb5, 0x01, 0x0a, 0x1a,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6c, 0x70, 0x62, 0x3b, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x6c, 0x70, 0x62, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x75,
	0x74, 0x6f, 0x4d, 0x6c, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x4d, 0x4c, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_automl_v1_classification_proto_rawDescOnce sync.Once
	file_google_cloud_automl_v1_classification_proto_rawDescData = file_google_cloud_automl_v1_classification_proto_rawDesc
)

func file_google_cloud_automl_v1_classification_proto_rawDescGZIP() []byte {
	file_google_cloud_automl_v1_classification_proto_rawDescOnce.Do(func() {
		file_google_cloud_automl_v1_classification_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_automl_v1_classification_proto_rawDescData)
	})
	return file_google_cloud_automl_v1_classification_proto_rawDescData
}

var file_google_cloud_automl_v1_classification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_automl_v1_classification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_automl_v1_classification_proto_goTypes = []interface{}{
	(ClassificationType)(0),                                        // 0: google.cloud.automl.v1.ClassificationType
	(*ClassificationAnnotation)(nil),                               // 1: google.cloud.automl.v1.ClassificationAnnotation
	(*ClassificationEvaluationMetrics)(nil),                        // 2: google.cloud.automl.v1.ClassificationEvaluationMetrics
	(*ClassificationEvaluationMetrics_ConfidenceMetricsEntry)(nil), // 3: google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry
	(*ClassificationEvaluationMetrics_ConfusionMatrix)(nil),        // 4: google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix
	(*ClassificationEvaluationMetrics_ConfusionMatrix_Row)(nil),    // 5: google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.Row
}
var file_google_cloud_automl_v1_classification_proto_depIdxs = []int32{
	3, // 0: google.cloud.automl.v1.ClassificationEvaluationMetrics.confidence_metrics_entry:type_name -> google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfidenceMetricsEntry
	4, // 1: google.cloud.automl.v1.ClassificationEvaluationMetrics.confusion_matrix:type_name -> google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix
	5, // 2: google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.row:type_name -> google.cloud.automl.v1.ClassificationEvaluationMetrics.ConfusionMatrix.Row
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_automl_v1_classification_proto_init() }
func file_google_cloud_automl_v1_classification_proto_init() {
	if File_google_cloud_automl_v1_classification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_automl_v1_classification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificationAnnotation); i {
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
		file_google_cloud_automl_v1_classification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificationEvaluationMetrics); i {
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
		file_google_cloud_automl_v1_classification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificationEvaluationMetrics_ConfidenceMetricsEntry); i {
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
		file_google_cloud_automl_v1_classification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificationEvaluationMetrics_ConfusionMatrix); i {
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
		file_google_cloud_automl_v1_classification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificationEvaluationMetrics_ConfusionMatrix_Row); i {
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
			RawDescriptor: file_google_cloud_automl_v1_classification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_automl_v1_classification_proto_goTypes,
		DependencyIndexes: file_google_cloud_automl_v1_classification_proto_depIdxs,
		EnumInfos:         file_google_cloud_automl_v1_classification_proto_enumTypes,
		MessageInfos:      file_google_cloud_automl_v1_classification_proto_msgTypes,
	}.Build()
	File_google_cloud_automl_v1_classification_proto = out.File
	file_google_cloud_automl_v1_classification_proto_rawDesc = nil
	file_google_cloud_automl_v1_classification_proto_goTypes = nil
	file_google_cloud_automl_v1_classification_proto_depIdxs = nil
}
