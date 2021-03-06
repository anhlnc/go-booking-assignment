// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: booking.proto

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

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	FlightId   string `protobuf:"bytes,3,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty"`
	Code       string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Status     string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	BookDate   string `protobuf:"bytes,6,opt,name=book_date,json=bookDate,proto3" json:"book_date,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Booking) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *Booking) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

func (x *Booking) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Booking) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Booking) GetBookDate() string {
	if x != nil {
		return x.BookDate
	}
	return ""
}

type GetBookingByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookingByIDRequest) Reset() {
	*x = GetBookingByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookingByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingByIDRequest) ProtoMessage() {}

func (x *GetBookingByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingByIDRequest.ProtoReflect.Descriptor instead.
func (*GetBookingByIDRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{1}
}

func (x *GetBookingByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AssignBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	FlightId   string `protobuf:"bytes,2,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty"`
	Status     string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Code       string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	BookDate   string `protobuf:"bytes,5,opt,name=book_date,json=bookDate,proto3" json:"book_date,omitempty"`
}

func (x *AssignBookingRequest) Reset() {
	*x = AssignBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignBookingRequest) ProtoMessage() {}

func (x *AssignBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignBookingRequest.ProtoReflect.Descriptor instead.
func (*AssignBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{2}
}

func (x *AssignBookingRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *AssignBookingRequest) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

func (x *AssignBookingRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AssignBookingRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AssignBookingRequest) GetBookDate() string {
	if x != nil {
		return x.BookDate
	}
	return ""
}

type ViewBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ViewBookingRequest) Reset() {
	*x = ViewBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewBookingRequest) ProtoMessage() {}

func (x *ViewBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewBookingRequest.ProtoReflect.Descriptor instead.
func (*ViewBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{3}
}

func (x *ViewBookingRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ViewBookingRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type ViewBookingReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string    `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	BookDate string    `protobuf:"bytes,2,opt,name=book_date,json=bookDate,proto3" json:"book_date,omitempty"`
	Customer *Customer `protobuf:"bytes,3,opt,name=customer,proto3" json:"customer,omitempty"`
	Flight   *Flight   `protobuf:"bytes,4,opt,name=flight,proto3" json:"flight,omitempty"`
}

func (x *ViewBookingReponse) Reset() {
	*x = ViewBookingReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewBookingReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewBookingReponse) ProtoMessage() {}

func (x *ViewBookingReponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewBookingReponse.ProtoReflect.Descriptor instead.
func (*ViewBookingReponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{4}
}

func (x *ViewBookingReponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ViewBookingReponse) GetBookDate() string {
	if x != nil {
		return x.BookDate
	}
	return ""
}

func (x *ViewBookingReponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *ViewBookingReponse) GetFlight() *Flight {
	if x != nil {
		return x.Flight
	}
	return nil
}

type ListBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ListBookingRequest) Reset() {
	*x = ListBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookingRequest) ProtoMessage() {}

func (x *ListBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookingRequest.ProtoReflect.Descriptor instead.
func (*ListBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{5}
}

func (x *ListBookingRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type ListBookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*Booking `protobuf:"bytes,1,rep,name=Bookings,proto3" json:"Bookings,omitempty"`
}

func (x *ListBookingResponse) Reset() {
	*x = ListBookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBookingResponse) ProtoMessage() {}

func (x *ListBookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBookingResponse.ProtoReflect.Descriptor instead.
func (*ListBookingResponse) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{6}
}

func (x *ListBookingResponse) GetBookings() []*Booking {
	if x != nil {
		return x.Bookings
	}
	return nil
}

type DeleteBookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteBookingRequest) Reset() {
	*x = DeleteBookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBookingRequest) ProtoMessage() {}

func (x *DeleteBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBookingRequest.ProtoReflect.Descriptor instead.
func (*DeleteBookingRequest) Descriptor() ([]byte, []int) {
	return file_booking_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteBookingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_booking_proto protoreflect.FileDescriptor

var file_booking_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f,
	0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f,
	0x6f, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x9d, 0x01, 0x0a, 0x14, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x49, 0x0a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x12, 0x56,
	0x69, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x12, 0x27, 0x0a, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x35, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x99,
	0x03, 0x0a, 0x08, 0x42, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x40, 0x0a, 0x0d, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x33, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x10,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x48, 0x0a, 0x0c, 0x56, 0x69, 0x65, 0x77, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x56, 0x69, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x69,
	0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x1d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x48, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12,
	0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_proto_rawDescOnce sync.Once
	file_booking_proto_rawDescData = file_booking_proto_rawDesc
)

func file_booking_proto_rawDescGZIP() []byte {
	file_booking_proto_rawDescOnce.Do(func() {
		file_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_proto_rawDescData)
	})
	return file_booking_proto_rawDescData
}

var file_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_booking_proto_goTypes = []interface{}{
	(*Booking)(nil),               // 0: booking.Booking
	(*GetBookingByIDRequest)(nil), // 1: booking.GetBookingByIDRequest
	(*AssignBookingRequest)(nil),  // 2: booking.AssignBookingRequest
	(*ViewBookingRequest)(nil),    // 3: booking.ViewBookingRequest
	(*ViewBookingReponse)(nil),    // 4: booking.ViewBookingReponse
	(*ListBookingRequest)(nil),    // 5: booking.ListBookingRequest
	(*ListBookingResponse)(nil),   // 6: booking.ListBookingResponse
	(*DeleteBookingRequest)(nil),  // 7: booking.DeleteBookingRequest
	(*Customer)(nil),              // 8: booking.Customer
	(*Flight)(nil),                // 9: booking.Flight
	(*Empty)(nil),                 // 10: booking.Empty
}
var file_booking_proto_depIdxs = []int32{
	8,  // 0: booking.ViewBookingReponse.customer:type_name -> booking.Customer
	9,  // 1: booking.ViewBookingReponse.flight:type_name -> booking.Flight
	0,  // 2: booking.ListBookingResponse.Bookings:type_name -> booking.Booking
	2,  // 3: booking.Bbooking.AssignBooking:input_type -> booking.AssignBookingRequest
	0,  // 4: booking.Bbooking.UpdateBooking:input_type -> booking.Booking
	1,  // 5: booking.Bbooking.GetBookingByID:input_type -> booking.GetBookingByIDRequest
	3,  // 6: booking.Bbooking.ViewBookings:input_type -> booking.ViewBookingRequest
	7,  // 7: booking.Bbooking.DeleteBooking:input_type -> booking.DeleteBookingRequest
	5,  // 8: booking.Bbooking.ListBooking:input_type -> booking.ListBookingRequest
	0,  // 9: booking.Bbooking.AssignBooking:output_type -> booking.Booking
	0,  // 10: booking.Bbooking.UpdateBooking:output_type -> booking.Booking
	0,  // 11: booking.Bbooking.GetBookingByID:output_type -> booking.Booking
	4,  // 12: booking.Bbooking.ViewBookings:output_type -> booking.ViewBookingReponse
	10, // 13: booking.Bbooking.DeleteBooking:output_type -> booking.Empty
	6,  // 14: booking.Bbooking.ListBooking:output_type -> booking.ListBookingResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_booking_proto_init() }
func file_booking_proto_init() {
	if File_booking_proto != nil {
		return
	}
	file_empty_proto_init()
	file_flight_proto_init()
	file_customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookingByIDRequest); i {
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
		file_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignBookingRequest); i {
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
		file_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewBookingRequest); i {
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
		file_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewBookingReponse); i {
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
		file_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookingRequest); i {
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
		file_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBookingResponse); i {
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
		file_booking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBookingRequest); i {
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
			RawDescriptor: file_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_proto_goTypes,
		DependencyIndexes: file_booking_proto_depIdxs,
		MessageInfos:      file_booking_proto_msgTypes,
	}.Build()
	File_booking_proto = out.File
	file_booking_proto_rawDesc = nil
	file_booking_proto_goTypes = nil
	file_booking_proto_depIdxs = nil
}
