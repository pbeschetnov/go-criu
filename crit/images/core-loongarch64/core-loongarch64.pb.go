// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: core-loongarch64.proto

package core_loongarch64

import (
	_ "github.com/checkpoint-restore/go-criu/v7/crit/images/opts"
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

type UserLoongarch64GpregsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regs []uint64 `protobuf:"varint,1,rep,name=regs" json:"regs,omitempty"`
	Pc   *uint64  `protobuf:"varint,2,req,name=pc" json:"pc,omitempty"`
}

func (x *UserLoongarch64GpregsEntry) Reset() {
	*x = UserLoongarch64GpregsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_loongarch64_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoongarch64GpregsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoongarch64GpregsEntry) ProtoMessage() {}

func (x *UserLoongarch64GpregsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_loongarch64_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoongarch64GpregsEntry.ProtoReflect.Descriptor instead.
func (*UserLoongarch64GpregsEntry) Descriptor() ([]byte, []int) {
	return file_core_loongarch64_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoongarch64GpregsEntry) GetRegs() []uint64 {
	if x != nil {
		return x.Regs
	}
	return nil
}

func (x *UserLoongarch64GpregsEntry) GetPc() uint64 {
	if x != nil && x.Pc != nil {
		return *x.Pc
	}
	return 0
}

type UserLoongarch64FpregsEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regs []uint64 `protobuf:"varint,1,rep,name=regs" json:"regs,omitempty"`
	Fcc  *uint64  `protobuf:"varint,2,req,name=fcc" json:"fcc,omitempty"`
	Fcsr *uint32  `protobuf:"varint,3,req,name=fcsr" json:"fcsr,omitempty"`
}

func (x *UserLoongarch64FpregsEntry) Reset() {
	*x = UserLoongarch64FpregsEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_loongarch64_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoongarch64FpregsEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoongarch64FpregsEntry) ProtoMessage() {}

func (x *UserLoongarch64FpregsEntry) ProtoReflect() protoreflect.Message {
	mi := &file_core_loongarch64_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoongarch64FpregsEntry.ProtoReflect.Descriptor instead.
func (*UserLoongarch64FpregsEntry) Descriptor() ([]byte, []int) {
	return file_core_loongarch64_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoongarch64FpregsEntry) GetRegs() []uint64 {
	if x != nil {
		return x.Regs
	}
	return nil
}

func (x *UserLoongarch64FpregsEntry) GetFcc() uint64 {
	if x != nil && x.Fcc != nil {
		return *x.Fcc
	}
	return 0
}

func (x *UserLoongarch64FpregsEntry) GetFcsr() uint32 {
	if x != nil && x.Fcsr != nil {
		return *x.Fcsr
	}
	return 0
}

type ThreadInfoLoongarch64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClearTidAddr *uint64                     `protobuf:"varint,1,req,name=clear_tid_addr,json=clearTidAddr" json:"clear_tid_addr,omitempty"`
	Tls          *uint64                     `protobuf:"varint,2,req,name=tls" json:"tls,omitempty"`
	Gpregs       *UserLoongarch64GpregsEntry `protobuf:"bytes,3,req,name=gpregs" json:"gpregs,omitempty"`
	Fpregs       *UserLoongarch64FpregsEntry `protobuf:"bytes,4,req,name=fpregs" json:"fpregs,omitempty"`
}

func (x *ThreadInfoLoongarch64) Reset() {
	*x = ThreadInfoLoongarch64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_loongarch64_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadInfoLoongarch64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadInfoLoongarch64) ProtoMessage() {}

func (x *ThreadInfoLoongarch64) ProtoReflect() protoreflect.Message {
	mi := &file_core_loongarch64_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadInfoLoongarch64.ProtoReflect.Descriptor instead.
func (*ThreadInfoLoongarch64) Descriptor() ([]byte, []int) {
	return file_core_loongarch64_proto_rawDescGZIP(), []int{2}
}

func (x *ThreadInfoLoongarch64) GetClearTidAddr() uint64 {
	if x != nil && x.ClearTidAddr != nil {
		return *x.ClearTidAddr
	}
	return 0
}

func (x *ThreadInfoLoongarch64) GetTls() uint64 {
	if x != nil && x.Tls != nil {
		return *x.Tls
	}
	return 0
}

func (x *ThreadInfoLoongarch64) GetGpregs() *UserLoongarch64GpregsEntry {
	if x != nil {
		return x.Gpregs
	}
	return nil
}

func (x *ThreadInfoLoongarch64) GetFpregs() *UserLoongarch64FpregsEntry {
	if x != nil {
		return x.Fpregs
	}
	return nil
}

var File_core_loongarch64_proto protoreflect.FileDescriptor

var file_core_loongarch64_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x6c, 0x6f, 0x6f, 0x6e, 0x67, 0x61, 0x72, 0x63, 0x68,
	0x36, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x6f,
	0x6e, 0x67, 0x61, 0x72, 0x63, 0x68, 0x36, 0x34, 0x5f, 0x67, 0x70, 0x72, 0x65, 0x67, 0x73, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x04, 0x72, 0x65, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x63, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x02, 0x70, 0x63, 0x22, 0x59, 0x0a, 0x1d, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x6f, 0x6f, 0x6e, 0x67, 0x61, 0x72, 0x63, 0x68, 0x36, 0x34, 0x5f, 0x66, 0x70,
	0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x72, 0x65, 0x67, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x66, 0x63, 0x63, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x03, 0x66, 0x63, 0x63,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x63, 0x73, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x04,
	0x66, 0x63, 0x73, 0x72, 0x22, 0xd6, 0x01, 0x0a, 0x17, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x6f, 0x6f, 0x6e, 0x67, 0x61, 0x72, 0x63, 0x68, 0x36, 0x34,
	0x12, 0x2b, 0x0a, 0x0e, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x5f, 0x74, 0x69, 0x64, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52,
	0x0c, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x54, 0x69, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x04, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12,
	0x3d, 0x0a, 0x06, 0x67, 0x70, 0x72, 0x65, 0x67, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x6f, 0x6e, 0x67, 0x61, 0x72, 0x63, 0x68,
	0x36, 0x34, 0x5f, 0x67, 0x70, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x06, 0x67, 0x70, 0x72, 0x65, 0x67, 0x73, 0x12, 0x3d,
	0x0a, 0x06, 0x66, 0x70, 0x72, 0x65, 0x67, 0x73, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x6f, 0x6e, 0x67, 0x61, 0x72, 0x63, 0x68, 0x36,
	0x34, 0x5f, 0x66, 0x70, 0x72, 0x65, 0x67, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x05,
	0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x06, 0x66, 0x70, 0x72, 0x65, 0x67, 0x73,
}

var (
	file_core_loongarch64_proto_rawDescOnce sync.Once
	file_core_loongarch64_proto_rawDescData = file_core_loongarch64_proto_rawDesc
)

func file_core_loongarch64_proto_rawDescGZIP() []byte {
	file_core_loongarch64_proto_rawDescOnce.Do(func() {
		file_core_loongarch64_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_loongarch64_proto_rawDescData)
	})
	return file_core_loongarch64_proto_rawDescData
}

var file_core_loongarch64_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_loongarch64_proto_goTypes = []interface{}{
	(*UserLoongarch64GpregsEntry)(nil), // 0: user_loongarch64_gpregs_entry
	(*UserLoongarch64FpregsEntry)(nil), // 1: user_loongarch64_fpregs_entry
	(*ThreadInfoLoongarch64)(nil),      // 2: thread_info_loongarch64
}
var file_core_loongarch64_proto_depIdxs = []int32{
	0, // 0: thread_info_loongarch64.gpregs:type_name -> user_loongarch64_gpregs_entry
	1, // 1: thread_info_loongarch64.fpregs:type_name -> user_loongarch64_fpregs_entry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_core_loongarch64_proto_init() }
func file_core_loongarch64_proto_init() {
	if File_core_loongarch64_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_loongarch64_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoongarch64GpregsEntry); i {
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
		file_core_loongarch64_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoongarch64FpregsEntry); i {
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
		file_core_loongarch64_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThreadInfoLoongarch64); i {
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
			RawDescriptor: file_core_loongarch64_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_loongarch64_proto_goTypes,
		DependencyIndexes: file_core_loongarch64_proto_depIdxs,
		MessageInfos:      file_core_loongarch64_proto_msgTypes,
	}.Build()
	File_core_loongarch64_proto = out.File
	file_core_loongarch64_proto_rawDesc = nil
	file_core_loongarch64_proto_goTypes = nil
	file_core_loongarch64_proto_depIdxs = nil
}
