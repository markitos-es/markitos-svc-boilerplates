// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: boilerplate.proto

package gapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Boilerplate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Boilerplate) Reset() {
	*x = Boilerplate{}
	mi := &file_boilerplate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Boilerplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Boilerplate) ProtoMessage() {}

func (x *Boilerplate) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Boilerplate.ProtoReflect.Descriptor instead.
func (*Boilerplate) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{0}
}

func (x *Boilerplate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Boilerplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Boilerplate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Boilerplate) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateBoilerplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBoilerplateRequest) Reset() {
	*x = CreateBoilerplateRequest{}
	mi := &file_boilerplate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBoilerplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoilerplateRequest) ProtoMessage() {}

func (x *CreateBoilerplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoilerplateRequest.ProtoReflect.Descriptor instead.
func (*CreateBoilerplateRequest) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBoilerplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateBoilerplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBoilerplateResponse) Reset() {
	*x = CreateBoilerplateResponse{}
	mi := &file_boilerplate_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBoilerplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoilerplateResponse) ProtoMessage() {}

func (x *CreateBoilerplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoilerplateResponse.ProtoReflect.Descriptor instead.
func (*CreateBoilerplateResponse) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBoilerplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateBoilerplateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetBoilerplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBoilerplateRequest) Reset() {
	*x = GetBoilerplateRequest{}
	mi := &file_boilerplate_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBoilerplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoilerplateRequest) ProtoMessage() {}

func (x *GetBoilerplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoilerplateRequest.ProtoReflect.Descriptor instead.
func (*GetBoilerplateRequest) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{3}
}

func (x *GetBoilerplateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBoilerplateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBoilerplateResponse) Reset() {
	*x = GetBoilerplateResponse{}
	mi := &file_boilerplate_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBoilerplateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoilerplateResponse) ProtoMessage() {}

func (x *GetBoilerplateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_boilerplate_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoilerplateResponse.ProtoReflect.Descriptor instead.
func (*GetBoilerplateResponse) Descriptor() ([]byte, []int) {
	return file_boilerplate_proto_rawDescGZIP(), []int{4}
}

func (x *GetBoilerplateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetBoilerplateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_boilerplate_proto protoreflect.FileDescriptor

var file_boilerplate_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2e, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x69, 0x6c,
	0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xd3, 0x01, 0x0a, 0x12, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12,
	0x25, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x69, 0x6c, 0x65,
	0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x69, 0x74, 0x6f, 0x73,
	0x2d, 0x65, 0x73, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x69, 0x74, 0x6f, 0x73, 0x2d, 0x73, 0x76, 0x63,
	0x2d, 0x62, 0x6f, 0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x67, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_boilerplate_proto_rawDescOnce sync.Once
	file_boilerplate_proto_rawDescData []byte
)

func file_boilerplate_proto_rawDescGZIP() []byte {
	file_boilerplate_proto_rawDescOnce.Do(func() {
		file_boilerplate_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_boilerplate_proto_rawDesc), len(file_boilerplate_proto_rawDesc)))
	})
	return file_boilerplate_proto_rawDescData
}

var file_boilerplate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_boilerplate_proto_goTypes = []any{
	(*Boilerplate)(nil),               // 0: boilerplate.Boilerplate
	(*CreateBoilerplateRequest)(nil),  // 1: boilerplate.CreateBoilerplateRequest
	(*CreateBoilerplateResponse)(nil), // 2: boilerplate.CreateBoilerplateResponse
	(*GetBoilerplateRequest)(nil),     // 3: boilerplate.GetBoilerplateRequest
	(*GetBoilerplateResponse)(nil),    // 4: boilerplate.GetBoilerplateResponse
	(*timestamppb.Timestamp)(nil),     // 5: google.protobuf.Timestamp
}
var file_boilerplate_proto_depIdxs = []int32{
	5, // 0: boilerplate.Boilerplate.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: boilerplate.Boilerplate.updated_at:type_name -> google.protobuf.Timestamp
	1, // 2: boilerplate.BoilerplateService.CreateBoilerplate:input_type -> boilerplate.CreateBoilerplateRequest
	3, // 3: boilerplate.BoilerplateService.GetBoilerplate:input_type -> boilerplate.GetBoilerplateRequest
	2, // 4: boilerplate.BoilerplateService.CreateBoilerplate:output_type -> boilerplate.CreateBoilerplateResponse
	4, // 5: boilerplate.BoilerplateService.GetBoilerplate:output_type -> boilerplate.GetBoilerplateResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_boilerplate_proto_init() }
func file_boilerplate_proto_init() {
	if File_boilerplate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_boilerplate_proto_rawDesc), len(file_boilerplate_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_boilerplate_proto_goTypes,
		DependencyIndexes: file_boilerplate_proto_depIdxs,
		MessageInfos:      file_boilerplate_proto_msgTypes,
	}.Build()
	File_boilerplate_proto = out.File
	file_boilerplate_proto_goTypes = nil
	file_boilerplate_proto_depIdxs = nil
}
