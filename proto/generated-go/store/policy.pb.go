// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: store/policy.proto

package store

import (
	expr "google.golang.org/genproto/googleapis/type/expr"
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

type IamPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Collection of binding.
	Bindings []*Binding `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
}

func (x *IamPolicy) Reset() {
	*x = IamPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IamPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IamPolicy) ProtoMessage() {}

func (x *IamPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IamPolicy.ProtoReflect.Descriptor instead.
func (*IamPolicy) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{0}
}

func (x *IamPolicy) GetBindings() []*Binding {
	if x != nil {
		return x.Bindings
	}
	return nil
}

// Reference: https://cloud.google.com/pubsub/docs/reference/rpc/google.iam.v1#binding
type Binding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Role that is assigned to the list of members.
	// Format: roles/{role}
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// Specifies the principals requesting access for a Bytebase resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone.
	// * `user:{emailid}`: An email address that represents a specific Bytebase account. For example, `alice@example.com`.
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// The condition that is associated with this binding.
	// If the condition evaluates to true, then this binding applies to the current request.
	// If the condition evaluates to false, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding.
	Condition *expr.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
}

func (x *Binding) Reset() {
	*x = Binding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Binding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Binding) ProtoMessage() {}

func (x *Binding) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Binding.ProtoReflect.Descriptor instead.
func (*Binding) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{1}
}

func (x *Binding) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Binding) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Binding) GetCondition() *expr.Expr {
	if x != nil {
		return x.Condition
	}
	return nil
}

type MaskingPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaskData []*MaskData `protobuf:"bytes,1,rep,name=mask_data,json=maskData,proto3" json:"mask_data,omitempty"`
}

func (x *MaskingPolicy) Reset() {
	*x = MaskingPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskingPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskingPolicy) ProtoMessage() {}

func (x *MaskingPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskingPolicy.ProtoReflect.Descriptor instead.
func (*MaskingPolicy) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{2}
}

func (x *MaskingPolicy) GetMaskData() []*MaskData {
	if x != nil {
		return x.MaskData
	}
	return nil
}

type MaskData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema             string       `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
	Table              string       `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Column             string       `protobuf:"bytes,3,opt,name=column,proto3" json:"column,omitempty"`
	SemanticCategoryId string       `protobuf:"bytes,4,opt,name=semantic_category_id,json=semanticCategoryId,proto3" json:"semantic_category_id,omitempty"`
	MaskingLevel       MaskingLevel `protobuf:"varint,5,opt,name=masking_level,json=maskingLevel,proto3,enum=bytebase.store.MaskingLevel" json:"masking_level,omitempty"`
}

func (x *MaskData) Reset() {
	*x = MaskData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskData) ProtoMessage() {}

func (x *MaskData) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskData.ProtoReflect.Descriptor instead.
func (*MaskData) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{3}
}

func (x *MaskData) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *MaskData) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *MaskData) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *MaskData) GetSemanticCategoryId() string {
	if x != nil {
		return x.SemanticCategoryId
	}
	return ""
}

func (x *MaskData) GetMaskingLevel() MaskingLevel {
	if x != nil {
		return x.MaskingLevel
	}
	return MaskingLevel_MASKING_LEVEL_UNSPECIFIED
}

type MaskingRulePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules []*MaskingRulePolicy_MaskingRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *MaskingRulePolicy) Reset() {
	*x = MaskingRulePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskingRulePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskingRulePolicy) ProtoMessage() {}

func (x *MaskingRulePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskingRulePolicy.ProtoReflect.Descriptor instead.
func (*MaskingRulePolicy) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{4}
}

func (x *MaskingRulePolicy) GetRules() []*MaskingRulePolicy_MaskingRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

type MaskingRulePolicy_MaskingRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition    *expr.Expr   `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	MaskingLevel MaskingLevel `protobuf:"varint,2,opt,name=masking_level,json=maskingLevel,proto3,enum=bytebase.store.MaskingLevel" json:"masking_level,omitempty"`
}

func (x *MaskingRulePolicy_MaskingRule) Reset() {
	*x = MaskingRulePolicy_MaskingRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_policy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaskingRulePolicy_MaskingRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaskingRulePolicy_MaskingRule) ProtoMessage() {}

func (x *MaskingRulePolicy_MaskingRule) ProtoReflect() protoreflect.Message {
	mi := &file_store_policy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaskingRulePolicy_MaskingRule.ProtoReflect.Descriptor instead.
func (*MaskingRulePolicy_MaskingRule) Descriptor() ([]byte, []int) {
	return file_store_policy_proto_rawDescGZIP(), []int{4, 0}
}

func (x *MaskingRulePolicy_MaskingRule) GetCondition() *expr.Expr {
	if x != nil {
		return x.Condition
	}
	return nil
}

func (x *MaskingRulePolicy_MaskingRule) GetMaskingLevel() MaskingLevel {
	if x != nil {
		return x.MaskingLevel
	}
	return MaskingLevel_MASKING_LEVEL_UNSPECIFIED
}

var File_store_policy_proto protoreflect.FileDescriptor

var file_store_policy_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x1a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x65, 0x78, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x09, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x33, 0x0a,
	0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x22, 0x68, 0x0a, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70,
	0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x0d,
	0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x35, 0x0a,
	0x09, 0x6d, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x61, 0x73, 0x6b,
	0x44, 0x61, 0x74, 0x61, 0x22, 0xc5, 0x01, 0x0a, 0x08, 0x4d, 0x61, 0x73, 0x6b, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x65, 0x6d, 0x61, 0x6e,
	0x74, 0x69, 0x63, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0d, 0x6d, 0x61, 0x73,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c,
	0x6d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0xdc, 0x01, 0x0a,
	0x11, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x43, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x81, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x73, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0d, 0x6d, 0x61, 0x73, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x4d, 0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c, 0x6d,
	0x61, 0x73, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x14, 0x5a, 0x12, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_policy_proto_rawDescOnce sync.Once
	file_store_policy_proto_rawDescData = file_store_policy_proto_rawDesc
)

func file_store_policy_proto_rawDescGZIP() []byte {
	file_store_policy_proto_rawDescOnce.Do(func() {
		file_store_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_policy_proto_rawDescData)
	})
	return file_store_policy_proto_rawDescData
}

var file_store_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_store_policy_proto_goTypes = []interface{}{
	(*IamPolicy)(nil),                     // 0: bytebase.store.IamPolicy
	(*Binding)(nil),                       // 1: bytebase.store.Binding
	(*MaskingPolicy)(nil),                 // 2: bytebase.store.MaskingPolicy
	(*MaskData)(nil),                      // 3: bytebase.store.MaskData
	(*MaskingRulePolicy)(nil),             // 4: bytebase.store.MaskingRulePolicy
	(*MaskingRulePolicy_MaskingRule)(nil), // 5: bytebase.store.MaskingRulePolicy.MaskingRule
	(*expr.Expr)(nil),                     // 6: google.type.Expr
	(MaskingLevel)(0),                     // 7: bytebase.store.MaskingLevel
}
var file_store_policy_proto_depIdxs = []int32{
	1, // 0: bytebase.store.IamPolicy.bindings:type_name -> bytebase.store.Binding
	6, // 1: bytebase.store.Binding.condition:type_name -> google.type.Expr
	3, // 2: bytebase.store.MaskingPolicy.mask_data:type_name -> bytebase.store.MaskData
	7, // 3: bytebase.store.MaskData.masking_level:type_name -> bytebase.store.MaskingLevel
	5, // 4: bytebase.store.MaskingRulePolicy.rules:type_name -> bytebase.store.MaskingRulePolicy.MaskingRule
	6, // 5: bytebase.store.MaskingRulePolicy.MaskingRule.condition:type_name -> google.type.Expr
	7, // 6: bytebase.store.MaskingRulePolicy.MaskingRule.masking_level:type_name -> bytebase.store.MaskingLevel
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_store_policy_proto_init() }
func file_store_policy_proto_init() {
	if File_store_policy_proto != nil {
		return
	}
	file_store_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_store_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IamPolicy); i {
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
		file_store_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Binding); i {
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
		file_store_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskingPolicy); i {
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
		file_store_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskData); i {
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
		file_store_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskingRulePolicy); i {
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
		file_store_policy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaskingRulePolicy_MaskingRule); i {
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
			RawDescriptor: file_store_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_policy_proto_goTypes,
		DependencyIndexes: file_store_policy_proto_depIdxs,
		MessageInfos:      file_store_policy_proto_msgTypes,
	}.Build()
	File_store_policy_proto = out.File
	file_store_policy_proto_rawDesc = nil
	file_store_policy_proto_goTypes = nil
	file_store_policy_proto_depIdxs = nil
}
