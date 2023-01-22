// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: location_services.proto

package v1

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

type PickupLocationService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"pickupLocationWithDock,omitempty"
	PickupLocationWithDock bool `protobuf:"varint,1,opt,name=pickupLocationWithDock,proto3" json:"pickupLocationWithDock,omitempty"`
	// @gotags: dynamodbav:"liftGatePickup,omitempty"
	LiftGatePickup bool `protobuf:"varint,2,opt,name=liftGatePickup,proto3" json:"liftGatePickup,omitempty"`
	// @gotags: dynamodbav:"pickupAppointment,omitempty"
	PickupAppointment bool `protobuf:"varint,3,opt,name=pickupAppointment,proto3" json:"pickupAppointment,omitempty"`
	// @gotags: dynamodbav:"insidePickup,omitempty"
	InsidePickup bool `protobuf:"varint,4,opt,name=insidePickup,proto3" json:"insidePickup,omitempty"`
	// @gotags: dynamodbav:"pickupNotification,omitempty"
	PickupNotification bool `protobuf:"varint,5,opt,name=pickupNotification,proto3" json:"pickupNotification,omitempty"`
	// @gotags: dynamodbav:"shipperDeliveryNotification,omitempty"
	ShipperDeliveryNotification bool `protobuf:"varint,6,opt,name=shipperDeliveryNotification,proto3" json:"shipperDeliveryNotification,omitempty"`
}

func (x *PickupLocationService) Reset() {
	*x = PickupLocationService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PickupLocationService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PickupLocationService) ProtoMessage() {}

func (x *PickupLocationService) ProtoReflect() protoreflect.Message {
	mi := &file_location_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PickupLocationService.ProtoReflect.Descriptor instead.
func (*PickupLocationService) Descriptor() ([]byte, []int) {
	return file_location_services_proto_rawDescGZIP(), []int{0}
}

func (x *PickupLocationService) GetPickupLocationWithDock() bool {
	if x != nil {
		return x.PickupLocationWithDock
	}
	return false
}

func (x *PickupLocationService) GetLiftGatePickup() bool {
	if x != nil {
		return x.LiftGatePickup
	}
	return false
}

func (x *PickupLocationService) GetPickupAppointment() bool {
	if x != nil {
		return x.PickupAppointment
	}
	return false
}

func (x *PickupLocationService) GetInsidePickup() bool {
	if x != nil {
		return x.InsidePickup
	}
	return false
}

func (x *PickupLocationService) GetPickupNotification() bool {
	if x != nil {
		return x.PickupNotification
	}
	return false
}

func (x *PickupLocationService) GetShipperDeliveryNotification() bool {
	if x != nil {
		return x.ShipperDeliveryNotification
	}
	return false
}

type DeliveryLocationService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"deliverLocationWithDock,omitempty"
	DeliverLocationWithDock bool `protobuf:"varint,1,opt,name=deliverLocationWithDock,proto3" json:"deliverLocationWithDock,omitempty"`
	// @gotags: dynamodbav:"liftGateDelivery,omitempty"
	LiftGateDelivery bool `protobuf:"varint,2,opt,name=liftGateDelivery,proto3" json:"liftGateDelivery,omitempty"`
	// @gotags: dynamodbav:"deliveryAppointment,omitempty"
	DeliveryAppointment bool `protobuf:"varint,3,opt,name=deliveryAppointment,proto3" json:"deliveryAppointment,omitempty"`
	// @gotags: dynamodbav:"insideDelivery,omitempty"
	InsideDelivery bool `protobuf:"varint,4,opt,name=insideDelivery,proto3" json:"insideDelivery,omitempty"`
	// @gotags: dynamodbav:"receiverPickupNotification,omitempty"
	ReceiverPickupNotification bool `protobuf:"varint,5,opt,name=receiverPickupNotification,proto3" json:"receiverPickupNotification,omitempty"`
	// @gotags: dynamodbav:"deliveryNotification,omitempty"
	DeliveryNotification bool `protobuf:"varint,6,opt,name=deliveryNotification,proto3" json:"deliveryNotification,omitempty"`
}

func (x *DeliveryLocationService) Reset() {
	*x = DeliveryLocationService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryLocationService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryLocationService) ProtoMessage() {}

func (x *DeliveryLocationService) ProtoReflect() protoreflect.Message {
	mi := &file_location_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryLocationService.ProtoReflect.Descriptor instead.
func (*DeliveryLocationService) Descriptor() ([]byte, []int) {
	return file_location_services_proto_rawDescGZIP(), []int{1}
}

func (x *DeliveryLocationService) GetDeliverLocationWithDock() bool {
	if x != nil {
		return x.DeliverLocationWithDock
	}
	return false
}

func (x *DeliveryLocationService) GetLiftGateDelivery() bool {
	if x != nil {
		return x.LiftGateDelivery
	}
	return false
}

func (x *DeliveryLocationService) GetDeliveryAppointment() bool {
	if x != nil {
		return x.DeliveryAppointment
	}
	return false
}

func (x *DeliveryLocationService) GetInsideDelivery() bool {
	if x != nil {
		return x.InsideDelivery
	}
	return false
}

func (x *DeliveryLocationService) GetReceiverPickupNotification() bool {
	if x != nil {
		return x.ReceiverPickupNotification
	}
	return false
}

func (x *DeliveryLocationService) GetDeliveryNotification() bool {
	if x != nil {
		return x.DeliveryNotification
	}
	return false
}

var File_location_services_proto protoreflect.FileDescriptor

var file_location_services_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xbb, 0x02,
	0x0a, 0x15, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x44, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x44, 0x6f, 0x63, 0x6b, 0x12,
	0x26, 0x0a, 0x0e, 0x6c, 0x69, 0x66, 0x74, 0x47, 0x61, 0x74, 0x65, 0x50, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6c, 0x69, 0x66, 0x74, 0x47, 0x61, 0x74,
	0x65, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x50,
	0x69, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x6e, 0x73,
	0x69, 0x64, 0x65, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x2e, 0x0a, 0x12, 0x70, 0x69, 0x63,
	0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x1b, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcd, 0x02, 0x0a, 0x17,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x44, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x44, 0x6f, 0x63,
	0x6b, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x69, 0x66, 0x74, 0x47, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6c, 0x69, 0x66,
	0x74, 0x47, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x30, 0x0a,
	0x13, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x41, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x12, 0x3e, 0x0a, 0x1a, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x50, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6a, 0x0a, 0x06, 0x63,
	0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x73, 0x66,
	0x6f, 0x72, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56,
	0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_location_services_proto_rawDescOnce sync.Once
	file_location_services_proto_rawDescData = file_location_services_proto_rawDesc
)

func file_location_services_proto_rawDescGZIP() []byte {
	file_location_services_proto_rawDescOnce.Do(func() {
		file_location_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_location_services_proto_rawDescData)
	})
	return file_location_services_proto_rawDescData
}

var file_location_services_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_location_services_proto_goTypes = []interface{}{
	(*PickupLocationService)(nil),   // 0: v1.pickupLocationService
	(*DeliveryLocationService)(nil), // 1: v1.deliveryLocationService
}
var file_location_services_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_location_services_proto_init() }
func file_location_services_proto_init() {
	if File_location_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_location_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PickupLocationService); i {
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
		file_location_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryLocationService); i {
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
			RawDescriptor: file_location_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_location_services_proto_goTypes,
		DependencyIndexes: file_location_services_proto_depIdxs,
		MessageInfos:      file_location_services_proto_msgTypes,
	}.Build()
	File_location_services_proto = out.File
	file_location_services_proto_rawDesc = nil
	file_location_services_proto_goTypes = nil
	file_location_services_proto_depIdxs = nil
}
