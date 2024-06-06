// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: protos/vote/election.proto

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

type Election struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OpenDate  string `protobuf:"bytes,3,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	EndDate   string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt int64  `protobuf:"varint,7,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Election) Reset() {
	*x = Election{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Election) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Election) ProtoMessage() {}

func (x *Election) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Election.ProtoReflect.Descriptor instead.
func (*Election) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{0}
}

func (x *Election) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Election) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Election) GetOpenDate() string {
	if x != nil {
		return x.OpenDate
	}
	return ""
}

func (x *Election) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Election) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Election) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Election) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

// Request for getting election reuslts
type GetCandidateVotesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCandidateVotesReq) Reset() {
	*x = GetCandidateVotesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCandidateVotesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandidateVotesReq) ProtoMessage() {}

func (x *GetCandidateVotesReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandidateVotesReq.ProtoReflect.Descriptor instead.
func (*GetCandidateVotesReq) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{1}
}

func (x *GetCandidateVotesReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Response for getting election rusults
type GetCandidateVotesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidateRes []*CandidateElectionRes `protobuf:"bytes,1,rep,name=candidateRes,proto3" json:"candidateRes,omitempty"`
}

func (x *GetCandidateVotesRes) Reset() {
	*x = GetCandidateVotesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCandidateVotesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandidateVotesRes) ProtoMessage() {}

func (x *GetCandidateVotesRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandidateVotesRes.ProtoReflect.Descriptor instead.
func (*GetCandidateVotesRes) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{2}
}

func (x *GetCandidateVotesRes) GetCandidateRes() []*CandidateElectionRes {
	if x != nil {
		return x.CandidateRes
	}
	return nil
}

type CandidateElectionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidateId string `protobuf:"bytes,1,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	Count       int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CandidateElectionRes) Reset() {
	*x = CandidateElectionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CandidateElectionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CandidateElectionRes) ProtoMessage() {}

func (x *CandidateElectionRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CandidateElectionRes.ProtoReflect.Descriptor instead.
func (*CandidateElectionRes) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{3}
}

func (x *CandidateElectionRes) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

func (x *CandidateElectionRes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// Request for creating a new election
type ElectionCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OpenDate string `protobuf:"bytes,2,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	EndDate  string `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *ElectionCreate) Reset() {
	*x = ElectionCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionCreate) ProtoMessage() {}

func (x *ElectionCreate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionCreate.ProtoReflect.Descriptor instead.
func (*ElectionCreate) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{4}
}

func (x *ElectionCreate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ElectionCreate) GetOpenDate() string {
	if x != nil {
		return x.OpenDate
	}
	return ""
}

func (x *ElectionCreate) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// Request for updating an existing election
type ElectionUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OpenDate string `protobuf:"bytes,3,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	EndDate  string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *ElectionUpdate) Reset() {
	*x = ElectionUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionUpdate) ProtoMessage() {}

func (x *ElectionUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionUpdate.ProtoReflect.Descriptor instead.
func (*ElectionUpdate) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{5}
}

func (x *ElectionUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ElectionUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ElectionUpdate) GetOpenDate() string {
	if x != nil {
		return x.OpenDate
	}
	return ""
}

func (x *ElectionUpdate) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// Request for deleting an existing election
type ElectionDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ElectionDelete) Reset() {
	*x = ElectionDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionDelete) ProtoMessage() {}

func (x *ElectionDelete) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionDelete.ProtoReflect.Descriptor instead.
func (*ElectionDelete) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{6}
}

func (x *ElectionDelete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request for getting an election by its ID
type ElectionById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ElectionById) Reset() {
	*x = ElectionById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ElectionById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ElectionById) ProtoMessage() {}

func (x *ElectionById) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ElectionById.ProtoReflect.Descriptor instead.
func (*ElectionById) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{7}
}

func (x *ElectionById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request for getting all elections
type GetAllElectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	OpenDate string `protobuf:"bytes,3,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	EndDate  string `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *GetAllElectionReq) Reset() {
	*x = GetAllElectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllElectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllElectionReq) ProtoMessage() {}

func (x *GetAllElectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllElectionReq.ProtoReflect.Descriptor instead.
func (*GetAllElectionReq) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllElectionReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllElectionReq) GetOpenDate() string {
	if x != nil {
		return x.OpenDate
	}
	return ""
}

func (x *GetAllElectionReq) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// Response for getting all elections
type GetAllElectionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elections []*Election `protobuf:"bytes,1,rep,name=elections,proto3" json:"elections,omitempty"`
	Count     int64       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllElectionRes) Reset() {
	*x = GetAllElectionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_election_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllElectionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllElectionRes) ProtoMessage() {}

func (x *GetAllElectionRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_election_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllElectionRes.ProtoReflect.Descriptor instead.
func (*GetAllElectionRes) Descriptor() ([]byte, []int) {
	return file_protos_vote_election_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllElectionRes) GetElections() []*Election {
	if x != nil {
		return x.Elections
	}
	return nil
}

func (x *GetAllElectionRes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_protos_vote_election_proto protoreflect.FileDescriptor

var file_protos_vote_election_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2f, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x74,
	0x65, 0x2f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x08, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x58, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x52, 0x0c, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x4f, 0x0a, 0x14, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x0e, 0x45, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x6c, 0x0a, 0x0e, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x59, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x2e,
	0x0a, 0x09, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xf4, 0x02, 0x0a, 0x0f, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x30,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x65, 0x73, 0x12, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f,
	0x76, 0x6f, 0x74, 0x65, 0x3b, 0x76, 0x6f, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_vote_election_proto_rawDescOnce sync.Once
	file_protos_vote_election_proto_rawDescData = file_protos_vote_election_proto_rawDesc
)

func file_protos_vote_election_proto_rawDescGZIP() []byte {
	file_protos_vote_election_proto_rawDescOnce.Do(func() {
		file_protos_vote_election_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_vote_election_proto_rawDescData)
	})
	return file_protos_vote_election_proto_rawDescData
}

var file_protos_vote_election_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_vote_election_proto_goTypes = []interface{}{
	(*Election)(nil),             // 0: protos.Election
	(*GetCandidateVotesReq)(nil), // 1: protos.GetCandidateVotesReq
	(*GetCandidateVotesRes)(nil), // 2: protos.GetCandidateVotesRes
	(*CandidateElectionRes)(nil), // 3: protos.CandidateElectionRes
	(*ElectionCreate)(nil),       // 4: protos.ElectionCreate
	(*ElectionUpdate)(nil),       // 5: protos.ElectionUpdate
	(*ElectionDelete)(nil),       // 6: protos.ElectionDelete
	(*ElectionById)(nil),         // 7: protos.ElectionById
	(*GetAllElectionReq)(nil),    // 8: protos.GetAllElectionReq
	(*GetAllElectionRes)(nil),    // 9: protos.GetAllElectionRes
	(*Void)(nil),                 // 10: protos.Void
}
var file_protos_vote_election_proto_depIdxs = []int32{
	3,  // 0: protos.GetCandidateVotesRes.candidateRes:type_name -> protos.CandidateElectionRes
	0,  // 1: protos.GetAllElectionRes.elections:type_name -> protos.Election
	4,  // 2: protos.ElectionService.Create:input_type -> protos.ElectionCreate
	5,  // 3: protos.ElectionService.Update:input_type -> protos.ElectionUpdate
	6,  // 4: protos.ElectionService.Delete:input_type -> protos.ElectionDelete
	7,  // 5: protos.ElectionService.GetById:input_type -> protos.ElectionById
	8,  // 6: protos.ElectionService.GetAll:input_type -> protos.GetAllElectionReq
	1,  // 7: protos.ElectionService.GetCandidateVoes:input_type -> protos.GetCandidateVotesReq
	0,  // 8: protos.ElectionService.Create:output_type -> protos.Election
	10, // 9: protos.ElectionService.Update:output_type -> protos.Void
	10, // 10: protos.ElectionService.Delete:output_type -> protos.Void
	0,  // 11: protos.ElectionService.GetById:output_type -> protos.Election
	9,  // 12: protos.ElectionService.GetAll:output_type -> protos.GetAllElectionRes
	2,  // 13: protos.ElectionService.GetCandidateVoes:output_type -> protos.GetCandidateVotesRes
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_protos_vote_election_proto_init() }
func file_protos_vote_election_proto_init() {
	if File_protos_vote_election_proto != nil {
		return
	}
	file_protos_vote_candidate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_vote_election_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Election); i {
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
		file_protos_vote_election_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCandidateVotesReq); i {
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
		file_protos_vote_election_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCandidateVotesRes); i {
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
		file_protos_vote_election_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CandidateElectionRes); i {
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
		file_protos_vote_election_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionCreate); i {
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
		file_protos_vote_election_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionUpdate); i {
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
		file_protos_vote_election_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionDelete); i {
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
		file_protos_vote_election_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ElectionById); i {
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
		file_protos_vote_election_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllElectionReq); i {
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
		file_protos_vote_election_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllElectionRes); i {
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
			RawDescriptor: file_protos_vote_election_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_vote_election_proto_goTypes,
		DependencyIndexes: file_protos_vote_election_proto_depIdxs,
		MessageInfos:      file_protos_vote_election_proto_msgTypes,
	}.Build()
	File_protos_vote_election_proto = out.File
	file_protos_vote_election_proto_rawDesc = nil
	file_protos_vote_election_proto_goTypes = nil
	file_protos_vote_election_proto_depIdxs = nil
}
