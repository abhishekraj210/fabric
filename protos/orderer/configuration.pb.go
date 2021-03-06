// Code generated by protoc-gen-go.
// source: orderer/configuration.proto
// DO NOT EDIT!

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type BatchSize struct {
	// Simply specified as messages for now, in the future we may want to allow
	// this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=maxMessageCount" json:"maxMessageCount,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// When submitting a new chain configuration transaction to create a new chain,
// the first configuration item must be of type Orderer with Key CreationPolicy
// and contents of a Marshaled CreationPolicy. The policy should be set to the
// policy which was supplied by the ordering service for the client's chain
// creation. The digest should be the hash of the concatenation of the remaining
// ConfigurationItem bytes. The signatures of the configuration item should
// satisfy the policy for chain creation.
type CreationPolicy struct {
	// The name of the policy which should be used to validate the creation of
	// this chain
	Policy string `protobuf:"bytes,1,opt,name=policy" json:"policy,omitempty"`
	// The hash of the concatenation of remaining configuration item bytes
	Digest []byte `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (m *CreationPolicy) Reset()                    { *m = CreationPolicy{} }
func (m *CreationPolicy) String() string            { return proto.CompactTextString(m) }
func (*CreationPolicy) ProtoMessage()               {}
func (*CreationPolicy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type ChainCreators struct {
	// A list of policies, any of which may be specified as the chain creation
	// policy in a chain creation request
	Policies []string `protobuf:"bytes,1,rep,name=policies" json:"policies,omitempty"`
}

func (m *ChainCreators) Reset()                    { *m = ChainCreators{} }
func (m *ChainCreators) String() string            { return proto.CompactTextString(m) }
func (*ChainCreators) ProtoMessage()               {}
func (*ChainCreators) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*CreationPolicy)(nil), "orderer.CreationPolicy")
	proto.RegisterType((*ChainCreators)(nil), "orderer.ChainCreators")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0xca, 0xae, 0x1b, 0x76, 0x15, 0x22, 0x48, 0x59, 0x2f, 0xa5, 0x5e, 0x0a, 0x4a,
	0x73, 0x10, 0xef, 0xd2, 0x1e, 0x45, 0x90, 0xea, 0xc9, 0x5b, 0x9a, 0x4e, 0xdb, 0xb0, 0xdb, 0x4c,
	0x99, 0xa4, 0x60, 0xfd, 0xf5, 0x62, 0x1a, 0x3d, 0x78, 0xca, 0xfb, 0x66, 0x5e, 0x1e, 0xc3, 0x63,
	0x37, 0x48, 0x0d, 0x10, 0x90, 0x50, 0x68, 0x5a, 0xdd, 0x4d, 0x24, 0x9d, 0x46, 0x93, 0x8f, 0x84,
	0x0e, 0xf9, 0x3a, 0x2c, 0xf7, 0x57, 0x0a, 0x87, 0x01, 0x8d, 0x58, 0x9e, 0x65, 0x9b, 0xde, 0xb2,
	0x5d, 0x89, 0xc6, 0x82, 0xb1, 0x93, 0x7d, 0x9f, 0x47, 0xe0, 0x9c, 0x9d, 0xb9, 0x79, 0x84, 0x38,
	0x4a, 0xa2, 0x6c, 0x53, 0x79, 0x9d, 0x3e, 0xb2, 0x4d, 0x21, 0x9d, 0xea, 0xdf, 0xf4, 0x17, 0xf0,
	0x8c, 0x5d, 0x0e, 0xf2, 0xf3, 0x05, 0xac, 0x95, 0x1d, 0x94, 0x38, 0x19, 0xe7, 0xbd, 0xbb, 0xea,
	0xff, 0x38, 0x7d, 0x62, 0x17, 0x25, 0x81, 0xbf, 0xe5, 0x15, 0x8f, 0x5a, 0xcd, 0xfc, 0x9a, 0xad,
	0x46, 0xaf, 0x42, 0x7c, 0xa0, 0x9f, 0x79, 0xa3, 0x3b, 0xb0, 0x2e, 0x3e, 0x49, 0xa2, 0x6c, 0x5b,
	0x05, 0x4a, 0xef, 0xd8, 0xae, 0xec, 0xa5, 0x36, 0x3e, 0x06, 0xc9, 0xf2, 0x3d, 0x3b, 0xf7, 0x5f,
	0x34, 0xd8, 0x38, 0x4a, 0x4e, 0xb3, 0x4d, 0xf5, 0xc7, 0x69, 0xc6, 0xb6, 0xcf, 0xb2, 0x3d, 0xc8,
	0x82, 0xf0, 0x00, 0x64, 0x79, 0xcc, 0xd6, 0xf5, 0x22, 0x83, 0xf5, 0x17, 0x8b, 0xfc, 0xe3, 0xbe,
	0xd3, 0xae, 0x9f, 0xea, 0x5c, 0xe1, 0x20, 0xfa, 0x79, 0x04, 0x3a, 0x42, 0xd3, 0x01, 0x89, 0x56,
	0xd6, 0xa4, 0x95, 0xf0, 0xdd, 0x58, 0x11, 0x9a, 0xab, 0x57, 0x9e, 0x1f, 0xbe, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0xc4, 0xe3, 0xe4, 0x68, 0x01, 0x00, 0x00,
}
