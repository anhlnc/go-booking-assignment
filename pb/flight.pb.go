// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: flight.proto

package pb

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

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	From          string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To            string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Date          string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	Status        string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	AvailableSlot int64  `protobuf:"varint,7,opt,name=available_slot,json=availableSlot,proto3" json:"available_slot,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{0}
}

func (x *Flight) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Flight) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flight) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Flight) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Flight) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Flight) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Flight) GetAvailableSlot() int64 {
	if x != nil {
		return x.AvailableSlot
	}
	return 0
}

type GetFlightIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFlightIDRequest) Reset() {
	*x = GetFlightIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlightIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlightIDRequest) ProtoMessage() {}

func (x *GetFlightIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlightIDRequest.ProtoReflect.Descriptor instead.
func (*GetFlightIDRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{1}
}

func (x *GetFlightIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Date string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *FindFlightRequest) Reset() {
	*x = FindFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindFlightRequest) ProtoMessage() {}

func (x *FindFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindFlightRequest.ProtoReflect.Descriptor instead.
func (*FindFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{2}
}

func (x *FindFlightRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FindFlightRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *FindFlightRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *FindFlightRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type ListFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListFlightRequest) Reset() {
	*x = ListFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlightRequest) ProtoMessage() {}

func (x *ListFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlightRequest.ProtoReflect.Descriptor instead.
func (*ListFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{3}
}

func (x *ListFlightRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListFlightResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flights []*Flight `protobuf:"bytes,1,rep,name=Flights,proto3" json:"Flights,omitempty"`
}

func (x *ListFlightResponse) Reset() {
	*x = ListFlightResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFlightResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFlightResponse) ProtoMessage() {}

func (x *ListFlightResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFlightResponse.ProtoReflect.Descriptor instead.
func (*ListFlightResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{4}
}

func (x *ListFlightResponse) GetFlights() []*Flight {
	if x != nil {
		return x.Flights
	}
	return nil
}

type DeleteFlightRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFlightRequest) Reset() {
	*x = DeleteFlightRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFlightRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFlightRequest) ProtoMessage() {}

func (x *DeleteFlightRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFlightRequest.ProtoReflect.Descriptor instead.
func (*DeleteFlightRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFlightRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_flight_proto protoreflect.FileDescriptor

var file_flight_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x5f, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x22, 0x27, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3f, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x07, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x07, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xf0, 0x02, 0x0a, 0x0d, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x46,
	0x6c, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x30, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x1a, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x49, 0x44, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x3c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flight_proto_rawDescOnce sync.Once
	file_flight_proto_rawDescData = file_flight_proto_rawDesc
)

func file_flight_proto_rawDescGZIP() []byte {
	file_flight_proto_rawDescOnce.Do(func() {
		file_flight_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_proto_rawDescData)
	})
	return file_flight_proto_rawDescData
}

var file_flight_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_flight_proto_goTypes = []interface{}{
	(*Flight)(nil),              // 0: booking.Flight
	(*GetFlightIDRequest)(nil),  // 1: booking.GetFlightIDRequest
	(*FindFlightRequest)(nil),   // 2: booking.FindFlightRequest
	(*ListFlightRequest)(nil),   // 3: booking.ListFlightRequest
	(*ListFlightResponse)(nil),  // 4: booking.ListFlightResponse
	(*DeleteFlightRequest)(nil), // 5: booking.DeleteFlightRequest
	(*Empty)(nil),               // 6: booking.Empty
}
var file_flight_proto_depIdxs = []int32{
	0, // 0: booking.ListFlightResponse.Flights:type_name -> booking.Flight
	0, // 1: booking.BookingFlight.CreateFlight:input_type -> booking.Flight
	0, // 2: booking.BookingFlight.UpdateFlight:input_type -> booking.Flight
	2, // 3: booking.BookingFlight.FindFlight:input_type -> booking.FindFlightRequest
	1, // 4: booking.BookingFlight.GetFlightID:input_type -> booking.GetFlightIDRequest
	5, // 5: booking.BookingFlight.DeleteFlight:input_type -> booking.DeleteFlightRequest
	3, // 6: booking.BookingFlight.ListFlight:input_type -> booking.ListFlightRequest
	0, // 7: booking.BookingFlight.CreateFlight:output_type -> booking.Flight
	0, // 8: booking.BookingFlight.UpdateFlight:output_type -> booking.Flight
	0, // 9: booking.BookingFlight.FindFlight:output_type -> booking.Flight
	0, // 10: booking.BookingFlight.GetFlightID:output_type -> booking.Flight
	6, // 11: booking.BookingFlight.DeleteFlight:output_type -> booking.Empty
	4, // 12: booking.BookingFlight.ListFlight:output_type -> booking.ListFlightResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_flight_proto_init() }
func file_flight_proto_init() {
	if File_flight_proto != nil {
		return
	}
	file_empty_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
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
		file_flight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlightIDRequest); i {
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
		file_flight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindFlightRequest); i {
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
		file_flight_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlightRequest); i {
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
		file_flight_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFlightResponse); i {
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
		file_flight_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFlightRequest); i {
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
			RawDescriptor: file_flight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flight_proto_goTypes,
		DependencyIndexes: file_flight_proto_depIdxs,
		MessageInfos:      file_flight_proto_msgTypes,
	}.Build()
	File_flight_proto = out.File
	file_flight_proto_rawDesc = nil
	file_flight_proto_goTypes = nil
	file_flight_proto_depIdxs = nil
}