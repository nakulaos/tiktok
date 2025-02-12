// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.2
// source: live.proto

package live

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

type StartLiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` // 用户id
}

func (x *StartLiveRequest) Reset() {
	*x = StartLiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartLiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartLiveRequest) ProtoMessage() {}

func (x *StartLiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_live_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartLiveRequest.ProtoReflect.Descriptor instead.
func (*StartLiveRequest) Descriptor() ([]byte, []int) {
	return file_live_proto_rawDescGZIP(), []int{0}
}

func (x *StartLiveRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type StartLiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamUrl string `protobuf:"bytes,1,opt,name=stream_url,json=streamUrl,proto3" json:"stream_url,omitempty"`
}

func (x *StartLiveResponse) Reset() {
	*x = StartLiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartLiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartLiveResponse) ProtoMessage() {}

func (x *StartLiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_live_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartLiveResponse.ProtoReflect.Descriptor instead.
func (*StartLiveResponse) Descriptor() ([]byte, []int) {
	return file_live_proto_rawDescGZIP(), []int{1}
}

func (x *StartLiveResponse) GetStreamUrl() string {
	if x != nil {
		return x.StreamUrl
	}
	return ""
}

type ListLiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` // 用户id
}

func (x *ListLiveRequest) Reset() {
	*x = ListLiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLiveRequest) ProtoMessage() {}

func (x *ListLiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_live_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLiveRequest.ProtoReflect.Descriptor instead.
func (*ListLiveRequest) Descriptor() ([]byte, []int) {
	return file_live_proto_rawDescGZIP(), []int{2}
}

func (x *ListLiveRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ListLiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserList []*User `protobuf:"bytes,1,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`
}

func (x *ListLiveResponse) Reset() {
	*x = ListLiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLiveResponse) ProtoMessage() {}

func (x *ListLiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_live_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLiveResponse.ProtoReflect.Descriptor instead.
func (*ListLiveResponse) Descriptor() ([]byte, []int) {
	return file_live_proto_rawDescGZIP(), []int{3}
}

func (x *ListLiveResponse) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                             // 用户id
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                          // 用户名称
	IsFollow     bool   `protobuf:"varint,3,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"` // true-已关注，false-未关注
	Avatar       string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`                      //用户头像
	Signature    string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`                //个人简介
	LiveUrl      string `protobuf:"bytes,6,opt,name=live_url,json=liveUrl,proto3" json:"live_url,omitempty"`
	LiveCoverUrl string `protobuf:"bytes,7,opt,name=live_cover_url,json=liveCoverUrl,proto3" json:"live_cover_url,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_live_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_live_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *User) GetLiveUrl() string {
	if x != nil {
		return x.LiveUrl
	}
	return ""
}

func (x *User) GetLiveCoverUrl() string {
	if x != nil {
		return x.LiveCoverUrl
	}
	return ""
}

var File_live_proto protoreflect.FileDescriptor

var file_live_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69,
	0x76, 0x65, 0x22, 0x24, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x55, 0x72, 0x6c, 0x22, 0x23, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0x3b, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xbe,
	0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6c, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x32,
	0x84, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x16, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x15, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c,
	0x69, 0x76, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x6c, 0x69, 0x76, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_live_proto_rawDescOnce sync.Once
	file_live_proto_rawDescData = file_live_proto_rawDesc
)

func file_live_proto_rawDescGZIP() []byte {
	file_live_proto_rawDescOnce.Do(func() {
		file_live_proto_rawDescData = protoimpl.X.CompressGZIP(file_live_proto_rawDescData)
	})
	return file_live_proto_rawDescData
}

var file_live_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_live_proto_goTypes = []interface{}{
	(*StartLiveRequest)(nil),  // 0: live.StartLiveRequest
	(*StartLiveResponse)(nil), // 1: live.StartLiveResponse
	(*ListLiveRequest)(nil),   // 2: live.ListLiveRequest
	(*ListLiveResponse)(nil),  // 3: live.ListLiveResponse
	(*User)(nil),              // 4: live.User
}
var file_live_proto_depIdxs = []int32{
	4, // 0: live.ListLiveResponse.user_list:type_name -> live.User
	0, // 1: live.Live.StartLive:input_type -> live.StartLiveRequest
	2, // 2: live.Live.ListVideo:input_type -> live.ListLiveRequest
	1, // 3: live.Live.StartLive:output_type -> live.StartLiveResponse
	3, // 4: live.Live.ListVideo:output_type -> live.ListLiveResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_live_proto_init() }
func file_live_proto_init() {
	if File_live_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_live_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartLiveRequest); i {
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
		file_live_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartLiveResponse); i {
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
		file_live_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLiveRequest); i {
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
		file_live_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLiveResponse); i {
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
		file_live_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_live_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_live_proto_goTypes,
		DependencyIndexes: file_live_proto_depIdxs,
		MessageInfos:      file_live_proto_msgTypes,
	}.Build()
	File_live_proto = out.File
	file_live_proto_rawDesc = nil
	file_live_proto_goTypes = nil
	file_live_proto_depIdxs = nil
}
