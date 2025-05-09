// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: rewards.proto

package proto

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

// Represents a reward rule with its benefits
type RewardSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reference string     `protobuf:"bytes,1,opt,name=reference,proto3" json:"reference,omitempty"`
	Benefits  []*Benefit `protobuf:"bytes,2,rep,name=benefits,proto3" json:"benefits,omitempty"`
}

func (x *RewardSummary) Reset() {
	*x = RewardSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardSummary) ProtoMessage() {}

func (x *RewardSummary) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardSummary.ProtoReflect.Descriptor instead.
func (*RewardSummary) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{0}
}

func (x *RewardSummary) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *RewardSummary) GetBenefits() []*Benefit {
	if x != nil {
		return x.Benefits
	}
	return nil
}

type FetchRewardSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	References []string `protobuf:"bytes,1,rep,name=references,proto3" json:"references,omitempty"`
}

func (x *FetchRewardSummaryRequest) Reset() {
	*x = FetchRewardSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRewardSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRewardSummaryRequest) ProtoMessage() {}

func (x *FetchRewardSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRewardSummaryRequest.ProtoReflect.Descriptor instead.
func (*FetchRewardSummaryRequest) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{1}
}

func (x *FetchRewardSummaryRequest) GetReferences() []string {
	if x != nil {
		return x.References
	}
	return nil
}

type FetchRewardSummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rewards []*RewardSummary `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *FetchRewardSummaryResponse) Reset() {
	*x = FetchRewardSummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRewardSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRewardSummaryResponse) ProtoMessage() {}

func (x *FetchRewardSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRewardSummaryResponse.ProtoReflect.Descriptor instead.
func (*FetchRewardSummaryResponse) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{2}
}

func (x *FetchRewardSummaryResponse) GetRewards() []*RewardSummary {
	if x != nil {
		return x.Rewards
	}
	return nil
}

type GetPotentialRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleIds []string `protobuf:"bytes,1,rep,name=rule_ids,json=ruleIds,proto3" json:"rule_ids,omitempty"`
}

func (x *GetPotentialRewardsRequest) Reset() {
	*x = GetPotentialRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPotentialRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPotentialRewardsRequest) ProtoMessage() {}

func (x *GetPotentialRewardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPotentialRewardsRequest.ProtoReflect.Descriptor instead.
func (*GetPotentialRewardsRequest) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{3}
}

func (x *GetPotentialRewardsRequest) GetRuleIds() []string {
	if x != nil {
		return x.RuleIds
	}
	return nil
}

type GetPotentialRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rewards []*RewardSummary `protobuf:"bytes,1,rep,name=rewards,proto3" json:"rewards,omitempty"`
}

func (x *GetPotentialRewardsResponse) Reset() {
	*x = GetPotentialRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rewards_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPotentialRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPotentialRewardsResponse) ProtoMessage() {}

func (x *GetPotentialRewardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rewards_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPotentialRewardsResponse.ProtoReflect.Descriptor instead.
func (*GetPotentialRewardsResponse) Descriptor() ([]byte, []int) {
	return file_rewards_proto_rawDescGZIP(), []int{4}
}

func (x *GetPotentialRewardsResponse) GetRewards() []*RewardSummary {
	if x != nil {
		return x.Rewards
	}
	return nil
}

var File_rewards_proto protoreflect.FileDescriptor

var file_rewards_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x6e, 0x75, 0x74, 0x6d, 0x65, 0x67, 0x1a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x0d, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69,
	0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x75, 0x74, 0x6d, 0x65, 0x67, 0x2e,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x52, 0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x73, 0x22, 0x3b, 0x0a, 0x19, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x5d,
	0x0a, 0x1a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6e, 0x75, 0x74, 0x6d, 0x65, 0x67, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x22, 0x37, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72,
	0x75, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72,
	0x75, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x22, 0x5e, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69,
	0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x75, 0x74, 0x6d, 0x65, 0x67, 0x2e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x32, 0x86, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x7b, 0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x31, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73,
	0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x75, 0x74, 0x6d, 0x65,
	0x67, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x75,
	0x74, 0x6d, 0x65, 0x67, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x7e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x32, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69,
	0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x75, 0x74, 0x6d, 0x65, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6f, 0x70, 0x74,
	0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x75, 0x74,
	0x6d, 0x65, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rewards_proto_rawDescOnce sync.Once
	file_rewards_proto_rawDescData = file_rewards_proto_rawDesc
)

func file_rewards_proto_rawDescGZIP() []byte {
	file_rewards_proto_rawDescOnce.Do(func() {
		file_rewards_proto_rawDescData = protoimpl.X.CompressGZIP(file_rewards_proto_rawDescData)
	})
	return file_rewards_proto_rawDescData
}

var file_rewards_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rewards_proto_goTypes = []any{
	(*RewardSummary)(nil),               // 0: optiswift.proto.nutmeg.RewardSummary
	(*FetchRewardSummaryRequest)(nil),   // 1: optiswift.proto.nutmeg.FetchRewardSummaryRequest
	(*FetchRewardSummaryResponse)(nil),  // 2: optiswift.proto.nutmeg.FetchRewardSummaryResponse
	(*GetPotentialRewardsRequest)(nil),  // 3: optiswift.proto.nutmeg.GetPotentialRewardsRequest
	(*GetPotentialRewardsResponse)(nil), // 4: optiswift.proto.nutmeg.GetPotentialRewardsResponse
	(*Benefit)(nil),                     // 5: optiswift.proto.nutmeg.Benefit
}
var file_rewards_proto_depIdxs = []int32{
	5, // 0: optiswift.proto.nutmeg.RewardSummary.benefits:type_name -> optiswift.proto.nutmeg.Benefit
	0, // 1: optiswift.proto.nutmeg.FetchRewardSummaryResponse.rewards:type_name -> optiswift.proto.nutmeg.RewardSummary
	0, // 2: optiswift.proto.nutmeg.GetPotentialRewardsResponse.rewards:type_name -> optiswift.proto.nutmeg.RewardSummary
	1, // 3: optiswift.proto.nutmeg.Rewards.FetchRewardSummary:input_type -> optiswift.proto.nutmeg.FetchRewardSummaryRequest
	3, // 4: optiswift.proto.nutmeg.Rewards.GetPotentialRewards:input_type -> optiswift.proto.nutmeg.GetPotentialRewardsRequest
	2, // 5: optiswift.proto.nutmeg.Rewards.FetchRewardSummary:output_type -> optiswift.proto.nutmeg.FetchRewardSummaryResponse
	4, // 6: optiswift.proto.nutmeg.Rewards.GetPotentialRewards:output_type -> optiswift.proto.nutmeg.GetPotentialRewardsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rewards_proto_init() }
func file_rewards_proto_init() {
	if File_rewards_proto != nil {
		return
	}
	file_events_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rewards_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RewardSummary); i {
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
		file_rewards_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*FetchRewardSummaryRequest); i {
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
		file_rewards_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FetchRewardSummaryResponse); i {
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
		file_rewards_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetPotentialRewardsRequest); i {
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
		file_rewards_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetPotentialRewardsResponse); i {
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
			RawDescriptor: file_rewards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rewards_proto_goTypes,
		DependencyIndexes: file_rewards_proto_depIdxs,
		MessageInfos:      file_rewards_proto_msgTypes,
	}.Build()
	File_rewards_proto = out.File
	file_rewards_proto_rawDesc = nil
	file_rewards_proto_goTypes = nil
	file_rewards_proto_depIdxs = nil
}
