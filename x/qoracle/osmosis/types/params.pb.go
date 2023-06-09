// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qoracle/osmosis/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	types "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Params struct {
	Enabled                bool         `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty" yaml:"enabled"`
	EpochIdentifier        string       `protobuf:"bytes,2,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty" yaml:"epoch_identifier"`
	AuthorizedChannel      string       `protobuf:"bytes,3,opt,name=authorized_channel,json=authorizedChannel,proto3" json:"authorized_channel,omitempty" yaml:"authorized_channel"`
	PacketTimeoutHeight    types.Height `protobuf:"bytes,4,opt,name=packet_timeout_height,json=packetTimeoutHeight,proto3" json:"packet_timeout_height" yaml:"packet_timeout_height"`
	PacketTimeoutTimestamp uint64       `protobuf:"varint,5,opt,name=packet_timeout_timestamp,json=packetTimeoutTimestamp,proto3" json:"packet_timeout_timestamp,omitempty" yaml:"packet_timeout_timestamp"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf7c77ca07c7bae6, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Params) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *Params) GetAuthorizedChannel() string {
	if m != nil {
		return m.AuthorizedChannel
	}
	return ""
}

func (m *Params) GetPacketTimeoutHeight() types.Height {
	if m != nil {
		return m.PacketTimeoutHeight
	}
	return types.Height{}
}

func (m *Params) GetPacketTimeoutTimestamp() uint64 {
	if m != nil {
		return m.PacketTimeoutTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "quasarlabs.quasarnode.qoracle.osmosis.Params")
}

func init() { proto.RegisterFile("qoracle/osmosis/params.proto", fileDescriptor_bf7c77ca07c7bae6) }

var fileDescriptor_bf7c77ca07c7bae6 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x8b, 0xd4, 0x30,
	0x18, 0x6e, 0xdd, 0x71, 0xd5, 0x0a, 0x7e, 0xd4, 0xaf, 0xee, 0xb8, 0x36, 0x43, 0x55, 0x98, 0x83,
	0x24, 0x8c, 0x5e, 0x64, 0x8f, 0x15, 0x44, 0xc1, 0x83, 0x94, 0x3d, 0x09, 0x32, 0x24, 0x99, 0x77,
	0xdb, 0x60, 0xdb, 0x74, 0x9b, 0x74, 0x70, 0xfc, 0x15, 0x1e, 0x3d, 0x7a, 0xf3, 0xaf, 0xec, 0x71,
	0x8f, 0x9e, 0x8a, 0xcc, 0xfc, 0x83, 0xfe, 0x02, 0x99, 0x26, 0xbb, 0x83, 0xe3, 0x9c, 0xf2, 0xe6,
	0xf9, 0x4a, 0x78, 0x79, 0xbc, 0xc3, 0x53, 0x59, 0x53, 0x9e, 0x03, 0x91, 0xaa, 0x90, 0x4a, 0x28,
	0x52, 0xd1, 0x9a, 0x16, 0x0a, 0x57, 0xb5, 0xd4, 0xd2, 0x7f, 0x7e, 0xda, 0x50, 0x45, 0xeb, 0x9c,
	0x32, 0x85, 0xcd, 0x58, 0xca, 0x19, 0x60, 0xeb, 0xc1, 0xd6, 0x33, 0xbc, 0x9f, 0xca, 0x54, 0xf6,
	0x0e, 0xb2, 0x9e, 0x8c, 0x79, 0x78, 0x90, 0x4a, 0x99, 0xe6, 0x40, 0xfa, 0x1b, 0x6b, 0x4e, 0x08,
	0x2d, 0x17, 0x96, 0x0a, 0x79, 0x6f, 0x25, 0x8c, 0x2a, 0x20, 0xf3, 0x09, 0x03, 0x4d, 0x27, 0x84,
	0x4b, 0x51, 0x5a, 0x1e, 0x09, 0xc6, 0x09, 0x97, 0x35, 0x10, 0x9e, 0x0b, 0x28, 0x35, 0x99, 0x4f,
	0xec, 0x64, 0x04, 0xd1, 0xaf, 0x3d, 0x6f, 0xff, 0x63, 0xff, 0x53, 0xff, 0x85, 0x77, 0x0d, 0x4a,
	0xca, 0x72, 0x98, 0x05, 0xee, 0xc8, 0x1d, 0x5f, 0x8f, 0xfd, 0xae, 0x45, 0xb7, 0x16, 0xb4, 0xc8,
	0x8f, 0x22, 0x4b, 0x44, 0xc9, 0x85, 0xc4, 0x7f, 0xeb, 0xdd, 0x81, 0x4a, 0xf2, 0x6c, 0x2a, 0x66,
	0x50, 0x6a, 0x71, 0x22, 0xa0, 0x0e, 0xae, 0x8c, 0xdc, 0xf1, 0x8d, 0xf8, 0x71, 0xd7, 0xa2, 0x47,
	0xd6, 0xb6, 0xa5, 0x88, 0x92, 0xdb, 0x3d, 0xf4, 0xfe, 0x12, 0xf1, 0x3f, 0x78, 0x3e, 0x6d, 0x74,
	0x26, 0x6b, 0xf1, 0x0d, 0x66, 0x53, 0x9e, 0xd1, 0xb2, 0x84, 0x3c, 0xd8, 0xeb, 0x93, 0x9e, 0x74,
	0x2d, 0x3a, 0x30, 0x49, 0xff, 0x6b, 0xa2, 0xe4, 0xee, 0x06, 0x7c, 0x63, 0x30, 0x5f, 0x7b, 0x0f,
	0x2a, 0xca, 0xbf, 0x80, 0x9e, 0x6a, 0x51, 0x80, 0x6c, 0xf4, 0x34, 0x03, 0x91, 0x66, 0x3a, 0x18,
	0x8c, 0xdc, 0xf1, 0xcd, 0x97, 0x43, 0x2c, 0x18, 0xc7, 0xeb, 0x7d, 0x60, 0xbb, 0x85, 0xf9, 0x04,
	0xbf, 0xeb, 0x15, 0xf1, 0xb3, 0xb3, 0x16, 0x39, 0x5d, 0x8b, 0x0e, 0xcd, 0x83, 0x3b, 0x63, 0xa2,
	0xe4, 0x9e, 0xc1, 0x8f, 0x0d, 0x6c, 0xac, 0xfe, 0x67, 0x2f, 0xd8, 0x92, 0xaf, 0x4f, 0xa5, 0x69,
	0x51, 0x05, 0x57, 0x47, 0xee, 0x78, 0x10, 0x3f, 0xed, 0x5a, 0x84, 0x76, 0x06, 0x5f, 0x2a, 0xa3,
	0xe4, 0xe1, 0x3f, 0xd9, 0xc7, 0x17, 0xc4, 0xd1, 0xe0, 0xc7, 0x4f, 0xe4, 0xc4, 0xc9, 0xd9, 0x32,
	0x74, 0xcf, 0x97, 0xa1, 0xfb, 0x67, 0x19, 0xba, 0xdf, 0x57, 0xa1, 0x73, 0xbe, 0x0a, 0x9d, 0xdf,
	0xab, 0xd0, 0xf9, 0xf4, 0x3a, 0x15, 0x3a, 0x6b, 0x18, 0xe6, 0xb2, 0x20, 0x9b, 0x9e, 0x91, 0x4d,
	0xcf, 0xc8, 0x57, 0xb2, 0xdd, 0x4e, 0xbd, 0xa8, 0x40, 0xb1, 0xfd, 0xbe, 0x04, 0xaf, 0xfe, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x17, 0x27, 0x9e, 0x1e, 0xbd, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PacketTimeoutTimestamp != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PacketTimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.PacketTimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.AuthorizedChannel) > 0 {
		i -= len(m.AuthorizedChannel)
		copy(dAtA[i:], m.AuthorizedChannel)
		i = encodeVarintParams(dAtA, i, uint64(len(m.AuthorizedChannel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x12
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.AuthorizedChannel)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.PacketTimeoutHeight.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.PacketTimeoutTimestamp != 0 {
		n += 1 + sovParams(uint64(m.PacketTimeoutTimestamp))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketTimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PacketTimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketTimeoutTimestamp", wireType)
			}
			m.PacketTimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PacketTimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowParams
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
