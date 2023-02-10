// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: business.proto

package v1

import (
	_ "github.com/ramsfords/types_gen/v1/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Business struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"type"
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty" dynamodbav:"type"`
	// @gotags: dynamodbav:"businessName"
	BusinessName string `protobuf:"bytes,2,opt,name=businessName,proto3" json:"businessName,omitempty" dynamodbav:"businessName"`
	// @gotags: dynamodbav:"businessId"
	BusinessId string `protobuf:"bytes,3,opt,name=businessId,proto3" json:"businessId,omitempty" dynamodbav:"businessId"`
	// @gotags: dynamodbav:"businessEmail"
	BusinessEmail string `protobuf:"bytes,4,opt,name=businessEmail,proto3" json:"businessEmail,omitempty" dynamodbav:"businessEmail"`
	// @gotags: dynamodbav:"accountingEmail"
	AccountingEmail string `protobuf:"bytes,5,opt,name=accountingEmail,proto3" json:"accountingEmail,omitempty" dynamodbav:"accountingEmail"`
	// @gotags: dynamodbav:"customerServiceEmail"
	CustomerServiceEmail string `protobuf:"bytes,6,opt,name=customerServiceEmail,proto3" json:"customerServiceEmail,omitempty" dynamodbav:"customerServiceEmail"`
	// @gotags: dynamodbav:"highPriorityEmail"
	HighPriorityEmail string `protobuf:"bytes,7,opt,name=highPriorityEmail,proto3" json:"highPriorityEmail,omitempty" dynamodbav:"highPriorityEmail"`
	// @gotags: dynamodbav:"avatarURL"
	AvatarURL string `protobuf:"bytes,8,opt,name=avatarURL,proto3" json:"avatarURL,omitempty" dynamodbav:"avatarURL"`
	// @gotags: dynamodbav:"adminEmail"
	AdminEmail string `protobuf:"bytes,9,opt,name=adminEmail,proto3" json:"adminEmail,omitempty" dynamodbav:"adminEmail"`
	// @gotags: dynamodbav:"createAt"
	CreateAt string `protobuf:"bytes,10,opt,name=createAt,proto3" json:"createAt,omitempty" dynamodbav:"createAt"`
	// @gotags: dynamodbav:"updatedUt"
	UpdatedUt string `protobuf:"bytes,11,opt,name=updatedUt,proto3" json:"updatedUt,omitempty" dynamodbav:"updatedUt"`
	// @gotags: dynamodbav:"deletedAt"
	DeletedAt string `protobuf:"bytes,12,opt,name=deletedAt,proto3" json:"deletedAt,omitempty" dynamodbav:"deletedAt"`
	// @gotags: dynamodbav:"phoneNumber"
	PhoneNumber *PhoneNumber `protobuf:"bytes,13,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty" dynamodbav:"phoneNumber"`
	// @gotags: dynamodbav:"needsAddressUpdate"
	NeedsAddressUpdate bool `protobuf:"varint,14,opt,name=needsAddressUpdate,proto3" json:"needsAddressUpdate,omitempty" dynamodbav:"needsAddressUpdate"`
	// @gotags: dynamodbav:"needsDefaultPickupAddressUpdate"
	NeedsDefaultPickupAddressUpdate bool `protobuf:"varint,15,opt,name=needsDefaultPickupAddressUpdate,proto3" json:"needsDefaultPickupAddressUpdate,omitempty" dynamodbav:"needsDefaultPickupAddressUpdate"`
	// @gotags: dynamodbav:"needsDefaultDeliveryAddressUpdate"
	NeedsDefaultDeliveryAddressUpdate bool `protobuf:"varint,16,opt,name=needsDefaultDeliveryAddressUpdate,proto3" json:"needsDefaultDeliveryAddressUpdate,omitempty" dynamodbav:"needsDefaultDeliveryAddressUpdate"`
	// @gotags: dynamodbav:"billingAddress"
	BillingAddress *Address `protobuf:"bytes,17,opt,name=billingAddress,proto3" json:"billingAddress,omitempty" dynamodbav:"billingAddress"`
	// @gotags: dynamodbav:"defaultPickupAddress"
	DefaultPickupAddress *Address `protobuf:"bytes,18,opt,name=defaultPickupAddress,proto3" json:"defaultPickupAddress,omitempty" dynamodbav:"defaultPickupAddress"`
	// @gotags: dynamodbav:"defaultDeliveryAddress"
	DefaultDeliveryAddress *Address `protobuf:"bytes,19,opt,name=defaultDeliveryAddress,proto3" json:"defaultDeliveryAddress,omitempty" dynamodbav:"defaultDeliveryAddress"`
	// @gotags: dynamodbav:"address"
	Address *Address `protobuf:"bytes,20,opt,name=address,proto3" json:"address,omitempty" dynamodbav:"address"`
	// @gotags: dynamodbav:"allowBooking"
	AllowBooking bool `protobuf:"varint,21,opt,name=allowBooking,proto3" json:"allowBooking,omitempty" dynamodbav:"allowBooking"`
	// @gotags: dynamodbav:"bookingLimit"
	BookingLimit float32 `protobuf:"fixed32,23,opt,name=bookingLimit,proto3" json:"bookingLimit,omitempty" dynamodbav:"bookingLimit"`
	// @gotags: dynamodbav:"disabled"
	Disabled bool `protobuf:"varint,24,opt,name=disabled,proto3" json:"disabled,omitempty" dynamodbav:"disabled"`
	// @gotags: dynamodbav:"referredBy"
	ReferredBy string `protobuf:"bytes,25,opt,name=referredBy,proto3" json:"referredBy,omitempty" dynamodbav:"referredBy"`
	// @gotags: dynamodbav:"accountCreditLimit"
	AccountCreditLimit int32 `protobuf:"varint,26,opt,name=accountCreditLimit,proto3" json:"accountCreditLimit,omitempty" dynamodbav:"accountCreditLimit"`
	// @gotags: dynamodbav:"priceTier"
	PriceTier int32 `protobuf:"varint,27,opt,name=priceTier,proto3" json:"priceTier,omitempty" dynamodbav:"priceTier"`
	// @gotags: dynamodbav:"booksOpened"
	BooksOpened bool `protobuf:"varint,28,opt,name=booksOpened,proto3" json:"booksOpened,omitempty" dynamodbav:"booksOpened"`
	// @gotags: dynamodbav:"adminUser"
	AdminUser *User `protobuf:"bytes,29,opt,name=adminUser,proto3" json:"adminUser,omitempty" dynamodbav:"adminUser"`
	// @gotags: dynamodbav:"booksContactId"
	BooksContactId string `protobuf:"bytes,30,opt,name=booksContactId,proto3" json:"booksContactId,omitempty" dynamodbav:"booksContactId"`
}

func (x *Business) Reset() {
	*x = Business{}
	if protoimpl.UnsafeEnabled {
		mi := &file_business_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Business) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Business) ProtoMessage() {}

func (x *Business) ProtoReflect() protoreflect.Message {
	mi := &file_business_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Business.ProtoReflect.Descriptor instead.
func (*Business) Descriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{0}
}

func (x *Business) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Business) GetBusinessName() string {
	if x != nil {
		return x.BusinessName
	}
	return ""
}

func (x *Business) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *Business) GetBusinessEmail() string {
	if x != nil {
		return x.BusinessEmail
	}
	return ""
}

func (x *Business) GetAccountingEmail() string {
	if x != nil {
		return x.AccountingEmail
	}
	return ""
}

func (x *Business) GetCustomerServiceEmail() string {
	if x != nil {
		return x.CustomerServiceEmail
	}
	return ""
}

func (x *Business) GetHighPriorityEmail() string {
	if x != nil {
		return x.HighPriorityEmail
	}
	return ""
}

func (x *Business) GetAvatarURL() string {
	if x != nil {
		return x.AvatarURL
	}
	return ""
}

func (x *Business) GetAdminEmail() string {
	if x != nil {
		return x.AdminEmail
	}
	return ""
}

func (x *Business) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *Business) GetUpdatedUt() string {
	if x != nil {
		return x.UpdatedUt
	}
	return ""
}

func (x *Business) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

func (x *Business) GetPhoneNumber() *PhoneNumber {
	if x != nil {
		return x.PhoneNumber
	}
	return nil
}

func (x *Business) GetNeedsAddressUpdate() bool {
	if x != nil {
		return x.NeedsAddressUpdate
	}
	return false
}

func (x *Business) GetNeedsDefaultPickupAddressUpdate() bool {
	if x != nil {
		return x.NeedsDefaultPickupAddressUpdate
	}
	return false
}

func (x *Business) GetNeedsDefaultDeliveryAddressUpdate() bool {
	if x != nil {
		return x.NeedsDefaultDeliveryAddressUpdate
	}
	return false
}

func (x *Business) GetBillingAddress() *Address {
	if x != nil {
		return x.BillingAddress
	}
	return nil
}

func (x *Business) GetDefaultPickupAddress() *Address {
	if x != nil {
		return x.DefaultPickupAddress
	}
	return nil
}

func (x *Business) GetDefaultDeliveryAddress() *Address {
	if x != nil {
		return x.DefaultDeliveryAddress
	}
	return nil
}

func (x *Business) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Business) GetAllowBooking() bool {
	if x != nil {
		return x.AllowBooking
	}
	return false
}

func (x *Business) GetBookingLimit() float32 {
	if x != nil {
		return x.BookingLimit
	}
	return 0
}

func (x *Business) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Business) GetReferredBy() string {
	if x != nil {
		return x.ReferredBy
	}
	return ""
}

func (x *Business) GetAccountCreditLimit() int32 {
	if x != nil {
		return x.AccountCreditLimit
	}
	return 0
}

func (x *Business) GetPriceTier() int32 {
	if x != nil {
		return x.PriceTier
	}
	return 0
}

func (x *Business) GetBooksOpened() bool {
	if x != nil {
		return x.BooksOpened
	}
	return false
}

func (x *Business) GetAdminUser() *User {
	if x != nil {
		return x.AdminUser
	}
	return nil
}

func (x *Business) GetBooksContactId() string {
	if x != nil {
		return x.BooksContactId
	}
	return ""
}

type Businesses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"businesses"
	Businesses []*Business `protobuf:"bytes,1,rep,name=businesses,proto3" json:"businesses,omitempty" dynamodbav:"businesses"`
}

func (x *Businesses) Reset() {
	*x = Businesses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_business_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Businesses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Businesses) ProtoMessage() {}

func (x *Businesses) ProtoReflect() protoreflect.Message {
	mi := &file_business_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Businesses.ProtoReflect.Descriptor instead.
func (*Businesses) Descriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{1}
}

func (x *Businesses) GetBusinesses() []*Business {
	if x != nil {
		return x.Businesses
	}
	return nil
}

var File_business_proto protoreflect.FileDescriptor

var file_business_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x09,
	0x0a, 0x08, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e,
	0x0a, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01,
	0x52, 0x0c, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x32,
	0x0a, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x11, 0x68, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x68,
	0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x55, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31,
	0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x65,
	0x64, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a, 0x1f, 0x6e, 0x65, 0x65,
	0x64, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1f, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x21, 0x6e, 0x65, 0x65, 0x64, 0x73, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x21,
	0x6e, 0x65, 0x65, 0x64, 0x73, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x33, 0x0a, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3f, 0x0a, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x43, 0x0a, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x18, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x42, 0x79, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x42, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x12, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x54, 0x69, 0x65, 0x72, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x54, 0x69, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x4f, 0x70,
	0x65, 0x6e, 0x65, 0x64, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x42, 0x62, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x73, 0x66,
	0x6f, 0x72, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56,
	0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_business_proto_rawDescOnce sync.Once
	file_business_proto_rawDescData = file_business_proto_rawDesc
)

func file_business_proto_rawDescGZIP() []byte {
	file_business_proto_rawDescOnce.Do(func() {
		file_business_proto_rawDescData = protoimpl.X.CompressGZIP(file_business_proto_rawDescData)
	})
	return file_business_proto_rawDescData
}

var file_business_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_business_proto_goTypes = []interface{}{
	(*Business)(nil),    // 0: v1.business
	(*Businesses)(nil),  // 1: v1.businesses
	(*PhoneNumber)(nil), // 2: v1.phoneNumber
	(*Address)(nil),     // 3: v1.address
	(*User)(nil),        // 4: v1.user
}
var file_business_proto_depIdxs = []int32{
	2, // 0: v1.business.phoneNumber:type_name -> v1.phoneNumber
	3, // 1: v1.business.billingAddress:type_name -> v1.address
	3, // 2: v1.business.defaultPickupAddress:type_name -> v1.address
	3, // 3: v1.business.defaultDeliveryAddress:type_name -> v1.address
	3, // 4: v1.business.address:type_name -> v1.address
	4, // 5: v1.business.adminUser:type_name -> v1.user
	0, // 6: v1.businesses.businesses:type_name -> v1.business
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_business_proto_init() }
func file_business_proto_init() {
	if File_business_proto != nil {
		return
	}
	file_location_proto_init()
	file_address_proto_init()
	file_user_proto_init()
	file_phone_number_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_business_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Business); i {
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
		file_business_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Businesses); i {
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
			RawDescriptor: file_business_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_business_proto_goTypes,
		DependencyIndexes: file_business_proto_depIdxs,
		MessageInfos:      file_business_proto_msgTypes,
	}.Build()
	File_business_proto = out.File
	file_business_proto_rawDesc = nil
	file_business_proto_goTypes = nil
	file_business_proto_depIdxs = nil
}
