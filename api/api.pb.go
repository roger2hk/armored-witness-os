//
// Copyright 2022 The Armored Witness OS authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorCode int32

const (
	ErrorCode_NONE ErrorCode = 0
	// GENERIC_ERROR is returned in case of a generic error, in this case
	// Payload might contain the error string.
	ErrorCode_GENERIC_ERROR ErrorCode = 1
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "NONE",
		1: "GENERIC_ERROR",
	}
	ErrorCode_value = map[string]int32{
		"NONE":          0,
		"GENERIC_ERROR": 1,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

//
//
//AppletUpdate
//
//A `AppletUpdate` represents an OTA sequence applet slice.
//
//The `TotalChunks` value indicates the total number of chunks for the update,
//`Seq` is the transmitted AppletUpdate chunk number:
//- `0` indicates that the struct contains verification data in `Header`.
//- `1` onwards identifies the first, second, ... chunk of firmware image data.
//
type AppletUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Seq   uint32 `protobuf:"varint,2,opt,name=Seq,proto3" json:"Seq,omitempty"`
	// Types that are assignable to Payload:
	//	*AppletUpdate_Data
	//	*AppletUpdate_Header
	Payload isAppletUpdate_Payload `protobuf_oneof:"Payload"`
}

func (x *AppletUpdate) Reset() {
	*x = AppletUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppletUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppletUpdate) ProtoMessage() {}

func (x *AppletUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppletUpdate.ProtoReflect.Descriptor instead.
func (*AppletUpdate) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *AppletUpdate) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AppletUpdate) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (m *AppletUpdate) GetPayload() isAppletUpdate_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *AppletUpdate) GetData() []byte {
	if x, ok := x.GetPayload().(*AppletUpdate_Data); ok {
		return x.Data
	}
	return nil
}

func (x *AppletUpdate) GetHeader() *AppletUpdateHeader {
	if x, ok := x.GetPayload().(*AppletUpdate_Header); ok {
		return x.Header
	}
	return nil
}

type isAppletUpdate_Payload interface {
	isAppletUpdate_Payload()
}

type AppletUpdate_Data struct {
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3,oneof"`
}

type AppletUpdate_Header struct {
	Header *AppletUpdateHeader `protobuf:"bytes,4,opt,name=Header,proto3,oneof"`
}

func (*AppletUpdate_Data) isAppletUpdate_Payload() {}

func (*AppletUpdate_Header) isAppletUpdate_Payload() {}

type AppletUpdateHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signature holds the signature over the applet.
	Signature []byte `protobuf:"bytes,1,opt,name=Signature,proto3" json:"Signature,omitempty"`
	// Checkpoint contains a note-formatted Log Checkpoint.
	Checkpoint []byte `protobuf:"bytes,2,opt,name=Checkpoint,proto3" json:"Checkpoint,omitempty"`
	// Manifest is metadata about the applet, which has been logged to a firmware
	// transparency log.
	Manifest []byte `protobuf:"bytes,3,opt,name=Manifest,proto3" json:"Manifest,omitempty"`
	// InclusionProof is a log inclusion proof for Manifest committed to by
	// Checkpoint.
	InclusionProof [][]byte `protobuf:"bytes,4,rep,name=InclusionProof,proto3" json:"InclusionProof,omitempty"`
	// LogIndex is the index of Manifest in the firmware transparency log.
	LogIndex uint64 `protobuf:"varint,5,opt,name=LogIndex,proto3" json:"LogIndex,omitempty"`
}

func (x *AppletUpdateHeader) Reset() {
	*x = AppletUpdateHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppletUpdateHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppletUpdateHeader) ProtoMessage() {}

func (x *AppletUpdateHeader) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppletUpdateHeader.ProtoReflect.Descriptor instead.
func (*AppletUpdateHeader) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *AppletUpdateHeader) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *AppletUpdateHeader) GetCheckpoint() []byte {
	if x != nil {
		return x.Checkpoint
	}
	return nil
}

func (x *AppletUpdateHeader) GetManifest() []byte {
	if x != nil {
		return x.Manifest
	}
	return nil
}

func (x *AppletUpdateHeader) GetInclusionProof() [][]byte {
	if x != nil {
		return x.InclusionProof
	}
	return nil
}

func (x *AppletUpdateHeader) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

//
//
//Status information
//
//The status information format is returned on any message sent with the
//`U2FHID_ARMORY_INF` vendor specific command.
//
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serial   string         `protobuf:"bytes,1,opt,name=Serial,proto3" json:"Serial,omitempty"`
	HAB      bool           `protobuf:"varint,2,opt,name=HAB,proto3" json:"HAB,omitempty"`
	Revision string         `protobuf:"bytes,3,opt,name=Revision,proto3" json:"Revision,omitempty"`
	Build    string         `protobuf:"bytes,4,opt,name=Build,proto3" json:"Build,omitempty"`
	Version  string         `protobuf:"bytes,5,opt,name=Version,proto3" json:"Version,omitempty"`
	Runtime  string         `protobuf:"bytes,6,opt,name=Runtime,proto3" json:"Runtime,omitempty"`
	Link     bool           `protobuf:"varint,7,opt,name=Link,proto3" json:"Link,omitempty"`
	Witness  *WitnessStatus `protobuf:"bytes,8,opt,name=Witness,proto3" json:"Witness,omitempty"`
	// IdentityCounter is incremented when the device is recovered and the device
	// needs a new witness identity.
	IdentityCounter uint32 `protobuf:"varint,9,opt,name=IdentityCounter,proto3" json:"IdentityCounter,omitempty"`
	SRKHash         string `protobuf:"bytes,10,opt,name=SRKHash,proto3" json:"SRKHash,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *Status) GetHAB() bool {
	if x != nil {
		return x.HAB
	}
	return false
}

func (x *Status) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *Status) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

func (x *Status) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Status) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *Status) GetLink() bool {
	if x != nil {
		return x.Link
	}
	return false
}

func (x *Status) GetWitness() *WitnessStatus {
	if x != nil {
		return x.Witness
	}
	return nil
}

func (x *Status) GetIdentityCounter() uint32 {
	if x != nil {
		return x.IdentityCounter
	}
	return 0
}

func (x *Status) GetSRKHash() string {
	if x != nil {
		return x.SRKHash
	}
	return ""
}

//
//
//WitnessStatus contains witness-applet specific status information.
//
//This is embedded in the general Status message if the applet has provided
//this information to the OS.
//
type WitnessStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identity is the note-formatted public key which can be used to verify
	// checkpoints cosigned by this witness.
	Identity string `protobuf:"bytes,1,opt,name=Identity,proto3" json:"Identity,omitempty"`
	// IP is a string representation of the witness applet's current IP address.
	IP string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	// IDAttestKey is the stable public key from this device, used to attest to all derived witness identities.
	IDAttestPublicKey string `protobuf:"bytes,3,opt,name=IDAttestPublicKey,proto3" json:"IDAttestPublicKey,omitempty"`
	// AttestedID is a note-formatted signed attestation for the current witness identity.
	// This attestation note contains:
	//   "ArmoredWitness ID attestation v1"
	//   <Device serial>
	//   <Witness identity counter in decimal>
	//   <Witness identity as a note verifier string>
	AttestedID string `protobuf:"bytes,4,opt,name=AttestedID,proto3" json:"AttestedID,omitempty"`
}

func (x *WitnessStatus) Reset() {
	*x = WitnessStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WitnessStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WitnessStatus) ProtoMessage() {}

func (x *WitnessStatus) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WitnessStatus.ProtoReflect.Descriptor instead.
func (*WitnessStatus) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *WitnessStatus) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *WitnessStatus) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *WitnessStatus) GetIDAttestPublicKey() string {
	if x != nil {
		return x.IDAttestPublicKey
	}
	return ""
}

func (x *WitnessStatus) GetAttestedID() string {
	if x != nil {
		return x.AttestedID
	}
	return ""
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DHCP      bool   `protobuf:"varint,1,opt,name=DHCP,proto3" json:"DHCP,omitempty"`
	IP        string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Netmask   string `protobuf:"bytes,3,opt,name=Netmask,proto3" json:"Netmask,omitempty"`
	Gateway   string `protobuf:"bytes,4,opt,name=Gateway,proto3" json:"Gateway,omitempty"`
	Resolver  string `protobuf:"bytes,5,opt,name=Resolver,proto3" json:"Resolver,omitempty"`
	NTPServer string `protobuf:"bytes,6,opt,name=NTPServer,proto3" json:"NTPServer,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Configuration) GetDHCP() bool {
	if x != nil {
		return x.DHCP
	}
	return false
}

func (x *Configuration) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *Configuration) GetNetmask() string {
	if x != nil {
		return x.Netmask
	}
	return ""
}

func (x *Configuration) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

func (x *Configuration) GetResolver() string {
	if x != nil {
		return x.Resolver
	}
	return ""
}

func (x *Configuration) GetNTPServer() string {
	if x != nil {
		return x.NTPServer
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error   ErrorCode `protobuf:"varint,1,opt,name=Error,proto3,enum=api.ErrorCode" json:"Error,omitempty"`
	Payload []byte    `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *Response) GetError() ErrorCode {
	if x != nil {
		return x.Error
	}
	return ErrorCode_NONE
}

func (x *Response) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x71, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x53, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x31, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xb2, 0x01,
	0x0a, 0x12, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x22, 0x9e, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x48, 0x41, 0x42, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x48, 0x41, 0x42, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x2c, 0x0a, 0x07, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x12,
	0x28, 0x0a, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x52, 0x4b,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x52, 0x4b, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x57, 0x69, 0x74, 0x6e, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x50, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x44, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x49, 0x44,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x65, 0x64, 0x49, 0x44, 0x22,
	0xa1, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x48, 0x43, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x44, 0x48, 0x43, 0x50, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x12,
	0x18, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x54, 0x50, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x54, 0x50, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2a,
	0x28, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x49,
	0x43, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []interface{}{
	(ErrorCode)(0),             // 0: api.ErrorCode
	(*AppletUpdate)(nil),       // 1: api.AppletUpdate
	(*AppletUpdateHeader)(nil), // 2: api.AppletUpdateHeader
	(*Status)(nil),             // 3: api.Status
	(*WitnessStatus)(nil),      // 4: api.WitnessStatus
	(*Configuration)(nil),      // 5: api.Configuration
	(*Response)(nil),           // 6: api.Response
}
var file_api_proto_depIdxs = []int32{
	2, // 0: api.AppletUpdate.Header:type_name -> api.AppletUpdateHeader
	4, // 1: api.Status.Witness:type_name -> api.WitnessStatus
	0, // 2: api.Response.Error:type_name -> api.ErrorCode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppletUpdate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppletUpdateHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WitnessStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configuration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AppletUpdate_Data)(nil),
		(*AppletUpdate_Header)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
