// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: app/ntp/config.proto

package ntp

import (
	net "github.com/v2fly/v2ray-core/v4/common/net"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ntp server
	Address      *net.Endpoint `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	SyncInterval uint32        `protobuf:"varint,2,opt,name=sync_interval,json=syncInterval,proto3" json:"sync_interval,omitempty"`
	InboundTag   string        `protobuf:"bytes,3,opt,name=inbound_tag,json=inboundTag,proto3" json:"inbound_tag,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ntp_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_ntp_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_ntp_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAddress() *net.Endpoint {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Config) GetSyncInterval() uint32 {
	if x != nil {
		return x.SyncInterval
	}
	return 0
}

func (x *Config) GetInboundTag() string {
	if x != nil {
		return x.InboundTag
	}
	return ""
}

var File_app_ntp_config_proto protoreflect.FileDescriptor

var file_app_ntp_config_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6e, 0x74, 0x70, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x54, 0x61, 0x67, 0x42, 0x57, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x6e, 0x74, 0x70, 0x50, 0x01,
	0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66,
	0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6e, 0x74, 0x70, 0xaa, 0x02, 0x12, 0x56, 0x32, 0x52, 0x61, 0x79,
	0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x4e, 0x54, 0x50, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_ntp_config_proto_rawDescOnce sync.Once
	file_app_ntp_config_proto_rawDescData = file_app_ntp_config_proto_rawDesc
)

func file_app_ntp_config_proto_rawDescGZIP() []byte {
	file_app_ntp_config_proto_rawDescOnce.Do(func() {
		file_app_ntp_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_ntp_config_proto_rawDescData)
	})
	return file_app_ntp_config_proto_rawDescData
}

var file_app_ntp_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_ntp_config_proto_goTypes = []interface{}{
	(*Config)(nil),       // 0: v2ray.core.app.ntp.Config
	(*net.Endpoint)(nil), // 1: v2ray.core.common.net.Endpoint
}
var file_app_ntp_config_proto_depIdxs = []int32{
	1, // 0: v2ray.core.app.ntp.Config.address:type_name -> v2ray.core.common.net.Endpoint
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_ntp_config_proto_init() }
func file_app_ntp_config_proto_init() {
	if File_app_ntp_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_ntp_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_app_ntp_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_ntp_config_proto_goTypes,
		DependencyIndexes: file_app_ntp_config_proto_depIdxs,
		MessageInfos:      file_app_ntp_config_proto_msgTypes,
	}.Build()
	File_app_ntp_config_proto = out.File
	file_app_ntp_config_proto_rawDesc = nil
	file_app_ntp_config_proto_goTypes = nil
	file_app_ntp_config_proto_depIdxs = nil
}