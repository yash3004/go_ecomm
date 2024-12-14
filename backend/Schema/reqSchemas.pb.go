// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: reqSchemas.proto

package Schema

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	PublicId   string `protobuf:"bytes,4,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	FirstName  string `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName string `protobuf:"bytes,6,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName   string `protobuf:"bytes,7,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email      string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNo    string `protobuf:"bytes,9,opt,name=phone_no,json=phoneNo,proto3" json:"phone_no,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_reqSchemas_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_reqSchemas_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_reqSchemas_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

type ValidateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ValidateUser) Reset() {
	*x = ValidateUser{}
	mi := &file_reqSchemas_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateUser) ProtoMessage() {}

func (x *ValidateUser) ProtoReflect() protoreflect.Message {
	mi := &file_reqSchemas_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateUser.ProtoReflect.Descriptor instead.
func (*ValidateUser) Descriptor() ([]byte, []int) {
	return file_reqSchemas_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ValidateUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ValidateUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PublicUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	PublicId   string `protobuf:"bytes,3,opt,name=public_id,json=publicId,proto3" json:"public_id,omitempty"`
	FirstName  string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName string `protobuf:"bytes,5,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName   string `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email      string `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNo    string `protobuf:"bytes,8,opt,name=phone_no,json=phoneNo,proto3" json:"phone_no,omitempty"`
}

func (x *PublicUser) Reset() {
	*x = PublicUser{}
	mi := &file_reqSchemas_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicUser) ProtoMessage() {}

func (x *PublicUser) ProtoReflect() protoreflect.Message {
	mi := &file_reqSchemas_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicUser.ProtoReflect.Descriptor instead.
func (*PublicUser) Descriptor() ([]byte, []int) {
	return file_reqSchemas_proto_rawDescGZIP(), []int{2}
}

func (x *PublicUser) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PublicUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PublicUser) GetPublicId() string {
	if x != nil {
		return x.PublicId
	}
	return ""
}

func (x *PublicUser) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PublicUser) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *PublicUser) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PublicUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PublicUser) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

type UserAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddrId     string `protobuf:"bytes,1,opt,name=addr_id,json=addrId,proto3" json:"addr_id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Country    string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	State      string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	City       string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Address1   string `protobuf:"bytes,6,opt,name=address1,proto3" json:"address1,omitempty"`
	Address2   string `protobuf:"bytes,7,opt,name=address2,proto3" json:"address2,omitempty"`
	Postalcode int32  `protobuf:"varint,8,opt,name=postalcode,proto3" json:"postalcode,omitempty"`
}

func (x *UserAddress) Reset() {
	*x = UserAddress{}
	mi := &file_reqSchemas_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAddress) ProtoMessage() {}

func (x *UserAddress) ProtoReflect() protoreflect.Message {
	mi := &file_reqSchemas_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAddress.ProtoReflect.Descriptor instead.
func (*UserAddress) Descriptor() ([]byte, []int) {
	return file_reqSchemas_proto_rawDescGZIP(), []int{3}
}

func (x *UserAddress) GetAddrId() string {
	if x != nil {
		return x.AddrId
	}
	return ""
}

func (x *UserAddress) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserAddress) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UserAddress) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *UserAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserAddress) GetAddress1() string {
	if x != nil {
		return x.Address1
	}
	return ""
}

func (x *UserAddress) GetAddress2() string {
	if x != nil {
		return x.Address2
	}
	return ""
}

func (x *UserAddress) GetPostalcode() int32 {
	if x != nil {
		return x.Postalcode
	}
	return 0
}

var File_reqSchemas_proto protoreflect.FileDescriptor

var file_reqSchemas_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x71, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x82, 0x02, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x6f,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x22,
	0x60, 0x0a, 0x0c, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0xec, 0x01, 0x0a, 0x0a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f,
	0x22, 0xdb, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x64, 0x64, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x1e,
	0x5a, 0x1c, 0x67, 0x6f, 0x5f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x2f, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reqSchemas_proto_rawDescOnce sync.Once
	file_reqSchemas_proto_rawDescData = file_reqSchemas_proto_rawDesc
)

func file_reqSchemas_proto_rawDescGZIP() []byte {
	file_reqSchemas_proto_rawDescOnce.Do(func() {
		file_reqSchemas_proto_rawDescData = protoimpl.X.CompressGZIP(file_reqSchemas_proto_rawDescData)
	})
	return file_reqSchemas_proto_rawDescData
}

var file_reqSchemas_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_reqSchemas_proto_goTypes = []any{
	(*User)(nil),         // 0: Schema.User
	(*ValidateUser)(nil), // 1: Schema.ValidateUser
	(*PublicUser)(nil),   // 2: Schema.PublicUser
	(*UserAddress)(nil),  // 3: Schema.UserAddress
}
var file_reqSchemas_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_reqSchemas_proto_init() }
func file_reqSchemas_proto_init() {
	if File_reqSchemas_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_reqSchemas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reqSchemas_proto_goTypes,
		DependencyIndexes: file_reqSchemas_proto_depIdxs,
		MessageInfos:      file_reqSchemas_proto_msgTypes,
	}.Build()
	File_reqSchemas_proto = out.File
	file_reqSchemas_proto_rawDesc = nil
	file_reqSchemas_proto_goTypes = nil
	file_reqSchemas_proto_depIdxs = nil
}
