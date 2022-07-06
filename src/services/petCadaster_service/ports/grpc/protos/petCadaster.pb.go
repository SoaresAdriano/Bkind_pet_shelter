// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: src/services/petCadaster_service/ports/grpc/protos/petCadaster.proto

package grpc

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
		mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[0]
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
	return file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescGZIP(), []int{0}
}

type PetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *ID     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Species      string  `protobuf:"bytes,2,opt,name=species,proto3" json:"species,omitempty"`
	Breed        string  `protobuf:"bytes,3,opt,name=breed,proto3" json:"breed,omitempty"`
	Sex          string  `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Age          int32   `protobuf:"varint,5,opt,name=age,proto3" json:"age,omitempty"`
	Weight       float32 `protobuf:"fixed32,6,opt,name=weight,proto3" json:"weight,omitempty"`
	Size         string  `protobuf:"bytes,7,opt,name=size,proto3" json:"size,omitempty"`
	Color        string  `protobuf:"bytes,8,opt,name=color,proto3" json:"color,omitempty"`
	Neutered     string  `protobuf:"bytes,9,opt,name=neutered,proto3" json:"neutered,omitempty"`
	Origin       string  `protobuf:"bytes,10,opt,name=origin,proto3" json:"origin,omitempty"`
	Situation    string  `protobuf:"bytes,11,opt,name=situation,proto3" json:"situation,omitempty"`
	Registeredat string  `protobuf:"bytes,12,opt,name=registeredat,proto3" json:"registeredat,omitempty"`
}

func (x *PetInfo) Reset() {
	*x = PetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetInfo) ProtoMessage() {}

func (x *PetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetInfo.ProtoReflect.Descriptor instead.
func (*PetInfo) Descriptor() ([]byte, []int) {
	return file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescGZIP(), []int{1}
}

func (x *PetInfo) GetId() *ID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *PetInfo) GetSpecies() string {
	if x != nil {
		return x.Species
	}
	return ""
}

func (x *PetInfo) GetBreed() string {
	if x != nil {
		return x.Breed
	}
	return ""
}

func (x *PetInfo) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *PetInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *PetInfo) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *PetInfo) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *PetInfo) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *PetInfo) GetNeutered() string {
	if x != nil {
		return x.Neutered
	}
	return ""
}

func (x *PetInfo) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *PetInfo) GetSituation() string {
	if x != nil {
		return x.Situation
	}
	return ""
}

func (x *PetInfo) GetRegisteredat() string {
	if x != nil {
		return x.Registeredat
	}
	return ""
}

type RenameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *ID    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RenameInfo) Reset() {
	*x = RenameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenameInfo) ProtoMessage() {}

func (x *RenameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenameInfo.ProtoReflect.Descriptor instead.
func (*RenameInfo) Descriptor() ([]byte, []int) {
	return file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescGZIP(), []int{2}
}

func (x *RenameInfo) GetId() *ID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *RenameInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CountInfo) Reset() {
	*x = CountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountInfo) ProtoMessage() {}

func (x *CountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountInfo.ProtoReflect.Descriptor instead.
func (*CountInfo) Descriptor() ([]byte, []int) {
	return file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescGZIP(), []int{3}
}

func (x *CountInfo) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *UUID `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescGZIP(), []int{4}
}

func (x *ID) GetValue() *UUID {
	if x != nil {
		return x.Value
	}
	return nil
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescGZIP(), []int{5}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[6]
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
	return file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescGZIP(), []int{6}
}

func (x *Status) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

var File_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto protoreflect.FileDescriptor

var file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDesc = []byte{
	0x0a, 0x44, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x65, 0x74, 0x43, 0x61, 0x64, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x65, 0x74, 0x43, 0x61, 0x64, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xaf, 0x02, 0x0a, 0x07, 0x50, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x65, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x75, 0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x75, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x74, 0x75, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x74, 0x75, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65,
	0x64, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x61, 0x74, 0x22, 0x3a, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1c,
	0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xc1, 0x02, 0x0a,
	0x0b, 0x70, 0x65, 0x74, 0x43, 0x61, 0x64, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x74, 0x73, 0x12, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01, 0x12, 0x21, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74,
	0x12, 0x08, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x74, 0x43, 0x61, 0x64, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0d,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x08, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x74, 0x43, 0x61, 0x64, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x52, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x50, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x65, 0x74, 0x43, 0x61, 0x64, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x74, 0x73,
	0x12, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01,
	0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescOnce sync.Once
	file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescData = file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDesc
)

func file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescGZIP() []byte {
	file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescOnce.Do(func() {
		file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescData)
	})
	return file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDescData
}

var file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_goTypes = []interface{}{
	(*Empty)(nil),      // 0: grpc.Empty
	(*PetInfo)(nil),    // 1: grpc.PetInfo
	(*RenameInfo)(nil), // 2: grpc.RenameInfo
	(*CountInfo)(nil),  // 3: grpc.CountInfo
	(*ID)(nil),         // 4: grpc.ID
	(*UUID)(nil),       // 5: grpc.UUID
	(*Status)(nil),     // 6: grpc.Status
}
var file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_depIdxs = []int32{
	4,  // 0: grpc.PetInfo.id:type_name -> grpc.ID
	4,  // 1: grpc.RenameInfo.id:type_name -> grpc.ID
	5,  // 2: grpc.ID.value:type_name -> grpc.UUID
	0,  // 3: grpc.petCadaster.ListPets:input_type -> grpc.Empty
	4,  // 4: grpc.petCadaster.GetPet:input_type -> grpc.ID
	1,  // 5: grpc.petCadaster.CreatePetCadaster:input_type -> grpc.PetInfo
	1,  // 6: grpc.petCadaster.UpdatePetCadaster:input_type -> grpc.PetInfo
	2,  // 7: grpc.petCadaster.RenamePet:input_type -> grpc.RenameInfo
	4,  // 8: grpc.petCadaster.DeletePetCadaster:input_type -> grpc.ID
	0,  // 9: grpc.petCadaster.CountPets:input_type -> grpc.Empty
	1,  // 10: grpc.petCadaster.ListPets:output_type -> grpc.PetInfo
	1,  // 11: grpc.petCadaster.GetPet:output_type -> grpc.PetInfo
	4,  // 12: grpc.petCadaster.CreatePetCadaster:output_type -> grpc.ID
	6,  // 13: grpc.petCadaster.UpdatePetCadaster:output_type -> grpc.Status
	6,  // 14: grpc.petCadaster.RenamePet:output_type -> grpc.Status
	6,  // 15: grpc.petCadaster.DeletePetCadaster:output_type -> grpc.Status
	3,  // 16: grpc.petCadaster.CountPets:output_type -> grpc.CountInfo
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_init() }
func file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_init() {
	if File_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PetInfo); i {
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
		file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenameInfo); i {
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
		file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountInfo); i {
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
		file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_goTypes,
		DependencyIndexes: file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_depIdxs,
		MessageInfos:      file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_msgTypes,
	}.Build()
	File_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto = out.File
	file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_rawDesc = nil
	file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_goTypes = nil
	file_src_services_petCadaster_service_ports_grpc_protos_petCadaster_proto_depIdxs = nil
}