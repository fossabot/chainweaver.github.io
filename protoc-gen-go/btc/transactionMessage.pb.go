// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transactionMessage.proto

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

type TransactionHashEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Txhash  string               `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	//Filters TXInputs/TXOutputs, if unset, default is 20.
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	//Filters TX to only include TXInputs from this input index and above.
	Instart int32 `protobuf:"varint,4,opt,name=instart,proto3" json:"instart,omitempty"`
	//Filters TX to only include TXOutputs from this output index and above.
	Outstart int32 `protobuf:"varint,5,opt,name=outstart,proto3" json:"outstart,omitempty"`
	//If true, includes hex-encoded raw transaction; false by default.
	IncludeHex bool `protobuf:"varint,6,opt,name=includeHex,proto3" json:"includeHex,omitempty"`
	//If true, includes the confidence attribute (useful for unconfirmed transactions). For more info about this figure, check the Confidence Factor documentation.
	IncludeConfidence    bool     `protobuf:"varint,7,opt,name=includeConfidence,proto3" json:"includeConfidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionHashEndpointRequest) Reset()         { *m = TransactionHashEndpointRequest{} }
func (m *TransactionHashEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionHashEndpointRequest) ProtoMessage()    {}
func (*TransactionHashEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a17a22c39ec5, []int{0}
}

func (m *TransactionHashEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionHashEndpointRequest.Unmarshal(m, b)
}
func (m *TransactionHashEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionHashEndpointRequest.Marshal(b, m, deterministic)
}
func (m *TransactionHashEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionHashEndpointRequest.Merge(m, src)
}
func (m *TransactionHashEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionHashEndpointRequest.Size(m)
}
func (m *TransactionHashEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionHashEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionHashEndpointRequest proto.InternalMessageInfo

func (m *TransactionHashEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *TransactionHashEndpointRequest) GetTxhash() string {
	if m != nil {
		return m.Txhash
	}
	return ""
}

func (m *TransactionHashEndpointRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *TransactionHashEndpointRequest) GetInstart() int32 {
	if m != nil {
		return m.Instart
	}
	return 0
}

func (m *TransactionHashEndpointRequest) GetOutstart() int32 {
	if m != nil {
		return m.Outstart
	}
	return 0
}

func (m *TransactionHashEndpointRequest) GetIncludeHex() bool {
	if m != nil {
		return m.IncludeHex
	}
	return false
}

func (m *TransactionHashEndpointRequest) GetIncludeConfidence() bool {
	if m != nil {
		return m.IncludeConfidence
	}
	return false
}

type UnconfirmedTransactionsEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UnconfirmedTransactionsEndpointRequest) Reset() {
	*m = UnconfirmedTransactionsEndpointRequest{}
}
func (m *UnconfirmedTransactionsEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*UnconfirmedTransactionsEndpointRequest) ProtoMessage()    {}
func (*UnconfirmedTransactionsEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a17a22c39ec5, []int{1}
}

func (m *UnconfirmedTransactionsEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnconfirmedTransactionsEndpointRequest.Unmarshal(m, b)
}
func (m *UnconfirmedTransactionsEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnconfirmedTransactionsEndpointRequest.Marshal(b, m, deterministic)
}
func (m *UnconfirmedTransactionsEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnconfirmedTransactionsEndpointRequest.Merge(m, src)
}
func (m *UnconfirmedTransactionsEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_UnconfirmedTransactionsEndpointRequest.Size(m)
}
func (m *UnconfirmedTransactionsEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnconfirmedTransactionsEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnconfirmedTransactionsEndpointRequest proto.InternalMessageInfo

func (m *UnconfirmedTransactionsEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

type NewTransactionEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	//	Height of the block that contains this transaction. If this is an unconfirmed transaction, it will equal -1.
	BlockHeight int32 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	//The hash of the transaction. While reasonably unique, using hashes as identifiers may be unsafe.
	Hash string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	//Array of bitcoin public addresses involved in the transaction.
	Addresses []string `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	//The total number of satoshis exchanged in this transaction.
	Total int32 `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
	//The total number of fees—in satoshis—collected by miners in this transaction.
	Fees int32 `protobuf:"varint,6,opt,name=fees,proto3" json:"fees,omitempty"`
	//The size of the transaction in bytes.
	Size int32 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	//The likelihood that this transaction will make it to the next block; reflects the preference level miners have to include this transaction. Can be high, medium or low.
	Preference string `protobuf:"bytes,8,opt,name=preference,proto3" json:"preference,omitempty"`
	//Address of the peer that sent Block’s servers this transaction.
	RelayedBy string `protobuf:"bytes,9,opt,name=relayed_by,json=relayedBy,proto3" json:"relayed_by,omitempty"`
	//Time this transaction was received by Block’s servers.
	Received string `protobuf:"bytes,10,opt,name=received,proto3" json:"received,omitempty"`
	//Version number, typically 1 for Bitcoin transactions.
	Ver int32 `protobuf:"varint,11,opt,name=ver,proto3" json:"ver,omitempty"`
	//Time when transaction can be valid. Can be interpreted in two ways: if less than 500 million, refers to block height. If more, refers to Unix epoch time.
	LockTime int32 `protobuf:"varint,12,opt,name=lock_time,json=lockTime,proto3" json:"lock_time,omitempty"`
	//true if this is an attempted double spend; false otherwise.
	DoubleSpend bool `protobuf:"varint,13,opt,name=double_spend,json=doubleSpend,proto3" json:"double_spend,omitempty"`
	//Total number of inputs in the transaction.
	VinSz int32 `protobuf:"varint,14,opt,name=vin_sz,json=vinSz,proto3" json:"vin_sz,omitempty"`
	//Total number of outputs in the transaction.
	VoutSz int32 `protobuf:"varint,15,opt,name=vout_sz,json=voutSz,proto3" json:"vout_sz,omitempty"`
	//Number of subsequent blocks, including the block the transaction is in. Unconfirmed transactions have 0 confirmations.
	Confirmations int32 `protobuf:"varint,16,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	//TXInput Array, limited to 20 by default.
	Inputs []*TXInput `protobuf:"bytes,17,rep,name=inputs,proto3" json:"inputs,omitempty"`
	//TXOutput Array, limited to 20 by default.
	Outputs []*TXOutput `protobuf:"bytes,18,rep,name=outputs,proto3" json:"outputs,omitempty"`
	//Optional Returns true if this transaction has opted in to Replace-By-Fee (RBF), either true or not present. You can read more about Opt-In RBF here.
	OptInRbf bool `protobuf:"varint,19,opt,name=opt_in_rbf,json=optInRbf,proto3" json:"opt_in_rbf,omitempty"`
	//Optional The percentage chance this transaction will not be double-spent against, if unconfirmed. For more information, check the section on Confidence Factor.
	Confidence float32 `protobuf:"fixed32,20,opt,name=confidence,proto3" json:"confidence,omitempty"`
	//Optional Time at which transaction was included in a block; only present for confirmed transactions.
	Confirmed string `protobuf:"bytes,21,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	//Optional Number of peers that have sent this transaction to Block; only present for unconfirmed transactions.
	ReceiveCount int32 `protobuf:"varint,22,opt,name=receive_count,json=receiveCount,proto3" json:"receive_count,omitempty"`
	//Optional Address Block will use to send back your change, if you constructed this transaction. If not set, defaults to the address from which the coins were originally sent.
	ChangeAddress string `protobuf:"bytes,23,opt,name=change_address,json=changeAddress,proto3" json:"change_address,omitempty"`
	//Optional Hash of the block that contains this transaction; only present for confirmed transactions.
	BlockHash string `protobuf:"bytes,24,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	//Optional Canonical, zero-indexed location of this transaction in a block; only present for confirmed transactions.
	BlockIndex int32 `protobuf:"varint,25,opt,name=block_index,json=blockIndex,proto3" json:"block_index,omitempty"`
	//Optional If this transaction is a double-spend (i.e. double_spend == true) then this is the hash of the transaction it’s double-spending.
	DoubleOf string `protobuf:"bytes,26,opt,name=double_of,json=doubleOf,proto3" json:"double_of,omitempty"`
	//Optional Returned if this transaction contains an OP_RETURN associated with a known data protocol. Data protocols currently detected: blockchainid ; openassets ; factom ; colu ; coinspark ; omni
	DataProtocol string `protobuf:"bytes,27,opt,name=data_protocol,json=dataProtocol,proto3" json:"data_protocol,omitempty"`
	//Optional Hex-encoded bytes of the transaction, as sent over the network.
	Hex string `protobuf:"bytes,28,opt,name=hex,proto3" json:"hex,omitempty"`
	//Optional If there are more transaction inptus that couldn’t fit into the TXInput array, this is the Block URL to query the next set of TXInputs (within a TX object).
	NextInputs string `protobuf:"bytes,29,opt,name=next_inputs,json=nextInputs,proto3" json:"next_inputs,omitempty"`
	//Optional If there are more transaction outputs that couldn’t fit into the TXOutput array, this is the Block URL to query the next set of TXOutputs(within a TX object).
	NextOutputs          string   `protobuf:"bytes,30,opt,name=next_outputs,json=nextOutputs,proto3" json:"next_outputs,omitempty"`
	IncludeToSignTx      bool     `protobuf:"varint,31,opt,name=includeToSignTx,proto3" json:"includeToSignTx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewTransactionEndpointRequest) Reset()         { *m = NewTransactionEndpointRequest{} }
func (m *NewTransactionEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*NewTransactionEndpointRequest) ProtoMessage()    {}
func (*NewTransactionEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a17a22c39ec5, []int{2}
}

func (m *NewTransactionEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewTransactionEndpointRequest.Unmarshal(m, b)
}
func (m *NewTransactionEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewTransactionEndpointRequest.Marshal(b, m, deterministic)
}
func (m *NewTransactionEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewTransactionEndpointRequest.Merge(m, src)
}
func (m *NewTransactionEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_NewTransactionEndpointRequest.Size(m)
}
func (m *NewTransactionEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewTransactionEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewTransactionEndpointRequest proto.InternalMessageInfo

func (m *NewTransactionEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *NewTransactionEndpointRequest) GetBlockHeight() int32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *NewTransactionEndpointRequest) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetFees() int32 {
	if m != nil {
		return m.Fees
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetPreference() string {
	if m != nil {
		return m.Preference
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetRelayedBy() string {
	if m != nil {
		return m.RelayedBy
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetReceived() string {
	if m != nil {
		return m.Received
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetVer() int32 {
	if m != nil {
		return m.Ver
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetLockTime() int32 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetDoubleSpend() bool {
	if m != nil {
		return m.DoubleSpend
	}
	return false
}

func (m *NewTransactionEndpointRequest) GetVinSz() int32 {
	if m != nil {
		return m.VinSz
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetVoutSz() int32 {
	if m != nil {
		return m.VoutSz
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetConfirmations() int32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetInputs() []*TXInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *NewTransactionEndpointRequest) GetOutputs() []*TXOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *NewTransactionEndpointRequest) GetOptInRbf() bool {
	if m != nil {
		return m.OptInRbf
	}
	return false
}

func (m *NewTransactionEndpointRequest) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetConfirmed() string {
	if m != nil {
		return m.Confirmed
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetReceiveCount() int32 {
	if m != nil {
		return m.ReceiveCount
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetChangeAddress() string {
	if m != nil {
		return m.ChangeAddress
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetBlockIndex() int32 {
	if m != nil {
		return m.BlockIndex
	}
	return 0
}

func (m *NewTransactionEndpointRequest) GetDoubleOf() string {
	if m != nil {
		return m.DoubleOf
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetDataProtocol() string {
	if m != nil {
		return m.DataProtocol
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetHex() string {
	if m != nil {
		return m.Hex
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetNextInputs() string {
	if m != nil {
		return m.NextInputs
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetNextOutputs() string {
	if m != nil {
		return m.NextOutputs
	}
	return ""
}

func (m *NewTransactionEndpointRequest) GetIncludeToSignTx() bool {
	if m != nil {
		return m.IncludeToSignTx
	}
	return false
}

type SendTransactionEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	//A temporary TX, usually returned fully filled but missing input scripts.
	Tx *TX `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	//Array of hex-encoded data for you to sign, one for each input.
	Tosign []string `protobuf:"bytes,3,rep,name=tosign,proto3" json:"tosign,omitempty"`
	//Array of signatures corresponding to all the data in tosign, typically provided by you.
	Signatures []string `protobuf:"bytes,4,rep,name=signatures,proto3" json:"signatures,omitempty"`
	//Array of public keys corresponding to each signature. In general, these are provided by you, and correspond to the signatures you provide.
	Pubkeys []string `protobuf:"bytes,5,rep,name=pubkeys,proto3" json:"pubkeys,omitempty"`
	//Optional Array of hex-encoded, work-in-progress transactions; optionally returned to validate the tosign data locally.
	TosignTx []string `protobuf:"bytes,6,rep,name=tosign_tx,json=tosignTx,proto3" json:"tosign_tx,omitempty"`
	//Optional Array of errors in the form “error”:“description-of-error”. This is only returned if there was an error in any stage of transaction generation, and is usually accompanied by a HTTP 400 code.
	Errors               []*TXerror `protobuf:"bytes,7,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SendTransactionEndpointRequest) Reset()         { *m = SendTransactionEndpointRequest{} }
func (m *SendTransactionEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*SendTransactionEndpointRequest) ProtoMessage()    {}
func (*SendTransactionEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a17a22c39ec5, []int{3}
}

func (m *SendTransactionEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendTransactionEndpointRequest.Unmarshal(m, b)
}
func (m *SendTransactionEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendTransactionEndpointRequest.Marshal(b, m, deterministic)
}
func (m *SendTransactionEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTransactionEndpointRequest.Merge(m, src)
}
func (m *SendTransactionEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_SendTransactionEndpointRequest.Size(m)
}
func (m *SendTransactionEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTransactionEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendTransactionEndpointRequest proto.InternalMessageInfo

func (m *SendTransactionEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *SendTransactionEndpointRequest) GetTx() *TX {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *SendTransactionEndpointRequest) GetTosign() []string {
	if m != nil {
		return m.Tosign
	}
	return nil
}

func (m *SendTransactionEndpointRequest) GetSignatures() []string {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *SendTransactionEndpointRequest) GetPubkeys() []string {
	if m != nil {
		return m.Pubkeys
	}
	return nil
}

func (m *SendTransactionEndpointRequest) GetTosignTx() []string {
	if m != nil {
		return m.TosignTx
	}
	return nil
}

func (m *SendTransactionEndpointRequest) GetErrors() []*TXerror {
	if m != nil {
		return m.Errors
	}
	return nil
}

type PushRawTransactionEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Tx                   string               `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PushRawTransactionEndpointRequest) Reset()         { *m = PushRawTransactionEndpointRequest{} }
func (m *PushRawTransactionEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*PushRawTransactionEndpointRequest) ProtoMessage()    {}
func (*PushRawTransactionEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a17a22c39ec5, []int{4}
}

func (m *PushRawTransactionEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRawTransactionEndpointRequest.Unmarshal(m, b)
}
func (m *PushRawTransactionEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRawTransactionEndpointRequest.Marshal(b, m, deterministic)
}
func (m *PushRawTransactionEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRawTransactionEndpointRequest.Merge(m, src)
}
func (m *PushRawTransactionEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_PushRawTransactionEndpointRequest.Size(m)
}
func (m *PushRawTransactionEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRawTransactionEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRawTransactionEndpointRequest proto.InternalMessageInfo

func (m *PushRawTransactionEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *PushRawTransactionEndpointRequest) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

type DecodeRawTransactionEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Tx                   string               `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DecodeRawTransactionEndpointRequest) Reset()         { *m = DecodeRawTransactionEndpointRequest{} }
func (m *DecodeRawTransactionEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*DecodeRawTransactionEndpointRequest) ProtoMessage()    {}
func (*DecodeRawTransactionEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a17a22c39ec5, []int{5}
}

func (m *DecodeRawTransactionEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecodeRawTransactionEndpointRequest.Unmarshal(m, b)
}
func (m *DecodeRawTransactionEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecodeRawTransactionEndpointRequest.Marshal(b, m, deterministic)
}
func (m *DecodeRawTransactionEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecodeRawTransactionEndpointRequest.Merge(m, src)
}
func (m *DecodeRawTransactionEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_DecodeRawTransactionEndpointRequest.Size(m)
}
func (m *DecodeRawTransactionEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecodeRawTransactionEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecodeRawTransactionEndpointRequest proto.InternalMessageInfo

func (m *DecodeRawTransactionEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *DecodeRawTransactionEndpointRequest) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

type DataEndpointRequest struct {
	Network NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	//The string representing the data to embed, can be either hex-encoded or plaintext.
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	//Optional Your BlockCypher API token, can either be included here or as a URL Parameter in your request.
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	//Optional The encoding of your data, can be either string (for plaintext) or hex (for hex-encoded). If not set, defaults to hex.
	Encoding string `protobuf:"bytes,4,opt,name=encoding,proto3" json:"encoding,omitempty"`
	//Optional The hash of the transaction containing your data; only part of return object.
	Hash                 string   `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataEndpointRequest) Reset()         { *m = DataEndpointRequest{} }
func (m *DataEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*DataEndpointRequest) ProtoMessage()    {}
func (*DataEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a17a22c39ec5, []int{6}
}

func (m *DataEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataEndpointRequest.Unmarshal(m, b)
}
func (m *DataEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataEndpointRequest.Marshal(b, m, deterministic)
}
func (m *DataEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataEndpointRequest.Merge(m, src)
}
func (m *DataEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_DataEndpointRequest.Size(m)
}
func (m *DataEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataEndpointRequest proto.InternalMessageInfo

func (m *DataEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *DataEndpointRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *DataEndpointRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DataEndpointRequest) GetEncoding() string {
	if m != nil {
		return m.Encoding
	}
	return ""
}

func (m *DataEndpointRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type TransactionPropagationEndpointRequest struct {
	Network              NetworkAllowingAlias `protobuf:"varint,1,opt,name=network,proto3,enum=fairwaycorp.blockchainprotobuf.btc.NetworkAllowingAlias" json:"network,omitempty"`
	Txhash               string               `protobuf:"bytes,2,opt,name=txhash,proto3" json:"txhash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransactionPropagationEndpointRequest) Reset()         { *m = TransactionPropagationEndpointRequest{} }
func (m *TransactionPropagationEndpointRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionPropagationEndpointRequest) ProtoMessage()    {}
func (*TransactionPropagationEndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dd3a17a22c39ec5, []int{7}
}

func (m *TransactionPropagationEndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionPropagationEndpointRequest.Unmarshal(m, b)
}
func (m *TransactionPropagationEndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionPropagationEndpointRequest.Marshal(b, m, deterministic)
}
func (m *TransactionPropagationEndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionPropagationEndpointRequest.Merge(m, src)
}
func (m *TransactionPropagationEndpointRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionPropagationEndpointRequest.Size(m)
}
func (m *TransactionPropagationEndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionPropagationEndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionPropagationEndpointRequest proto.InternalMessageInfo

func (m *TransactionPropagationEndpointRequest) GetNetwork() NetworkAllowingAlias {
	if m != nil {
		return m.Network
	}
	return NetworkAllowingAlias_MAINNET
}

func (m *TransactionPropagationEndpointRequest) GetTxhash() string {
	if m != nil {
		return m.Txhash
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionHashEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.TransactionHashEndpointRequest")
	proto.RegisterType((*UnconfirmedTransactionsEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.UnconfirmedTransactionsEndpointRequest")
	proto.RegisterType((*NewTransactionEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.NewTransactionEndpointRequest")
	proto.RegisterType((*SendTransactionEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.SendTransactionEndpointRequest")
	proto.RegisterType((*PushRawTransactionEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.PushRawTransactionEndpointRequest")
	proto.RegisterType((*DecodeRawTransactionEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.DecodeRawTransactionEndpointRequest")
	proto.RegisterType((*DataEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.DataEndpointRequest")
	proto.RegisterType((*TransactionPropagationEndpointRequest)(nil), "fairwaycorp.blockchainprotobuf.btc.TransactionPropagationEndpointRequest")
}

func init() { proto.RegisterFile("transactionMessage.proto", fileDescriptor_8dd3a17a22c39ec5) }

var fileDescriptor_8dd3a17a22c39ec5 = []byte{
	// 976 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xd6, 0xd8, 0xf1, 0x5f, 0xe7, 0x67, 0x77, 0x3b, 0xfb, 0xd3, 0x64, 0x93, 0xac, 0xe3, 0x65,
	0x57, 0x96, 0xd8, 0x38, 0x52, 0x90, 0x10, 0x17, 0x84, 0xb2, 0x59, 0x50, 0x72, 0x60, 0x13, 0x4d,
	0x8c, 0x84, 0xb8, 0x8c, 0x7a, 0x66, 0xca, 0x33, 0xad, 0x8c, 0xbb, 0x87, 0xe9, 0x9e, 0xc4, 0x8e,
	0xb8, 0x23, 0xce, 0xbc, 0x03, 0xcf, 0x81, 0xc4, 0x8b, 0xf0, 0x28, 0xa8, 0xab, 0xc7, 0x3f, 0x2c,
	0x42, 0x84, 0x8b, 0x39, 0xb9, 0xeb, 0xab, 0xae, 0x2a, 0xbb, 0xbe, 0xaf, 0xab, 0x4c, 0x98, 0x29,
	0xb8, 0xd4, 0x3c, 0x32, 0x42, 0xc9, 0x6f, 0x40, 0x6b, 0x9e, 0xc0, 0x20, 0x2f, 0x94, 0x51, 0xb4,
	0x37, 0xe2, 0xa2, 0xb8, 0xe5, 0xd3, 0x48, 0x15, 0xf9, 0x20, 0xcc, 0x54, 0x74, 0x1d, 0xa5, 0x5c,
	0x48, 0x74, 0x86, 0xe5, 0x68, 0x10, 0x9a, 0x68, 0x67, 0x3b, 0x52, 0xe3, 0xf1, 0x07, 0x81, 0xbd,
	0x5f, 0x6b, 0x64, 0x7f, 0xb8, 0xc8, 0x7a, 0xc6, 0x75, 0xfa, 0x95, 0x8c, 0x73, 0x25, 0xa4, 0xf1,
	0xe1, 0x87, 0x12, 0xb4, 0xa1, 0x3e, 0x69, 0x49, 0x30, 0xb7, 0xaa, 0xb8, 0x66, 0x5e, 0xd7, 0xeb,
	0x6f, 0x1d, 0x7f, 0x3e, 0xf8, 0xf7, 0x6a, 0x83, 0xf7, 0x2e, 0xe4, 0x24, 0xcb, 0xd4, 0xad, 0x90,
	0xc9, 0x49, 0x26, 0xb8, 0xf6, 0x67, 0x89, 0xe8, 0x53, 0xd2, 0x34, 0x93, 0x94, 0xeb, 0x94, 0xd5,
	0xba, 0x5e, 0xbf, 0xe3, 0x57, 0x16, 0x7d, 0x4c, 0x1a, 0x99, 0x18, 0x0b, 0xc3, 0xea, 0x5d, 0xaf,
	0xdf, 0xf0, 0x9d, 0x41, 0x19, 0x69, 0x09, 0xa9, 0x0d, 0x2f, 0x0c, 0x5b, 0x43, 0x7c, 0x66, 0xd2,
	0x1d, 0xd2, 0x56, 0xa5, 0x71, 0xae, 0x06, 0xba, 0xe6, 0x36, 0xdd, 0x27, 0x44, 0xc8, 0x28, 0x2b,
	0x63, 0x38, 0x83, 0x09, 0x6b, 0x76, 0xbd, 0x7e, 0xdb, 0x5f, 0x42, 0xe8, 0x1b, 0xf2, 0xa8, 0xb2,
	0x4e, 0x95, 0x1c, 0x89, 0x18, 0x64, 0x04, 0xac, 0x85, 0xd7, 0xfe, 0xee, 0xe8, 0xfd, 0x48, 0x5e,
	0x7f, 0x2b, 0x23, 0x6b, 0x17, 0x63, 0x88, 0x97, 0x5a, 0xa6, 0x57, 0xd0, 0xaf, 0xde, 0x6f, 0x6d,
	0xb2, 0xf7, 0x1e, 0x6e, 0x97, 0xca, 0xae, 0x82, 0xa5, 0x03, 0xb2, 0x81, 0x71, 0x41, 0x0a, 0x22,
	0x49, 0x0d, 0x72, 0xd5, 0xf0, 0xd7, 0x11, 0x3b, 0x43, 0x88, 0x52, 0xb2, 0x86, 0x34, 0xd6, 0x91,
	0x46, 0x3c, 0xd3, 0x5d, 0xd2, 0xe1, 0x71, 0x5c, 0x80, 0xd6, 0xa0, 0xd9, 0x5a, 0xb7, 0xde, 0xef,
	0xf8, 0x0b, 0xc0, 0x52, 0x6c, 0x94, 0xe1, 0x59, 0xc5, 0x97, 0x33, 0x6c, 0x9e, 0x11, 0x80, 0x46,
	0x9a, 0x1a, 0x3e, 0x9e, 0x2d, 0xa6, 0xc5, 0x9d, 0xe3, 0xa4, 0xe1, 0xe3, 0xd9, 0x92, 0x9a, 0x17,
	0x30, 0x82, 0x02, 0xd9, 0x6a, 0x63, 0xd5, 0x25, 0x84, 0xee, 0x11, 0x52, 0x40, 0xc6, 0xa7, 0x10,
	0x07, 0xe1, 0x94, 0x75, 0xd0, 0xdf, 0xa9, 0x90, 0xb7, 0x53, 0xab, 0x97, 0x02, 0x22, 0x10, 0x37,
	0x10, 0x33, 0x82, 0xce, 0xb9, 0x4d, 0x1f, 0x92, 0xfa, 0x0d, 0x14, 0x6c, 0x1d, 0xab, 0xd9, 0x23,
	0x7d, 0x4e, 0x3a, 0xf8, 0xf3, 0x8d, 0x18, 0x03, 0xdb, 0x70, 0xf2, 0xb2, 0xc0, 0x50, 0x8c, 0xc1,
	0x36, 0x27, 0x56, 0x65, 0x98, 0x41, 0xa0, 0x73, 0x90, 0x31, 0xdb, 0x44, 0xe5, 0xac, 0x3b, 0xec,
	0xca, 0x42, 0xf4, 0x09, 0x69, 0xde, 0x08, 0x19, 0xe8, 0x3b, 0xb6, 0xe5, 0x7e, 0xeb, 0x8d, 0x90,
	0x57, 0x77, 0xf4, 0x19, 0x69, 0xdd, 0xa8, 0xd2, 0x58, 0xfc, 0x01, 0xe2, 0x4d, 0x6b, 0x5e, 0xdd,
	0xd1, 0x8f, 0xc9, 0x66, 0xa5, 0x30, 0x8e, 0xca, 0x62, 0x0f, 0xd1, 0xfd, 0x57, 0x90, 0x9e, 0x92,
	0xa6, 0x90, 0x79, 0x69, 0x34, 0x7b, 0xd4, 0xad, 0xf7, 0xd7, 0x8f, 0x3f, 0xb9, 0x0f, 0xd1, 0xc3,
	0xef, 0xce, 0x6d, 0x8c, 0x5f, 0x85, 0xd2, 0xaf, 0x49, 0x4b, 0x95, 0x06, 0xb3, 0x50, 0xcc, 0xf2,
	0xe6, 0x7e, 0x59, 0x2e, 0x30, 0xc8, 0x9f, 0x05, 0xd3, 0x5d, 0x42, 0x54, 0x6e, 0x02, 0x21, 0x83,
	0x22, 0x1c, 0xb1, 0x6d, 0xec, 0x41, 0x5b, 0xe5, 0xe6, 0x5c, 0xfa, 0xe1, 0xc8, 0xb2, 0x15, 0x2d,
	0xde, 0xd6, 0xe3, 0xae, 0xd7, 0xaf, 0xf9, 0x4b, 0x88, 0x55, 0xca, 0xfc, 0x49, 0xb1, 0x27, 0x8e,
	0xac, 0x39, 0x40, 0x5f, 0x92, 0xcd, 0x8a, 0x9c, 0x20, 0x52, 0xa5, 0x34, 0xec, 0x29, 0xb6, 0x63,
	0xa3, 0x02, 0x4f, 0x2d, 0x46, 0x5f, 0x91, 0xad, 0x28, 0xe5, 0x32, 0x81, 0xa0, 0x92, 0x18, 0x7b,
	0x86, 0x79, 0x36, 0x1d, 0x7a, 0xe2, 0x40, 0xab, 0x8b, 0x4a, 0xca, 0x56, 0xad, 0xcc, 0x95, 0x72,
	0x42, 0xb6, 0x92, 0x7d, 0x41, 0x9c, 0xaa, 0x03, 0x21, 0x63, 0x98, 0xb0, 0x8f, 0xb0, 0x90, 0x8b,
	0x38, 0xb7, 0x88, 0x95, 0x42, 0xc5, 0xb6, 0x1a, 0xb1, 0x1d, 0xa7, 0x1c, 0x07, 0x5c, 0x8c, 0xec,
	0x17, 0x8d, 0xb9, 0xe1, 0x01, 0xf6, 0x2a, 0x52, 0x19, 0x7b, 0x8e, 0x17, 0x36, 0x2c, 0x78, 0x59,
	0x61, 0x56, 0x5e, 0x29, 0x4c, 0xd8, 0x2e, 0xba, 0xec, 0xd1, 0x16, 0x95, 0x30, 0xb1, 0xcd, 0x43,
	0x1e, 0xf6, 0x9c, 0x98, 0x2d, 0x74, 0xee, 0x48, 0x3a, 0x20, 0x1b, 0x78, 0x61, 0xc6, 0xd4, 0x3e,
	0xde, 0xc0, 0xa0, 0x8b, 0xaa, 0xff, 0x7d, 0xf2, 0xa0, 0x9a, 0x55, 0x43, 0x75, 0x25, 0x12, 0x39,
	0x9c, 0xb0, 0x17, 0x48, 0xc2, 0x87, 0x70, 0xef, 0x8f, 0x1a, 0xd9, 0xbf, 0x02, 0x19, 0xaf, 0x78,
	0x86, 0x7c, 0x46, 0x6a, 0x66, 0x82, 0x93, 0x63, 0xfd, 0xf8, 0xf5, 0xfd, 0x34, 0xe6, 0xd7, 0xcc,
	0x04, 0x37, 0x84, 0xd2, 0x22, 0x91, 0xac, 0x8e, 0x13, 0xa4, 0xb2, 0xac, 0xa4, 0xec, 0x27, 0x37,
	0x65, 0x31, 0x9f, 0x2e, 0x4b, 0x88, 0xdd, 0x15, 0x79, 0x19, 0x5e, 0xc3, 0x54, 0xb3, 0x06, 0x3a,
	0x67, 0xa6, 0xa5, 0xd0, 0xe5, 0x08, 0x8c, 0x5d, 0x07, 0xd6, 0xd7, 0x76, 0xc0, 0x70, 0x62, 0x1f,
	0x15, 0x14, 0x85, 0x2a, 0x34, 0x6b, 0xfd, 0x97, 0x47, 0x85, 0x31, 0x7e, 0x15, 0xda, 0xfb, 0xc9,
	0x23, 0x07, 0x97, 0xa5, 0x4e, 0x7d, 0xbe, 0xea, 0x49, 0xbd, 0x35, 0xef, 0x72, 0xc7, 0x76, 0xaf,
	0xf7, 0xb3, 0x47, 0x5e, 0xbe, 0x83, 0x48, 0xc5, 0xf0, 0xff, 0x7f, 0x97, 0xdf, 0x3d, 0xb2, 0xfd,
	0x8e, 0x1b, 0xbe, 0x8a, 0xda, 0x94, 0xac, 0xd9, 0x47, 0x57, 0x55, 0xc7, 0xb3, 0x5b, 0x38, 0xd7,
	0x20, 0xab, 0x1d, 0xe5, 0x0c, 0xbb, 0x09, 0x40, 0x46, 0x2a, 0x16, 0x32, 0xc1, 0x3f, 0x15, 0x1d,
	0x7f, 0x6e, 0xcf, 0x97, 0x5a, 0x63, 0xb1, 0xd4, 0x7a, 0xbf, 0x78, 0xe4, 0xd5, 0x52, 0x23, 0x2f,
	0x0b, 0x95, 0xf3, 0x84, 0xaf, 0xaa, 0xa7, 0xff, 0xf0, 0x7f, 0xe9, 0xed, 0x97, 0xdf, 0x7f, 0x91,
	0x08, 0x93, 0x96, 0xe1, 0x20, 0x52, 0xe3, 0xa3, 0xaa, 0xcc, 0xa1, 0xad, 0x73, 0xb4, 0xa8, 0x73,
	0x38, 0x2b, 0x74, 0xe4, 0x06, 0xd4, 0x61, 0x02, 0xf2, 0x30, 0x51, 0x47, 0xa1, 0x89, 0xc2, 0x26,
	0x42, 0x9f, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0x30, 0x4b, 0x9d, 0x6f, 0x5b, 0x0a, 0x00, 0x00,
}
