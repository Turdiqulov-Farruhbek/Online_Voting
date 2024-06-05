// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: protos/vote.proto

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

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CandidateId string `protobuf:"bytes,2,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
	CreatedAt   string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_protos_vote_proto_rawDescGZIP(), []int{0}
}

func (x *Vote) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Vote) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

func (x *Vote) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

// Request for creating a new vote
type VoteCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CandidateId string `protobuf:"bytes,1,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
}

func (x *VoteCreate) Reset() {
	*x = VoteCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteCreate) ProtoMessage() {}

func (x *VoteCreate) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteCreate.ProtoReflect.Descriptor instead.
func (*VoteCreate) Descriptor() ([]byte, []int) {
	return file_protos_vote_proto_rawDescGZIP(), []int{1}
}

func (x *VoteCreate) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

// Request for getting a vote by its ID
type VoteById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VoteById) Reset() {
	*x = VoteById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteById) ProtoMessage() {}

func (x *VoteById) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteById.ProtoReflect.Descriptor instead.
func (*VoteById) Descriptor() ([]byte, []int) {
	return file_protos_vote_proto_rawDescGZIP(), []int{2}
}

func (x *VoteById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Request for getting all votes
type GetAllVoteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional filter for candidate ID
	CandidateId string `protobuf:"bytes,1,opt,name=candidate_id,json=candidateId,proto3" json:"candidate_id,omitempty"`
}

func (x *GetAllVoteReq) Reset() {
	*x = GetAllVoteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllVoteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllVoteReq) ProtoMessage() {}

func (x *GetAllVoteReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllVoteReq.ProtoReflect.Descriptor instead.
func (*GetAllVoteReq) Descriptor() ([]byte, []int) {
	return file_protos_vote_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllVoteReq) GetCandidateId() string {
	if x != nil {
		return x.CandidateId
	}
	return ""
}

// Response for getting all votes
type GetAllVoteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Votes []*Vote `protobuf:"bytes,1,rep,name=votes,proto3" json:"votes,omitempty"`
	Count int32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllVoteRes) Reset() {
	*x = GetAllVoteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vote_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllVoteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllVoteRes) ProtoMessage() {}

func (x *GetAllVoteRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vote_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllVoteRes.ProtoReflect.Descriptor instead.
func (*GetAllVoteRes) Descriptor() ([]byte, []int) {
	return file_protos_vote_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllVoteRes) GetVotes() []*Vote {
	if x != nil {
		return x.Votes
	}
	return nil
}

func (x *GetAllVoteRes) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_protos_vote_proto protoreflect.FileDescriptor

var file_protos_vote_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x16, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2f, 0x0a,
	0x0a, 0x56, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x1a,
	0x0a, 0x08, 0x56, 0x6f, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x22, 0x49,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12,
	0x22, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x76, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xa2, 0x01, 0x0a, 0x0b, 0x56, 0x6f,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x42, 0x79, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x10,
	0x5a, 0x0e, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x76, 0x6f, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_vote_proto_rawDescOnce sync.Once
	file_protos_vote_proto_rawDescData = file_protos_vote_proto_rawDesc
)

func file_protos_vote_proto_rawDescGZIP() []byte {
	file_protos_vote_proto_rawDescOnce.Do(func() {
		file_protos_vote_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_vote_proto_rawDescData)
	})
	return file_protos_vote_proto_rawDescData
}

var file_protos_vote_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_vote_proto_goTypes = []interface{}{
	(*Vote)(nil),          // 0: protos.Vote
	(*VoteCreate)(nil),    // 1: protos.VoteCreate
	(*VoteById)(nil),      // 2: protos.VoteById
	(*GetAllVoteReq)(nil), // 3: protos.GetAllVoteReq
	(*GetAllVoteRes)(nil), // 4: protos.GetAllVoteRes
	(*Void)(nil),          // 5: protos.Void
}
var file_protos_vote_proto_depIdxs = []int32{
	0, // 0: protos.GetAllVoteRes.votes:type_name -> protos.Vote
	1, // 1: protos.VoteService.Create:input_type -> protos.VoteCreate
	2, // 2: protos.VoteService.GetById:input_type -> protos.VoteById
	3, // 3: protos.VoteService.GetAll:input_type -> protos.GetAllVoteReq
	0, // 4: protos.VoteService.Create:output_type -> protos.Vote
	5, // 5: protos.VoteService.GetById:output_type -> protos.Void
	4, // 6: protos.VoteService.GetAll:output_type -> protos.GetAllVoteRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_vote_proto_init() }
func file_protos_vote_proto_init() {
	if File_protos_vote_proto != nil {
		return
	}
	file_protos_candidate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_vote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
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
		file_protos_vote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteCreate); i {
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
		file_protos_vote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteById); i {
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
		file_protos_vote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllVoteReq); i {
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
		file_protos_vote_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllVoteRes); i {
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
			RawDescriptor: file_protos_vote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_vote_proto_goTypes,
		DependencyIndexes: file_protos_vote_proto_depIdxs,
		MessageInfos:      file_protos_vote_proto_msgTypes,
	}.Build()
	File_protos_vote_proto = out.File
	file_protos_vote_proto_rawDesc = nil
	file_protos_vote_proto_goTypes = nil
	file_protos_vote_proto_depIdxs = nil
}
