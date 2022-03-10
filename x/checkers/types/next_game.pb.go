// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/next_game.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type NextGame struct {
	IdValue  uint64 `protobuf:"varint,1,opt,name=idValue,proto3" json:"idValue,omitempty"`
	FifoHead string `protobuf:"bytes,2,opt,name=fifoHead,proto3" json:"fifoHead,omitempty"`
	FifoTail string `protobuf:"bytes,3,opt,name=fifoTail,proto3" json:"fifoTail,omitempty"`
}

func (m *NextGame) Reset()         { *m = NextGame{} }
func (m *NextGame) String() string { return proto.CompactTextString(m) }
func (*NextGame) ProtoMessage()    {}
func (*NextGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_c12fafd6c218bc9c, []int{0}
}
func (m *NextGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NextGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NextGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NextGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NextGame.Merge(m, src)
}
func (m *NextGame) XXX_Size() int {
	return m.Size()
}
func (m *NextGame) XXX_DiscardUnknown() {
	xxx_messageInfo_NextGame.DiscardUnknown(m)
}

var xxx_messageInfo_NextGame proto.InternalMessageInfo

func (m *NextGame) GetIdValue() uint64 {
	if m != nil {
		return m.IdValue
	}
	return 0
}

func (m *NextGame) GetFifoHead() string {
	if m != nil {
		return m.FifoHead
	}
	return ""
}

func (m *NextGame) GetFifoTail() string {
	if m != nil {
		return m.FifoTail
	}
	return ""
}

func init() {
	proto.RegisterType((*NextGame)(nil), "minhhung123.checkers.checkers.NextGame")
}

func init() { proto.RegisterFile("checkers/next_game.proto", fileDescriptor_c12fafd6c218bc9c) }

var fileDescriptor_c12fafd6c218bc9c = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0xce, 0x48, 0x4d,
	0xce, 0x4e, 0x2d, 0x2a, 0xd6, 0xcf, 0x4b, 0xad, 0x28, 0x89, 0x4f, 0x4f, 0xcc, 0x4d, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcd, 0xcd, 0xcc, 0xcb, 0xc8, 0x28, 0xcd, 0x4b, 0x37, 0x34,
	0x32, 0xd6, 0x83, 0xa9, 0x82, 0x33, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x2a, 0xf5, 0x41,
	0x2c, 0x88, 0x26, 0xa5, 0x18, 0x2e, 0x0e, 0xbf, 0xd4, 0x8a, 0x12, 0xf7, 0xc4, 0xdc, 0x54, 0x21,
	0x09, 0x2e, 0xf6, 0xcc, 0x94, 0xb0, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96,
	0x20, 0x18, 0x57, 0x48, 0x8a, 0x8b, 0x23, 0x2d, 0x33, 0x2d, 0xdf, 0x23, 0x35, 0x31, 0x45, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xce, 0x87, 0xc9, 0x85, 0x24, 0x66, 0xe6, 0x48, 0x30, 0x23,
	0xe4, 0x40, 0x7c, 0x27, 0xaf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32,
	0x48, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0x72, 0xb7, 0x3e, 0xdc,
	0x77, 0x15, 0x08, 0x66, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xc1, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x4e, 0x45, 0x16, 0x48, 0x01, 0x01, 0x00, 0x00,
}

func (m *NextGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NextGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NextGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FifoTail) > 0 {
		i -= len(m.FifoTail)
		copy(dAtA[i:], m.FifoTail)
		i = encodeVarintNextGame(dAtA, i, uint64(len(m.FifoTail)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.FifoHead) > 0 {
		i -= len(m.FifoHead)
		copy(dAtA[i:], m.FifoHead)
		i = encodeVarintNextGame(dAtA, i, uint64(len(m.FifoHead)))
		i--
		dAtA[i] = 0x12
	}
	if m.IdValue != 0 {
		i = encodeVarintNextGame(dAtA, i, uint64(m.IdValue))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNextGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovNextGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NextGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IdValue != 0 {
		n += 1 + sovNextGame(uint64(m.IdValue))
	}
	l = len(m.FifoHead)
	if l > 0 {
		n += 1 + l + sovNextGame(uint64(l))
	}
	l = len(m.FifoTail)
	if l > 0 {
		n += 1 + l + sovNextGame(uint64(l))
	}
	return n
}

func sovNextGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNextGame(x uint64) (n int) {
	return sovNextGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NextGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNextGame
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
			return fmt.Errorf("proto: NextGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NextGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdValue", wireType)
			}
			m.IdValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNextGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IdValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoHead", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNextGame
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
				return ErrInvalidLengthNextGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNextGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FifoHead = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FifoTail", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNextGame
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
				return ErrInvalidLengthNextGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNextGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FifoTail = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNextGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNextGame
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
func skipNextGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNextGame
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
					return 0, ErrIntOverflowNextGame
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
					return 0, ErrIntOverflowNextGame
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
				return 0, ErrInvalidLengthNextGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNextGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNextGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNextGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNextGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNextGame = fmt.Errorf("proto: unexpected end of group")
)
