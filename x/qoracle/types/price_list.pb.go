// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qoracle/price_list.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("qoracle/price_list.proto", fileDescriptor_735f6ba1ec1eb2a7) }

var fileDescriptor_735f6ba1ec1eb2a7 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x31, 0x8e, 0xc2, 0x30,
	0x10, 0x45, 0x93, 0x66, 0x8b, 0x2d, 0x57, 0x5b, 0xac, 0x22, 0xad, 0xf7, 0x06, 0x9e, 0x8d, 0xb8,
	0x01, 0x25, 0x47, 0xa0, 0x41, 0xb6, 0x31, 0xc6, 0x52, 0x9c, 0x71, 0x3c, 0x13, 0x04, 0xb7, 0xe0,
	0x58, 0x94, 0x29, 0x29, 0x51, 0x72, 0x11, 0x44, 0x12, 0x94, 0xee, 0x4b, 0xf3, 0xe6, 0xeb, 0xfd,
	0xcf, 0x9f, 0x06, 0x93, 0x32, 0x95, 0x85, 0x98, 0xbc, 0xb1, 0xbb, 0xca, 0x13, 0xcb, 0x98, 0x90,
	0xf1, 0xeb, 0xb7, 0x69, 0x15, 0xa9, 0x54, 0x29, 0x4d, 0x72, 0x8a, 0x35, 0xee, 0xad, 0x9c, 0xf9,
	0xe2, 0xdb, 0xa1, 0xc3, 0x91, 0x84, 0x57, 0x9a, 0x9e, 0x8a, 0x3f, 0x87, 0xe8, 0xc6, 0x36, 0x64,
	0xd4, 0xed, 0x01, 0xd8, 0x07, 0x4b, 0xac, 0x42, 0x9c, 0x01, 0x61, 0x90, 0x02, 0x12, 0x68, 0x45,
	0x16, 0x4e, 0xa5, 0xb6, 0xac, 0x4a, 0x30, 0xe8, 0xeb, 0xe9, 0xbe, 0xde, 0xdc, 0x7a, 0x91, 0x77,
	0xbd, 0xc8, 0x1f, 0xbd, 0xc8, 0xaf, 0x83, 0xc8, 0xba, 0x41, 0x64, 0xf7, 0x41, 0x64, 0xdb, 0x7f,
	0xe7, 0xf9, 0xd8, 0x6a, 0x69, 0x30, 0xc0, 0xa2, 0x06, 0x8b, 0x1a, 0x9c, 0xe1, 0x3d, 0x86, 0x2f,
	0xd1, 0x92, 0xfe, 0x18, 0x2b, 0x57, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xff, 0xc6, 0x67,
	0xe4, 0x00, 0x00, 0x00,
}