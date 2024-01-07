// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: protos/notes/v1/notes.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Content *string `protobuf:"bytes,3,opt,name=content,proto3,oneof" json:"content,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notes_v1_notes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notes_v1_notes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_protos_notes_v1_notes_proto_rawDescGZIP(), []int{0}
}

func (x *Note) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Note) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Note) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

// Get
type GetNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetNoteRequest) Reset() {
	*x = GetNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notes_v1_notes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNoteRequest) ProtoMessage() {}

func (x *GetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notes_v1_notes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNoteRequest.ProtoReflect.Descriptor instead.
func (*GetNoteRequest) Descriptor() ([]byte, []int) {
	return file_protos_notes_v1_notes_proto_rawDescGZIP(), []int{1}
}

func (x *GetNoteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// List
type ListNotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter    *string `protobuf:"bytes,1,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
	PageSize  int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string  `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListNotesRequest) Reset() {
	*x = ListNotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notes_v1_notes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesRequest) ProtoMessage() {}

func (x *ListNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notes_v1_notes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesRequest.ProtoReflect.Descriptor instead.
func (*ListNotesRequest) Descriptor() ([]byte, []int) {
	return file_protos_notes_v1_notes_proto_rawDescGZIP(), []int{2}
}

func (x *ListNotesRequest) GetFilter() string {
	if x != nil && x.Filter != nil {
		return *x.Filter
	}
	return ""
}

func (x *ListNotesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListNotesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListNotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes         []*Note `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	NextPageToken string  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListNotesResponse) Reset() {
	*x = ListNotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notes_v1_notes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesResponse) ProtoMessage() {}

func (x *ListNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notes_v1_notes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesResponse.ProtoReflect.Descriptor instead.
func (*ListNotesResponse) Descriptor() ([]byte, []int) {
	return file_protos_notes_v1_notes_proto_rawDescGZIP(), []int{3}
}

func (x *ListNotesResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *ListNotesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Create
type CreateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content *string `protobuf:"bytes,2,opt,name=content,proto3,oneof" json:"content,omitempty"`
}

func (x *CreateNoteRequest) Reset() {
	*x = CreateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notes_v1_notes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteRequest) ProtoMessage() {}

func (x *CreateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notes_v1_notes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteRequest.ProtoReflect.Descriptor instead.
func (*CreateNoteRequest) Descriptor() ([]byte, []int) {
	return file_protos_notes_v1_notes_proto_rawDescGZIP(), []int{4}
}

func (x *CreateNoteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNoteRequest) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

// Update
type UpdateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *UpdateNoteRequest) Reset() {
	*x = UpdateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notes_v1_notes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteRequest) ProtoMessage() {}

func (x *UpdateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notes_v1_notes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return file_protos_notes_v1_notes_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateNoteRequest) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

// Delete
type DeleteNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteNoteRequest) Reset() {
	*x = DeleteNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_notes_v1_notes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteRequest) ProtoMessage() {}

func (x *DeleteNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_notes_v1_notes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteNoteRequest) Descriptor() ([]byte, []int) {
	return file_protos_notes_v1_notes_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteNoteRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_protos_notes_v1_notes_proto protoreflect.FileDescriptor

var file_protos_notes_v1_notes_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73,
	0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x76, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0x6d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x77,
	0x61, 0x73, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x52, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65,
	0x5f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0xbf,
	0x04, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x71, 0x6c, 0x69,
	0x74, 0x65, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x73, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x6f, 0x74, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x77, 0x61,
	0x73, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73,
	0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x6b, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x71, 0x6c,
	0x69, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x73,
	0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x7a, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65,
	0x5f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x73, 0x6d, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x3a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x1a, 0x19, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x6f, 0x74, 0x65, 0x2e,
	0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x69, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x77, 0x61, 0x73,
	0x6d, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49,
	0x45, 0x50, 0x2f, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x2d, 0x77, 0x61, 0x73, 0x6d, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_notes_v1_notes_proto_rawDescOnce sync.Once
	file_protos_notes_v1_notes_proto_rawDescData = file_protos_notes_v1_notes_proto_rawDesc
)

func file_protos_notes_v1_notes_proto_rawDescGZIP() []byte {
	file_protos_notes_v1_notes_proto_rawDescOnce.Do(func() {
		file_protos_notes_v1_notes_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_notes_v1_notes_proto_rawDescData)
	})
	return file_protos_notes_v1_notes_proto_rawDescData
}

var file_protos_notes_v1_notes_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_notes_v1_notes_proto_goTypes = []interface{}{
	(*Note)(nil),              // 0: sqlite_wasm.notes.v1.Note
	(*GetNoteRequest)(nil),    // 1: sqlite_wasm.notes.v1.GetNoteRequest
	(*ListNotesRequest)(nil),  // 2: sqlite_wasm.notes.v1.ListNotesRequest
	(*ListNotesResponse)(nil), // 3: sqlite_wasm.notes.v1.ListNotesResponse
	(*CreateNoteRequest)(nil), // 4: sqlite_wasm.notes.v1.CreateNoteRequest
	(*UpdateNoteRequest)(nil), // 5: sqlite_wasm.notes.v1.UpdateNoteRequest
	(*DeleteNoteRequest)(nil), // 6: sqlite_wasm.notes.v1.DeleteNoteRequest
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_protos_notes_v1_notes_proto_depIdxs = []int32{
	0, // 0: sqlite_wasm.notes.v1.ListNotesResponse.notes:type_name -> sqlite_wasm.notes.v1.Note
	0, // 1: sqlite_wasm.notes.v1.UpdateNoteRequest.note:type_name -> sqlite_wasm.notes.v1.Note
	1, // 2: sqlite_wasm.notes.v1.NoteService.GetNote:input_type -> sqlite_wasm.notes.v1.GetNoteRequest
	2, // 3: sqlite_wasm.notes.v1.NoteService.ListNotes:input_type -> sqlite_wasm.notes.v1.ListNotesRequest
	4, // 4: sqlite_wasm.notes.v1.NoteService.CreateNote:input_type -> sqlite_wasm.notes.v1.CreateNoteRequest
	5, // 5: sqlite_wasm.notes.v1.NoteService.UpdateNote:input_type -> sqlite_wasm.notes.v1.UpdateNoteRequest
	6, // 6: sqlite_wasm.notes.v1.NoteService.DeleteNote:input_type -> sqlite_wasm.notes.v1.DeleteNoteRequest
	0, // 7: sqlite_wasm.notes.v1.NoteService.GetNote:output_type -> sqlite_wasm.notes.v1.Note
	3, // 8: sqlite_wasm.notes.v1.NoteService.ListNotes:output_type -> sqlite_wasm.notes.v1.ListNotesResponse
	0, // 9: sqlite_wasm.notes.v1.NoteService.CreateNote:output_type -> sqlite_wasm.notes.v1.Note
	0, // 10: sqlite_wasm.notes.v1.NoteService.UpdateNote:output_type -> sqlite_wasm.notes.v1.Note
	7, // 11: sqlite_wasm.notes.v1.NoteService.DeleteNote:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_notes_v1_notes_proto_init() }
func file_protos_notes_v1_notes_proto_init() {
	if File_protos_notes_v1_notes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_notes_v1_notes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_protos_notes_v1_notes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNoteRequest); i {
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
		file_protos_notes_v1_notes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotesRequest); i {
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
		file_protos_notes_v1_notes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotesResponse); i {
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
		file_protos_notes_v1_notes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteRequest); i {
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
		file_protos_notes_v1_notes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteRequest); i {
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
		file_protos_notes_v1_notes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNoteRequest); i {
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
	file_protos_notes_v1_notes_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_protos_notes_v1_notes_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_protos_notes_v1_notes_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_notes_v1_notes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_notes_v1_notes_proto_goTypes,
		DependencyIndexes: file_protos_notes_v1_notes_proto_depIdxs,
		MessageInfos:      file_protos_notes_v1_notes_proto_msgTypes,
	}.Build()
	File_protos_notes_v1_notes_proto = out.File
	file_protos_notes_v1_notes_proto_rawDesc = nil
	file_protos_notes_v1_notes_proto_goTypes = nil
	file_protos_notes_v1_notes_proto_depIdxs = nil
}
