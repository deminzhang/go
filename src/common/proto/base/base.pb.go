// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: base.proto

package base

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{0}
}

type PlayerId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId int64 `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"` // 用户id
}

func (x *PlayerId) Reset() {
	*x = PlayerId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerId) ProtoMessage() {}

func (x *PlayerId) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerId.ProtoReflect.Descriptor instead.
func (*PlayerId) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerId) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

// 通用Val
type ValInt32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V int32 `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *ValInt32) Reset() {
	*x = ValInt32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValInt32) ProtoMessage() {}

func (x *ValInt32) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValInt32.ProtoReflect.Descriptor instead.
func (*ValInt32) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{2}
}

func (x *ValInt32) GetV() int32 {
	if x != nil {
		return x.V
	}
	return 0
}

type ValInt64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V int64 `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *ValInt64) Reset() {
	*x = ValInt64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValInt64) ProtoMessage() {}

func (x *ValInt64) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValInt64.ProtoReflect.Descriptor instead.
func (*ValInt64) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{3}
}

func (x *ValInt64) GetV() int64 {
	if x != nil {
		return x.V
	}
	return 0
}

type ValBool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V bool `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
}

func (x *ValBool) Reset() {
	*x = ValBool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValBool) ProtoMessage() {}

func (x *ValBool) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValBool.ProtoReflect.Descriptor instead.
func (*ValBool) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{4}
}

func (x *ValBool) GetV() bool {
	if x != nil {
		return x.V
	}
	return false
}

// 通用Id
type IdInt32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdInt32) Reset() {
	*x = IdInt32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdInt32) ProtoMessage() {}

func (x *IdInt32) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdInt32.ProtoReflect.Descriptor instead.
func (*IdInt32) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{5}
}

func (x *IdInt32) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdsInt32 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsInt32) Reset() {
	*x = IdsInt32{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsInt32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsInt32) ProtoMessage() {}

func (x *IdsInt32) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsInt32.ProtoReflect.Descriptor instead.
func (*IdsInt32) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{6}
}

func (x *IdsInt32) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type IdInt64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdInt64) Reset() {
	*x = IdInt64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdInt64) ProtoMessage() {}

func (x *IdInt64) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdInt64.ProtoReflect.Descriptor instead.
func (*IdInt64) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{7}
}

func (x *IdInt64) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdsInt64 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsInt64) Reset() {
	*x = IdsInt64{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsInt64) ProtoMessage() {}

func (x *IdsInt64) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsInt64.ProtoReflect.Descriptor instead.
func (*IdsInt64) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{8}
}

func (x *IdsInt64) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 通用ModifyNumById
type ModifyNum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`         // 库id
	Modify int32 `protobuf:"varint,2,opt,name=modify,proto3" json:"modify,omitempty"` // 正加负减
}

func (x *ModifyNum) Reset() {
	*x = ModifyNum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyNum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyNum) ProtoMessage() {}

func (x *ModifyNum) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyNum.ProtoReflect.Descriptor instead.
func (*ModifyNum) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{9}
}

func (x *ModifyNum) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ModifyNum) GetModify() int32 {
	if x != nil {
		return x.Modify
	}
	return 0
}

type Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameType int32  `protobuf:"varint,1,opt,name=frame_type,json=frameType,proto3" json:"frame_type,omitempty"`
	Data      []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Frame) Reset() {
	*x = Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frame) ProtoMessage() {}

func (x *Frame) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frame.ProtoReflect.Descriptor instead.
func (*Frame) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{10}
}

func (x *Frame) GetFrameType() int32 {
	if x != nil {
		return x.FrameType
	}
	return 0
}

func (x *Frame) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Host         string  `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"` // 域名
	ClientIp     string  `protobuf:"bytes,3,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	BusinessType string  `protobuf:"bytes,4,opt,name=business_type,json=businessType,proto3" json:"business_type,omitempty"` // ugc/homestead
	ClientType   string  `protobuf:"bytes,5,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`       // app/web
	Platform     string  `protobuf:"bytes,6,opt,name=platform,proto3" json:"platform,omitempty"`
	DistinctId   string  `protobuf:"bytes,7,opt,name=distinct_id,json=distinctId,proto3" json:"distinct_id,omitempty"`
	VprsCurrent  string  `protobuf:"bytes,8,opt,name=vprs_current,json=vprsCurrent,proto3" json:"vprs_current,omitempty"`
	Country      string  `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`      // 数数.国家
	Timezone     float64 `protobuf:"fixed64,10,opt,name=timezone,proto3" json:"timezone,omitempty"` // 数数.前端时区
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{11}
}

func (x *ClientInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ClientInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ClientInfo) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *ClientInfo) GetBusinessType() string {
	if x != nil {
		return x.BusinessType
	}
	return ""
}

func (x *ClientInfo) GetClientType() string {
	if x != nil {
		return x.ClientType
	}
	return ""
}

func (x *ClientInfo) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *ClientInfo) GetDistinctId() string {
	if x != nil {
		return x.DistinctId
	}
	return ""
}

func (x *ClientInfo) GetVprsCurrent() string {
	if x != nil {
		return x.VprsCurrent
	}
	return ""
}

func (x *ClientInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *ClientInfo) GetTimezone() float64 {
	if x != nil {
		return x.Timezone
	}
	return 0
}

type ClientInfoEx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *ClientInfo `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	PlayerId int64       `protobuf:"varint,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	HostId   int32       `protobuf:"varint,3,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
}

func (x *ClientInfoEx) Reset() {
	*x = ClientInfoEx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfoEx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfoEx) ProtoMessage() {}

func (x *ClientInfoEx) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfoEx.ProtoReflect.Descriptor instead.
func (*ClientInfoEx) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{12}
}

func (x *ClientInfoEx) GetBase() *ClientInfo {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *ClientInfoEx) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *ClientInfoEx) GetHostId() int32 {
	if x != nil {
		return x.HostId
	}
	return 0
}

type PlayerIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *PlayerIds) Reset() {
	*x = PlayerIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerIds) ProtoMessage() {}

func (x *PlayerIds) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerIds.ProtoReflect.Descriptor instead.
func (*PlayerIds) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{13}
}

func (x *PlayerIds) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Id32ToId32Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map map[int32]int32 `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *Id32ToId32Map) Reset() {
	*x = Id32ToId32Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id32ToId32Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id32ToId32Map) ProtoMessage() {}

func (x *Id32ToId32Map) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id32ToId32Map.ProtoReflect.Descriptor instead.
func (*Id32ToId32Map) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{14}
}

func (x *Id32ToId32Map) GetMap() map[int32]int32 {
	if x != nil {
		return x.Map
	}
	return nil
}

type Id32ToStringMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map map[int32]string `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Id32ToStringMap) Reset() {
	*x = Id32ToStringMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id32ToStringMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id32ToStringMap) ProtoMessage() {}

func (x *Id32ToStringMap) ProtoReflect() protoreflect.Message {
	mi := &file_base_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id32ToStringMap.ProtoReflect.Descriptor instead.
func (*Id32ToStringMap) Descriptor() ([]byte, []int) {
	return file_base_proto_rawDescGZIP(), []int{15}
}

func (x *Id32ToStringMap) GetMap() map[int32]string {
	if x != nil {
		return x.Map
	}
	return nil
}

var File_base_proto protoreflect.FileDescriptor

var file_base_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x76, 0x22, 0x18,
	0x0a, 0x08, 0x56, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x76, 0x22, 0x17, 0x0a, 0x07, 0x56, 0x61, 0x6c, 0x42,
	0x6f, 0x6f, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x01,
	0x76, 0x22, 0x19, 0x0a, 0x07, 0x49, 0x64, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x08,
	0x49, 0x64, 0x73, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x49, 0x64,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x08, 0x49, 0x64, 0x73, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0x33, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4e, 0x75, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x22, 0x3a, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xb2, 0x02, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x23, 0x0a,
	0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x76, 0x70, 0x72, 0x73, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x70, 0x72, 0x73, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0x6a, 0x0a, 0x0c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x78, 0x12, 0x24, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x03, 0x69, 0x64, 0x73, 0x22, 0x77, 0x0a, 0x0d, 0x49, 0x64, 0x33, 0x32, 0x54, 0x6f, 0x49, 0x64,
	0x33, 0x32, 0x4d, 0x61, 0x70, 0x12, 0x2e, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x49, 0x64, 0x33, 0x32, 0x54, 0x6f,
	0x49, 0x64, 0x33, 0x32, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7b, 0x0a,
	0x0f, 0x49, 0x64, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x70,
	0x12, 0x30, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x49, 0x64, 0x33, 0x32, 0x54, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d,
	0x61, 0x70, 0x1a, 0x36, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_proto_rawDescOnce sync.Once
	file_base_proto_rawDescData = file_base_proto_rawDesc
)

func file_base_proto_rawDescGZIP() []byte {
	file_base_proto_rawDescOnce.Do(func() {
		file_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_proto_rawDescData)
	})
	return file_base_proto_rawDescData
}

var file_base_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_base_proto_goTypes = []interface{}{
	(*Empty)(nil),           // 0: base.Empty
	(*PlayerId)(nil),        // 1: base.PlayerId
	(*ValInt32)(nil),        // 2: base.ValInt32
	(*ValInt64)(nil),        // 3: base.ValInt64
	(*ValBool)(nil),         // 4: base.ValBool
	(*IdInt32)(nil),         // 5: base.IdInt32
	(*IdsInt32)(nil),        // 6: base.IdsInt32
	(*IdInt64)(nil),         // 7: base.IdInt64
	(*IdsInt64)(nil),        // 8: base.IdsInt64
	(*ModifyNum)(nil),       // 9: base.ModifyNum
	(*Frame)(nil),           // 10: base.Frame
	(*ClientInfo)(nil),      // 11: base.ClientInfo
	(*ClientInfoEx)(nil),    // 12: base.ClientInfoEx
	(*PlayerIds)(nil),       // 13: base.PlayerIds
	(*Id32ToId32Map)(nil),   // 14: base.Id32ToId32Map
	(*Id32ToStringMap)(nil), // 15: base.Id32ToStringMap
	nil,                     // 16: base.Id32ToId32Map.MapEntry
	nil,                     // 17: base.Id32ToStringMap.MapEntry
}
var file_base_proto_depIdxs = []int32{
	11, // 0: base.ClientInfoEx.base:type_name -> base.ClientInfo
	16, // 1: base.Id32ToId32Map.map:type_name -> base.Id32ToId32Map.MapEntry
	17, // 2: base.Id32ToStringMap.map:type_name -> base.Id32ToStringMap.MapEntry
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_base_proto_init() }
func file_base_proto_init() {
	if File_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerId); i {
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
		file_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValInt32); i {
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
		file_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValInt64); i {
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
		file_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValBool); i {
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
		file_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdInt32); i {
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
		file_base_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdsInt32); i {
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
		file_base_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdInt64); i {
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
		file_base_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdsInt64); i {
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
		file_base_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyNum); i {
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
		file_base_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Frame); i {
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
		file_base_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
		file_base_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfoEx); i {
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
		file_base_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerIds); i {
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
		file_base_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id32ToId32Map); i {
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
		file_base_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id32ToStringMap); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_proto_goTypes,
		DependencyIndexes: file_base_proto_depIdxs,
		MessageInfos:      file_base_proto_msgTypes,
	}.Build()
	File_base_proto = out.File
	file_base_proto_rawDesc = nil
	file_base_proto_goTypes = nil
	file_base_proto_depIdxs = nil
}
