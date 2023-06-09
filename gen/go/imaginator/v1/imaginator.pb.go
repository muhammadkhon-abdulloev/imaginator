// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/imaginator.proto

package v1

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

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imaginator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imaginator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_imaginator_proto_rawDescGZIP(), []int{0}
}

func (x *UploadFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadFileRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Filesize int64  `protobuf:"varint,2,opt,name=filesize,proto3" json:"filesize,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imaginator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imaginator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_imaginator_proto_rawDescGZIP(), []int{1}
}

func (x *UploadFileResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UploadFileResponse) GetFilesize() int64 {
	if x != nil {
		return x.Filesize
	}
	return 0
}

type DownloadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename      string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	ChunkBuffSize uint64 `protobuf:"varint,2,opt,name=chunkBuffSize,proto3" json:"chunkBuffSize,omitempty"`
}

func (x *DownloadFileRequest) Reset() {
	*x = DownloadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imaginator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileRequest) ProtoMessage() {}

func (x *DownloadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imaginator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_imaginator_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DownloadFileRequest) GetChunkBuffSize() uint64 {
	if x != nil {
		return x.ChunkBuffSize
	}
	return 0
}

type DownloadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Offset   int64  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Checksum string `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *DownloadFileResponse) Reset() {
	*x = DownloadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imaginator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileResponse) ProtoMessage() {}

func (x *DownloadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imaginator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_imaginator_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadFileResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DownloadFileResponse) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *DownloadFileResponse) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

type Files struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename  string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Files) Reset() {
	*x = Files{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imaginator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Files) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Files) ProtoMessage() {}

func (x *Files) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imaginator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Files.ProtoReflect.Descriptor instead.
func (*Files) Descriptor() ([]byte, []int) {
	return file_proto_imaginator_proto_rawDescGZIP(), []int{4}
}

func (x *Files) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Files) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Files) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type ListAllFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListAllFilesRequest) Reset() {
	*x = ListAllFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imaginator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllFilesRequest) ProtoMessage() {}

func (x *ListAllFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imaginator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllFilesRequest.ProtoReflect.Descriptor instead.
func (*ListAllFilesRequest) Descriptor() ([]byte, []int) {
	return file_proto_imaginator_proto_rawDescGZIP(), []int{5}
}

func (x *ListAllFilesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListAllFilesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListAllFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files    []*Files `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	TotalLen int64    `protobuf:"varint,2,opt,name=totalLen,proto3" json:"totalLen,omitempty"`
}

func (x *ListAllFilesResponse) Reset() {
	*x = ListAllFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_imaginator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllFilesResponse) ProtoMessage() {}

func (x *ListAllFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_imaginator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllFilesResponse.ProtoReflect.Descriptor instead.
func (*ListAllFilesResponse) Descriptor() ([]byte, []int) {
	return file_proto_imaginator_proto_rawDescGZIP(), []int{6}
}

func (x *ListAllFilesResponse) GetFiles() []*Files {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *ListAllFilesResponse) GetTotalLen() int64 {
	if x != nil {
		return x.TotalLen
	}
	return 0
}

var File_proto_imaginator_proto protoreflect.FileDescriptor

var file_proto_imaginator_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x43, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x12,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x57, 0x0a, 0x13, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x75, 0x66, 0x66, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x75, 0x66, 0x66, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x5e, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x73, 0x75, 0x6d, 0x22, 0x5f, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x5e, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x32, 0xd9, 0x03, 0x0a, 0x0a, 0x49, 0x6d,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x53, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x22, 0x2e,
	0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x42, 0x79, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x20, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x12, 0x62, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x42, 0x79, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x22, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69,
	0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_imaginator_proto_rawDescOnce sync.Once
	file_proto_imaginator_proto_rawDescData = file_proto_imaginator_proto_rawDesc
)

func file_proto_imaginator_proto_rawDescGZIP() []byte {
	file_proto_imaginator_proto_rawDescOnce.Do(func() {
		file_proto_imaginator_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_imaginator_proto_rawDescData)
	})
	return file_proto_imaginator_proto_rawDescData
}

var file_proto_imaginator_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_imaginator_proto_goTypes = []interface{}{
	(*UploadFileRequest)(nil),    // 0: imaginator.v1.UploadFileRequest
	(*UploadFileResponse)(nil),   // 1: imaginator.v1.UploadFileResponse
	(*DownloadFileRequest)(nil),  // 2: imaginator.v1.DownloadFileRequest
	(*DownloadFileResponse)(nil), // 3: imaginator.v1.DownloadFileResponse
	(*Files)(nil),                // 4: imaginator.v1.Files
	(*ListAllFilesRequest)(nil),  // 5: imaginator.v1.ListAllFilesRequest
	(*ListAllFilesResponse)(nil), // 6: imaginator.v1.ListAllFilesResponse
}
var file_proto_imaginator_proto_depIdxs = []int32{
	4, // 0: imaginator.v1.ListAllFilesResponse.files:type_name -> imaginator.v1.Files
	0, // 1: imaginator.v1.Imaginator.UploadFile:input_type -> imaginator.v1.UploadFileRequest
	2, // 2: imaginator.v1.Imaginator.DownloadFile:input_type -> imaginator.v1.DownloadFileRequest
	5, // 3: imaginator.v1.Imaginator.ListAllFiles:input_type -> imaginator.v1.ListAllFilesRequest
	0, // 4: imaginator.v1.Imaginator.UploadFileByChunk:input_type -> imaginator.v1.UploadFileRequest
	2, // 5: imaginator.v1.Imaginator.DownloadFileByChunk:input_type -> imaginator.v1.DownloadFileRequest
	1, // 6: imaginator.v1.Imaginator.UploadFile:output_type -> imaginator.v1.UploadFileResponse
	3, // 7: imaginator.v1.Imaginator.DownloadFile:output_type -> imaginator.v1.DownloadFileResponse
	6, // 8: imaginator.v1.Imaginator.ListAllFiles:output_type -> imaginator.v1.ListAllFilesResponse
	1, // 9: imaginator.v1.Imaginator.UploadFileByChunk:output_type -> imaginator.v1.UploadFileResponse
	3, // 10: imaginator.v1.Imaginator.DownloadFileByChunk:output_type -> imaginator.v1.DownloadFileResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_imaginator_proto_init() }
func file_proto_imaginator_proto_init() {
	if File_proto_imaginator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_imaginator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
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
		file_proto_imaginator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
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
		file_proto_imaginator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileRequest); i {
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
		file_proto_imaginator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileResponse); i {
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
		file_proto_imaginator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Files); i {
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
		file_proto_imaginator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllFilesRequest); i {
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
		file_proto_imaginator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAllFilesResponse); i {
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
			RawDescriptor: file_proto_imaginator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_imaginator_proto_goTypes,
		DependencyIndexes: file_proto_imaginator_proto_depIdxs,
		MessageInfos:      file_proto_imaginator_proto_msgTypes,
	}.Build()
	File_proto_imaginator_proto = out.File
	file_proto_imaginator_proto_rawDesc = nil
	file_proto_imaginator_proto_goTypes = nil
	file_proto_imaginator_proto_depIdxs = nil
}
