//
//     Copyright 2023 The Dragonfly Authors
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
// source: pkg/apis/trainer/v1/trainer.proto

package trainer

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TrainGNNRequest represents to train GNN model request of TrainRequest.
type TrainGNNRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dataset of training GNN.
	Dataset []byte `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
}

func (x *TrainGNNRequest) Reset() {
	*x = TrainGNNRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_trainer_v1_trainer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainGNNRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainGNNRequest) ProtoMessage() {}

func (x *TrainGNNRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_trainer_v1_trainer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainGNNRequest.ProtoReflect.Descriptor instead.
func (*TrainGNNRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_trainer_v1_trainer_proto_rawDescGZIP(), []int{0}
}

func (x *TrainGNNRequest) GetDataset() []byte {
	if x != nil {
		return x.Dataset
	}
	return nil
}

// TrainMLPRequest represents to train MLP model request of TrainRequest.
type TrainMLPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Dataset of training MLP.
	Dataset []byte `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
}

func (x *TrainMLPRequest) Reset() {
	*x = TrainMLPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_trainer_v1_trainer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainMLPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainMLPRequest) ProtoMessage() {}

func (x *TrainMLPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_trainer_v1_trainer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainMLPRequest.ProtoReflect.Descriptor instead.
func (*TrainMLPRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_trainer_v1_trainer_proto_rawDescGZIP(), []int{1}
}

func (x *TrainMLPRequest) GetDataset() []byte {
	if x != nil {
		return x.Dataset
	}
	return nil
}

// TrainRequest represents request of Train.
type TrainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Scheduler hostname.
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Scheduler ip.
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	// Scheduler cluster id.
	ClusterId uint64 `protobuf:"varint,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Types that are assignable to Request:
	//
	//	*TrainRequest_TrainGnnRequest
	//	*TrainRequest_TrainMlpRequest
	Request isTrainRequest_Request `protobuf_oneof:"request"`
}

func (x *TrainRequest) Reset() {
	*x = TrainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_trainer_v1_trainer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainRequest) ProtoMessage() {}

func (x *TrainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_trainer_v1_trainer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainRequest.ProtoReflect.Descriptor instead.
func (*TrainRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_trainer_v1_trainer_proto_rawDescGZIP(), []int{2}
}

func (x *TrainRequest) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *TrainRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *TrainRequest) GetClusterId() uint64 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

func (m *TrainRequest) GetRequest() isTrainRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *TrainRequest) GetTrainGnnRequest() *TrainGNNRequest {
	if x, ok := x.GetRequest().(*TrainRequest_TrainGnnRequest); ok {
		return x.TrainGnnRequest
	}
	return nil
}

func (x *TrainRequest) GetTrainMlpRequest() *TrainMLPRequest {
	if x, ok := x.GetRequest().(*TrainRequest_TrainMlpRequest); ok {
		return x.TrainMlpRequest
	}
	return nil
}

type isTrainRequest_Request interface {
	isTrainRequest_Request()
}

type TrainRequest_TrainGnnRequest struct {
	TrainGnnRequest *TrainGNNRequest `protobuf:"bytes,4,opt,name=train_gnn_request,json=trainGnnRequest,proto3,oneof"`
}

type TrainRequest_TrainMlpRequest struct {
	TrainMlpRequest *TrainMLPRequest `protobuf:"bytes,5,opt,name=train_mlp_request,json=trainMlpRequest,proto3,oneof"`
}

func (*TrainRequest_TrainGnnRequest) isTrainRequest_Request() {}

func (*TrainRequest_TrainMlpRequest) isTrainRequest_Request() {}

var File_pkg_apis_trainer_v1_trainer_proto protoreflect.FileDescriptor

var file_pkg_apis_trainer_v1_trainer_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x47, 0x4e,
	0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x22, 0x34, 0x0a, 0x0f, 0x54,
	0x72, 0x61, 0x69, 0x6e, 0x4d, 0x4c, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x10, 0x01, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65,
	0x74, 0x22, 0x9a, 0x02, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x08, 0x68,
	0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x70, 0x01, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x26, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x49, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x5f, 0x67, 0x6e, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x47, 0x4e, 0x4e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x00, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x47, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x6c, 0x70,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x69,
	0x6e, 0x4d, 0x4c, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0f, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x4d, 0x6c, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x0e,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x32, 0x46,
	0x0a, 0x07, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x05, 0x54, 0x72, 0x61,
	0x69, 0x6e, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x64, 0x37, 0x79, 0x2e, 0x69, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_trainer_v1_trainer_proto_rawDescOnce sync.Once
	file_pkg_apis_trainer_v1_trainer_proto_rawDescData = file_pkg_apis_trainer_v1_trainer_proto_rawDesc
)

func file_pkg_apis_trainer_v1_trainer_proto_rawDescGZIP() []byte {
	file_pkg_apis_trainer_v1_trainer_proto_rawDescOnce.Do(func() {
		file_pkg_apis_trainer_v1_trainer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_trainer_v1_trainer_proto_rawDescData)
	})
	return file_pkg_apis_trainer_v1_trainer_proto_rawDescData
}

var file_pkg_apis_trainer_v1_trainer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_apis_trainer_v1_trainer_proto_goTypes = []interface{}{
	(*TrainGNNRequest)(nil), // 0: trainer.v1.TrainGNNRequest
	(*TrainMLPRequest)(nil), // 1: trainer.v1.TrainMLPRequest
	(*TrainRequest)(nil),    // 2: trainer.v1.TrainRequest
	(*emptypb.Empty)(nil),   // 3: google.protobuf.Empty
}
var file_pkg_apis_trainer_v1_trainer_proto_depIdxs = []int32{
	0, // 0: trainer.v1.TrainRequest.train_gnn_request:type_name -> trainer.v1.TrainGNNRequest
	1, // 1: trainer.v1.TrainRequest.train_mlp_request:type_name -> trainer.v1.TrainMLPRequest
	2, // 2: trainer.v1.Trainer.Train:input_type -> trainer.v1.TrainRequest
	3, // 3: trainer.v1.Trainer.Train:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_apis_trainer_v1_trainer_proto_init() }
func file_pkg_apis_trainer_v1_trainer_proto_init() {
	if File_pkg_apis_trainer_v1_trainer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_trainer_v1_trainer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainGNNRequest); i {
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
		file_pkg_apis_trainer_v1_trainer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainMLPRequest); i {
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
		file_pkg_apis_trainer_v1_trainer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainRequest); i {
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
	file_pkg_apis_trainer_v1_trainer_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TrainRequest_TrainGnnRequest)(nil),
		(*TrainRequest_TrainMlpRequest)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_apis_trainer_v1_trainer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_trainer_v1_trainer_proto_goTypes,
		DependencyIndexes: file_pkg_apis_trainer_v1_trainer_proto_depIdxs,
		MessageInfos:      file_pkg_apis_trainer_v1_trainer_proto_msgTypes,
	}.Build()
	File_pkg_apis_trainer_v1_trainer_proto = out.File
	file_pkg_apis_trainer_v1_trainer_proto_rawDesc = nil
	file_pkg_apis_trainer_v1_trainer_proto_goTypes = nil
	file_pkg_apis_trainer_v1_trainer_proto_depIdxs = nil
}