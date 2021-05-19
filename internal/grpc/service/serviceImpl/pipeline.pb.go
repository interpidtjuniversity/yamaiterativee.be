// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: pipeline.proto

package serviceImpl

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StartYaMaPipeLineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int64  `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName   string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Repository string `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	Branch     string `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *StartYaMaPipeLineRequest) Reset() {
	*x = StartYaMaPipeLineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartYaMaPipeLineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartYaMaPipeLineRequest) ProtoMessage() {}

func (x *StartYaMaPipeLineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartYaMaPipeLineRequest.ProtoReflect.Descriptor instead.
func (*StartYaMaPipeLineRequest) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *StartYaMaPipeLineRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *StartYaMaPipeLineRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *StartYaMaPipeLineRequest) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *StartYaMaPipeLineRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

type StartYaMaPipeLineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StartYaMaPipeLineResponse) Reset() {
	*x = StartYaMaPipeLineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartYaMaPipeLineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartYaMaPipeLineResponse) ProtoMessage() {}

func (x *StartYaMaPipeLineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartYaMaPipeLineResponse.ProtoReflect.Descriptor instead.
func (*StartYaMaPipeLineResponse) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *StartYaMaPipeLineResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PassMergeRequestCodeReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionId int64 `protobuf:"varint,1,opt,name=actionId,proto3" json:"actionId,omitempty"`
	StageId  int64 `protobuf:"varint,2,opt,name=stageId,proto3" json:"stageId,omitempty"`
	StepId   int64 `protobuf:"varint,3,opt,name=stepId,proto3" json:"stepId,omitempty"`
}

func (x *PassMergeRequestCodeReviewRequest) Reset() {
	*x = PassMergeRequestCodeReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassMergeRequestCodeReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassMergeRequestCodeReviewRequest) ProtoMessage() {}

func (x *PassMergeRequestCodeReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassMergeRequestCodeReviewRequest.ProtoReflect.Descriptor instead.
func (*PassMergeRequestCodeReviewRequest) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *PassMergeRequestCodeReviewRequest) GetActionId() int64 {
	if x != nil {
		return x.ActionId
	}
	return 0
}

func (x *PassMergeRequestCodeReviewRequest) GetStageId() int64 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *PassMergeRequestCodeReviewRequest) GetStepId() int64 {
	if x != nil {
		return x.StepId
	}
	return 0
}

type PassMergeRequestCodeReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PassMergeRequestCodeReviewResponse) Reset() {
	*x = PassMergeRequestCodeReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassMergeRequestCodeReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassMergeRequestCodeReviewResponse) ProtoMessage() {}

func (x *PassMergeRequestCodeReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassMergeRequestCodeReviewResponse.ProtoReflect.Descriptor instead.
func (*PassMergeRequestCodeReviewResponse) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{3}
}

func (x *PassMergeRequestCodeReviewResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RestartYaMaPipeLineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PipelineId    int64    `protobuf:"varint,1,opt,name=pipelineId,proto3" json:"pipelineId,omitempty"`
	IterationId   int64    `protobuf:"varint,2,opt,name=iterationId,proto3" json:"iterationId,omitempty"`
	ActorName     string   `protobuf:"bytes,3,opt,name=actorName,proto3" json:"actorName,omitempty"`
	SourceBranch  string   `protobuf:"bytes,4,opt,name=sourceBranch,proto3" json:"sourceBranch,omitempty"`
	TargetBranch  string   `protobuf:"bytes,5,opt,name=targetBranch,proto3" json:"targetBranch,omitempty"`
	MrCodeReviews []string `protobuf:"bytes,6,rep,name=mrCodeReviews,proto3" json:"mrCodeReviews,omitempty"`
	Env           string   `protobuf:"bytes,7,opt,name=env,proto3" json:"env,omitempty"`
	MrInfo        string   `protobuf:"bytes,8,opt,name=mrInfo,proto3" json:"mrInfo,omitempty"`
	AppOwner      string   `protobuf:"bytes,9,opt,name=appOwner,proto3" json:"appOwner,omitempty"`
	AppName       string   `protobuf:"bytes,10,opt,name=appName,proto3" json:"appName,omitempty"`
	ActionId      int64    `protobuf:"varint,11,opt,name=actionId,proto3" json:"actionId,omitempty"`
}

func (x *RestartYaMaPipeLineRequest) Reset() {
	*x = RestartYaMaPipeLineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartYaMaPipeLineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartYaMaPipeLineRequest) ProtoMessage() {}

func (x *RestartYaMaPipeLineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartYaMaPipeLineRequest.ProtoReflect.Descriptor instead.
func (*RestartYaMaPipeLineRequest) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{4}
}

func (x *RestartYaMaPipeLineRequest) GetPipelineId() int64 {
	if x != nil {
		return x.PipelineId
	}
	return 0
}

func (x *RestartYaMaPipeLineRequest) GetIterationId() int64 {
	if x != nil {
		return x.IterationId
	}
	return 0
}

func (x *RestartYaMaPipeLineRequest) GetActorName() string {
	if x != nil {
		return x.ActorName
	}
	return ""
}

func (x *RestartYaMaPipeLineRequest) GetSourceBranch() string {
	if x != nil {
		return x.SourceBranch
	}
	return ""
}

func (x *RestartYaMaPipeLineRequest) GetTargetBranch() string {
	if x != nil {
		return x.TargetBranch
	}
	return ""
}

func (x *RestartYaMaPipeLineRequest) GetMrCodeReviews() []string {
	if x != nil {
		return x.MrCodeReviews
	}
	return nil
}

func (x *RestartYaMaPipeLineRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *RestartYaMaPipeLineRequest) GetMrInfo() string {
	if x != nil {
		return x.MrInfo
	}
	return ""
}

func (x *RestartYaMaPipeLineRequest) GetAppOwner() string {
	if x != nil {
		return x.AppOwner
	}
	return ""
}

func (x *RestartYaMaPipeLineRequest) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *RestartYaMaPipeLineRequest) GetActionId() int64 {
	if x != nil {
		return x.ActionId
	}
	return 0
}

type RestartYaMaPipeLineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RestartYaMaPipeLineResponse) Reset() {
	*x = RestartYaMaPipeLineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartYaMaPipeLineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartYaMaPipeLineResponse) ProtoMessage() {}

func (x *RestartYaMaPipeLineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartYaMaPipeLineResponse.ProtoReflect.Descriptor instead.
func (*RestartYaMaPipeLineResponse) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{5}
}

func (x *RestartYaMaPipeLineResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_pipeline_proto protoreflect.FileDescriptor

var file_pipeline_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x18, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x59, 0x61, 0x4d, 0x61, 0x50, 0x69, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x22, 0x35, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61, 0x4d, 0x61, 0x50, 0x69, 0x70,
	0x65, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x71, 0x0a, 0x21, 0x50, 0x61, 0x73, 0x73, 0x4d,
	0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x74, 0x65, 0x70, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x22, 0x50, 0x61,
	0x73, 0x73, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xe6, 0x02, 0x0a, 0x1a, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61, 0x4d, 0x61, 0x50, 0x69, 0x70, 0x65, 0x4c, 0x69,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x74, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x22, 0x0a,
	0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x1b, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61,
	0x4d, 0x61, 0x50, 0x69, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xc4, 0x02, 0x0a,
	0x13, 0x59, 0x61, 0x4d, 0x61, 0x50, 0x69, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61, 0x4d,
	0x61, 0x50, 0x69, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61, 0x4d, 0x61, 0x50, 0x69, 0x70, 0x65, 0x4c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61, 0x4d, 0x61, 0x50, 0x69, 0x70, 0x65,
	0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73,
	0x0a, 0x1a, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x28, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x61, 0x73, 0x73, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61,
	0x4d, 0x61, 0x50, 0x69, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61, 0x4d, 0x61, 0x50, 0x69,
	0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x59, 0x61, 0x4d,
	0x61, 0x50, 0x69, 0x70, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6d, 0x70, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pipeline_proto_rawDescOnce sync.Once
	file_pipeline_proto_rawDescData = file_pipeline_proto_rawDesc
)

func file_pipeline_proto_rawDescGZIP() []byte {
	file_pipeline_proto_rawDescOnce.Do(func() {
		file_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipeline_proto_rawDescData)
	})
	return file_pipeline_proto_rawDescData
}

var file_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pipeline_proto_goTypes = []interface{}{
	(*StartYaMaPipeLineRequest)(nil),           // 0: proto.StartYaMaPipeLineRequest
	(*StartYaMaPipeLineResponse)(nil),          // 1: proto.StartYaMaPipeLineResponse
	(*PassMergeRequestCodeReviewRequest)(nil),  // 2: proto.PassMergeRequestCodeReviewRequest
	(*PassMergeRequestCodeReviewResponse)(nil), // 3: proto.PassMergeRequestCodeReviewResponse
	(*RestartYaMaPipeLineRequest)(nil),         // 4: proto.RestartYaMaPipeLineRequest
	(*RestartYaMaPipeLineResponse)(nil),        // 5: proto.RestartYaMaPipeLineResponse
}
var file_pipeline_proto_depIdxs = []int32{
	0, // 0: proto.YaMaPipeLineService.StartYaMaPipeLine:input_type -> proto.StartYaMaPipeLineRequest
	2, // 1: proto.YaMaPipeLineService.PassMergeRequestCodeReview:input_type -> proto.PassMergeRequestCodeReviewRequest
	4, // 2: proto.YaMaPipeLineService.RestartYaMaPipeLine:input_type -> proto.RestartYaMaPipeLineRequest
	1, // 3: proto.YaMaPipeLineService.StartYaMaPipeLine:output_type -> proto.StartYaMaPipeLineResponse
	3, // 4: proto.YaMaPipeLineService.PassMergeRequestCodeReview:output_type -> proto.PassMergeRequestCodeReviewResponse
	5, // 5: proto.YaMaPipeLineService.RestartYaMaPipeLine:output_type -> proto.RestartYaMaPipeLineResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pipeline_proto_init() }
func file_pipeline_proto_init() {
	if File_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartYaMaPipeLineRequest); i {
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
		file_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartYaMaPipeLineResponse); i {
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
		file_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassMergeRequestCodeReviewRequest); i {
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
		file_pipeline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassMergeRequestCodeReviewResponse); i {
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
		file_pipeline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartYaMaPipeLineRequest); i {
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
		file_pipeline_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartYaMaPipeLineResponse); i {
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
			RawDescriptor: file_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pipeline_proto_goTypes,
		DependencyIndexes: file_pipeline_proto_depIdxs,
		MessageInfos:      file_pipeline_proto_msgTypes,
	}.Build()
	File_pipeline_proto = out.File
	file_pipeline_proto_rawDesc = nil
	file_pipeline_proto_goTypes = nil
	file_pipeline_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// YaMaPipeLineServiceClient is the client API for YaMaPipeLineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YaMaPipeLineServiceClient interface {
	StartYaMaPipeLine(ctx context.Context, in *StartYaMaPipeLineRequest, opts ...grpc.CallOption) (*StartYaMaPipeLineResponse, error)
	PassMergeRequestCodeReview(ctx context.Context, in *PassMergeRequestCodeReviewRequest, opts ...grpc.CallOption) (*PassMergeRequestCodeReviewResponse, error)
	RestartYaMaPipeLine(ctx context.Context, in *RestartYaMaPipeLineRequest, opts ...grpc.CallOption) (*RestartYaMaPipeLineResponse, error)
}

type yaMaPipeLineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYaMaPipeLineServiceClient(cc grpc.ClientConnInterface) YaMaPipeLineServiceClient {
	return &yaMaPipeLineServiceClient{cc}
}

func (c *yaMaPipeLineServiceClient) StartYaMaPipeLine(ctx context.Context, in *StartYaMaPipeLineRequest, opts ...grpc.CallOption) (*StartYaMaPipeLineResponse, error) {
	out := new(StartYaMaPipeLineResponse)
	err := c.cc.Invoke(ctx, "/proto.YaMaPipeLineService/StartYaMaPipeLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yaMaPipeLineServiceClient) PassMergeRequestCodeReview(ctx context.Context, in *PassMergeRequestCodeReviewRequest, opts ...grpc.CallOption) (*PassMergeRequestCodeReviewResponse, error) {
	out := new(PassMergeRequestCodeReviewResponse)
	err := c.cc.Invoke(ctx, "/proto.YaMaPipeLineService/PassMergeRequestCodeReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yaMaPipeLineServiceClient) RestartYaMaPipeLine(ctx context.Context, in *RestartYaMaPipeLineRequest, opts ...grpc.CallOption) (*RestartYaMaPipeLineResponse, error) {
	out := new(RestartYaMaPipeLineResponse)
	err := c.cc.Invoke(ctx, "/proto.YaMaPipeLineService/RestartYaMaPipeLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YaMaPipeLineServiceServer is the server API for YaMaPipeLineService service.
type YaMaPipeLineServiceServer interface {
	StartYaMaPipeLine(context.Context, *StartYaMaPipeLineRequest) (*StartYaMaPipeLineResponse, error)
	PassMergeRequestCodeReview(context.Context, *PassMergeRequestCodeReviewRequest) (*PassMergeRequestCodeReviewResponse, error)
	RestartYaMaPipeLine(context.Context, *RestartYaMaPipeLineRequest) (*RestartYaMaPipeLineResponse, error)
}

// UnimplementedYaMaPipeLineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedYaMaPipeLineServiceServer struct {
}

func (*UnimplementedYaMaPipeLineServiceServer) StartYaMaPipeLine(context.Context, *StartYaMaPipeLineRequest) (*StartYaMaPipeLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartYaMaPipeLine not implemented")
}
func (*UnimplementedYaMaPipeLineServiceServer) PassMergeRequestCodeReview(context.Context, *PassMergeRequestCodeReviewRequest) (*PassMergeRequestCodeReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PassMergeRequestCodeReview not implemented")
}
func (*UnimplementedYaMaPipeLineServiceServer) RestartYaMaPipeLine(context.Context, *RestartYaMaPipeLineRequest) (*RestartYaMaPipeLineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartYaMaPipeLine not implemented")
}

func RegisterYaMaPipeLineServiceServer(s *grpc.Server, srv YaMaPipeLineServiceServer) {
	s.RegisterService(&_YaMaPipeLineService_serviceDesc, srv)
}

func _YaMaPipeLineService_StartYaMaPipeLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartYaMaPipeLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YaMaPipeLineServiceServer).StartYaMaPipeLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YaMaPipeLineService/StartYaMaPipeLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YaMaPipeLineServiceServer).StartYaMaPipeLine(ctx, req.(*StartYaMaPipeLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YaMaPipeLineService_PassMergeRequestCodeReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassMergeRequestCodeReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YaMaPipeLineServiceServer).PassMergeRequestCodeReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YaMaPipeLineService/PassMergeRequestCodeReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YaMaPipeLineServiceServer).PassMergeRequestCodeReview(ctx, req.(*PassMergeRequestCodeReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YaMaPipeLineService_RestartYaMaPipeLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartYaMaPipeLineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YaMaPipeLineServiceServer).RestartYaMaPipeLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.YaMaPipeLineService/RestartYaMaPipeLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YaMaPipeLineServiceServer).RestartYaMaPipeLine(ctx, req.(*RestartYaMaPipeLineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _YaMaPipeLineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.YaMaPipeLineService",
	HandlerType: (*YaMaPipeLineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartYaMaPipeLine",
			Handler:    _YaMaPipeLineService_StartYaMaPipeLine_Handler,
		},
		{
			MethodName: "PassMergeRequestCodeReview",
			Handler:    _YaMaPipeLineService_PassMergeRequestCodeReview_Handler,
		},
		{
			MethodName: "RestartYaMaPipeLine",
			Handler:    _YaMaPipeLineService_RestartYaMaPipeLine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipeline.proto",
}
