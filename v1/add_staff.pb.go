// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: add_staff.proto

package v1

import (
	_ "github.com/ramsfords/types_gen/v1/validate"
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

type AddStaff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Roles         []Role `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=v1.Role" json:"roles,omitempty"`
	NewStaffEmail string `protobuf:"bytes,3,opt,name=newStaffEmail,proto3" json:"newStaffEmail,omitempty"`
	Password      string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	FirstName     string `protobuf:"bytes,5,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string `protobuf:"bytes,6,opt,name=lastName,proto3" json:"lastName,omitempty"`
	PhoneNumber   string `protobuf:"bytes,7,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	BusinessId    string `protobuf:"bytes,8,opt,name=businessId,proto3" json:"businessId,omitempty"`
}

func (x *AddStaff) Reset() {
	*x = AddStaff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_add_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStaff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStaff) ProtoMessage() {}

func (x *AddStaff) ProtoReflect() protoreflect.Message {
	mi := &file_add_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStaff.ProtoReflect.Descriptor instead.
func (*AddStaff) Descriptor() ([]byte, []int) {
	return file_add_staff_proto_rawDescGZIP(), []int{0}
}

func (x *AddStaff) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AddStaff) GetRoles() []Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AddStaff) GetNewStaffEmail() string {
	if x != nil {
		return x.NewStaffEmail
	}
	return ""
}

func (x *AddStaff) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddStaff) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *AddStaff) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *AddStaff) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *AddStaff) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

var File_add_staff_proto protoreflect.FileDescriptor

var file_add_staff_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x64, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x02, 0x0a, 0x08, 0x61,
	0x64, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9,
	0xc0, 0x03, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72,
	0x02, 0x10, 0x08, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x42, 0x62, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x31, 0x42, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x74, 0x61, 0x66, 0x66, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x61, 0x6d, 0x73, 0x66, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02,
	0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_add_staff_proto_rawDescOnce sync.Once
	file_add_staff_proto_rawDescData = file_add_staff_proto_rawDesc
)

func file_add_staff_proto_rawDescGZIP() []byte {
	file_add_staff_proto_rawDescOnce.Do(func() {
		file_add_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_add_staff_proto_rawDescData)
	})
	return file_add_staff_proto_rawDescData
}

var file_add_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_add_staff_proto_goTypes = []interface{}{
	(*AddStaff)(nil), // 0: v1.addStaff
	(Role)(0),        // 1: v1.role
}
var file_add_staff_proto_depIdxs = []int32{
	1, // 0: v1.addStaff.roles:type_name -> v1.role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_add_staff_proto_init() }
func file_add_staff_proto_init() {
	if File_add_staff_proto != nil {
		return
	}
	file_role_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_add_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStaff); i {
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
			RawDescriptor: file_add_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_add_staff_proto_goTypes,
		DependencyIndexes: file_add_staff_proto_depIdxs,
		MessageInfos:      file_add_staff_proto_msgTypes,
	}.Build()
	File_add_staff_proto = out.File
	file_add_staff_proto_rawDesc = nil
	file_add_staff_proto_goTypes = nil
	file_add_staff_proto_depIdxs = nil
}
