// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transactionMessage_Eth.proto

package eth

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

type TxArray struct {
	Result               []*TX    `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxArray) Reset()         { *m = TxArray{} }
func (m *TxArray) String() string { return proto.CompactTextString(m) }
func (*TxArray) ProtoMessage()    {}
func (*TxArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b52dc89669b02, []int{0}
}

func (m *TxArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxArray.Unmarshal(m, b)
}
func (m *TxArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxArray.Marshal(b, m, deterministic)
}
func (m *TxArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxArray.Merge(m, src)
}
func (m *TxArray) XXX_Size() int {
	return xxx_messageInfo_TxArray.Size(m)
}
func (m *TxArray) XXX_DiscardUnknown() {
	xxx_messageInfo_TxArray.DiscardUnknown(m)
}

var xxx_messageInfo_TxArray proto.InternalMessageInfo

func (m *TxArray) GetResult() []*TX {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetTransactionHashEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	Txhash               string               `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetTransactionHashEndpointRequest) Reset()         { *m = GetTransactionHashEndpointRequest{} }
func (m *GetTransactionHashEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionHashEndpointRequest) ProtoMessage()    {}
func (*GetTransactionHashEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b52dc89669b02, []int{1}
}

func (m *GetTransactionHashEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionHashEndpointRequest.Unmarshal(m, b)
}
func (m *GetTransactionHashEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionHashEndpointRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionHashEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionHashEndpointRequest.Merge(m, src)
}
func (m *GetTransactionHashEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionHashEndpointRequest.Size(m)
}
func (m *GetTransactionHashEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionHashEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionHashEndpointRequest proto.InternalMessageInfo

func (m *GetTransactionHashEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *GetTransactionHashEndpointRequest) GetTxhash() string {
	if m != nil {
		return m.Txhash
	}
	return ""
}

type GetUnconfirmedTransactionsEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetUnconfirmedTransactionsEndpointRequest) Reset() {
	*m = GetUnconfirmedTransactionsEndpointRequest{}
}
func (m *GetUnconfirmedTransactionsEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*GetUnconfirmedTransactionsEndpointRequest) ProtoMessage()    {}
func (*GetUnconfirmedTransactionsEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b52dc89669b02, []int{2}
}

func (m *GetUnconfirmedTransactionsEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUnconfirmedTransactionsEndpointRequest.Unmarshal(m, b)
}
func (m *GetUnconfirmedTransactionsEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUnconfirmedTransactionsEndpointRequest.Marshal(b, m, deterministic)
}
func (m *GetUnconfirmedTransactionsEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUnconfirmedTransactionsEndpointRequest.Merge(m, src)
}
func (m *GetUnconfirmedTransactionsEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_GetUnconfirmedTransactionsEndpointRequest.Size(m)
}
func (m *GetUnconfirmedTransactionsEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUnconfirmedTransactionsEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUnconfirmedTransactionsEndpointRequest proto.InternalMessageInfo

func (m *GetUnconfirmedTransactionsEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

type PostCreateTransactionsEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	// Height of the block that contains this transaction. If this is an unconfirmed transaction, it will equal -1.
	BlockHeight int32 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// The hash of the transaction.
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	// Array of Ethereum addresses involved in the transaction.
	Addresses []string `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	// The total number of wei exchanged in this transaction.
	Total int32 `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	// The total number of fees—in wei—collected by miners in this transaction. Equal to gas_price * gas_used.
	Fees int32 `protobuf:"varint,6,opt,name=fees,proto3" json:"fees,omitempty"`
	// The size of the transaction in bytes.
	Size int32 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	// The amount of gas used by this transaction.
	GasUsed int32 `protobuf:"varint,8,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// The price of gas—in wei—in this transaction.
	GasPrice int32 `protobuf:"varint,9,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// Address of the peer that sent BlockCypher’s servers this transaction. May be empty.
	RelayedBy string `protobuf:"bytes,10,opt,name=relayed_by,json=relayedBy,proto3" json:"relayed_by,omitempty"`
	// Time this transaction was received by BlockCypher’s servers.
	Received string `protobuf:"bytes,11,opt,name=received,proto3" json:"received,omitempty"`
	// Version number of this transaction.
	Ver int32 `protobuf:"varint,12,opt,name=ver,proto3" json:"ver,omitempty"`
	// true if this is an attempted double spend; false otherwise.
	DoubleSpend bool `protobuf:"varint,13,opt,name=double_spend,json=doubleSpend,proto3" json:"double_spend,omitempty"`
	// Total number of inputs in the transaction.
	VinSz int32 `protobuf:"varint,14,opt,name=vin_sz,json=vinSz,proto3" json:"vin_sz,omitempty"`
	// Total number of outputs in the transaction.
	VoutSz int32 `protobuf:"varint,15,opt,name=vout_sz,json=voutSz,proto3" json:"vout_sz,omitempty"`
	// Number of subsequent blocks, including the block the transaction is in. Unconfirmed transactions have 0 confirmations.
	Confirmations int32 `protobuf:"varint,16,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	// An array object containing a single input with a sequence number (used as a nonce for account balances) and an Ethereum account address. Only contains one input in the array; we still use an array to maintain parity with the Bitcoin API.
	Inputs string `protobuf:"bytes,17,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// An array object containing a single output with value (in wei), script, and an Ethereum account address. Only contains one output in the array; we still use an array to maintain parity with the Bitcoin API.
	Outputs string `protobuf:"bytes,18,opt,name=outputs,proto3" json:"outputs,omitempty"`
	// Optional If this transaction executed a contract which propagated its own subsequent transactions, then this array will be present, containing the hashes of those subsequent internal transactions.
	InternalTxids []string `protobuf:"bytes,19,rep,name=internal_txids,json=internalTxids,proto3" json:"internal_txids,omitempty"`
	// Optional If this transaction was initiated by a contract, this field will be present, conveying the hash of the parent transaction that executed the contract resulting in this transaction (the inverse of an internal_txids hash).
	ParentTx string `protobuf:"bytes,20,opt,name=parent_tx,json=parentTx,proto3" json:"parent_tx,omitempty"`
	// Optional Time at which transaction was included in a block; only present for confirmed transactions.
	Confirmed string `protobuf:"bytes,21,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	// Optional If creating a transaction, can optionally set a higher default gas limit (useful if your recepient is a contract). If not set, default is 21000 gas for external accounts and 80000 for contract accounts.
	GasLimit int32 `protobuf:"varint,22,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// Optional If true, this transaction was used to create a contract and contract account. Note that the contract address (in the outputs field) will be blank until the transaction is confirmed.
	ContractCreation bool `protobuf:"varint,23,opt,name=contract_creation,json=contractCreation,proto3" json:"contract_creation,omitempty"`
	// Optional Number of peers that have sent this transaction to BlockCypher; only present for unconfirmed transactions.
	ReceiveCount int32 `protobuf:"varint,24,opt,name=receive_count,json=receiveCount,proto3" json:"receive_count,omitempty"`
	// Optional Hash of the block that contains this transaction; only present for confirmed transactions.
	BlockHash string `protobuf:"bytes,25,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// Optional Canonical, zero-indexed location of this transaction in a block; only present for confirmed transactions.
	BlockIndex int32 `protobuf:"varint,26,opt,name=block_index,json=blockIndex,proto3" json:"block_index,omitempty"`
	// Optional If this transaction is a double-spend (i.e. double_spend == true) then this is the hash of the transaction it’s double-spending.
	DoubleOf string `protobuf:"bytes,27,opt,name=double_of,json=doubleOf,proto3" json:"double_of,omitempty"`
	// Optional If this transaction has an execution error, then this field will be included (e.g. “out of gas”).
	ExecutionError       string   `protobuf:"bytes,28,opt,name=execution_error,json=executionError,proto3" json:"execution_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostCreateTransactionsEndpointRequest) Reset()         { *m = PostCreateTransactionsEndpointRequest{} }
func (m *PostCreateTransactionsEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*PostCreateTransactionsEndpointRequest) ProtoMessage()    {}
func (*PostCreateTransactionsEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b52dc89669b02, []int{3}
}

func (m *PostCreateTransactionsEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostCreateTransactionsEndpointRequest.Unmarshal(m, b)
}
func (m *PostCreateTransactionsEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostCreateTransactionsEndpointRequest.Marshal(b, m, deterministic)
}
func (m *PostCreateTransactionsEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostCreateTransactionsEndpointRequest.Merge(m, src)
}
func (m *PostCreateTransactionsEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_PostCreateTransactionsEndpointRequest.Size(m)
}
func (m *PostCreateTransactionsEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostCreateTransactionsEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostCreateTransactionsEndpointRequest proto.InternalMessageInfo

func (m *PostCreateTransactionsEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *PostCreateTransactionsEndpointRequest) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *PostCreateTransactionsEndpointRequest) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetFees() int32 {
	if m != nil {
		return m.Fees
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetGasUsed() int32 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetGasPrice() int32 {
	if m != nil {
		return m.GasPrice
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetRelayedBy() string {
	if m != nil {
		return m.RelayedBy
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetReceived() string {
	if m != nil {
		return m.Received
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetVer() int32 {
	if m != nil {
		return m.Ver
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetDoubleSpend() bool {
	if m != nil {
		return m.DoubleSpend
	}
	return false
}

func (m *PostCreateTransactionsEndpointRequest) GetVinSz() int32 {
	if m != nil {
		return m.VinSz
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetVoutSz() int32 {
	if m != nil {
		return m.VoutSz
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetConfirmations() int32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetInputs() string {
	if m != nil {
		return m.Inputs
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetOutputs() string {
	if m != nil {
		return m.Outputs
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetInternalTxids() []string {
	if m != nil {
		return m.InternalTxids
	}
	return nil
}

func (m *PostCreateTransactionsEndpointRequest) GetParentTx() string {
	if m != nil {
		return m.ParentTx
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetConfirmed() string {
	if m != nil {
		return m.Confirmed
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetGasLimit() int32 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetContractCreation() bool {
	if m != nil {
		return m.ContractCreation
	}
	return false
}

func (m *PostCreateTransactionsEndpointRequest) GetReceiveCount() int32 {
	if m != nil {
		return m.ReceiveCount
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetBlockIndex() int32 {
	if m != nil {
		return m.BlockIndex
	}
	return 0
}

func (m *PostCreateTransactionsEndpointRequest) GetDoubleOf() string {
	if m != nil {
		return m.DoubleOf
	}
	return ""
}

func (m *PostCreateTransactionsEndpointRequest) GetExecutionError() string {
	if m != nil {
		return m.ExecutionError
	}
	return ""
}

type PostSendTransactionEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	// A temporary TX, usually returned fully filled.
	Tx *TX `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	// Array of hex-encoded data for you to sign, containing one element for you to sign. Still an array to maintain parity with the Bitcoin API.
	Tosign []string `protobuf:"bytes,3,rep,name=tosign,proto3" json:"tosign,omitempty"`
	// Array of signatures corresponding to all the data in tosign, typically provided by you. Only one signature is required.
	Signatures []string `protobuf:"bytes,4,rep,name=signatures,proto3" json:"signatures,omitempty"`
	// Optional Array of errors in the form “error”:“description-of-error”. This is only returned if there was an error in any stage of transaction generation, and is usually accompanied by a HTTP 400 code.
	Errors               []string `protobuf:"bytes,5,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostSendTransactionEndpointRequest) Reset()         { *m = PostSendTransactionEndpointRequest{} }
func (m *PostSendTransactionEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*PostSendTransactionEndpointRequest) ProtoMessage()    {}
func (*PostSendTransactionEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b52dc89669b02, []int{4}
}

func (m *PostSendTransactionEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostSendTransactionEndpointRequest.Unmarshal(m, b)
}
func (m *PostSendTransactionEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostSendTransactionEndpointRequest.Marshal(b, m, deterministic)
}
func (m *PostSendTransactionEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostSendTransactionEndpointRequest.Merge(m, src)
}
func (m *PostSendTransactionEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_PostSendTransactionEndpointRequest.Size(m)
}
func (m *PostSendTransactionEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostSendTransactionEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostSendTransactionEndpointRequest proto.InternalMessageInfo

func (m *PostSendTransactionEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *PostSendTransactionEndpointRequest) GetTx() *TX {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *PostSendTransactionEndpointRequest) GetTosign() []string {
	if m != nil {
		return m.Tosign
	}
	return nil
}

func (m *PostSendTransactionEndpointRequest) GetSignatures() []string {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *PostSendTransactionEndpointRequest) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

type PostDecodeRawTransactionEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	Tx                   string               `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PostDecodeRawTransactionEndpointRequest) Reset() {
	*m = PostDecodeRawTransactionEndpointRequest{}
}
func (m *PostDecodeRawTransactionEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*PostDecodeRawTransactionEndpointRequest) ProtoMessage()    {}
func (*PostDecodeRawTransactionEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b52dc89669b02, []int{5}
}

func (m *PostDecodeRawTransactionEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostDecodeRawTransactionEndpointRequest.Unmarshal(m, b)
}
func (m *PostDecodeRawTransactionEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostDecodeRawTransactionEndpointRequest.Marshal(b, m, deterministic)
}
func (m *PostDecodeRawTransactionEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostDecodeRawTransactionEndpointRequest.Merge(m, src)
}
func (m *PostDecodeRawTransactionEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_PostDecodeRawTransactionEndpointRequest.Size(m)
}
func (m *PostDecodeRawTransactionEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostDecodeRawTransactionEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostDecodeRawTransactionEndpointRequest proto.InternalMessageInfo

func (m *PostDecodeRawTransactionEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *PostDecodeRawTransactionEndpointRequest) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

type PostPushRawTransactionEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.eth.NetworkAllowingAlias" json:"network,omitempty"`
	Tx                   string               `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PostPushRawTransactionEndpointRequest) Reset()         { *m = PostPushRawTransactionEndpointRequest{} }
func (m *PostPushRawTransactionEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*PostPushRawTransactionEndpointRequest) ProtoMessage()    {}
func (*PostPushRawTransactionEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_690b52dc89669b02, []int{6}
}

func (m *PostPushRawTransactionEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostPushRawTransactionEndpointRequest.Unmarshal(m, b)
}
func (m *PostPushRawTransactionEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostPushRawTransactionEndpointRequest.Marshal(b, m, deterministic)
}
func (m *PostPushRawTransactionEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostPushRawTransactionEndpointRequest.Merge(m, src)
}
func (m *PostPushRawTransactionEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_PostPushRawTransactionEndpointRequest.Size(m)
}
func (m *PostPushRawTransactionEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostPushRawTransactionEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostPushRawTransactionEndpointRequest proto.InternalMessageInfo

func (m *PostPushRawTransactionEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAIN
}

func (m *PostPushRawTransactionEndpointRequest) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

func init() {
	proto.RegisterType((*TxArray)(nil), "fairwaycorp.blockchainprotobuf.eth.TxArray")
	proto.RegisterType((*GetTransactionHashEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.GetTransactionHashEndpointRequest")
	proto.RegisterType((*GetUnconfirmedTransactionsEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.GetUnconfirmedTransactionsEndpointRequest")
	proto.RegisterType((*PostCreateTransactionsEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.PostCreateTransactionsEndpointRequest")
	proto.RegisterType((*PostSendTransactionEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.PostSendTransactionEndpointRequest")
	proto.RegisterType((*PostDecodeRawTransactionEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.PostDecodeRawTransactionEndpointRequest")
	proto.RegisterType((*PostPushRawTransactionEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.eth.PostPushRawTransactionEndpointRequest")
}

func init() { proto.RegisterFile("transactionMessage_Eth.proto", fileDescriptor_690b52dc89669b02) }

var fileDescriptor_690b52dc89669b02 = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xdf, 0x6f, 0x1b, 0x45,
	0x10, 0xc7, 0xe5, 0xb8, 0xb1, 0xe3, 0x49, 0xe2, 0xa6, 0x4b, 0xdb, 0x6c, 0xd3, 0x00, 0xee, 0x41,
	0xa9, 0x11, 0x8a, 0x23, 0x15, 0x09, 0xf1, 0x02, 0x28, 0x0d, 0x51, 0x5b, 0x89, 0x1f, 0xd1, 0xc5,
	0x95, 0x10, 0x2f, 0xa7, 0xf5, 0xdd, 0xf8, 0x6e, 0xd5, 0xf3, 0xae, 0xd9, 0x9d, 0x73, 0xce, 0x79,
	0x81, 0x77, 0x84, 0x78, 0xe2, 0xff, 0x45, 0xbb, 0x7b, 0x76, 0x82, 0x84, 0x44, 0x9f, 0xc2, 0xd3,
	0xed, 0x7c, 0xe6, 0x66, 0x76, 0xf6, 0xbb, 0x33, 0x0b, 0x87, 0x64, 0x84, 0xb2, 0x22, 0x25, 0xa9,
	0xd5, 0xf7, 0x68, 0xad, 0xc8, 0x31, 0x39, 0xa3, 0x62, 0x34, 0x37, 0x9a, 0x34, 0x8b, 0xa6, 0x42,
	0x9a, 0x4b, 0xb1, 0x4c, 0xb5, 0x99, 0x8f, 0x26, 0xa5, 0x4e, 0xdf, 0xa6, 0x85, 0x90, 0xca, 0x3b,
	0x27, 0xd5, 0x74, 0x84, 0x54, 0x1c, 0xec, 0xa7, 0x7a, 0x36, 0xfb, 0x97, 0xe0, 0xe8, 0x35, 0x74,
	0xc7, 0xf5, 0x89, 0x31, 0x62, 0xc9, 0xbe, 0x86, 0x8e, 0x41, 0x5b, 0x95, 0xc4, 0x5b, 0x83, 0xf6,
	0x70, 0xfb, 0xf9, 0x27, 0xa3, 0xff, 0x4e, 0x3c, 0x1a, 0xff, 0x14, 0x37, 0x51, 0xd1, 0x9f, 0x2d,
	0x78, 0xf2, 0x12, 0x69, 0x7c, 0x5d, 0xeb, 0x2b, 0x61, 0x8b, 0x33, 0x95, 0xcd, 0xb5, 0x54, 0x14,
	0xe3, 0x2f, 0x15, 0x5a, 0x62, 0x31, 0x74, 0x15, 0xd2, 0xa5, 0x36, 0x6f, 0x79, 0x6b, 0xd0, 0x1a,
	0xf6, 0x9f, 0x7f, 0xf9, 0x2e, 0xdb, 0xfc, 0x10, 0x42, 0x4e, 0xca, 0x52, 0x5f, 0x4a, 0x95, 0x9f,
	0x94, 0x52, 0xd8, 0x78, 0x95, 0x88, 0x3d, 0x84, 0x0e, 0xd5, 0x85, 0xb0, 0x05, 0xdf, 0x18, 0xb4,
	0x86, 0xbd, 0xb8, 0xb1, 0xa2, 0x5f, 0xe1, 0xd3, 0x97, 0x48, 0x6f, 0x54, 0xaa, 0xd5, 0x54, 0x9a,
	0x19, 0x66, 0x37, 0x6a, 0xb3, 0xb7, 0x50, 0x58, 0xf4, 0x57, 0x17, 0x9e, 0x9e, 0x6b, 0x4b, 0xa7,
	0x06, 0x05, 0xe1, 0x2d, 0xef, 0xce, 0x9e, 0xc0, 0x8e, 0x8f, 0x4b, 0x0a, 0x94, 0x79, 0x41, 0x5e,
	0x9c, 0xcd, 0x78, 0xdb, 0xb3, 0x57, 0x1e, 0x31, 0x06, 0x77, 0xbc, 0x6e, 0x6d, 0xaf, 0x9b, 0x5f,
	0xb3, 0x43, 0xe8, 0x89, 0x2c, 0x33, 0x68, 0x2d, 0x5a, 0x7e, 0x67, 0xd0, 0x1e, 0xf6, 0xe2, 0x6b,
	0xc0, 0xee, 0xc3, 0x26, 0x69, 0x12, 0x25, 0xdf, 0xf4, 0xd9, 0x82, 0xe1, 0xf2, 0x4c, 0x11, 0x2d,
	0xef, 0x78, 0xe8, 0xd7, 0x8e, 0x59, 0x79, 0x85, 0xbc, 0x1b, 0x98, 0x5b, 0xb3, 0x47, 0xb0, 0x95,
	0x0b, 0x9b, 0x54, 0x16, 0x33, 0xbe, 0xe5, 0x79, 0x37, 0x17, 0xf6, 0x8d, 0xc5, 0x8c, 0x3d, 0x86,
	0x9e, 0x73, 0xcd, 0x8d, 0x4c, 0x91, 0xf7, 0xbc, 0xcf, 0xfd, 0x7b, 0xee, 0x6c, 0xf6, 0x3e, 0x80,
	0xc1, 0x52, 0x2c, 0x31, 0x4b, 0x26, 0x4b, 0x0e, 0xbe, 0xda, 0x5e, 0x43, 0x5e, 0x2c, 0xd9, 0x01,
	0x6c, 0x19, 0x4c, 0x51, 0x2e, 0x30, 0xe3, 0xdb, 0xde, 0xb9, 0xb6, 0xd9, 0x1e, 0xb4, 0x17, 0x68,
	0xf8, 0x8e, 0xcf, 0xe8, 0x96, 0x4e, 0x97, 0x4c, 0x57, 0x93, 0x12, 0x13, 0x3b, 0x47, 0x95, 0xf1,
	0xdd, 0x41, 0x6b, 0xb8, 0x15, 0x6f, 0x07, 0x76, 0xe1, 0x10, 0x7b, 0x00, 0x9d, 0x85, 0x54, 0x89,
	0xbd, 0xe2, 0xfd, 0x70, 0xcc, 0x85, 0x54, 0x17, 0x57, 0x6c, 0x1f, 0xba, 0x0b, 0x5d, 0x91, 0xe3,
	0x77, 0x3d, 0xef, 0x38, 0xf3, 0xe2, 0x8a, 0x7d, 0x0c, 0xbb, 0x4d, 0x93, 0x09, 0x7f, 0xbd, 0x7c,
	0xcf, 0xbb, 0xff, 0x09, 0x5d, 0x9f, 0x4a, 0x35, 0xaf, 0xc8, 0xf2, 0x7b, 0xa1, 0x4f, 0x83, 0xc5,
	0x38, 0x74, 0x75, 0x45, 0xde, 0xc1, 0xbc, 0x63, 0x65, 0xb2, 0xa7, 0xd0, 0x97, 0x8a, 0xd0, 0x28,
	0x51, 0x26, 0x54, 0xcb, 0xcc, 0xf2, 0xf7, 0xfc, 0x85, 0xec, 0xae, 0xe8, 0xd8, 0x41, 0xa7, 0xdd,
	0x5c, 0x18, 0x54, 0x94, 0x50, 0xcd, 0xef, 0x07, 0x01, 0x02, 0x18, 0xd7, 0xee, 0x3e, 0xd7, 0x03,
	0xc0, 0x1f, 0x04, 0xe9, 0xd6, 0x60, 0x25, 0x7b, 0x29, 0x67, 0x92, 0xf8, 0xc3, 0xb5, 0xec, 0xdf,
	0x39, 0x9b, 0x7d, 0x06, 0xf7, 0x52, 0xad, 0xc8, 0x88, 0x94, 0x92, 0xd4, 0xf5, 0xb0, 0xd4, 0x8a,
	0xef, 0x7b, 0xb9, 0xf6, 0x56, 0x8e, 0xd3, 0x86, 0xb3, 0x8f, 0x60, 0xb7, 0x11, 0x3d, 0x49, 0x75,
	0xa5, 0x88, 0x73, 0x9f, 0x6d, 0xa7, 0x81, 0xa7, 0x8e, 0xb9, 0x8b, 0x6c, 0x7a, 0xd2, 0xb5, 0xdd,
	0xa3, 0x50, 0x4d, 0xe8, 0x48, 0xd7, 0x7b, 0x1f, 0x42, 0x68, 0xcf, 0x44, 0xaa, 0x0c, 0x6b, 0x7e,
	0xe0, 0x33, 0x84, 0x88, 0xd7, 0x8e, 0xb8, 0x72, 0x9b, 0xbb, 0xd3, 0x53, 0xfe, 0x38, 0x9c, 0x34,
	0x80, 0x1f, 0xa7, 0xec, 0x19, 0xdc, 0xc5, 0x1a, 0xd3, 0xca, 0x95, 0x93, 0xa0, 0x31, 0xda, 0xf0,
	0x43, 0xff, 0x4b, 0x7f, 0x8d, 0xcf, 0x1c, 0x8d, 0x7e, 0xdb, 0x80, 0xc8, 0xcd, 0xe5, 0x05, 0xaa,
	0x9b, 0x6f, 0xc2, 0x6d, 0x0c, 0xe5, 0x17, 0xb0, 0x41, 0xb5, 0x1f, 0xc5, 0x77, 0x7f, 0x61, 0x37,
	0xa8, 0xf6, 0x6f, 0x9c, 0xb6, 0x32, 0x57, 0xbc, 0xed, 0x3b, 0xa0, 0xb1, 0xd8, 0x07, 0x00, 0xee,
	0x2b, 0xa8, 0x32, 0xeb, 0x71, 0xbd, 0x41, 0x5c, 0x9c, 0x57, 0xc2, 0xf2, 0xcd, 0x10, 0x17, 0xac,
	0xe8, 0x8f, 0x16, 0x3c, 0x73, 0x12, 0x7c, 0x8b, 0xa9, 0xce, 0x30, 0x16, 0x97, 0xb7, 0xac, 0x43,
	0x7f, 0xad, 0x43, 0xcf, 0x9d, 0x2f, 0xfa, 0xbd, 0x15, 0x9e, 0xca, 0xf3, 0xca, 0x16, 0xff, 0x7b,
	0x35, 0x2f, 0xbe, 0xf9, 0xf9, 0xab, 0x5c, 0x52, 0x51, 0x4d, 0x46, 0xa9, 0x9e, 0x1d, 0x37, 0xe9,
	0x8f, 0x5c, 0xfe, 0xe3, 0xeb, 0xfc, 0x47, 0xab, 0x0d, 0x8e, 0xfd, 0x22, 0x3d, 0xca, 0x51, 0x1d,
	0xe5, 0xfa, 0x18, 0xa9, 0x98, 0x74, 0x3c, 0xfa, 0xfc, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8,
	0x5d, 0x7b, 0x09, 0xbb, 0x07, 0x00, 0x00,
}
