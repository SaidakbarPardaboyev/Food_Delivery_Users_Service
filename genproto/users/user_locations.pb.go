// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: user_locations.proto

package users

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

type CreateUserLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Address         string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	HomeNumber      string  `protobuf:"bytes,4,opt,name=home_number,json=homeNumber,proto3" json:"home_number,omitempty"`
	FloorNumber     int32   `protobuf:"varint,5,opt,name=floor_number,json=floorNumber,proto3" json:"floor_number,omitempty"`
	ApartmentNumber int32   `protobuf:"varint,6,opt,name=apartment_number,json=apartmentNumber,proto3" json:"apartment_number,omitempty"`
	EntranceNumber  int32   `protobuf:"varint,7,opt,name=entrance_number,json=entranceNumber,proto3" json:"entrance_number,omitempty"`
	Latitude        float64 `protobuf:"fixed64,8,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude       float64 `protobuf:"fixed64,9,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *CreateUserLocation) Reset() {
	*x = CreateUserLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_locations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserLocation) ProtoMessage() {}

func (x *CreateUserLocation) ProtoReflect() protoreflect.Message {
	mi := &file_user_locations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserLocation.ProtoReflect.Descriptor instead.
func (*CreateUserLocation) Descriptor() ([]byte, []int) {
	return file_user_locations_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserLocation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateUserLocation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateUserLocation) GetHomeNumber() string {
	if x != nil {
		return x.HomeNumber
	}
	return ""
}

func (x *CreateUserLocation) GetFloorNumber() int32 {
	if x != nil {
		return x.FloorNumber
	}
	return 0
}

func (x *CreateUserLocation) GetApartmentNumber() int32 {
	if x != nil {
		return x.ApartmentNumber
	}
	return 0
}

func (x *CreateUserLocation) GetEntranceNumber() int32 {
	if x != nil {
		return x.EntranceNumber
	}
	return 0
}

func (x *CreateUserLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CreateUserLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type UserLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId          string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Address         string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	HomeNumber      string  `protobuf:"bytes,4,opt,name=home_number,json=homeNumber,proto3" json:"home_number,omitempty"`
	FloorNumber     int32   `protobuf:"varint,5,opt,name=floor_number,json=floorNumber,proto3" json:"floor_number,omitempty"`
	ApartmentNumber int32   `protobuf:"varint,6,opt,name=apartment_number,json=apartmentNumber,proto3" json:"apartment_number,omitempty"`
	EntranceNumber  int32   `protobuf:"varint,7,opt,name=entrance_number,json=entranceNumber,proto3" json:"entrance_number,omitempty"`
	Latitude        float64 `protobuf:"fixed64,8,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude       float64 `protobuf:"fixed64,9,opt,name=longitude,proto3" json:"longitude,omitempty"`
	CreatedAt       string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string  `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserLocation) Reset() {
	*x = UserLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_locations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLocation) ProtoMessage() {}

func (x *UserLocation) ProtoReflect() protoreflect.Message {
	mi := &file_user_locations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLocation.ProtoReflect.Descriptor instead.
func (*UserLocation) Descriptor() ([]byte, []int) {
	return file_user_locations_proto_rawDescGZIP(), []int{1}
}

func (x *UserLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserLocation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLocation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserLocation) GetHomeNumber() string {
	if x != nil {
		return x.HomeNumber
	}
	return ""
}

func (x *UserLocation) GetFloorNumber() int32 {
	if x != nil {
		return x.FloorNumber
	}
	return 0
}

func (x *UserLocation) GetApartmentNumber() int32 {
	if x != nil {
		return x.ApartmentNumber
	}
	return 0
}

func (x *UserLocation) GetEntranceNumber() int32 {
	if x != nil {
		return x.EntranceNumber
	}
	return 0
}

func (x *UserLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UserLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *UserLocation) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserLocation) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UserLocationFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address         string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	HomeNumber      string `protobuf:"bytes,2,opt,name=home_number,json=homeNumber,proto3" json:"home_number,omitempty"`
	FloorNumber     int32  `protobuf:"varint,3,opt,name=floor_number,json=floorNumber,proto3" json:"floor_number,omitempty"`
	ApartmentNumber int32  `protobuf:"varint,4,opt,name=apartment_number,json=apartmentNumber,proto3" json:"apartment_number,omitempty"`
	EntranceNumber  int32  `protobuf:"varint,5,opt,name=entrance_number,json=entranceNumber,proto3" json:"entrance_number,omitempty"`
}

func (x *UserLocationFilter) Reset() {
	*x = UserLocationFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_locations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLocationFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLocationFilter) ProtoMessage() {}

func (x *UserLocationFilter) ProtoReflect() protoreflect.Message {
	mi := &file_user_locations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLocationFilter.ProtoReflect.Descriptor instead.
func (*UserLocationFilter) Descriptor() ([]byte, []int) {
	return file_user_locations_proto_rawDescGZIP(), []int{2}
}

func (x *UserLocationFilter) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserLocationFilter) GetHomeNumber() string {
	if x != nil {
		return x.HomeNumber
	}
	return ""
}

func (x *UserLocationFilter) GetFloorNumber() int32 {
	if x != nil {
		return x.FloorNumber
	}
	return 0
}

func (x *UserLocationFilter) GetApartmentNumber() int32 {
	if x != nil {
		return x.ApartmentNumber
	}
	return 0
}

func (x *UserLocationFilter) GetEntranceNumber() int32 {
	if x != nil {
		return x.EntranceNumber
	}
	return 0
}

type UserLocations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLocations []*UserLocation `protobuf:"bytes,1,rep,name=user_locations,json=userLocations,proto3" json:"user_locations,omitempty"`
}

func (x *UserLocations) Reset() {
	*x = UserLocations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_locations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLocations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLocations) ProtoMessage() {}

func (x *UserLocations) ProtoReflect() protoreflect.Message {
	mi := &file_user_locations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLocations.ProtoReflect.Descriptor instead.
func (*UserLocations) Descriptor() ([]byte, []int) {
	return file_user_locations_proto_rawDescGZIP(), []int{3}
}

func (x *UserLocations) GetUserLocations() []*UserLocation {
	if x != nil {
		return x.UserLocations
	}
	return nil
}

type UpdateUserLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address         string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	HomeNumber      string  `protobuf:"bytes,3,opt,name=home_number,json=homeNumber,proto3" json:"home_number,omitempty"`
	FloorNumber     int32   `protobuf:"varint,4,opt,name=floor_number,json=floorNumber,proto3" json:"floor_number,omitempty"`
	ApartmentNumber int32   `protobuf:"varint,5,opt,name=apartment_number,json=apartmentNumber,proto3" json:"apartment_number,omitempty"`
	EntranceNumber  int32   `protobuf:"varint,6,opt,name=entrance_number,json=entranceNumber,proto3" json:"entrance_number,omitempty"`
	Latitude        float64 `protobuf:"fixed64,7,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude       float64 `protobuf:"fixed64,8,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *UpdateUserLocation) Reset() {
	*x = UpdateUserLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_locations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserLocation) ProtoMessage() {}

func (x *UpdateUserLocation) ProtoReflect() protoreflect.Message {
	mi := &file_user_locations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserLocation.ProtoReflect.Descriptor instead.
func (*UpdateUserLocation) Descriptor() ([]byte, []int) {
	return file_user_locations_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserLocation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateUserLocation) GetHomeNumber() string {
	if x != nil {
		return x.HomeNumber
	}
	return ""
}

func (x *UpdateUserLocation) GetFloorNumber() int32 {
	if x != nil {
		return x.FloorNumber
	}
	return 0
}

func (x *UpdateUserLocation) GetApartmentNumber() int32 {
	if x != nil {
		return x.ApartmentNumber
	}
	return 0
}

func (x *UpdateUserLocation) GetEntranceNumber() int32 {
	if x != nil {
		return x.EntranceNumber
	}
	return 0
}

func (x *UpdateUserLocation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *UpdateUserLocation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_user_locations_proto protoreflect.FileDescriptor

var file_user_locations_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x6c, 0x6f,
	0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xe1, 0x02, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6f,
	0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66,
	0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x10, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0e, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x12, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68,
	0x6f, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x10, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x61, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x90, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68, 0x6f, 0x6d, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x61, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x32, 0x86, 0x03, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x4b, 0x65, 0x79, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x39, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2c,
	0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a,
	0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x10, 0x5a, 0x0e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_locations_proto_rawDescOnce sync.Once
	file_user_locations_proto_rawDescData = file_user_locations_proto_rawDesc
)

func file_user_locations_proto_rawDescGZIP() []byte {
	file_user_locations_proto_rawDescOnce.Do(func() {
		file_user_locations_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_locations_proto_rawDescData)
	})
	return file_user_locations_proto_rawDescData
}

var file_user_locations_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_locations_proto_goTypes = []interface{}{
	(*CreateUserLocation)(nil), // 0: users.CreateUserLocation
	(*UserLocation)(nil),       // 1: users.UserLocation
	(*UserLocationFilter)(nil), // 2: users.UserLocationFilter
	(*UserLocations)(nil),      // 3: users.UserLocations
	(*UpdateUserLocation)(nil), // 4: users.UpdateUserLocation
	(*PrimaryKey)(nil),         // 5: users.PrimaryKey
	(*Void)(nil),               // 6: users.Void
}
var file_user_locations_proto_depIdxs = []int32{
	1, // 0: users.UserLocations.user_locations:type_name -> users.UserLocation
	0, // 1: users.UserLocationService.Create:input_type -> users.CreateUserLocation
	5, // 2: users.UserLocationService.GetById:input_type -> users.PrimaryKey
	5, // 3: users.UserLocationService.GetByUserId:input_type -> users.PrimaryKey
	2, // 4: users.UserLocationService.GetAll:input_type -> users.UserLocationFilter
	4, // 5: users.UserLocationService.Update:input_type -> users.UpdateUserLocation
	5, // 6: users.UserLocationService.Delete:input_type -> users.PrimaryKey
	5, // 7: users.UserLocationService.ValidateId:input_type -> users.PrimaryKey
	1, // 8: users.UserLocationService.Create:output_type -> users.UserLocation
	1, // 9: users.UserLocationService.GetById:output_type -> users.UserLocation
	1, // 10: users.UserLocationService.GetByUserId:output_type -> users.UserLocation
	3, // 11: users.UserLocationService.GetAll:output_type -> users.UserLocations
	1, // 12: users.UserLocationService.Update:output_type -> users.UserLocation
	6, // 13: users.UserLocationService.Delete:output_type -> users.Void
	6, // 14: users.UserLocationService.ValidateId:output_type -> users.Void
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_locations_proto_init() }
func file_user_locations_proto_init() {
	if File_user_locations_proto != nil {
		return
	}
	file_users_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_locations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserLocation); i {
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
		file_user_locations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLocation); i {
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
		file_user_locations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLocationFilter); i {
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
		file_user_locations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLocations); i {
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
		file_user_locations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserLocation); i {
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
			RawDescriptor: file_user_locations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_locations_proto_goTypes,
		DependencyIndexes: file_user_locations_proto_depIdxs,
		MessageInfos:      file_user_locations_proto_msgTypes,
	}.Build()
	File_user_locations_proto = out.File
	file_user_locations_proto_rawDesc = nil
	file_user_locations_proto_goTypes = nil
	file_user_locations_proto_depIdxs = nil
}
