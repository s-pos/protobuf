// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: pb/auth/email.proto

package auth

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

type Verification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Otp   string `protobuf:"bytes,3,opt,name=otp,proto3" json:"otp,omitempty"`
	// Types that are assignable to TokenVerification:
	//	*Verification_Token
	TokenVerification isVerification_TokenVerification `protobuf_oneof:"tokenVerification"`
}

func (x *Verification) Reset() {
	*x = Verification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Verification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verification) ProtoMessage() {}

func (x *Verification) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verification.ProtoReflect.Descriptor instead.
func (*Verification) Descriptor() ([]byte, []int) {
	return file_pb_auth_email_proto_rawDescGZIP(), []int{0}
}

func (x *Verification) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Verification) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Verification) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (m *Verification) GetTokenVerification() isVerification_TokenVerification {
	if m != nil {
		return m.TokenVerification
	}
	return nil
}

func (x *Verification) GetToken() string {
	if x, ok := x.GetTokenVerification().(*Verification_Token); ok {
		return x.Token
	}
	return ""
}

type isVerification_TokenVerification interface {
	isVerification_TokenVerification()
}

type Verification_Token struct {
	Token string `protobuf:"bytes,4,opt,name=token,proto3,oneof"`
}

func (*Verification_Token) isVerification_TokenVerification() {}

type ResetPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ResetPassword) Reset() {
	*x = ResetPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPassword) ProtoMessage() {}

func (x *ResetPassword) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPassword.ProtoReflect.Descriptor instead.
func (*ResetPassword) Descriptor() ([]byte, []int) {
	return file_pb_auth_email_proto_rawDescGZIP(), []int{1}
}

func (x *ResetPassword) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResetPassword) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResetPassword) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerificationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *VerificationReply) Reset() {
	*x = VerificationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationReply) ProtoMessage() {}

func (x *VerificationReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationReply.ProtoReflect.Descriptor instead.
func (*VerificationReply) Descriptor() ([]byte, []int) {
	return file_pb_auth_email_proto_rawDescGZIP(), []int{2}
}

func (x *VerificationReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResetPasswordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResetPasswordReply) Reset() {
	*x = ResetPasswordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordReply) ProtoMessage() {}

func (x *ResetPasswordReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordReply.ProtoReflect.Descriptor instead.
func (*ResetPasswordReply) Descriptor() ([]byte, []int) {
	return file_pb_auth_email_proto_rawDescGZIP(), []int{3}
}

func (x *ResetPasswordReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_auth_email_proto protoreflect.FileDescriptor

var file_pb_auth_email_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x77, 0x0a, 0x0c, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x13, 0x0a, 0x11, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2d, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xa4, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a,
	0x18, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_auth_email_proto_rawDescOnce sync.Once
	file_pb_auth_email_proto_rawDescData = file_pb_auth_email_proto_rawDesc
)

func file_pb_auth_email_proto_rawDescGZIP() []byte {
	file_pb_auth_email_proto_rawDescOnce.Do(func() {
		file_pb_auth_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_auth_email_proto_rawDescData)
	})
	return file_pb_auth_email_proto_rawDescData
}

var file_pb_auth_email_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_auth_email_proto_goTypes = []interface{}{
	(*Verification)(nil),       // 0: auth.Verification
	(*ResetPassword)(nil),      // 1: auth.ResetPassword
	(*VerificationReply)(nil),  // 2: auth.VerificationReply
	(*ResetPasswordReply)(nil), // 3: auth.ResetPasswordReply
}
var file_pb_auth_email_proto_depIdxs = []int32{
	0, // 0: auth.UserAuthService.SendEmailVerification:input_type -> auth.Verification
	1, // 1: auth.UserAuthService.SendEmailResetPassword:input_type -> auth.ResetPassword
	2, // 2: auth.UserAuthService.SendEmailVerification:output_type -> auth.VerificationReply
	3, // 3: auth.UserAuthService.SendEmailResetPassword:output_type -> auth.ResetPasswordReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_auth_email_proto_init() }
func file_pb_auth_email_proto_init() {
	if File_pb_auth_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_auth_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Verification); i {
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
		file_pb_auth_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPassword); i {
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
		file_pb_auth_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationReply); i {
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
		file_pb_auth_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordReply); i {
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
	file_pb_auth_email_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Verification_Token)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_auth_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_auth_email_proto_goTypes,
		DependencyIndexes: file_pb_auth_email_proto_depIdxs,
		MessageInfos:      file_pb_auth_email_proto_msgTypes,
	}.Build()
	File_pb_auth_email_proto = out.File
	file_pb_auth_email_proto_rawDesc = nil
	file_pb_auth_email_proto_goTypes = nil
	file_pb_auth_email_proto_depIdxs = nil
}
