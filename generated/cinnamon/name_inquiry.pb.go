// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: name_inquiry.proto

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

type NameInquiryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstitutionId   string  `protobuf:"bytes,1,opt,name=institution_id,json=institutionId,proto3" json:"institution_id,omitempty"`
	AccountId       string  `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Country         string  `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	PaymentMethod   string  `protobuf:"bytes,4,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	UserId          *string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
	SenderFirstname *string `protobuf:"bytes,6,opt,name=sender_firstname,json=senderFirstname,proto3,oneof" json:"sender_firstname,omitempty"`
	SenderLastname  *string `protobuf:"bytes,7,opt,name=sender_lastname,json=senderLastname,proto3,oneof" json:"sender_lastname,omitempty"`
	AccountName     *string `protobuf:"bytes,8,opt,name=account_name,json=accountName,proto3,oneof" json:"account_name,omitempty"`
	BranchCode      *string `protobuf:"bytes,9,opt,name=branch_code,json=branchCode,proto3,oneof" json:"branch_code,omitempty"`
}

func (x *NameInquiryRequest) Reset() {
	*x = NameInquiryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_name_inquiry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameInquiryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameInquiryRequest) ProtoMessage() {}

func (x *NameInquiryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_name_inquiry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameInquiryRequest.ProtoReflect.Descriptor instead.
func (*NameInquiryRequest) Descriptor() ([]byte, []int) {
	return file_name_inquiry_proto_rawDescGZIP(), []int{0}
}

func (x *NameInquiryRequest) GetInstitutionId() string {
	if x != nil {
		return x.InstitutionId
	}
	return ""
}

func (x *NameInquiryRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *NameInquiryRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *NameInquiryRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *NameInquiryRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *NameInquiryRequest) GetSenderFirstname() string {
	if x != nil && x.SenderFirstname != nil {
		return *x.SenderFirstname
	}
	return ""
}

func (x *NameInquiryRequest) GetSenderLastname() string {
	if x != nil && x.SenderLastname != nil {
		return *x.SenderLastname
	}
	return ""
}

func (x *NameInquiryRequest) GetAccountName() string {
	if x != nil && x.AccountName != nil {
		return *x.AccountName
	}
	return ""
}

func (x *NameInquiryRequest) GetBranchCode() string {
	if x != nil && x.BranchCode != nil {
		return *x.BranchCode
	}
	return ""
}

type NameInquiryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	Reference   string `protobuf:"bytes,2,opt,name=reference,proto3" json:"reference,omitempty"`
}

func (x *NameInquiryResponse) Reset() {
	*x = NameInquiryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_name_inquiry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameInquiryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameInquiryResponse) ProtoMessage() {}

func (x *NameInquiryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_name_inquiry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameInquiryResponse.ProtoReflect.Descriptor instead.
func (*NameInquiryResponse) Descriptor() ([]byte, []int) {
	return file_name_inquiry_proto_rawDescGZIP(), []int{1}
}

func (x *NameInquiryResponse) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *NameInquiryResponse) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

var File_name_inquiry_proto protoreflect.FileDescriptor

var file_name_inquiry_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x6f, 0x6e, 0x22, 0xbb,
	0x03, 0x0a, 0x12, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69,
	0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4c, 0x61, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x24, 0x0a, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x43,
	0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x56, 0x0a, 0x13,
	0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x32, 0x74, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x12, 0x65, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x2c, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x70,
	0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69,
	0x6e, 0x6e, 0x61, 0x6d, 0x6f, 0x6e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x71, 0x75, 0x69,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_name_inquiry_proto_rawDescOnce sync.Once
	file_name_inquiry_proto_rawDescData = file_name_inquiry_proto_rawDesc
)

func file_name_inquiry_proto_rawDescGZIP() []byte {
	file_name_inquiry_proto_rawDescOnce.Do(func() {
		file_name_inquiry_proto_rawDescData = protoimpl.X.CompressGZIP(file_name_inquiry_proto_rawDescData)
	})
	return file_name_inquiry_proto_rawDescData
}

var file_name_inquiry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_name_inquiry_proto_goTypes = []any{
	(*NameInquiryRequest)(nil),  // 0: optiswift.proto.cinnamon.NameInquiryRequest
	(*NameInquiryResponse)(nil), // 1: optiswift.proto.cinnamon.NameInquiryResponse
}
var file_name_inquiry_proto_depIdxs = []int32{
	0, // 0: optiswift.proto.cinnamon.NameInquiry.Lookup:input_type -> optiswift.proto.cinnamon.NameInquiryRequest
	1, // 1: optiswift.proto.cinnamon.NameInquiry.Lookup:output_type -> optiswift.proto.cinnamon.NameInquiryResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_name_inquiry_proto_init() }
func file_name_inquiry_proto_init() {
	if File_name_inquiry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_name_inquiry_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NameInquiryRequest); i {
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
		file_name_inquiry_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NameInquiryResponse); i {
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
	file_name_inquiry_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_name_inquiry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_name_inquiry_proto_goTypes,
		DependencyIndexes: file_name_inquiry_proto_depIdxs,
		MessageInfos:      file_name_inquiry_proto_msgTypes,
	}.Build()
	File_name_inquiry_proto = out.File
	file_name_inquiry_proto_rawDesc = nil
	file_name_inquiry_proto_goTypes = nil
	file_name_inquiry_proto_depIdxs = nil
}
