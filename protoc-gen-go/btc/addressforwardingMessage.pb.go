// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressforwardingMessage.proto

package btc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateForwardEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	//Identifier of the address forwarding request; generated when a new request is created.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	//The mandatory user token.
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	//The required destination address for address forwarding.
	Destination string `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	//The address which will automatically forward to destination; generated when a new request is created.
	InputAddress string `protobuf:"bytes,5,opt,name=input_address,json=inputAddress,proto3" json:"input_address,omitempty"`
	//Optional Address to forward processing fees, if specified. Allows you to receive a fee for your own services.
	ProcessFeesAddress string `protobuf:"bytes,6,opt,name=process_fees_address,json=processFeesAddress,proto3" json:"process_fees_address,omitempty"`
	//Optional Fixed processing fee amount to be sent to the fee address. A fixed satoshi amount or a percentage is required if a process_fees_address has been specified.
	ProcessFeesSatoshis int32 `protobuf:"varint,7,opt,name=process_fees_satoshis,json=processFeesSatoshis,proto3" json:"process_fees_satoshis,omitempty"`
	//Optional Percentage of the transaction to be sent to the fee address. A fixed satoshi amount or a percentage is required if a process_fees_address has been specified.
	ProcessFeesPercent float32 `protobuf:"fixed32,8,opt,name=process_fees_percent,json=processFeesPercent,proto3" json:"process_fees_percent,omitempty"`
	//Optional The URL to call anytime a new is forwarded.
	CallbackUrl string `protobuf:"bytes,9,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	//Optional Whether to also call the callback_url with subsequent confirmations of the forwarding transactions. Automatically sets up a WebHook.
	EnableConfirmations bool `protobuf:"varint,10,opt,name=enable_confirmations,json=enableConfirmations,proto3" json:"enable_confirmations,omitempty"`
	//Optional Mining fee amount to include in the forwarding transaction, in satoshis. If not set, defaults to 10,000.
	MiningFeesSatoshis int32 `protobuf:"varint,11,opt,name=mining_fees_satoshis,json=miningFeesSatoshis,proto3" json:"mining_fees_satoshis,omitempty"`
	//Optional History of forwarding transaction hashes for this address forward; not present if this request has yet to forward any transactions.
	Txs                  []string `protobuf:"bytes,12,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateForwardEndpointRequest) Reset()         { *m = CreateForwardEndpointRequest{} }
func (m *CreateForwardEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*CreateForwardEndpointRequest) ProtoMessage()    {}
func (*CreateForwardEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11170982bd173f0c, []int{0}
}

func (m *CreateForwardEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateForwardEndpointRequest.Unmarshal(m, b)
}
func (m *CreateForwardEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateForwardEndpointRequest.Marshal(b, m, deterministic)
}
func (m *CreateForwardEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateForwardEndpointRequest.Merge(m, src)
}
func (m *CreateForwardEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_CreateForwardEndpointRequest.Size(m)
}
func (m *CreateForwardEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateForwardEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateForwardEndpointRequest proto.InternalMessageInfo

func (m *CreateForwardEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *CreateForwardEndpointRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateForwardEndpointRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CreateForwardEndpointRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *CreateForwardEndpointRequest) GetInputAddress() string {
	if m != nil {
		return m.InputAddress
	}
	return ""
}

func (m *CreateForwardEndpointRequest) GetProcessFeesAddress() string {
	if m != nil {
		return m.ProcessFeesAddress
	}
	return ""
}

func (m *CreateForwardEndpointRequest) GetProcessFeesSatoshis() int32 {
	if m != nil {
		return m.ProcessFeesSatoshis
	}
	return 0
}

func (m *CreateForwardEndpointRequest) GetProcessFeesPercent() float32 {
	if m != nil {
		return m.ProcessFeesPercent
	}
	return 0
}

func (m *CreateForwardEndpointRequest) GetCallbackUrl() string {
	if m != nil {
		return m.CallbackUrl
	}
	return ""
}

func (m *CreateForwardEndpointRequest) GetEnableConfirmations() bool {
	if m != nil {
		return m.EnableConfirmations
	}
	return false
}

func (m *CreateForwardEndpointRequest) GetMiningFeesSatoshis() int32 {
	if m != nil {
		return m.MiningFeesSatoshis
	}
	return 0
}

func (m *CreateForwardEndpointRequest) GetTxs() []string {
	if m != nil {
		return m.Txs
	}
	return nil
}

type ListForwardsEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	//Returns list of address forwards starting at the start index; useful for paging beyond the limit of 200 address forwards.
	Start                int32    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListForwardsEndpointRequest) Reset()         { *m = ListForwardsEndpointRequest{} }
func (m *ListForwardsEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*ListForwardsEndpointRequest) ProtoMessage()    {}
func (*ListForwardsEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11170982bd173f0c, []int{1}
}

func (m *ListForwardsEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListForwardsEndpointRequest.Unmarshal(m, b)
}
func (m *ListForwardsEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListForwardsEndpointRequest.Marshal(b, m, deterministic)
}
func (m *ListForwardsEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListForwardsEndpointRequest.Merge(m, src)
}
func (m *ListForwardsEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_ListForwardsEndpointRequest.Size(m)
}
func (m *ListForwardsEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListForwardsEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListForwardsEndpointRequest proto.InternalMessageInfo

func (m *ListForwardsEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *ListForwardsEndpointRequest) GetStart() int32 {
	if m != nil {
		return m.Start
	}
	return 0
}

type ListForwardsEndpointResponse struct {
	Forwards             []*AddressForward `protobuf:"bytes,1,rep,name=forwards,proto3" json:"forwards,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListForwardsEndpointResponse) Reset()         { *m = ListForwardsEndpointResponse{} }
func (m *ListForwardsEndpointResponse) String() string { return proto.CompactTextString(m) }
func (*ListForwardsEndpointResponse) ProtoMessage()    {}
func (*ListForwardsEndpointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11170982bd173f0c, []int{2}
}

func (m *ListForwardsEndpointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListForwardsEndpointResponse.Unmarshal(m, b)
}
func (m *ListForwardsEndpointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListForwardsEndpointResponse.Marshal(b, m, deterministic)
}
func (m *ListForwardsEndpointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListForwardsEndpointResponse.Merge(m, src)
}
func (m *ListForwardsEndpointResponse) XXX_Size() int {
	return xxx_messageInfo_ListForwardsEndpointResponse.Size(m)
}
func (m *ListForwardsEndpointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListForwardsEndpointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListForwardsEndpointResponse proto.InternalMessageInfo

func (m *ListForwardsEndpointResponse) GetForwards() []*AddressForward {
	if m != nil {
		return m.Forwards
	}
	return nil
}

type DeleteForwardEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Payid                string               `protobuf:"bytes,2,opt,name=payid,proto3" json:"payid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeleteForwardEndpointRequest) Reset()         { *m = DeleteForwardEndpointRequest{} }
func (m *DeleteForwardEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteForwardEndpointRequest) ProtoMessage()    {}
func (*DeleteForwardEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11170982bd173f0c, []int{3}
}

func (m *DeleteForwardEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteForwardEndpointRequest.Unmarshal(m, b)
}
func (m *DeleteForwardEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteForwardEndpointRequest.Marshal(b, m, deterministic)
}
func (m *DeleteForwardEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteForwardEndpointRequest.Merge(m, src)
}
func (m *DeleteForwardEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteForwardEndpointRequest.Size(m)
}
func (m *DeleteForwardEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteForwardEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteForwardEndpointRequest proto.InternalMessageInfo

func (m *DeleteForwardEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *DeleteForwardEndpointRequest) GetPayid() string {
	if m != nil {
		return m.Payid
	}
	return ""
}

type DeleteForwardEndpointResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteForwardEndpointResponse) Reset()         { *m = DeleteForwardEndpointResponse{} }
func (m *DeleteForwardEndpointResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteForwardEndpointResponse) ProtoMessage()    {}
func (*DeleteForwardEndpointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11170982bd173f0c, []int{4}
}

func (m *DeleteForwardEndpointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteForwardEndpointResponse.Unmarshal(m, b)
}
func (m *DeleteForwardEndpointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteForwardEndpointResponse.Marshal(b, m, deterministic)
}
func (m *DeleteForwardEndpointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteForwardEndpointResponse.Merge(m, src)
}
func (m *DeleteForwardEndpointResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteForwardEndpointResponse.Size(m)
}
func (m *DeleteForwardEndpointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteForwardEndpointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteForwardEndpointResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateForwardEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.CreateForwardEndpointRequest")
	proto.RegisterType((*ListForwardsEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointRequest")
	proto.RegisterType((*ListForwardsEndpointResponse)(nil), "fairwaycorp.blockchainprotobuf.btc.ListForwardsEndpointResponse")
	proto.RegisterType((*DeleteForwardEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointRequest")
	proto.RegisterType((*DeleteForwardEndpointResponse)(nil), "fairwaycorp.blockchainprotobuf.btc.DeleteForwardEndpointResponse")
}

func init() { proto.RegisterFile("addressforwardingMessage.proto", fileDescriptor_11170982bd173f0c) }

var fileDescriptor_11170982bd173f0c = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x86, 0x95, 0x76, 0xb3, 0x1f, 0x6e, 0x59, 0x21, 0xb7, 0x48, 0x16, 0x14, 0x08, 0xe1, 0x92,
	0x4b, 0x53, 0x28, 0x17, 0x2e, 0x08, 0x95, 0x85, 0x3d, 0xc1, 0x0a, 0x05, 0x71, 0xe1, 0x52, 0x39,
	0xce, 0x34, 0xb5, 0x9a, 0xda, 0xc1, 0xe3, 0xaa, 0xec, 0x2f, 0x80, 0x2b, 0xff, 0x18, 0xd5, 0x4e,
	0x4b, 0xcb, 0x87, 0xe0, 0xb4, 0x37, 0xfb, 0x7d, 0x3d, 0x33, 0xcf, 0x24, 0x33, 0xe4, 0x01, 0x2f,
	0x0a, 0x03, 0x88, 0x33, 0x6d, 0xd6, 0xdc, 0x14, 0x52, 0x95, 0xef, 0x00, 0x91, 0x97, 0x90, 0xd6,
	0x46, 0x5b, 0x4d, 0xe3, 0x19, 0x97, 0x66, 0xcd, 0xaf, 0x85, 0x36, 0x75, 0x9a, 0x57, 0x5a, 0x2c,
	0xc4, 0x9c, 0x4b, 0xe5, 0xcc, 0x7c, 0x35, 0x4b, 0x73, 0x2b, 0xee, 0xf6, 0x84, 0x5e, 0x2e, 0xb5,
	0x3a, 0x08, 0x8c, 0xbf, 0x1f, 0x91, 0xc1, 0x85, 0x01, 0x6e, 0xe1, 0xd2, 0xa7, 0x7e, 0xa3, 0x8a,
	0x5a, 0x4b, 0x65, 0x33, 0xf8, 0xbc, 0x02, 0xb4, 0x34, 0x23, 0x27, 0x0a, 0xec, 0x5a, 0x9b, 0x05,
	0x0b, 0xa2, 0x20, 0x39, 0x1f, 0x3f, 0x4f, 0xff, 0x5d, 0x2b, 0xbd, 0xf2, 0x21, 0x93, 0xaa, 0xd2,
	0x6b, 0xa9, 0xca, 0x49, 0x25, 0x39, 0x66, 0xdb, 0x44, 0xf4, 0x9c, 0xb4, 0x64, 0xc1, 0x5a, 0x51,
	0x90, 0x9c, 0x65, 0x2d, 0x59, 0xd0, 0x3e, 0x09, 0xad, 0x5e, 0x80, 0x62, 0x6d, 0x27, 0xf9, 0x0b,
	0x8d, 0x48, 0xa7, 0x00, 0xb4, 0x52, 0x71, 0x2b, 0xb5, 0x62, 0x47, 0xce, 0xdb, 0x97, 0xe8, 0x63,
	0x72, 0x4b, 0xaa, 0x7a, 0x65, 0xa7, 0xcd, 0xd7, 0x61, 0xa1, 0x7b, 0xd3, 0x75, 0xe2, 0xc4, 0x6b,
	0xf4, 0x09, 0xe9, 0xd7, 0x46, 0x0b, 0x40, 0x9c, 0xce, 0x00, 0x70, 0xf7, 0xf6, 0xd8, 0xbd, 0xa5,
	0x8d, 0x77, 0x09, 0x80, 0xdb, 0x88, 0x31, 0xb9, 0x73, 0x10, 0x81, 0xdc, 0x6a, 0x9c, 0x4b, 0x64,
	0x27, 0x51, 0x90, 0x84, 0x59, 0x6f, 0x2f, 0xe4, 0x43, 0x63, 0xfd, 0x56, 0xa5, 0x06, 0x23, 0x40,
	0x59, 0x76, 0x1a, 0x05, 0x49, 0xeb, 0xa0, 0xca, 0x7b, 0xef, 0xd0, 0x47, 0xa4, 0x2b, 0x78, 0x55,
	0xe5, 0x5c, 0x2c, 0xa6, 0x2b, 0x53, 0xb1, 0x33, 0xdf, 0xdf, 0x56, 0xfb, 0x68, 0x2a, 0xfa, 0x94,
	0xf4, 0x41, 0xf1, 0xbc, 0x82, 0xa9, 0xd0, 0x6a, 0x26, 0xcd, 0xd2, 0xb5, 0x8d, 0x8c, 0x44, 0x41,
	0x72, 0x9a, 0xf5, 0xbc, 0x77, 0xb1, 0x6f, 0x6d, 0x38, 0x96, 0x52, 0x49, 0x55, 0xfe, 0x82, 0xde,
	0x71, 0xe8, 0xd4, 0x7b, 0x07, 0xe4, 0xb7, 0x49, 0xdb, 0x7e, 0x41, 0xd6, 0x8d, 0xda, 0xc9, 0x59,
	0xb6, 0x39, 0xc6, 0x5f, 0x03, 0x72, 0xef, 0xad, 0x44, 0xdb, 0x4c, 0x04, 0xde, 0xc4, 0x48, 0xf4,
	0x49, 0x88, 0x96, 0x1b, 0xeb, 0xa6, 0x22, 0xcc, 0xfc, 0x25, 0x56, 0x64, 0xf0, 0x67, 0x10, 0xac,
	0xb5, 0x42, 0xa0, 0x57, 0xe4, 0xb4, 0xd9, 0x08, 0x64, 0x41, 0xd4, 0x4e, 0x3a, 0xe3, 0xf1, 0xff,
	0xa0, 0x34, 0x3f, 0xba, 0x49, 0x9b, 0xed, 0x72, 0xc4, 0xdf, 0x02, 0x32, 0x78, 0x0d, 0x15, 0xdc,
	0xe8, 0x36, 0xf4, 0x49, 0x58, 0xf3, 0xeb, 0xdd, 0x42, 0xf8, 0x4b, 0xfc, 0x90, 0xdc, 0xff, 0x0b,
	0x89, 0xef, 0xfd, 0xd5, 0xcb, 0x4f, 0x2f, 0x4a, 0x69, 0xe7, 0xab, 0x3c, 0x15, 0x7a, 0x39, 0x6a,
	0x28, 0x86, 0x1b, 0x8c, 0xd1, 0x4f, 0x8c, 0xe1, 0x96, 0x63, 0xe4, 0x0e, 0x62, 0x58, 0x82, 0x1a,
	0x96, 0x7a, 0x94, 0x5b, 0x91, 0x1f, 0x3b, 0xe9, 0xd9, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7,
	0x3d, 0x52, 0x35, 0x5c, 0x04, 0x00, 0x00,
}
