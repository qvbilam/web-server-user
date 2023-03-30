// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.19.4
// source: file.proto

package publicV1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	v1 "user/api/qvbilam/page/v1"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Channel     string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Sha1        string `protobuf:"bytes,3,opt,name=sha1,proto3" json:"sha1,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	ContentType string `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Size        int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Extra       string `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`
	Callback    string `protobuf:"bytes,8,opt,name=callback,proto3" json:"callback,omitempty"`
	CreatedAt   string `protobuf:"bytes,9,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *FileResponse) Reset() {
	*x = FileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileResponse) ProtoMessage() {}

func (x *FileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileResponse.ProtoReflect.Descriptor instead.
func (*FileResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileResponse) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *FileResponse) GetSha1() string {
	if x != nil {
		return x.Sha1
	}
	return ""
}

func (x *FileResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *FileResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *FileResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileResponse) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *FileResponse) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

func (x *FileResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type FilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64           `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Files []*FileResponse `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *FilesResponse) Reset() {
	*x = FilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesResponse) ProtoMessage() {}

func (x *FilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesResponse.ProtoReflect.Descriptor instead.
func (*FilesResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{1}
}

func (x *FilesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *FilesResponse) GetFiles() []*FileResponse {
	if x != nil {
		return x.Files
	}
	return nil
}

type UpdateFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Channel     string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Sha1        string `protobuf:"bytes,3,opt,name=sha1,proto3" json:"sha1,omitempty"`
	Url         string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	ContentType string `protobuf:"bytes,5,opt,name=contentType,proto3" json:"contentType,omitempty"`
	Size        int64  `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Extra       string `protobuf:"bytes,7,opt,name=extra,proto3" json:"extra,omitempty"`
	Callback    string `protobuf:"bytes,8,opt,name=callback,proto3" json:"callback,omitempty"`
}

func (x *UpdateFileRequest) Reset() {
	*x = UpdateFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileRequest) ProtoMessage() {}

func (x *UpdateFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFileRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFileRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *UpdateFileRequest) GetSha1() string {
	if x != nil {
		return x.Sha1
	}
	return ""
}

func (x *UpdateFileRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateFileRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *UpdateFileRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UpdateFileRequest) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *UpdateFileRequest) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

type FileDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sha1 string `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
}

func (x *FileDetailRequest) Reset() {
	*x = FileDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDetailRequest) ProtoMessage() {}

func (x *FileDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDetailRequest.ProtoReflect.Descriptor instead.
func (*FileDetailRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{3}
}

func (x *FileDetailRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileDetailRequest) GetSha1() string {
	if x != nil {
		return x.Sha1
	}
	return ""
}

type SearchFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sha1  string          `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
	Ids   []int64         `protobuf:"varint,3,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Sha1S []string        `protobuf:"bytes,4,rep,name=sha1s,proto3" json:"sha1s,omitempty"`
	Page  *v1.PageRequest `protobuf:"bytes,5,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchFileRequest) Reset() {
	*x = SearchFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFileRequest) ProtoMessage() {}

func (x *SearchFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFileRequest.ProtoReflect.Descriptor instead.
func (*SearchFileRequest) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{4}
}

func (x *SearchFileRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SearchFileRequest) GetSha1() string {
	if x != nil {
		return x.Sha1
	}
	return ""
}

func (x *SearchFileRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *SearchFileRequest) GetSha1S() []string {
	if x != nil {
		return x.Sha1S
	}
	return nil
}

func (x *SearchFileRequest) GetPage() *v1.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type ExistsFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExists bool          `protobuf:"varint,1,opt,name=isExists,proto3" json:"isExists,omitempty"`
	File     *FileResponse `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *ExistsFileResponse) Reset() {
	*x = ExistsFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsFileResponse) ProtoMessage() {}

func (x *ExistsFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsFileResponse.ProtoReflect.Descriptor instead.
func (*ExistsFileResponse) Descriptor() ([]byte, []int) {
	return file_file_proto_rawDescGZIP(), []int{5}
}

func (x *ExistsFileResponse) GetIsExists() bool {
	if x != nil {
		return x.IsExists
	}
	return false
}

func (x *ExistsFileResponse) GetFile() *FileResponse {
	if x != nil {
		return x.File
	}
	return nil
}

var File_file_proto protoreflect.FileDescriptor

var file_file_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x76, 0x62, 0x69,
	0x6c, 0x61, 0x6d, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x61, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x68, 0x61, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x56, 0x0a,
	0x0d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x61, 0x31, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x61, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x22, 0x37, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x61, 0x31,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x68, 0x61, 0x31, 0x22, 0x8b, 0x01, 0x0a,
	0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x61, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x68, 0x61, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x31,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x61, 0x31, 0x73, 0x12, 0x2a,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x61, 0x67, 0x65, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x5f, 0x0a, 0x12, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xa5, 0x03, 0x0a, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x71, 0x76, 0x62, 0x69, 0x6c, 0x61, 0x6d, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_file_proto_rawDescOnce sync.Once
	file_file_proto_rawDescData = file_file_proto_rawDesc
)

func file_file_proto_rawDescGZIP() []byte {
	file_file_proto_rawDescOnce.Do(func() {
		file_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_proto_rawDescData)
	})
	return file_file_proto_rawDescData
}

var file_file_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_file_proto_goTypes = []interface{}{
	(*FileResponse)(nil),       // 0: publicPb.v1.FileResponse
	(*FilesResponse)(nil),      // 1: publicPb.v1.FilesResponse
	(*UpdateFileRequest)(nil),  // 2: publicPb.v1.UpdateFileRequest
	(*FileDetailRequest)(nil),  // 3: publicPb.v1.FileDetailRequest
	(*SearchFileRequest)(nil),  // 4: publicPb.v1.SearchFileRequest
	(*ExistsFileResponse)(nil), // 5: publicPb.v1.ExistsFileResponse
	(*v1.PageRequest)(nil),     // 6: pagePb.v1.PageRequest
	(*emptypb.Empty)(nil),      // 7: google.protobuf.Empty
}
var file_file_proto_depIdxs = []int32{
	0, // 0: publicPb.v1.FilesResponse.files:type_name -> publicPb.v1.FileResponse
	6, // 1: publicPb.v1.SearchFileRequest.page:type_name -> pagePb.v1.PageRequest
	0, // 2: publicPb.v1.ExistsFileResponse.file:type_name -> publicPb.v1.FileResponse
	2, // 3: publicPb.v1.File.Create:input_type -> publicPb.v1.UpdateFileRequest
	2, // 4: publicPb.v1.File.Update:input_type -> publicPb.v1.UpdateFileRequest
	2, // 5: publicPb.v1.File.Delete:input_type -> publicPb.v1.UpdateFileRequest
	4, // 6: publicPb.v1.File.Get:input_type -> publicPb.v1.SearchFileRequest
	3, // 7: publicPb.v1.File.GetDetail:input_type -> publicPb.v1.FileDetailRequest
	3, // 8: publicPb.v1.File.Exists:input_type -> publicPb.v1.FileDetailRequest
	0, // 9: publicPb.v1.File.Create:output_type -> publicPb.v1.FileResponse
	7, // 10: publicPb.v1.File.Update:output_type -> google.protobuf.Empty
	7, // 11: publicPb.v1.File.Delete:output_type -> google.protobuf.Empty
	1, // 12: publicPb.v1.File.Get:output_type -> publicPb.v1.FilesResponse
	0, // 13: publicPb.v1.File.GetDetail:output_type -> publicPb.v1.FileResponse
	5, // 14: publicPb.v1.File.Exists:output_type -> publicPb.v1.ExistsFileResponse
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_file_proto_init() }
func file_file_proto_init() {
	if File_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileResponse); i {
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
		file_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesResponse); i {
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
		file_file_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileRequest); i {
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
		file_file_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDetailRequest); i {
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
		file_file_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFileRequest); i {
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
		file_file_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistsFileResponse); i {
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
			RawDescriptor: file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_proto_goTypes,
		DependencyIndexes: file_file_proto_depIdxs,
		MessageInfos:      file_file_proto_msgTypes,
	}.Build()
	File_file_proto = out.File
	file_file_proto_rawDesc = nil
	file_file_proto_goTypes = nil
	file_file_proto_depIdxs = nil
}
