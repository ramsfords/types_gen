// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: book.proto

package v1

import (
	_ "github.com/ramsfords/types_gen/v1/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FreightChargeTerm int32

const (
	FreightChargeTerm_CHARGE_NONE     FreightChargeTerm = 0
	FreightChargeTerm_THIRD_PARTY     FreightChargeTerm = 1
	FreightChargeTerm_INBOUND_COLLECT FreightChargeTerm = 2
	FreightChargeTerm_PREPAID         FreightChargeTerm = 3
)

// Enum value maps for FreightChargeTerm.
var (
	FreightChargeTerm_name = map[int32]string{
		0: "CHARGE_NONE",
		1: "THIRD_PARTY",
		2: "INBOUND_COLLECT",
		3: "PREPAID",
	}
	FreightChargeTerm_value = map[string]int32{
		"CHARGE_NONE":     0,
		"THIRD_PARTY":     1,
		"INBOUND_COLLECT": 2,
		"PREPAID":         3,
	}
)

func (x FreightChargeTerm) Enum() *FreightChargeTerm {
	p := new(FreightChargeTerm)
	*p = x
	return p
}

func (x FreightChargeTerm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FreightChargeTerm) Descriptor() protoreflect.EnumDescriptor {
	return file_book_proto_enumTypes[0].Descriptor()
}

func (FreightChargeTerm) Type() protoreflect.EnumType {
	return &file_book_proto_enumTypes[0]
}

func (x FreightChargeTerm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FreightChargeTerm.Descriptor instead.
func (FreightChargeTerm) EnumDescriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"shipment_details,omitempty"
	ShipmentDetails *ShipmentDetails `protobuf:"bytes,1,opt,name=shipment_details,json=shipmentDetails,proto3" json:"shipment_details,omitempty" dynamodbav:"shipment_details,omitempty"`
	// @gotags: dynamodbav:"commodities,omitempty"
	Commodities []*Commodity `protobuf:"bytes,2,rep,name=commodities,proto3" json:"commodities,omitempty" dynamodbav:"commodities,omitempty"`
	// @gotags: dynamodbav:"pickup,omitempty"
	Pickup *Location `protobuf:"bytes,3,opt,name=pickup,proto3" json:"pickup,omitempty" dynamodbav:"pickup,omitempty"`
	// @gotags: dynamodbav:"delivery,omitempty"
	Delivery *Location `protobuf:"bytes,4,opt,name=delivery,proto3" json:"delivery,omitempty" dynamodbav:"delivery,omitempty"`
	// @gotags: dynamodbav:"bid,omitempty"
	Bid *Bid `protobuf:"bytes,5,opt,name=bid,proto3" json:"bid,omitempty" dynamodbav:"bid,omitempty"`
	// @gotags: dynamodbav:"booking_success,omitempty"
	BookingSuccess bool `protobuf:"varint,6,opt,name=booking_success,json=bookingSuccess,proto3" json:"booking_success,omitempty" dynamodbav:"booking_success,omitempty"`
	// @gotags: dynamodbav:"rapid_booking_confirmation,omitempty"
	DispatchResponse *DispatchResponse `protobuf:"bytes,7,opt,name=DispatchResponse,proto3" json:"DispatchResponse,omitempty" dynamodbav:"rapid_booking_confirmation,omitempty"`
	// @gotags: dynamodbav:"booking_info,omitempty"
	BookingInfo *BookingInfo `protobuf:"bytes,8,opt,name=booking_info,json=bookingInfo,proto3" json:"booking_info,omitempty" dynamodbav:"booking_info,omitempty"`
	// @gotags: dynamodbav:"business_id,omitempty"
	BusinessId string `protobuf:"bytes,9,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty" dynamodbav:"business_id,omitempty"`
	// @gotags: dynamodbav:"quote_id,omitempty"
	QuoteId string `protobuf:"bytes,10,opt,name=quote_id,json=quoteId,proto3" json:"quote_id,omitempty" dynamodbav:"quote_id,omitempty"`
	// @gotags: dynamodbav:"svg_data,omitempty"
	SvgData string `protobuf:"bytes,11,opt,name=svg_data,json=svgData,proto3" json:"svg_data,omitempty" dynamodbav:"svg_data,omitempty"`
	// @gotags: dynamodbav:"bol_url,omitempty"
	BolUrl string `protobuf:"bytes,12,opt,name=bol_url,json=bolUrl,proto3" json:"bol_url,omitempty" dynamodbav:"bol_url,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *BookingResponse) GetShipmentDetails() *ShipmentDetails {
	if x != nil {
		return x.ShipmentDetails
	}
	return nil
}

func (x *BookingResponse) GetCommodities() []*Commodity {
	if x != nil {
		return x.Commodities
	}
	return nil
}

func (x *BookingResponse) GetPickup() *Location {
	if x != nil {
		return x.Pickup
	}
	return nil
}

func (x *BookingResponse) GetDelivery() *Location {
	if x != nil {
		return x.Delivery
	}
	return nil
}

func (x *BookingResponse) GetBid() *Bid {
	if x != nil {
		return x.Bid
	}
	return nil
}

func (x *BookingResponse) GetBookingSuccess() bool {
	if x != nil {
		return x.BookingSuccess
	}
	return false
}

func (x *BookingResponse) GetDispatchResponse() *DispatchResponse {
	if x != nil {
		return x.DispatchResponse
	}
	return nil
}

func (x *BookingResponse) GetBookingInfo() *BookingInfo {
	if x != nil {
		return x.BookingInfo
	}
	return nil
}

func (x *BookingResponse) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *BookingResponse) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

func (x *BookingResponse) GetSvgData() string {
	if x != nil {
		return x.SvgData
	}
	return ""
}

func (x *BookingResponse) GetBolUrl() string {
	if x != nil {
		return x.BolUrl
	}
	return ""
}

type BookingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstShipperBolNumber        string            `protobuf:"bytes,1,opt,name=first_shipper_bol_number,json=firstShipperBolNumber,proto3" json:"first_shipper_bol_number,omitempty"`
	FreightTerm                  FreightChargeTerm `protobuf:"varint,2,opt,name=freight_term,json=freightTerm,proto3,enum=v1.FreightChargeTerm" json:"freight_term,omitempty"`
	SpecialInstructions          string            `protobuf:"bytes,3,opt,name=special_instructions,json=specialInstructions,proto3" json:"special_instructions,omitempty"`
	CarrierName                  string            `protobuf:"bytes,4,opt,name=carrier_name,json=carrierName,proto3" json:"carrier_name,omitempty"`
	CarrierProNumber             string            `protobuf:"bytes,5,opt,name=carrier_pro_number,json=carrierProNumber,proto3" json:"carrier_pro_number,omitempty"`
	CarrierLogoUrl               string            `protobuf:"bytes,6,opt,name=carrier_logo_url,json=carrierLogoUrl,proto3" json:"carrier_logo_url,omitempty"`
	CarrierDispatchContactNumber string            `protobuf:"bytes,7,opt,name=carrier_dispatch_contact_number,json=carrierDispatchContactNumber,proto3" json:"carrier_dispatch_contact_number,omitempty"`
	CarrierDispatchEmail         string            `protobuf:"bytes,8,opt,name=carrier_dispatch_email,json=carrierDispatchEmail,proto3" json:"carrier_dispatch_email,omitempty"`
	CarrierBolNumber             string            `protobuf:"bytes,9,opt,name=carrier_bol_number,json=carrierBolNumber,proto3" json:"carrier_bol_number,omitempty"`
	CarrierReference             string            `protobuf:"bytes,10,opt,name=carrier_reference,json=carrierReference,proto3" json:"carrier_reference,omitempty"`
}

func (x *BookingInfo) Reset() {
	*x = BookingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingInfo) ProtoMessage() {}

func (x *BookingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingInfo.ProtoReflect.Descriptor instead.
func (*BookingInfo) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{1}
}

func (x *BookingInfo) GetFirstShipperBolNumber() string {
	if x != nil {
		return x.FirstShipperBolNumber
	}
	return ""
}

func (x *BookingInfo) GetFreightTerm() FreightChargeTerm {
	if x != nil {
		return x.FreightTerm
	}
	return FreightChargeTerm_CHARGE_NONE
}

func (x *BookingInfo) GetSpecialInstructions() string {
	if x != nil {
		return x.SpecialInstructions
	}
	return ""
}

func (x *BookingInfo) GetCarrierName() string {
	if x != nil {
		return x.CarrierName
	}
	return ""
}

func (x *BookingInfo) GetCarrierProNumber() string {
	if x != nil {
		return x.CarrierProNumber
	}
	return ""
}

func (x *BookingInfo) GetCarrierLogoUrl() string {
	if x != nil {
		return x.CarrierLogoUrl
	}
	return ""
}

func (x *BookingInfo) GetCarrierDispatchContactNumber() string {
	if x != nil {
		return x.CarrierDispatchContactNumber
	}
	return ""
}

func (x *BookingInfo) GetCarrierDispatchEmail() string {
	if x != nil {
		return x.CarrierDispatchEmail
	}
	return ""
}

func (x *BookingInfo) GetCarrierBolNumber() string {
	if x != nil {
		return x.CarrierBolNumber
	}
	return ""
}

func (x *BookingInfo) GetCarrierReference() string {
	if x != nil {
		return x.CarrierReference
	}
	return ""
}

type Bookings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"bookings,omitempty"
	BookingResponses []*BookingResponse `protobuf:"bytes,1,rep,name=booking_responses,json=bookingResponses,proto3" json:"booking_responses,omitempty" dynamodbav:"bookings,omitempty"`
}

func (x *Bookings) Reset() {
	*x = Bookings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookings) ProtoMessage() {}

func (x *Bookings) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookings.ProtoReflect.Descriptor instead.
func (*Bookings) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{2}
}

func (x *Bookings) GetBookingResponses() []*BookingResponse {
	if x != nil {
		return x.BookingResponses
	}
	return nil
}

type DispatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"shipmentId,omitempty"
	ShipmentID int32 `protobuf:"varint,1,opt,name=ShipmentID,proto3" json:"ShipmentID,omitempty" dynamodbav:"shipmentId,omitempty"`
	// @gotags: dynamodbav:"securityKey,omitempty"
	SecurityKey string `protobuf:"bytes,2,opt,name=SecurityKey,proto3" json:"SecurityKey,omitempty" dynamodbav:"securityKey,omitempty"`
	// @gotags: dynamodbav:"pickupNumber,omitempty"
	PickupNumber string `protobuf:"bytes,3,opt,name=PickupNumber,proto3" json:"PickupNumber,omitempty" dynamodbav:"pickupNumber,omitempty"`
	// @gotags: dynamodbav:"carrierName,omitempty"
	CarrierName string `protobuf:"bytes,4,opt,name=CarrierName,proto3" json:"CarrierName,omitempty" dynamodbav:"carrierName,omitempty"`
	// @gotags: dynamodbav:"carrierPhone,omitempty"
	CarrierPhone string `protobuf:"bytes,5,opt,name=CarrierPhone,proto3" json:"CarrierPhone,omitempty" dynamodbav:"carrierPhone,omitempty"`
	// @gotags: dynamodbav:"carrierPRONumber,omitempty"
	CarrierPRONumber string `protobuf:"bytes,6,opt,name=CarrierPRONumber,proto3" json:"CarrierPRONumber,omitempty" dynamodbav:"carrierPRONumber,omitempty"`
	// @gotags: dynamodbav:"handlingUnitTotal,omitempty"
	HandlingUnitTotal float64 `protobuf:"fixed64,7,opt,name=HandlingUnitTotal,proto3" json:"HandlingUnitTotal,omitempty" dynamodbav:"handlingUnitTotal,omitempty"`
	// @gotags: dynamodbav:"isShipmentEdit,omitempty"
	IsShipmentEdit bool `protobuf:"varint,8,opt,name=IsShipmentEdit,proto3" json:"IsShipmentEdit,omitempty" dynamodbav:"isShipmentEdit,omitempty"`
	// @gotags: dynamodbav:"isShipmentManual,omitempty"
	IsShipmentManual bool `protobuf:"varint,9,opt,name=IsShipmentManual,proto3" json:"IsShipmentManual,omitempty" dynamodbav:"isShipmentManual,omitempty"`
	// @gotags: dynamodbav:"serviceType,omitempty"
	ServiceType int32 `protobuf:"varint,10,opt,name=ServiceType,proto3" json:"ServiceType,omitempty" dynamodbav:"serviceType,omitempty"`
	// @gotags: dynamodbav:"isTrackingEmailSend,omitempty"
	IsTrackingEmailSend bool `protobuf:"varint,11,opt,name=IsTrackingEmailSend,proto3" json:"IsTrackingEmailSend,omitempty" dynamodbav:"isTrackingEmailSend,omitempty"`
	// @gotags: dynamodbav:"isTrackingAPIEnabled,omitempty"
	IsTrackingAPIEnabled bool `protobuf:"varint,12,opt,name=IsTrackingAPIEnabled,proto3" json:"IsTrackingAPIEnabled,omitempty" dynamodbav:"isTrackingAPIEnabled,omitempty"`
	// @gotags: dynamodbav:"customerBOLNumber,omitempty"
	CustomerBOLNumber string `protobuf:"bytes,13,opt,name=CustomerBOLNumber,proto3" json:"CustomerBOLNumber,omitempty" dynamodbav:"customerBOLNumber,omitempty"`
	// @gotags: dynamodbav:"shipperEmail,omitempty"
	ShipperEmail string `protobuf:"bytes,14,opt,name=ShipperEmail,proto3" json:"ShipperEmail,omitempty" dynamodbav:"shipperEmail,omitempty"`
	// @gotags: dynamodbav:"consigneeEmail,omitempty"
	ConsigneeEmail string `protobuf:"bytes,15,opt,name=ConsigneeEmail,proto3" json:"ConsigneeEmail,omitempty" dynamodbav:"consigneeEmail,omitempty"`
	// @gotags: dynamodbav:"result,omitempty"
	Result *Result `protobuf:"bytes,16,opt,name=Result,proto3" json:"Result,omitempty" dynamodbav:"result,omitempty"`
}

func (x *DispatchResponse) Reset() {
	*x = DispatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchResponse) ProtoMessage() {}

func (x *DispatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchResponse.ProtoReflect.Descriptor instead.
func (*DispatchResponse) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{3}
}

func (x *DispatchResponse) GetShipmentID() int32 {
	if x != nil {
		return x.ShipmentID
	}
	return 0
}

func (x *DispatchResponse) GetSecurityKey() string {
	if x != nil {
		return x.SecurityKey
	}
	return ""
}

func (x *DispatchResponse) GetPickupNumber() string {
	if x != nil {
		return x.PickupNumber
	}
	return ""
}

func (x *DispatchResponse) GetCarrierName() string {
	if x != nil {
		return x.CarrierName
	}
	return ""
}

func (x *DispatchResponse) GetCarrierPhone() string {
	if x != nil {
		return x.CarrierPhone
	}
	return ""
}

func (x *DispatchResponse) GetCarrierPRONumber() string {
	if x != nil {
		return x.CarrierPRONumber
	}
	return ""
}

func (x *DispatchResponse) GetHandlingUnitTotal() float64 {
	if x != nil {
		return x.HandlingUnitTotal
	}
	return 0
}

func (x *DispatchResponse) GetIsShipmentEdit() bool {
	if x != nil {
		return x.IsShipmentEdit
	}
	return false
}

func (x *DispatchResponse) GetIsShipmentManual() bool {
	if x != nil {
		return x.IsShipmentManual
	}
	return false
}

func (x *DispatchResponse) GetServiceType() int32 {
	if x != nil {
		return x.ServiceType
	}
	return 0
}

func (x *DispatchResponse) GetIsTrackingEmailSend() bool {
	if x != nil {
		return x.IsTrackingEmailSend
	}
	return false
}

func (x *DispatchResponse) GetIsTrackingAPIEnabled() bool {
	if x != nil {
		return x.IsTrackingAPIEnabled
	}
	return false
}

func (x *DispatchResponse) GetCustomerBOLNumber() string {
	if x != nil {
		return x.CustomerBOLNumber
	}
	return ""
}

func (x *DispatchResponse) GetShipperEmail() string {
	if x != nil {
		return x.ShipperEmail
	}
	return ""
}

func (x *DispatchResponse) GetConsigneeEmail() string {
	if x != nil {
		return x.ConsigneeEmail
	}
	return ""
}

func (x *DispatchResponse) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"capacityProviderBolUrl,omitempty"
	CapacityProviderBolUrl string `protobuf:"bytes,1,opt,name=CapacityProviderBolUrl,proto3" json:"CapacityProviderBolUrl,omitempty" dynamodbav:"capacityProviderBolUrl,omitempty"`
	// @gotags: dynamodbav:"shipmentIdentifier,omitempty"
	ShipmentIdentifier string `protobuf:"bytes,2,opt,name=ShipmentIdentifier,proto3" json:"ShipmentIdentifier,omitempty" dynamodbav:"shipmentIdentifier,omitempty"`
	// @gotags: dynamodbav:"pickupNote,omitempty"
	PickupNote string `protobuf:"bytes,3,opt,name=PickupNote,proto3" json:"PickupNote,omitempty" dynamodbav:"pickupNote,omitempty"`
	// @gotags: dynamodbav:"pickupDateTime,omitempty"
	PickupDateTime string `protobuf:"bytes,4,opt,name=PickupDateTime,proto3" json:"PickupDateTime,omitempty" dynamodbav:"pickupDateTime,omitempty"`
	// @gotags: dynamodbav:"errors,omitempty"
	Errors []string `protobuf:"bytes,5,rep,name=Errors,proto3" json:"Errors,omitempty" dynamodbav:"errors,omitempty"`
	// @gotags: dynamodbav:"infoMessages,omitempty"
	InfoMessages []string `protobuf:"bytes,6,rep,name=InfoMessages,proto3" json:"InfoMessages,omitempty" dynamodbav:"infoMessages,omitempty"`
	// @gotags: dynamodbav:"type,omitempty"
	Type string `protobuf:"bytes,7,opt,name=Type,proto3" json:"Type,omitempty" dynamodbav:"type,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetCapacityProviderBolUrl() string {
	if x != nil {
		return x.CapacityProviderBolUrl
	}
	return ""
}

func (x *Result) GetShipmentIdentifier() string {
	if x != nil {
		return x.ShipmentIdentifier
	}
	return ""
}

func (x *Result) GetPickupNote() string {
	if x != nil {
		return x.PickupNote
	}
	return ""
}

func (x *Result) GetPickupDateTime() string {
	if x != nil {
		return x.PickupDateTime
	}
	return ""
}

func (x *Result) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *Result) GetInfoMessages() []string {
	if x != nil {
		return x.InfoMessages
	}
	return nil
}

func (x *Result) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x72, 0x65, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x62, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x04, 0x0a, 0x10, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x10, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x0a,
	0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x73, 0x68, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x42,
	0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x06, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x34, 0x0a, 0x08, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76,
	0x31, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03,
	0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x12, 0x19, 0x0a, 0x03, 0x62, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x76, 0x31, 0x2e, 0x62, 0x69, 0x64, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x40, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76,
	0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x0b,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x76, 0x67, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x76, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x6f, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6f, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x89, 0x04, 0x0a, 0x0c,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x37, 0x0a, 0x18,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f,
	0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x6c, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x0c, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x76, 0x31,
	0x2e, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f,
	0x74, 0x65, 0x72, 0x6d, 0x52, 0x0b, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x54, 0x65, 0x72,
	0x6d, 0x12, 0x31, 0x0a, 0x14, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72,
	0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12,
	0x45, 0x0a, 0x1f, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x16, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x44,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2c, 0x0a, 0x12,
	0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x42, 0x6f, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x41, 0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x76, 0x31, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x92, 0x05, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a,
	0x0c, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x50, 0x52, 0x4f, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x50, 0x52, 0x4f, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x55,
	0x6e, 0x69, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x69, 0x74, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x73, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x45,
	0x64, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x49, 0x73, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x64, 0x69, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x49, 0x73, 0x53,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x10, 0x49, 0x73, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x49, 0x73, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x49, 0x73, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x49, 0x73, 0x54,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x50, 0x49, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x49, 0x73, 0x54, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x41, 0x50, 0x49, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2c, 0x0a,
	0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x4f, 0x4c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x42, 0x4f, 0x4c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x88, 0x02, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6c, 0x55, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x6f, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x2e,
	0x0a, 0x12, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x53, 0x68, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x44, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x2a, 0x59, 0x0a, 0x13, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x0f, 0x0a,
	0x0b, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x59, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x49, 0x4e, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45,
	0x43, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x45, 0x50, 0x41, 0x49, 0x44, 0x10,
	0x03, 0x42, 0x5e, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x42, 0x6f, 0x6f,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x73, 0x66, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58,
	0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_proto_rawDescOnce sync.Once
	file_book_proto_rawDescData = file_book_proto_rawDesc
)

func file_book_proto_rawDescGZIP() []byte {
	file_book_proto_rawDescOnce.Do(func() {
		file_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_proto_rawDescData)
	})
	return file_book_proto_rawDescData
}

var file_book_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_book_proto_goTypes = []interface{}{
	(FreightChargeTerm)(0),   // 0: v1.freight_charge_term
	(*BookingResponse)(nil),  // 1: v1.booking_response
	(*BookingInfo)(nil),      // 2: v1.booking_info
	(*Bookings)(nil),         // 3: v1.bookings
	(*DispatchResponse)(nil), // 4: v1.DispatchResponse
	(*Result)(nil),           // 5: v1.Result
	(*ShipmentDetails)(nil),  // 6: v1.shipment_details
	(*Commodity)(nil),        // 7: v1.commodity
	(*Location)(nil),         // 8: v1.location
	(*Bid)(nil),              // 9: v1.bid
}
var file_book_proto_depIdxs = []int32{
	6,  // 0: v1.booking_response.shipment_details:type_name -> v1.shipment_details
	7,  // 1: v1.booking_response.commodities:type_name -> v1.commodity
	8,  // 2: v1.booking_response.pickup:type_name -> v1.location
	8,  // 3: v1.booking_response.delivery:type_name -> v1.location
	9,  // 4: v1.booking_response.bid:type_name -> v1.bid
	4,  // 5: v1.booking_response.DispatchResponse:type_name -> v1.DispatchResponse
	2,  // 6: v1.booking_response.booking_info:type_name -> v1.booking_info
	0,  // 7: v1.booking_info.freight_term:type_name -> v1.freight_charge_term
	1,  // 8: v1.bookings.booking_responses:type_name -> v1.booking_response
	5,  // 9: v1.DispatchResponse.Result:type_name -> v1.Result
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	file_role_proto_init()
	file_amount_proto_init()
	file_ref_proto_init()
	file_quote_proto_init()
	file_bid_proto_init()
	file_address_proto_init()
	file_shipment_details_proto_init()
	file_location_proto_init()
	file_commodity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingResponse); i {
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
		file_book_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingInfo); i {
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
		file_book_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bookings); i {
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
		file_book_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchResponse); i {
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
		file_book_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_book_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		EnumInfos:         file_book_proto_enumTypes,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_rawDesc = nil
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}
