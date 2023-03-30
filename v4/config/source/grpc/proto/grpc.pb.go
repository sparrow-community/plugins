// protoc --proto_path=. --micro_out=. --go_out=. grpc.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: grpc.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChangeSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Checksum  string `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Format    string `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	Source    string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ChangeSet) Reset() {
	*x = ChangeSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeSet) ProtoMessage() {}

func (x *ChangeSet) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeSet.ProtoReflect.Descriptor instead.
func (*ChangeSet) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeSet) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChangeSet) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

func (x *ChangeSet) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *ChangeSet) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ChangeSet) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangeSet *ChangeSet `protobuf:"bytes,1,opt,name=change_set,json=changeSet,proto3" json:"change_set,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *ReadResponse) GetChangeSet() *ChangeSet {
	if x != nil {
		return x.ChangeSet
	}
	return nil
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string     `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	ChangeSet *ChangeSet `protobuf:"bytes,2,opt,name=change_set,json=changeSet,proto3" json:"change_set,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *WriteRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *WriteRequest) GetChangeSet() *ChangeSet {
	if x != nil {
		return x.ChangeSet
	}
	return nil
}

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *WatchRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangeSet *ChangeSet `protobuf:"bytes,1,opt,name=change_set,json=changeSet,proto3" json:"change_set,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *WatchResponse) GetChangeSet() *ChangeSet {
	if x != nil {
		return x.ChangeSet
	}
	return nil
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72,
	0x70, 0x63, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x21,
	0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x22, 0x3e, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65,
	0x74, 0x22, 0x52, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x65, 0x74, 0x22, 0x22, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x3f, 0x0a, 0x0d, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x32, 0xaa, 0x01, 0x0a, 0x06, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x11, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_proto_goTypes = []interface{}{
	(*ChangeSet)(nil),            // 0: grpc.ChangeSet
	(*ReadRequest)(nil),          // 1: grpc.ReadRequest
	(*ReadResponse)(nil),         // 2: grpc.ReadResponse
	(*WriteRequest)(nil),         // 3: grpc.WriteRequest
	(*WatchRequest)(nil),         // 4: grpc.WatchRequest
	(*WatchResponse)(nil),        // 5: grpc.WatchResponse
	(*wrapperspb.BoolValue)(nil), // 6: google.protobuf.BoolValue
}
var file_grpc_proto_depIdxs = []int32{
	0, // 0: grpc.ReadResponse.change_set:type_name -> grpc.ChangeSet
	0, // 1: grpc.WriteRequest.change_set:type_name -> grpc.ChangeSet
	0, // 2: grpc.WatchResponse.change_set:type_name -> grpc.ChangeSet
	1, // 3: grpc.Source.Read:input_type -> grpc.ReadRequest
	3, // 4: grpc.Source.Write:input_type -> grpc.WriteRequest
	4, // 5: grpc.Source.Watch:input_type -> grpc.WatchRequest
	2, // 6: grpc.Source.Read:output_type -> grpc.ReadResponse
	6, // 7: grpc.Source.Write:output_type -> google.protobuf.BoolValue
	5, // 8: grpc.Source.Watch:output_type -> grpc.WatchResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeSet); i {
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
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
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
		file_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}