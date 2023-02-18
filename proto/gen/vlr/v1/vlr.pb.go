// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: vlr/v1/vlr.proto

package v1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Score int32  `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	Icon  string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_v1_vlr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_v1_vlr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_vlr_v1_vlr_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Team) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamOne        *Team                `protobuf:"bytes,2,opt,name=team_one,json=teamOne,proto3" json:"team_one,omitempty"`
	TeamTwo        *Team                `protobuf:"bytes,3,opt,name=team_two,json=teamTwo,proto3" json:"team_two,omitempty"`
	Maps           []string             `protobuf:"bytes,4,rep,name=maps,proto3" json:"maps,omitempty"`
	TournamentName string               `protobuf:"bytes,5,opt,name=tournament_name,json=tournamentName,proto3" json:"tournament_name,omitempty"`
	Date           *timestamp.Timestamp `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_v1_vlr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_v1_vlr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_vlr_v1_vlr_proto_rawDescGZIP(), []int{1}
}

func (x *Match) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Match) GetTeamOne() *Team {
	if x != nil {
		return x.TeamOne
	}
	return nil
}

func (x *Match) GetTeamTwo() *Team {
	if x != nil {
		return x.TeamTwo
	}
	return nil
}

func (x *Match) GetMaps() []string {
	if x != nil {
		return x.Maps
	}
	return nil
}

func (x *Match) GetTournamentName() string {
	if x != nil {
		return x.TournamentName
	}
	return ""
}

func (x *Match) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

type MatchesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matches []*Match `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
}

func (x *MatchesResponse) Reset() {
	*x = MatchesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_v1_vlr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchesResponse) ProtoMessage() {}

func (x *MatchesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_v1_vlr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchesResponse.ProtoReflect.Descriptor instead.
func (*MatchesResponse) Descriptor() ([]byte, []int) {
	return file_vlr_v1_vlr_proto_rawDescGZIP(), []int{2}
}

func (x *MatchesResponse) GetMatches() []*Match {
	if x != nil {
		return x.Matches
	}
	return nil
}

type MatchesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *MatchesRequest) Reset() {
	*x = MatchesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vlr_v1_vlr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchesRequest) ProtoMessage() {}

func (x *MatchesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vlr_v1_vlr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchesRequest.ProtoReflect.Descriptor instead.
func (*MatchesRequest) Descriptor() ([]byte, []int) {
	return file_vlr_v1_vlr_proto_rawDescGZIP(), []int{3}
}

func (x *MatchesRequest) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

var File_vlr_v1_vlr_proto protoreflect.FileDescriptor

var file_vlr_v1_vlr_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x6c, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6c, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x76, 0x6c, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x04, 0x54,
	0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x22, 0xd6, 0x01, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x76, 0x6c, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x07, 0x74, 0x65, 0x61,
	0x6d, 0x4f, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x74, 0x77, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x6c, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x54, 0x77, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x61, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x70,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x3a, 0x0a, 0x0f, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x76, 0x6c, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x07, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x22, 0x40, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32, 0x41, 0x0a, 0x03, 0x56, 0x6c, 0x72, 0x12,
	0x3a, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x76, 0x6c, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x6c, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x69, 0x6f, 0x53,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x2f, 0x76, 0x6c, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x6c, 0x72, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vlr_v1_vlr_proto_rawDescOnce sync.Once
	file_vlr_v1_vlr_proto_rawDescData = file_vlr_v1_vlr_proto_rawDesc
)

func file_vlr_v1_vlr_proto_rawDescGZIP() []byte {
	file_vlr_v1_vlr_proto_rawDescOnce.Do(func() {
		file_vlr_v1_vlr_proto_rawDescData = protoimpl.X.CompressGZIP(file_vlr_v1_vlr_proto_rawDescData)
	})
	return file_vlr_v1_vlr_proto_rawDescData
}

var file_vlr_v1_vlr_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_vlr_v1_vlr_proto_goTypes = []interface{}{
	(*Team)(nil),                // 0: vlr.v1.Team
	(*Match)(nil),               // 1: vlr.v1.Match
	(*MatchesResponse)(nil),     // 2: vlr.v1.MatchesResponse
	(*MatchesRequest)(nil),      // 3: vlr.v1.MatchesRequest
	(*timestamp.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_vlr_v1_vlr_proto_depIdxs = []int32{
	0, // 0: vlr.v1.Match.team_one:type_name -> vlr.v1.Team
	0, // 1: vlr.v1.Match.team_two:type_name -> vlr.v1.Team
	4, // 2: vlr.v1.Match.date:type_name -> google.protobuf.Timestamp
	1, // 3: vlr.v1.MatchesResponse.matches:type_name -> vlr.v1.Match
	4, // 4: vlr.v1.MatchesRequest.date:type_name -> google.protobuf.Timestamp
	3, // 5: vlr.v1.Vlr.Matches:input_type -> vlr.v1.MatchesRequest
	2, // 6: vlr.v1.Vlr.Matches:output_type -> vlr.v1.MatchesResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_vlr_v1_vlr_proto_init() }
func file_vlr_v1_vlr_proto_init() {
	if File_vlr_v1_vlr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vlr_v1_vlr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Team); i {
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
		file_vlr_v1_vlr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
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
		file_vlr_v1_vlr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchesResponse); i {
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
		file_vlr_v1_vlr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchesRequest); i {
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
			RawDescriptor: file_vlr_v1_vlr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vlr_v1_vlr_proto_goTypes,
		DependencyIndexes: file_vlr_v1_vlr_proto_depIdxs,
		MessageInfos:      file_vlr_v1_vlr_proto_msgTypes,
	}.Build()
	File_vlr_v1_vlr_proto = out.File
	file_vlr_v1_vlr_proto_rawDesc = nil
	file_vlr_v1_vlr_proto_goTypes = nil
	file_vlr_v1_vlr_proto_depIdxs = nil
}
