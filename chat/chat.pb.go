// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: chat.proto

package grpc_api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The request message containing the user's name
type ClaimRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClaimId string `protobuf:"bytes,1,opt,name=claim_id,json=claimId,proto3" json:"claim_id,omitempty"`
}

func (x *ClaimRequest) Reset() {
	*x = ClaimRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimRequest) ProtoMessage() {}

func (x *ClaimRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimRequest.ProtoReflect.Descriptor instead.
func (*ClaimRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ClaimRequest) GetClaimId() string {
	if x != nil {
		return x.ClaimId
	}
	return ""
}

// The response message containing the greetings
type ClaimResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Claim *Claim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *ClaimResponse) Reset() {
	*x = ClaimResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimResponse) ProtoMessage() {}

func (x *ClaimResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimResponse.ProtoReflect.Descriptor instead.
func (*ClaimResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ClaimResponse) GetClaim() *Claim {
	if x != nil {
		return x.Claim
	}
	return nil
}

type Claim struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClaimId       string `protobuf:"bytes,1,opt,name=claim_id,json=claimId,proto3" json:"claim_id,omitempty"`
	ClaimAmount   string `protobuf:"bytes,2,opt,name=claim_amount,json=claimAmount,proto3" json:"claim_amount,omitempty"`
	ClaimCoverage string `protobuf:"bytes,3,opt,name=claim_coverage,json=claimCoverage,proto3" json:"claim_coverage,omitempty"`
	ClaimMessage  string `protobuf:"bytes,4,opt,name=claim_message,json=claimMessage,proto3" json:"claim_message,omitempty"`
}

func (x *Claim) Reset() {
	*x = Claim{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claim) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claim) ProtoMessage() {}

func (x *Claim) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claim.ProtoReflect.Descriptor instead.
func (*Claim) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *Claim) GetClaimId() string {
	if x != nil {
		return x.ClaimId
	}
	return ""
}

func (x *Claim) GetClaimAmount() string {
	if x != nil {
		return x.ClaimAmount
	}
	return ""
}

func (x *Claim) GetClaimCoverage() string {
	if x != nil {
		return x.ClaimCoverage
	}
	return ""
}

func (x *Claim) GetClaimMessage() string {
	if x != nil {
		return x.ClaimMessage
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x29, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0d, 0x43,
	0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x22,
	0x91, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x32, 0x61, 0x0a, 0x0c, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x12,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12,
	0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2f, 0x7b, 0x63, 0x6c, 0x61, 0x69,
	0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x32, 0x39, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x6d, 0x61, 0x6c, 0x38, 0x39, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chat_proto_goTypes = []interface{}{
	(*ClaimRequest)(nil),  // 0: chat.ClaimRequest
	(*ClaimResponse)(nil), // 1: chat.ClaimResponse
	(*Claim)(nil),         // 2: chat.Claim
	(*Message)(nil),       // 3: chat.Message
}
var file_chat_proto_depIdxs = []int32{
	2, // 0: chat.ClaimResponse.claim:type_name -> chat.Claim
	0, // 1: chat.ClaimService.GetClaim:input_type -> chat.ClaimRequest
	3, // 2: chat.ChatService.SayHello:input_type -> chat.Message
	1, // 3: chat.ClaimService.GetClaim:output_type -> chat.ClaimResponse
	3, // 4: chat.ChatService.SayHello:output_type -> chat.Message
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimRequest); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimResponse); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claim); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
