// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identity/delegation.proto

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

type Delegation struct {
	Index   string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	DsnpId  uint64 `protobuf:"varint,2,opt,name=dsnpId,proto3" json:"dsnpId,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Role    uint32 `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	Id      uint64 `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *Delegation) Reset()         { *m = Delegation{} }
func (m *Delegation) String() string { return proto.CompactTextString(m) }
func (*Delegation) ProtoMessage()    {}
func (*Delegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f28619112beef76, []int{0}
}
func (m *Delegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegation.Merge(m, src)
}
func (m *Delegation) XXX_Size() int {
	return m.Size()
}
func (m *Delegation) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegation.DiscardUnknown(m)
}

var xxx_messageInfo_Delegation proto.InternalMessageInfo

func (m *Delegation) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Delegation) GetDsnpId() uint64 {
	if m != nil {
		return m.DsnpId
	}
	return 0
}

func (m *Delegation) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Delegation) GetRole() uint32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *Delegation) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Delegation)(nil), "amparks100.registry.identity.Delegation")
}

func init() { proto.RegisterFile("identity/delegation.proto", fileDescriptor_6f28619112beef76) }

var fileDescriptor_6f28619112beef76 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0xb3, 0x31, 0x77, 0xe2, 0x80, 0x16, 0x8b, 0xc8, 0x0a, 0xb2, 0x04, 0xab, 0x54, 0xd9,
	0x13, 0xdf, 0x40, 0x6c, 0xae, 0x4d, 0x69, 0x97, 0x73, 0x86, 0x38, 0x78, 0x97, 0x0d, 0xbb, 0x2b,
	0x24, 0x6f, 0xe1, 0x63, 0x59, 0x5e, 0x69, 0x29, 0xc9, 0x8b, 0x08, 0xab, 0xc9, 0x75, 0xf3, 0xc1,
	0x3f, 0x3f, 0xff, 0x07, 0xb7, 0x8c, 0xd4, 0x06, 0x0e, 0x83, 0x41, 0xda, 0x53, 0x53, 0x07, 0xb6,
	0x6d, 0xd9, 0x39, 0x1b, 0xac, 0xbc, 0xab, 0x0f, 0x5d, 0xed, 0xde, 0xfd, 0xc3, 0x66, 0x53, 0x3a,
	0x6a, 0xd8, 0x07, 0x37, 0x94, 0x73, 0xfc, 0xbe, 0x07, 0x78, 0x5e, 0x3e, 0xe4, 0x35, 0xac, 0xb8,
	0x45, 0xea, 0x95, 0xc8, 0x45, 0x71, 0x51, 0xfd, 0x81, 0xbc, 0x81, 0x35, 0xfa, 0xb6, 0xdb, 0xa2,
	0x4a, 0x73, 0x51, 0x64, 0xd5, 0x3f, 0x49, 0x05, 0xe7, 0x35, 0xa2, 0x23, 0xef, 0xd5, 0x59, 0xcc,
	0xcf, 0x28, 0x25, 0x64, 0xce, 0xee, 0x49, 0x65, 0xb9, 0x28, 0x2e, 0xab, 0x78, 0xcb, 0x2b, 0x48,
	0x19, 0xd5, 0x2a, 0x36, 0xa4, 0x8c, 0x4f, 0xdb, 0xaf, 0x51, 0x8b, 0xe3, 0xa8, 0xc5, 0xcf, 0xa8,
	0xc5, 0xe7, 0xa4, 0x93, 0xe3, 0xa4, 0x93, 0xef, 0x49, 0x27, 0x2f, 0xa6, 0xe1, 0xf0, 0xf6, 0xb1,
	0x2b, 0x5f, 0xed, 0xc1, 0x9c, 0xc6, 0x9b, 0x79, 0xbc, 0xe9, 0xcd, 0x62, 0x1b, 0x86, 0x8e, 0xfc,
	0x6e, 0x1d, 0x4d, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xe2, 0x6d, 0xae, 0x06, 0x01,
	0x00, 0x00,
}

func (m *Delegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintDelegation(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x28
	}
	if m.Role != 0 {
		i = encodeVarintDelegation(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintDelegation(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DsnpId != 0 {
		i = encodeVarintDelegation(dAtA, i, uint64(m.DsnpId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintDelegation(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegation(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Delegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovDelegation(uint64(l))
	}
	if m.DsnpId != 0 {
		n += 1 + sovDelegation(uint64(m.DsnpId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovDelegation(uint64(l))
	}
	if m.Role != 0 {
		n += 1 + sovDelegation(uint64(m.Role))
	}
	if m.Id != 0 {
		n += 1 + sovDelegation(uint64(m.Id))
	}
	return n
}

func sovDelegation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegation(x uint64) (n int) {
	return sovDelegation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Delegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegation
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
			return fmt.Errorf("proto: Delegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DsnpId", wireType)
			}
			m.DsnpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DsnpId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegation
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
func skipDelegation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
				return 0, ErrInvalidLengthDelegation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegation = fmt.Errorf("proto: unexpected end of group")
)
