// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: unavailable_foods_of_branches.proto

package orders_service

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

type CreateUnavailableFood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	FoodId   string `protobuf:"bytes,2,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
}

func (x *CreateUnavailableFood) Reset() {
	*x = CreateUnavailableFood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unavailable_foods_of_branches_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUnavailableFood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUnavailableFood) ProtoMessage() {}

func (x *CreateUnavailableFood) ProtoReflect() protoreflect.Message {
	mi := &file_unavailable_foods_of_branches_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUnavailableFood.ProtoReflect.Descriptor instead.
func (*CreateUnavailableFood) Descriptor() ([]byte, []int) {
	return file_unavailable_foods_of_branches_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUnavailableFood) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *CreateUnavailableFood) GetFoodId() string {
	if x != nil {
		return x.FoodId
	}
	return ""
}

type UnavailableFood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId    string `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	BranchTitle string `protobuf:"bytes,3,opt,name=branch_title,json=branchTitle,proto3" json:"branch_title,omitempty"`
	FoodId      string `protobuf:"bytes,4,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	FoodName    string `protobuf:"bytes,5,opt,name=food_name,json=foodName,proto3" json:"food_name,omitempty"`
	CreatedAt   string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UnavailableFood) Reset() {
	*x = UnavailableFood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unavailable_foods_of_branches_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnavailableFood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnavailableFood) ProtoMessage() {}

func (x *UnavailableFood) ProtoReflect() protoreflect.Message {
	mi := &file_unavailable_foods_of_branches_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnavailableFood.ProtoReflect.Descriptor instead.
func (*UnavailableFood) Descriptor() ([]byte, []int) {
	return file_unavailable_foods_of_branches_proto_rawDescGZIP(), []int{1}
}

func (x *UnavailableFood) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnavailableFood) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UnavailableFood) GetBranchTitle() string {
	if x != nil {
		return x.BranchTitle
	}
	return ""
}

func (x *UnavailableFood) GetFoodId() string {
	if x != nil {
		return x.FoodId
	}
	return ""
}

func (x *UnavailableFood) GetFoodName() string {
	if x != nil {
		return x.FoodName
	}
	return ""
}

func (x *UnavailableFood) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UnavailableFood) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateUnavailableFood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId  string `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	FoodId    string `protobuf:"bytes,3,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UpdateUnavailableFood) Reset() {
	*x = UpdateUnavailableFood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unavailable_foods_of_branches_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUnavailableFood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUnavailableFood) ProtoMessage() {}

func (x *UpdateUnavailableFood) ProtoReflect() protoreflect.Message {
	mi := &file_unavailable_foods_of_branches_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUnavailableFood.ProtoReflect.Descriptor instead.
func (*UpdateUnavailableFood) Descriptor() ([]byte, []int) {
	return file_unavailable_foods_of_branches_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUnavailableFood) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUnavailableFood) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UpdateUnavailableFood) GetFoodId() string {
	if x != nil {
		return x.FoodId
	}
	return ""
}

func (x *UpdateUnavailableFood) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UpdateUnavailableFood) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UnavailableFoods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnavailableFoods []*UnavailableFood `protobuf:"bytes,1,rep,name=unavailableFoods,proto3" json:"unavailableFoods,omitempty"`
	Count            int64              `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *UnavailableFoods) Reset() {
	*x = UnavailableFoods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unavailable_foods_of_branches_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnavailableFoods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnavailableFoods) ProtoMessage() {}

func (x *UnavailableFoods) ProtoReflect() protoreflect.Message {
	mi := &file_unavailable_foods_of_branches_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnavailableFoods.ProtoReflect.Descriptor instead.
func (*UnavailableFoods) Descriptor() ([]byte, []int) {
	return file_unavailable_foods_of_branches_proto_rawDescGZIP(), []int{3}
}

func (x *UnavailableFoods) GetUnavailableFoods() []*UnavailableFood {
	if x != nil {
		return x.UnavailableFoods
	}
	return nil
}

func (x *UnavailableFoods) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type UnavailableFoodFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	FoodId   string `protobuf:"bytes,2,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	Page     int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit    int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *UnavailableFoodFilter) Reset() {
	*x = UnavailableFoodFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_unavailable_foods_of_branches_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnavailableFoodFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnavailableFoodFilter) ProtoMessage() {}

func (x *UnavailableFoodFilter) ProtoReflect() protoreflect.Message {
	mi := &file_unavailable_foods_of_branches_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnavailableFoodFilter.ProtoReflect.Descriptor instead.
func (*UnavailableFoodFilter) Descriptor() ([]byte, []int) {
	return file_unavailable_foods_of_branches_proto_rawDescGZIP(), []int{4}
}

func (x *UnavailableFoodFilter) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UnavailableFoodFilter) GetFoodId() string {
	if x != nil {
		return x.FoodId
	}
	return ""
}

func (x *UnavailableFoodFilter) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *UnavailableFoodFilter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_unavailable_foods_of_branches_proto protoreflect.FileDescriptor

var file_unavailable_foods_of_branches_proto_rawDesc = []byte{
	0x0a, 0x23, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6f,
	0x6f, 0x64, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64,
	0x49, 0x64, 0x22, 0xd5, 0x01, 0x0a, 0x0f, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x9b, 0x01, 0x0a, 0x15, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x75, 0x0a, 0x10, 0x75, 0x6e, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x4b, 0x0a, 0x10,
	0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x10, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x78, 0x0a, 0x16, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f,
	0x6f, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x32, 0xf1, 0x03, 0x0a, 0x25, 0x75, 0x6e,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x5f,
	0x6f, 0x66, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x46, 0x6f, 0x6f, 0x64, 0x1a, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x46, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x12, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1f, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x6e,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x52, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x26, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x20, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f, 0x6f, 0x64,
	0x73, 0x12, 0x50, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x6f,
	0x6f, 0x64, 0x1a, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x75, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x46,
	0x6f, 0x6f, 0x64, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x4c, 0x0a, 0x18, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x6e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42, 0x19, 0x5a,
	0x17, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_unavailable_foods_of_branches_proto_rawDescOnce sync.Once
	file_unavailable_foods_of_branches_proto_rawDescData = file_unavailable_foods_of_branches_proto_rawDesc
)

func file_unavailable_foods_of_branches_proto_rawDescGZIP() []byte {
	file_unavailable_foods_of_branches_proto_rawDescOnce.Do(func() {
		file_unavailable_foods_of_branches_proto_rawDescData = protoimpl.X.CompressGZIP(file_unavailable_foods_of_branches_proto_rawDescData)
	})
	return file_unavailable_foods_of_branches_proto_rawDescData
}

var file_unavailable_foods_of_branches_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_unavailable_foods_of_branches_proto_goTypes = []interface{}{
	(*CreateUnavailableFood)(nil), // 0: orders_service.createUnavailableFood
	(*UnavailableFood)(nil),       // 1: orders_service.unavailableFood
	(*UpdateUnavailableFood)(nil), // 2: orders_service.updateUnavailableFood
	(*UnavailableFoods)(nil),      // 3: orders_service.unavailableFoods
	(*UnavailableFoodFilter)(nil), // 4: orders_service.unavailableFood_filter
	(*PrimaryKey)(nil),            // 5: orders_service.PrimaryKey
	(*Void)(nil),                  // 6: orders_service.Void
}
var file_unavailable_foods_of_branches_proto_depIdxs = []int32{
	1, // 0: orders_service.unavailableFoods.unavailableFoods:type_name -> orders_service.unavailableFood
	0, // 1: orders_service.unavailable_foods_of_branches_service.Create:input_type -> orders_service.createUnavailableFood
	5, // 2: orders_service.unavailable_foods_of_branches_service.getById:input_type -> orders_service.PrimaryKey
	4, // 3: orders_service.unavailable_foods_of_branches_service.GetAll:input_type -> orders_service.unavailableFood_filter
	2, // 4: orders_service.unavailable_foods_of_branches_service.Update:input_type -> orders_service.updateUnavailableFood
	5, // 5: orders_service.unavailable_foods_of_branches_service.Delete:input_type -> orders_service.PrimaryKey
	5, // 6: orders_service.unavailable_foods_of_branches_service.CheckUnavailabeFoodExist:input_type -> orders_service.PrimaryKey
	1, // 7: orders_service.unavailable_foods_of_branches_service.Create:output_type -> orders_service.unavailableFood
	1, // 8: orders_service.unavailable_foods_of_branches_service.getById:output_type -> orders_service.unavailableFood
	3, // 9: orders_service.unavailable_foods_of_branches_service.GetAll:output_type -> orders_service.unavailableFoods
	1, // 10: orders_service.unavailable_foods_of_branches_service.Update:output_type -> orders_service.unavailableFood
	6, // 11: orders_service.unavailable_foods_of_branches_service.Delete:output_type -> orders_service.Void
	6, // 12: orders_service.unavailable_foods_of_branches_service.CheckUnavailabeFoodExist:output_type -> orders_service.Void
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_unavailable_foods_of_branches_proto_init() }
func file_unavailable_foods_of_branches_proto_init() {
	if File_unavailable_foods_of_branches_proto != nil {
		return
	}
	file_orders_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_unavailable_foods_of_branches_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUnavailableFood); i {
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
		file_unavailable_foods_of_branches_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnavailableFood); i {
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
		file_unavailable_foods_of_branches_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUnavailableFood); i {
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
		file_unavailable_foods_of_branches_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnavailableFoods); i {
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
		file_unavailable_foods_of_branches_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnavailableFoodFilter); i {
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
			RawDescriptor: file_unavailable_foods_of_branches_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_unavailable_foods_of_branches_proto_goTypes,
		DependencyIndexes: file_unavailable_foods_of_branches_proto_depIdxs,
		MessageInfos:      file_unavailable_foods_of_branches_proto_msgTypes,
	}.Build()
	File_unavailable_foods_of_branches_proto = out.File
	file_unavailable_foods_of_branches_proto_rawDesc = nil
	file_unavailable_foods_of_branches_proto_goTypes = nil
	file_unavailable_foods_of_branches_proto_depIdxs = nil
}
