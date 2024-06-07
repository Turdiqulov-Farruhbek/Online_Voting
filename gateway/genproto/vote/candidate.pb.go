// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: protos/vote/candidate.proto

package vote

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

type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  int64  `protobuf:"varint,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_candidate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_candidate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_protos_vote_candidate_proto_rawDescGZIP(), []int{0}
}

func (x *Candidate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Candidate) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *Candidate) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *Candidate) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Candidate) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Candidate) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_candidate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_candidate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_protos_vote_candidate_proto_rawDescGZIP(), []int{1}
}

type CandidateCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElectionId string `protobuf:"bytes,1,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,2,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
}

func (x *CandidateCreate) Reset() {
	*x = CandidateCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_candidate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateCreate) ProtoMessage() {}

func (x *CandidateCreate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_candidate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateCreate.ProtoReflect.Descriptor instead.
func (*CandidateCreate) Descriptor() ([]byte, []int) {
	return file_protos_vote_candidate_proto_rawDescGZIP(), []int{2}
}

func (x *CandidateCreate) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *CandidateCreate) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

type CandidateUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ElectionId string `protobuf:"bytes,2,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
}

func (x *CandidateUpdate) Reset() {
	*x = CandidateUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_candidate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateUpdate) ProtoMessage() {}

func (x *CandidateUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_candidate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateUpdate.ProtoReflect.Descriptor instead.
func (*CandidateUpdate) Descriptor() ([]byte, []int) {
	return file_protos_vote_candidate_proto_rawDescGZIP(), []int{3}
}

func (x *CandidateUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CandidateUpdate) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *CandidateUpdate) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

type CandidateDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CandidateDelete) Reset() {
	*x = CandidateDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_candidate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateDelete) ProtoMessage() {}

func (x *CandidateDelete) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_candidate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateDelete.ProtoReflect.Descriptor instead.
func (*CandidateDelete) Descriptor() ([]byte, []int) {
	return file_protos_vote_candidate_proto_rawDescGZIP(), []int{4}
}

func (x *CandidateDelete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CandidateById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CandidateById) Reset() {
	*x = CandidateById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_candidate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateById) ProtoMessage() {}

func (x *CandidateById) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_candidate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateById.ProtoReflect.Descriptor instead.
func (*CandidateById) Descriptor() ([]byte, []int) {
	return file_protos_vote_candidate_proto_rawDescGZIP(), []int{5}
}

func (x *CandidateById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllCandidateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElectionId string `protobuf:"bytes,1,opt,name=election_id,json=electionId,proto3" json:"election_id,omitempty"`
	PublicId   string `protobuf:"bytes,2,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
}

func (x *GetAllCandidateReq) Reset() {
	*x = GetAllCandidateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_candidate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCandidateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCandidateReq) ProtoMessage() {}

func (x *GetAllCandidateReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_candidate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCandidateReq.ProtoReflect.Descriptor instead.
func (*GetAllCandidateReq) Descriptor() ([]byte, []int) {
	return file_protos_vote_candidate_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllCandidateReq) GetElectionId() string {
	if x != nil {
		return x.ElectionId
	}
	return ""
}

func (x *GetAllCandidateReq) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

type GetAllCandidateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candidates []*Candidate `protobuf:"bytes,1,rep,name=candidates,proto3" json:"candidates,omitempty"`
	Count      int64        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllCandidateRes) Reset() {
	*x = GetAllCandidateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_candidate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCandidateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCandidateRes) ProtoMessage() {}

func (x *GetAllCandidateRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_candidate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCandidateRes.ProtoReflect.Descriptor instead.
func (*GetAllCandidateRes) Descriptor() ([]byte, []int) {
	return file_protos_vote_candidate_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllCandidateRes) GetCandidates() []*Candidate {
	if x != nil {
		return x.Candidates
	}
	return nil
}

func (x *GetAllCandidateRes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_protos_vote_candidate_proto protoreflect.FileDescriptor

var file_protos_vote_candidate_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x63, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x06,
	0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x64,
	0x22, 0x5d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0xab, 0x02, 0x0a, 0x10, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x3b, 0x76, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_vote_candidate_proto_rawDescOnce sync.Once
	file_protos_vote_candidate_proto_rawDescData = file_protos_vote_candidate_proto_rawDesc
)

func file_protos_vote_candidate_proto_rawDescGZIP() []byte {
	file_protos_vote_candidate_proto_rawDescOnce.Do(func() {
		file_protos_vote_candidate_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_vote_candidate_proto_rawDescData)
	})
	return file_protos_vote_candidate_proto_rawDescData
}

var file_protos_vote_candidate_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_vote_candidate_proto_goTypes = []interface{}{
	(*Candidate)(nil),          // 0: protos.Candidate
	(*Void)(nil),               // 1: protos.Void
	(*CandidateCreate)(nil),    // 2: protos.CandidateCreate
	(*CandidateUpdate)(nil),    // 3: protos.CandidateUpdate
	(*CandidateDelete)(nil),    // 4: protos.CandidateDelete
	(*CandidateById)(nil),      // 5: protos.CandidateById
	(*GetAllCandidateReq)(nil), // 6: protos.GetAllCandidateReq
	(*GetAllCandidateRes)(nil), // 7: protos.GetAllCandidateRes
}
var file_protos_vote_candidate_proto_depIdxs = []int32{
	0, // 0: protos.GetAllCandidateRes.candidates:type_name -> protos.Candidate
	2, // 1: protos.CandidateService.Create:input_type -> protos.CandidateCreate
	3, // 2: protos.CandidateService.Update:input_type -> protos.CandidateUpdate
	4, // 3: protos.CandidateService.Delete:input_type -> protos.CandidateDelete
	5, // 4: protos.CandidateService.GetById:input_type -> protos.CandidateById
	6, // 5: protos.CandidateService.GetAll:input_type -> protos.GetAllCandidateReq
	0, // 6: protos.CandidateService.Create:output_type -> protos.Candidate
	1, // 7: protos.CandidateService.Update:output_type -> protos.Void
	1, // 8: protos.CandidateService.Delete:output_type -> protos.Void
	0, // 9: protos.CandidateService.GetById:output_type -> protos.Candidate
	7, // 10: protos.CandidateService.GetAll:output_type -> protos.GetAllCandidateRes
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_vote_candidate_proto_init() }
func file_protos_vote_candidate_proto_init() {
	if File_protos_vote_candidate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_vote_candidate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
		file_protos_vote_candidate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_protos_vote_candidate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateCreate); i {
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
		file_protos_vote_candidate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateUpdate); i {
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
		file_protos_vote_candidate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateDelete); i {
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
		file_protos_vote_candidate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateById); i {
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
		file_protos_vote_candidate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCandidateReq); i {
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
		file_protos_vote_candidate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCandidateRes); i {
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
			RawDescriptor: file_protos_vote_candidate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_vote_candidate_proto_goTypes,
		DependencyIndexes: file_protos_vote_candidate_proto_depIdxs,
		MessageInfos:      file_protos_vote_candidate_proto_msgTypes,
	}.Build()
	File_protos_vote_candidate_proto = out.File
	file_protos_vote_candidate_proto_rawDesc = nil
	file_protos_vote_candidate_proto_goTypes = nil
	file_protos_vote_candidate_proto_depIdxs = nil
}