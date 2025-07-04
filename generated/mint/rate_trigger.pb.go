// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: rate_trigger.proto

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

type CreateRateTriggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuoteCurrency      string  `protobuf:"bytes,1,opt,name=quote_currency,json=quoteCurrency,proto3" json:"quote_currency,omitempty"`
	BaseCurrency       string  `protobuf:"bytes,2,opt,name=base_currency,json=baseCurrency,proto3" json:"base_currency,omitempty"`
	QuoteCurrencyId    int32   `protobuf:"varint,3,opt,name=quote_currency_id,json=quoteCurrencyId,proto3" json:"quote_currency_id,omitempty"`
	BaseCurrencyId     int32   `protobuf:"varint,4,opt,name=base_currency_id,json=baseCurrencyId,proto3" json:"base_currency_id,omitempty"`
	UserId             string  `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ThresholdAmount    float64 `protobuf:"fixed64,6,opt,name=threshold_amount,json=thresholdAmount,proto3" json:"threshold_amount,omitempty"`
	BaseCurrencyAmount float64 `protobuf:"fixed64,7,opt,name=base_currency_amount,json=baseCurrencyAmount,proto3" json:"base_currency_amount,omitempty"`
}

func (x *CreateRateTriggerRequest) Reset() {
	*x = CreateRateTriggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_trigger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRateTriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRateTriggerRequest) ProtoMessage() {}

func (x *CreateRateTriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_trigger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRateTriggerRequest.ProtoReflect.Descriptor instead.
func (*CreateRateTriggerRequest) Descriptor() ([]byte, []int) {
	return file_rate_trigger_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRateTriggerRequest) GetQuoteCurrency() string {
	if x != nil {
		return x.QuoteCurrency
	}
	return ""
}

func (x *CreateRateTriggerRequest) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

func (x *CreateRateTriggerRequest) GetQuoteCurrencyId() int32 {
	if x != nil {
		return x.QuoteCurrencyId
	}
	return 0
}

func (x *CreateRateTriggerRequest) GetBaseCurrencyId() int32 {
	if x != nil {
		return x.BaseCurrencyId
	}
	return 0
}

func (x *CreateRateTriggerRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateRateTriggerRequest) GetThresholdAmount() float64 {
	if x != nil {
		return x.ThresholdAmount
	}
	return 0
}

func (x *CreateRateTriggerRequest) GetBaseCurrencyAmount() float64 {
	if x != nil {
		return x.BaseCurrencyAmount
	}
	return 0
}

type CreateRateTriggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Expression         string  `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
	Triggered          bool    `protobuf:"varint,3,opt,name=triggered,proto3" json:"triggered,omitempty"`
	BaseCurrency       string  `protobuf:"bytes,4,opt,name=base_currency,json=baseCurrency,proto3" json:"base_currency,omitempty"`
	QuoteCurrency      string  `protobuf:"bytes,5,opt,name=quote_currency,json=quoteCurrency,proto3" json:"quote_currency,omitempty"`
	ThresholdAmount    float64 `protobuf:"fixed64,6,opt,name=threshold_amount,json=thresholdAmount,proto3" json:"threshold_amount,omitempty"`
	BaseCurrencyAmount float64 `protobuf:"fixed64,7,opt,name=base_currency_amount,json=baseCurrencyAmount,proto3" json:"base_currency_amount,omitempty"`
	QuoteCurrencyId    int32   `protobuf:"varint,8,opt,name=quote_currency_id,json=quoteCurrencyId,proto3" json:"quote_currency_id,omitempty"`
	BaseCurrencyId     int32   `protobuf:"varint,9,opt,name=base_currency_id,json=baseCurrencyId,proto3" json:"base_currency_id,omitempty"`
	UserId             string  `protobuf:"bytes,10,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateRateTriggerResponse) Reset() {
	*x = CreateRateTriggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_trigger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRateTriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRateTriggerResponse) ProtoMessage() {}

func (x *CreateRateTriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rate_trigger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRateTriggerResponse.ProtoReflect.Descriptor instead.
func (*CreateRateTriggerResponse) Descriptor() ([]byte, []int) {
	return file_rate_trigger_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRateTriggerResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateRateTriggerResponse) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *CreateRateTriggerResponse) GetTriggered() bool {
	if x != nil {
		return x.Triggered
	}
	return false
}

func (x *CreateRateTriggerResponse) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

func (x *CreateRateTriggerResponse) GetQuoteCurrency() string {
	if x != nil {
		return x.QuoteCurrency
	}
	return ""
}

func (x *CreateRateTriggerResponse) GetThresholdAmount() float64 {
	if x != nil {
		return x.ThresholdAmount
	}
	return 0
}

func (x *CreateRateTriggerResponse) GetBaseCurrencyAmount() float64 {
	if x != nil {
		return x.BaseCurrencyAmount
	}
	return 0
}

func (x *CreateRateTriggerResponse) GetQuoteCurrencyId() int32 {
	if x != nil {
		return x.QuoteCurrencyId
	}
	return 0
}

func (x *CreateRateTriggerResponse) GetBaseCurrencyId() int32 {
	if x != nil {
		return x.BaseCurrencyId
	}
	return 0
}

func (x *CreateRateTriggerResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CelEvaluationInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseCurrency    string  `protobuf:"bytes,1,opt,name=base_currency,json=baseCurrency,proto3" json:"base_currency,omitempty"`
	QuoteCurrency   string  `protobuf:"bytes,2,opt,name=quote_currency,json=quoteCurrency,proto3" json:"quote_currency,omitempty"`
	ThresholdAmount float64 `protobuf:"fixed64,3,opt,name=threshold_amount,json=thresholdAmount,proto3" json:"threshold_amount,omitempty"`
	UserId          string  `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CelEvaluationInput) Reset() {
	*x = CelEvaluationInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_trigger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CelEvaluationInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CelEvaluationInput) ProtoMessage() {}

func (x *CelEvaluationInput) ProtoReflect() protoreflect.Message {
	mi := &file_rate_trigger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CelEvaluationInput.ProtoReflect.Descriptor instead.
func (*CelEvaluationInput) Descriptor() ([]byte, []int) {
	return file_rate_trigger_proto_rawDescGZIP(), []int{2}
}

func (x *CelEvaluationInput) GetBaseCurrency() string {
	if x != nil {
		return x.BaseCurrency
	}
	return ""
}

func (x *CelEvaluationInput) GetQuoteCurrency() string {
	if x != nil {
		return x.QuoteCurrency
	}
	return ""
}

func (x *CelEvaluationInput) GetThresholdAmount() float64 {
	if x != nil {
		return x.ThresholdAmount
	}
	return 0
}

func (x *CelEvaluationInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetPendingRateTriggersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetPendingRateTriggersRequest) Reset() {
	*x = GetPendingRateTriggersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_trigger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPendingRateTriggersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingRateTriggersRequest) ProtoMessage() {}

func (x *GetPendingRateTriggersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_trigger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingRateTriggersRequest.ProtoReflect.Descriptor instead.
func (*GetPendingRateTriggersRequest) Descriptor() ([]byte, []int) {
	return file_rate_trigger_proto_rawDescGZIP(), []int{3}
}

func (x *GetPendingRateTriggersRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetPendingRateTriggersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RateTriggers []*CreateRateTriggerResponse `protobuf:"bytes,1,rep,name=rate_triggers,json=rateTriggers,proto3" json:"rate_triggers,omitempty"`
}

func (x *GetPendingRateTriggersResponse) Reset() {
	*x = GetPendingRateTriggersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_trigger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPendingRateTriggersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingRateTriggersResponse) ProtoMessage() {}

func (x *GetPendingRateTriggersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rate_trigger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingRateTriggersResponse.ProtoReflect.Descriptor instead.
func (*GetPendingRateTriggersResponse) Descriptor() ([]byte, []int) {
	return file_rate_trigger_proto_rawDescGZIP(), []int{4}
}

func (x *GetPendingRateTriggersResponse) GetRateTriggers() []*CreateRateTriggerResponse {
	if x != nil {
		return x.RateTriggers
	}
	return nil
}

type DeleteRateTriggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteRateTriggerRequest) Reset() {
	*x = DeleteRateTriggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_trigger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRateTriggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRateTriggerRequest) ProtoMessage() {}

func (x *DeleteRateTriggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rate_trigger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRateTriggerRequest.ProtoReflect.Descriptor instead.
func (*DeleteRateTriggerRequest) Descriptor() ([]byte, []int) {
	return file_rate_trigger_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRateTriggerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteRateTriggerRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteRateTriggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteRateTriggerResponse) Reset() {
	*x = DeleteRateTriggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rate_trigger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRateTriggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRateTriggerResponse) ProtoMessage() {}

func (x *DeleteRateTriggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rate_trigger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRateTriggerResponse.ProtoReflect.Descriptor instead.
func (*DeleteRateTriggerResponse) Descriptor() ([]byte, []int) {
	return file_rate_trigger_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRateTriggerResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_rate_trigger_proto protoreflect.FileDescriptor

var file_rate_trigger_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x22, 0xb2, 0x02, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x2a, 0x0a, 0x11, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a,
	0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x62, 0x61, 0x73,
	0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x81, 0x03, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x25, 0x0a, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x12, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x62, 0x61, 0x73, 0x65,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xa4, 0x01, 0x0a, 0x12, 0x43, 0x65, 0x6c, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x25, 0x0a, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x1d, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x22, 0x43, 0x0a, 0x18,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x86, 0x03, 0x0a, 0x12, 0x52, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x74, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73,
	0x12, 0x33, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x2e, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x6f, 0x70, 0x74, 0x69, 0x73, 0x77, 0x69, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rate_trigger_proto_rawDescOnce sync.Once
	file_rate_trigger_proto_rawDescData = file_rate_trigger_proto_rawDesc
)

func file_rate_trigger_proto_rawDescGZIP() []byte {
	file_rate_trigger_proto_rawDescOnce.Do(func() {
		file_rate_trigger_proto_rawDescData = protoimpl.X.CompressGZIP(file_rate_trigger_proto_rawDescData)
	})
	return file_rate_trigger_proto_rawDescData
}

var file_rate_trigger_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rate_trigger_proto_goTypes = []any{
	(*CreateRateTriggerRequest)(nil),       // 0: optiswift.proto.mint.CreateRateTriggerRequest
	(*CreateRateTriggerResponse)(nil),      // 1: optiswift.proto.mint.CreateRateTriggerResponse
	(*CelEvaluationInput)(nil),             // 2: optiswift.proto.mint.CelEvaluationInput
	(*GetPendingRateTriggersRequest)(nil),  // 3: optiswift.proto.mint.GetPendingRateTriggersRequest
	(*GetPendingRateTriggersResponse)(nil), // 4: optiswift.proto.mint.GetPendingRateTriggersResponse
	(*DeleteRateTriggerRequest)(nil),       // 5: optiswift.proto.mint.DeleteRateTriggerRequest
	(*DeleteRateTriggerResponse)(nil),      // 6: optiswift.proto.mint.DeleteRateTriggerResponse
}
var file_rate_trigger_proto_depIdxs = []int32{
	1, // 0: optiswift.proto.mint.GetPendingRateTriggersResponse.rate_triggers:type_name -> optiswift.proto.mint.CreateRateTriggerResponse
	0, // 1: optiswift.proto.mint.RateTriggerService.CreateRateTrigger:input_type -> optiswift.proto.mint.CreateRateTriggerRequest
	3, // 2: optiswift.proto.mint.RateTriggerService.GetPendingRateTriggers:input_type -> optiswift.proto.mint.GetPendingRateTriggersRequest
	5, // 3: optiswift.proto.mint.RateTriggerService.DeleteRateTrigger:input_type -> optiswift.proto.mint.DeleteRateTriggerRequest
	1, // 4: optiswift.proto.mint.RateTriggerService.CreateRateTrigger:output_type -> optiswift.proto.mint.CreateRateTriggerResponse
	4, // 5: optiswift.proto.mint.RateTriggerService.GetPendingRateTriggers:output_type -> optiswift.proto.mint.GetPendingRateTriggersResponse
	6, // 6: optiswift.proto.mint.RateTriggerService.DeleteRateTrigger:output_type -> optiswift.proto.mint.DeleteRateTriggerResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rate_trigger_proto_init() }
func file_rate_trigger_proto_init() {
	if File_rate_trigger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rate_trigger_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRateTriggerRequest); i {
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
		file_rate_trigger_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRateTriggerResponse); i {
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
		file_rate_trigger_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CelEvaluationInput); i {
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
		file_rate_trigger_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetPendingRateTriggersRequest); i {
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
		file_rate_trigger_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetPendingRateTriggersResponse); i {
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
		file_rate_trigger_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRateTriggerRequest); i {
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
		file_rate_trigger_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRateTriggerResponse); i {
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
			RawDescriptor: file_rate_trigger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rate_trigger_proto_goTypes,
		DependencyIndexes: file_rate_trigger_proto_depIdxs,
		MessageInfos:      file_rate_trigger_proto_msgTypes,
	}.Build()
	File_rate_trigger_proto = out.File
	file_rate_trigger_proto_rawDesc = nil
	file_rate_trigger_proto_goTypes = nil
	file_rate_trigger_proto_depIdxs = nil
}
