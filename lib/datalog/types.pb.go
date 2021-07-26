// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package datalog

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

type Facts_PredicateType int32

const (
	Facts_HASROLE             Facts_PredicateType = 0
	Facts_HASTRAIT            Facts_PredicateType = 1
	Facts_ROLEALLOWSLOGIN     Facts_PredicateType = 2
	Facts_ROLEDENIESLOGIN     Facts_PredicateType = 3
	Facts_ROLEALLOWSNODELABEL Facts_PredicateType = 4
	Facts_ROLEDENIESNODELABEL Facts_PredicateType = 5
	Facts_NODEHASLABEL        Facts_PredicateType = 6
	// Outputs
	Facts_HASACCESS  Facts_PredicateType = 7
	Facts_DENYACCESS Facts_PredicateType = 8
	Facts_DENYLOGINS Facts_PredicateType = 9
)

var Facts_PredicateType_name = map[int32]string{
	0: "HASROLE",
	1: "HASTRAIT",
	2: "ROLEALLOWSLOGIN",
	3: "ROLEDENIESLOGIN",
	4: "ROLEALLOWSNODELABEL",
	5: "ROLEDENIESNODELABEL",
	6: "NODEHASLABEL",
	7: "HASACCESS",
	8: "DENYACCESS",
	9: "DENYLOGINS",
}

var Facts_PredicateType_value = map[string]int32{
	"HASROLE":             0,
	"HASTRAIT":            1,
	"ROLEALLOWSLOGIN":     2,
	"ROLEDENIESLOGIN":     3,
	"ROLEALLOWSNODELABEL": 4,
	"ROLEDENIESNODELABEL": 5,
	"NODEHASLABEL":        6,
	"HASACCESS":           7,
	"DENYACCESS":          8,
	"DENYLOGINS":          9,
}

func (x Facts_PredicateType) String() string {
	return proto.EnumName(Facts_PredicateType_name, int32(x))
}

func (Facts_PredicateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0, 0}
}

// Facts represents a datalog database of predicates
type Facts struct {
	Predicates           []*Facts_Predicate `protobuf:"bytes,1,rep,name=predicates,proto3" json:"predicates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Facts) Reset()         { *m = Facts{} }
func (m *Facts) String() string { return proto.CompactTextString(m) }
func (*Facts) ProtoMessage()    {}
func (*Facts) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}
func (m *Facts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Facts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Facts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Facts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Facts.Merge(m, src)
}
func (m *Facts) XXX_Size() int {
	return m.Size()
}
func (m *Facts) XXX_DiscardUnknown() {
	xxx_messageInfo_Facts.DiscardUnknown(m)
}

var xxx_messageInfo_Facts proto.InternalMessageInfo

func (m *Facts) GetPredicates() []*Facts_Predicate {
	if m != nil {
		return m.Predicates
	}
	return nil
}

type Facts_Predicate struct {
	Name                 Facts_PredicateType `protobuf:"varint,1,opt,name=name,proto3,enum=datalog.Facts_PredicateType" json:"name,omitempty"`
	Atoms                []uint32            `protobuf:"varint,2,rep,packed,name=atoms,proto3" json:"atoms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Facts_Predicate) Reset()         { *m = Facts_Predicate{} }
func (m *Facts_Predicate) String() string { return proto.CompactTextString(m) }
func (*Facts_Predicate) ProtoMessage()    {}
func (*Facts_Predicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0, 0}
}
func (m *Facts_Predicate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Facts_Predicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Facts_Predicate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Facts_Predicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Facts_Predicate.Merge(m, src)
}
func (m *Facts_Predicate) XXX_Size() int {
	return m.Size()
}
func (m *Facts_Predicate) XXX_DiscardUnknown() {
	xxx_messageInfo_Facts_Predicate.DiscardUnknown(m)
}

var xxx_messageInfo_Facts_Predicate proto.InternalMessageInfo

func (m *Facts_Predicate) GetName() Facts_PredicateType {
	if m != nil {
		return m.Name
	}
	return Facts_HASROLE
}

func (m *Facts_Predicate) GetAtoms() []uint32 {
	if m != nil {
		return m.Atoms
	}
	return nil
}

func init() {
	proto.RegisterEnum("datalog.Facts_PredicateType", Facts_PredicateType_name, Facts_PredicateType_value)
	proto.RegisterType((*Facts)(nil), "datalog.Facts")
	proto.RegisterType((*Facts_Predicate)(nil), "datalog.Facts.Predicate")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0xcd, 0x4a, 0xf3, 0x40,
	0x18, 0x05, 0xe0, 0x6f, 0xfa, 0x9f, 0x37, 0x49, 0xbf, 0x61, 0x2a, 0x18, 0x44, 0x42, 0xe8, 0x2a,
	0xab, 0x20, 0x75, 0xe3, 0x76, 0xda, 0x8c, 0x26, 0x30, 0x24, 0x92, 0x09, 0x88, 0xcb, 0xb1, 0x0d,
	0x22, 0x58, 0x13, 0x9a, 0xd9, 0xf4, 0x0a, 0x75, 0xe9, 0x25, 0x94, 0x5c, 0x89, 0x34, 0x3f, 0x56,
	0x17, 0x2e, 0xcf, 0x39, 0x0f, 0x33, 0x8b, 0x17, 0x74, 0xb5, 0x2f, 0xb2, 0xd2, 0x2b, 0x76, 0xb9,
	0xca, 0xc9, 0x78, 0x23, 0x95, 0x7c, 0xcd, 0x9f, 0xe7, 0x87, 0x1e, 0x0c, 0x6f, 0xe5, 0x5a, 0x95,
	0xe4, 0x06, 0xa0, 0xd8, 0x65, 0x9b, 0x97, 0xb5, 0x54, 0x59, 0x69, 0x21, 0xa7, 0xef, 0xea, 0x0b,
	0xcb, 0x6b, 0x9d, 0x57, 0x1b, 0xef, 0xbe, 0x03, 0xc9, 0x0f, 0x7b, 0x21, 0x40, 0xfb, 0x1e, 0xc8,
	0x15, 0x0c, 0xde, 0xe4, 0x36, 0xb3, 0x90, 0x83, 0xdc, 0xe9, 0xe2, 0xf2, 0xaf, 0x07, 0xd2, 0x7d,
	0x91, 0x25, 0xb5, 0x24, 0x67, 0x30, 0x94, 0x2a, 0xdf, 0x96, 0x56, 0xcf, 0xe9, 0xbb, 0x66, 0xd2,
	0x84, 0xf9, 0x3b, 0x02, 0xf3, 0x97, 0x26, 0x3a, 0x8c, 0x03, 0x2a, 0x92, 0x98, 0x33, 0xfc, 0x8f,
	0x18, 0x30, 0x09, 0xa8, 0x48, 0x13, 0x1a, 0xa6, 0x18, 0x91, 0x19, 0xfc, 0x3f, 0xf6, 0x94, 0xf3,
	0xf8, 0x41, 0xf0, 0xf8, 0x2e, 0x8c, 0x70, 0xaf, 0x2b, 0x7d, 0x16, 0x85, 0xac, 0x2d, 0xfb, 0xe4,
	0x1c, 0x66, 0x27, 0x19, 0xc5, 0x3e, 0xe3, 0x74, 0xc9, 0x38, 0x1e, 0x74, 0x43, 0xa3, 0x4f, 0xc3,
	0x90, 0x60, 0x30, 0x8e, 0x31, 0xa0, 0xa2, 0x69, 0x46, 0xc4, 0x04, 0x2d, 0xa0, 0x82, 0xae, 0x56,
	0x4c, 0x08, 0x3c, 0x26, 0x53, 0x00, 0x9f, 0x45, 0x8f, 0x6d, 0x9e, 0x74, 0xb9, 0xfe, 0x51, 0x60,
	0x6d, 0x69, 0x7c, 0x54, 0x36, 0xfa, 0xac, 0x6c, 0x74, 0xa8, 0x6c, 0xf4, 0x34, 0xaa, 0x0f, 0x70,
	0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x62, 0x53, 0x79, 0x8f, 0x01, 0x00, 0x00,
}

func (m *Facts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Facts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Facts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Predicates) > 0 {
		for iNdEx := len(m.Predicates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Predicates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Facts_Predicate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Facts_Predicate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Facts_Predicate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Atoms) > 0 {
		dAtA2 := make([]byte, len(m.Atoms)*10)
		var j1 int
		for _, num := range m.Atoms {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTypes(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.Name != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Name))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Facts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Predicates) > 0 {
		for _, e := range m.Predicates {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Facts_Predicate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Name != 0 {
		n += 1 + sovTypes(uint64(m.Name))
	}
	if len(m.Atoms) > 0 {
		l = 0
		for _, e := range m.Atoms {
			l += sovTypes(uint64(e))
		}
		n += 1 + sovTypes(uint64(l)) + l
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Facts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Facts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Facts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Predicates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Predicates = append(m.Predicates, &Facts_Predicate{})
			if err := m.Predicates[len(m.Predicates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Facts_Predicate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Predicate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Predicate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			m.Name = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Name |= Facts_PredicateType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Atoms = append(m.Atoms, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypes
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Atoms) == 0 {
					m.Atoms = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Atoms = append(m.Atoms, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Atoms", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
