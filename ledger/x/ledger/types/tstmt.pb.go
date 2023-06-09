// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ledger/ledger/tstmt.proto

package types

import (
	fmt "fmt"
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

type Tstmt struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Payor string `protobuf:"bytes,2,opt,name=payor,proto3" json:"payor,omitempty"`
	Payee string `protobuf:"bytes,3,opt,name=payee,proto3" json:"payee,omitempty"`
	Denom string `protobuf:"bytes,4,opt,name=denom,proto3" json:"denom,omitempty"`
	Q     uint64 `protobuf:"varint,5,opt,name=q,proto3" json:"q,omitempty"`
	Ts    uint64 `protobuf:"varint,6,opt,name=ts,proto3" json:"ts,omitempty"`
	Cid   string `protobuf:"bytes,7,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (m *Tstmt) Reset()         { *m = Tstmt{} }
func (m *Tstmt) String() string { return proto.CompactTextString(m) }
func (*Tstmt) ProtoMessage()    {}
func (*Tstmt) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cc50fdf8013edde, []int{0}
}
func (m *Tstmt) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tstmt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tstmt.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tstmt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tstmt.Merge(m, src)
}
func (m *Tstmt) XXX_Size() int {
	return m.Size()
}
func (m *Tstmt) XXX_DiscardUnknown() {
	xxx_messageInfo_Tstmt.DiscardUnknown(m)
}

var xxx_messageInfo_Tstmt proto.InternalMessageInfo

func (m *Tstmt) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Tstmt) GetPayor() string {
	if m != nil {
		return m.Payor
	}
	return ""
}

func (m *Tstmt) GetPayee() string {
	if m != nil {
		return m.Payee
	}
	return ""
}

func (m *Tstmt) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Tstmt) GetQ() uint64 {
	if m != nil {
		return m.Q
	}
	return 0
}

func (m *Tstmt) GetTs() uint64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *Tstmt) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func init() {
	proto.RegisterType((*Tstmt)(nil), "ledger.ledger.Tstmt")
}

func init() { proto.RegisterFile("ledger/ledger/tstmt.proto", fileDescriptor_8cc50fdf8013edde) }

var fileDescriptor_8cc50fdf8013edde = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0x49, 0x4d, 0x49,
	0x4f, 0x2d, 0xd2, 0x87, 0x52, 0x25, 0xc5, 0x25, 0xb9, 0x25, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xbc, 0x10, 0x31, 0x3d, 0x08, 0xa5, 0xd4, 0xcf, 0xc8, 0xc5, 0x1a, 0x02, 0x92, 0x16, 0x12,
	0xe1, 0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70,
	0x40, 0xa2, 0x05, 0x89, 0x95, 0xf9, 0x45, 0x12, 0x4c, 0x10, 0x51, 0x30, 0x07, 0x2a, 0x9a, 0x9a,
	0x2a, 0xc1, 0x0c, 0x17, 0x4d, 0x4d, 0x05, 0x89, 0xa6, 0xa4, 0xe6, 0xe5, 0xe7, 0x4a, 0xb0, 0x40,
	0x44, 0xc1, 0x1c, 0x21, 0x1e, 0x2e, 0xc6, 0x42, 0x09, 0x56, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xc6,
	0x42, 0x21, 0x3e, 0x2e, 0xa6, 0x92, 0x62, 0x09, 0x36, 0x30, 0x97, 0xa9, 0xa4, 0x58, 0x48, 0x80,
	0x8b, 0x39, 0x39, 0x33, 0x45, 0x82, 0x1d, 0xac, 0x03, 0xc4, 0x74, 0xd2, 0x3f, 0xf1, 0x48, 0x8e,
	0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58,
	0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x51, 0xa8, 0x77, 0x2a, 0xe0, 0xfe, 0xaa, 0x2c, 0x48,
	0x2d, 0x4e, 0x62, 0x03, 0x7b, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x31, 0xca, 0x97, 0x2e,
	0xf5, 0x00, 0x00, 0x00,
}

func (m *Tstmt) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tstmt) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tstmt) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cid) > 0 {
		i -= len(m.Cid)
		copy(dAtA[i:], m.Cid)
		i = encodeVarintTstmt(dAtA, i, uint64(len(m.Cid)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Ts != 0 {
		i = encodeVarintTstmt(dAtA, i, uint64(m.Ts))
		i--
		dAtA[i] = 0x30
	}
	if m.Q != 0 {
		i = encodeVarintTstmt(dAtA, i, uint64(m.Q))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTstmt(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Payee) > 0 {
		i -= len(m.Payee)
		copy(dAtA[i:], m.Payee)
		i = encodeVarintTstmt(dAtA, i, uint64(len(m.Payee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Payor) > 0 {
		i -= len(m.Payor)
		copy(dAtA[i:], m.Payor)
		i = encodeVarintTstmt(dAtA, i, uint64(len(m.Payor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTstmt(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTstmt(dAtA []byte, offset int, v uint64) int {
	offset -= sovTstmt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tstmt) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTstmt(uint64(l))
	}
	l = len(m.Payor)
	if l > 0 {
		n += 1 + l + sovTstmt(uint64(l))
	}
	l = len(m.Payee)
	if l > 0 {
		n += 1 + l + sovTstmt(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTstmt(uint64(l))
	}
	if m.Q != 0 {
		n += 1 + sovTstmt(uint64(m.Q))
	}
	if m.Ts != 0 {
		n += 1 + sovTstmt(uint64(m.Ts))
	}
	l = len(m.Cid)
	if l > 0 {
		n += 1 + l + sovTstmt(uint64(l))
	}
	return n
}

func sovTstmt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTstmt(x uint64) (n int) {
	return sovTstmt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tstmt) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTstmt
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
			return fmt.Errorf("proto: Tstmt: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tstmt: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstmt
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
				return ErrInvalidLengthTstmt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTstmt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstmt
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
				return ErrInvalidLengthTstmt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTstmt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstmt
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
				return ErrInvalidLengthTstmt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTstmt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstmt
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
				return ErrInvalidLengthTstmt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTstmt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Q", wireType)
			}
			m.Q = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstmt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Q |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ts", wireType)
			}
			m.Ts = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstmt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ts |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTstmt
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
				return ErrInvalidLengthTstmt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTstmt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTstmt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTstmt
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
func skipTstmt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTstmt
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
					return 0, ErrIntOverflowTstmt
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
					return 0, ErrIntOverflowTstmt
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
				return 0, ErrInvalidLengthTstmt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTstmt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTstmt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTstmt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTstmt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTstmt = fmt.Errorf("proto: unexpected end of group")
)
