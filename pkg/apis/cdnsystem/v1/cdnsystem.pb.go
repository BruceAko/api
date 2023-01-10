//
//     Copyright 2022 The Dragonfly Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: pkg/apis/cdnsystem/v1/cdnsystem.proto

package cdnsystem

import (
	v1 "d7y.io/api/pkg/apis/common/v1"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type SeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId  string      `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Url     string      `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	UrlMeta *v1.UrlMeta `protobuf:"bytes,3,opt,name=url_meta,json=urlMeta,proto3" json:"url_meta,omitempty"`
}

func (x *SeedRequest) Reset() {
	*x = SeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_cdnsystem_v1_cdnsystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeedRequest) ProtoMessage() {}

func (x *SeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_cdnsystem_v1_cdnsystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeedRequest.ProtoReflect.Descriptor instead.
func (*SeedRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescGZIP(), []int{0}
}

func (x *SeedRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *SeedRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SeedRequest) GetUrlMeta() *v1.UrlMeta {
	if x != nil {
		return x.UrlMeta
	}
	return nil
}

// keep piece meta and data separately
// check piece md5, md5s sign and total content length
type PieceSeed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reuse already downloaded peer
	Reuse bool `protobuf:"varint,1,opt,name=reuse,proto3" json:"reuse,omitempty"`
	// peer id for cdn node, need suffix with _CDN
	PeerId string `protobuf:"bytes,2,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	// cdn host id
	HostId string `protobuf:"bytes,3,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	// piece info
	PieceInfo *v1.PieceInfo `protobuf:"bytes,4,opt,name=piece_info,json=pieceInfo,proto3" json:"piece_info,omitempty"`
	// whether or not all seeds are downloaded
	Done bool `protobuf:"varint,5,opt,name=done,proto3" json:"done,omitempty"`
	// content total length for the url, content_length < 0 represent content length is unknown
	ContentLength int64 `protobuf:"varint,6,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	// total piece count, -1 represents task is downloading or failed
	TotalPieceCount int32 `protobuf:"varint,7,opt,name=total_piece_count,json=totalPieceCount,proto3" json:"total_piece_count,omitempty"`
	// begin time for the piece downloading
	BeginTime uint64 `protobuf:"varint,8,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	// end time for the piece downloading
	EndTime uint64 `protobuf:"varint,9,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// task extend attribute
	ExtendAttribute *v1.ExtendAttribute `protobuf:"bytes,10,opt,name=extend_attribute,json=extendAttribute,proto3" json:"extend_attribute,omitempty"`
}

func (x *PieceSeed) Reset() {
	*x = PieceSeed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_cdnsystem_v1_cdnsystem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PieceSeed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PieceSeed) ProtoMessage() {}

func (x *PieceSeed) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_cdnsystem_v1_cdnsystem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PieceSeed.ProtoReflect.Descriptor instead.
func (*PieceSeed) Descriptor() ([]byte, []int) {
	return file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescGZIP(), []int{1}
}

func (x *PieceSeed) GetReuse() bool {
	if x != nil {
		return x.Reuse
	}
	return false
}

func (x *PieceSeed) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *PieceSeed) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *PieceSeed) GetPieceInfo() *v1.PieceInfo {
	if x != nil {
		return x.PieceInfo
	}
	return nil
}

func (x *PieceSeed) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *PieceSeed) GetContentLength() int64 {
	if x != nil {
		return x.ContentLength
	}
	return 0
}

func (x *PieceSeed) GetTotalPieceCount() int32 {
	if x != nil {
		return x.TotalPieceCount
	}
	return 0
}

func (x *PieceSeed) GetBeginTime() uint64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *PieceSeed) GetEndTime() uint64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *PieceSeed) GetExtendAttribute() *v1.ExtendAttribute {
	if x != nil {
		return x.ExtendAttribute
	}
	return nil
}

var File_pkg_apis_cdnsystem_v1_cdnsystem_proto protoreflect.FileDescriptor

var file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x64, 0x6e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x64, 0x6e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x64, 0x6e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7a, 0x0a, 0x0b, 0x53, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0x88, 0x01, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x2d, 0x0a, 0x08,
	0x75, 0x72, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x72, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x07, 0x75, 0x72, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x22, 0x82, 0x03, 0x0a, 0x09,
	0x50, 0x69, 0x65, 0x63, 0x65, 0x53, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x75,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x75, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x68, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0a, 0x70, 0x69, 0x65, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70,
	0x69, 0x65, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x69, 0x65,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x69, 0x65, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52,
	0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x32, 0xde, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x65, 0x64, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0b, 0x4f,
	0x62, 0x74, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x65, 0x64, 0x73, 0x12, 0x19, 0x2e, 0x63, 0x64, 0x6e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x64, 0x6e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x53, 0x65, 0x65, 0x64, 0x30, 0x01,
	0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x69, 0x65, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69,
	0x65, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x49, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x69,
	0x65, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x69, 0x65, 0x63, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x64, 0x37, 0x79, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x64, 0x6e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x64, 0x6e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescOnce sync.Once
	file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescData = file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDesc
)

func file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescGZIP() []byte {
	file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescOnce.Do(func() {
		file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescData)
	})
	return file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDescData
}

var file_pkg_apis_cdnsystem_v1_cdnsystem_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_apis_cdnsystem_v1_cdnsystem_proto_goTypes = []interface{}{
	(*SeedRequest)(nil),         // 0: cdnsystem.v1.SeedRequest
	(*PieceSeed)(nil),           // 1: cdnsystem.v1.PieceSeed
	(*v1.UrlMeta)(nil),          // 2: common.v1.UrlMeta
	(*v1.PieceInfo)(nil),        // 3: common.v1.PieceInfo
	(*v1.ExtendAttribute)(nil),  // 4: common.v1.ExtendAttribute
	(*v1.PieceTaskRequest)(nil), // 5: common.v1.PieceTaskRequest
	(*v1.PiecePacket)(nil),      // 6: common.v1.PiecePacket
}
var file_pkg_apis_cdnsystem_v1_cdnsystem_proto_depIdxs = []int32{
	2, // 0: cdnsystem.v1.SeedRequest.url_meta:type_name -> common.v1.UrlMeta
	3, // 1: cdnsystem.v1.PieceSeed.piece_info:type_name -> common.v1.PieceInfo
	4, // 2: cdnsystem.v1.PieceSeed.extend_attribute:type_name -> common.v1.ExtendAttribute
	0, // 3: cdnsystem.v1.Seeder.ObtainSeeds:input_type -> cdnsystem.v1.SeedRequest
	5, // 4: cdnsystem.v1.Seeder.GetPieceTasks:input_type -> common.v1.PieceTaskRequest
	5, // 5: cdnsystem.v1.Seeder.SyncPieceTasks:input_type -> common.v1.PieceTaskRequest
	1, // 6: cdnsystem.v1.Seeder.ObtainSeeds:output_type -> cdnsystem.v1.PieceSeed
	6, // 7: cdnsystem.v1.Seeder.GetPieceTasks:output_type -> common.v1.PiecePacket
	6, // 8: cdnsystem.v1.Seeder.SyncPieceTasks:output_type -> common.v1.PiecePacket
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_apis_cdnsystem_v1_cdnsystem_proto_init() }
func file_pkg_apis_cdnsystem_v1_cdnsystem_proto_init() {
	if File_pkg_apis_cdnsystem_v1_cdnsystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_cdnsystem_v1_cdnsystem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeedRequest); i {
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
		file_pkg_apis_cdnsystem_v1_cdnsystem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PieceSeed); i {
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
			RawDescriptor: file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_cdnsystem_v1_cdnsystem_proto_goTypes,
		DependencyIndexes: file_pkg_apis_cdnsystem_v1_cdnsystem_proto_depIdxs,
		MessageInfos:      file_pkg_apis_cdnsystem_v1_cdnsystem_proto_msgTypes,
	}.Build()
	File_pkg_apis_cdnsystem_v1_cdnsystem_proto = out.File
	file_pkg_apis_cdnsystem_v1_cdnsystem_proto_rawDesc = nil
	file_pkg_apis_cdnsystem_v1_cdnsystem_proto_goTypes = nil
	file_pkg_apis_cdnsystem_v1_cdnsystem_proto_depIdxs = nil
}
