// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BbookingClient is the client API for Bbooking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BbookingClient interface {
	AssignBooking(ctx context.Context, in *AssignBookingRequest, opts ...grpc.CallOption) (*Booking, error)
	UpdateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error)
	GetBookingByID(ctx context.Context, in *GetBookingByIDRequest, opts ...grpc.CallOption) (*Booking, error)
	ViewBookings(ctx context.Context, in *ViewBookingRequest, opts ...grpc.CallOption) (*ViewBookingReponse, error)
	DeleteBooking(ctx context.Context, in *DeleteBookingRequest, opts ...grpc.CallOption) (*Empty, error)
	ListBooking(ctx context.Context, in *ListBookingRequest, opts ...grpc.CallOption) (*ListBookingResponse, error)
}

type bbookingClient struct {
	cc grpc.ClientConnInterface
}

func NewBbookingClient(cc grpc.ClientConnInterface) BbookingClient {
	return &bbookingClient{cc}
}

func (c *bbookingClient) AssignBooking(ctx context.Context, in *AssignBookingRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/booking.Bbooking/AssignBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bbookingClient) UpdateBooking(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/booking.Bbooking/UpdateBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bbookingClient) GetBookingByID(ctx context.Context, in *GetBookingByIDRequest, opts ...grpc.CallOption) (*Booking, error) {
	out := new(Booking)
	err := c.cc.Invoke(ctx, "/booking.Bbooking/GetBookingByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bbookingClient) ViewBookings(ctx context.Context, in *ViewBookingRequest, opts ...grpc.CallOption) (*ViewBookingReponse, error) {
	out := new(ViewBookingReponse)
	err := c.cc.Invoke(ctx, "/booking.Bbooking/ViewBookings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bbookingClient) DeleteBooking(ctx context.Context, in *DeleteBookingRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/booking.Bbooking/DeleteBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bbookingClient) ListBooking(ctx context.Context, in *ListBookingRequest, opts ...grpc.CallOption) (*ListBookingResponse, error) {
	out := new(ListBookingResponse)
	err := c.cc.Invoke(ctx, "/booking.Bbooking/ListBooking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BbookingServer is the server API for Bbooking service.
// All implementations must embed UnimplementedBbookingServer
// for forward compatibility
type BbookingServer interface {
	AssignBooking(context.Context, *AssignBookingRequest) (*Booking, error)
	UpdateBooking(context.Context, *Booking) (*Booking, error)
	GetBookingByID(context.Context, *GetBookingByIDRequest) (*Booking, error)
	ViewBookings(context.Context, *ViewBookingRequest) (*ViewBookingReponse, error)
	DeleteBooking(context.Context, *DeleteBookingRequest) (*Empty, error)
	ListBooking(context.Context, *ListBookingRequest) (*ListBookingResponse, error)
	mustEmbedUnimplementedBbookingServer()
}

// UnimplementedBbookingServer must be embedded to have forward compatible implementations.
type UnimplementedBbookingServer struct {
}

func (UnimplementedBbookingServer) AssignBooking(context.Context, *AssignBookingRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignBooking not implemented")
}
func (UnimplementedBbookingServer) UpdateBooking(context.Context, *Booking) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBooking not implemented")
}
func (UnimplementedBbookingServer) GetBookingByID(context.Context, *GetBookingByIDRequest) (*Booking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookingByID not implemented")
}
func (UnimplementedBbookingServer) ViewBookings(context.Context, *ViewBookingRequest) (*ViewBookingReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewBookings not implemented")
}
func (UnimplementedBbookingServer) DeleteBooking(context.Context, *DeleteBookingRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBooking not implemented")
}
func (UnimplementedBbookingServer) ListBooking(context.Context, *ListBookingRequest) (*ListBookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBooking not implemented")
}
func (UnimplementedBbookingServer) mustEmbedUnimplementedBbookingServer() {}

// UnsafeBbookingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BbookingServer will
// result in compilation errors.
type UnsafeBbookingServer interface {
	mustEmbedUnimplementedBbookingServer()
}

func RegisterBbookingServer(s grpc.ServiceRegistrar, srv BbookingServer) {
	s.RegisterService(&Bbooking_ServiceDesc, srv)
}

func _Bbooking_AssignBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BbookingServer).AssignBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bbooking/AssignBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BbookingServer).AssignBooking(ctx, req.(*AssignBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bbooking_UpdateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BbookingServer).UpdateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bbooking/UpdateBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BbookingServer).UpdateBooking(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bbooking_GetBookingByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookingByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BbookingServer).GetBookingByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bbooking/GetBookingByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BbookingServer).GetBookingByID(ctx, req.(*GetBookingByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bbooking_ViewBookings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BbookingServer).ViewBookings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bbooking/ViewBookings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BbookingServer).ViewBookings(ctx, req.(*ViewBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bbooking_DeleteBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BbookingServer).DeleteBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bbooking/DeleteBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BbookingServer).DeleteBooking(ctx, req.(*DeleteBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bbooking_ListBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BbookingServer).ListBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/booking.Bbooking/ListBooking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BbookingServer).ListBooking(ctx, req.(*ListBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Bbooking_ServiceDesc is the grpc.ServiceDesc for Bbooking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bbooking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.Bbooking",
	HandlerType: (*BbookingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignBooking",
			Handler:    _Bbooking_AssignBooking_Handler,
		},
		{
			MethodName: "UpdateBooking",
			Handler:    _Bbooking_UpdateBooking_Handler,
		},
		{
			MethodName: "GetBookingByID",
			Handler:    _Bbooking_GetBookingByID_Handler,
		},
		{
			MethodName: "ViewBookings",
			Handler:    _Bbooking_ViewBookings_Handler,
		},
		{
			MethodName: "DeleteBooking",
			Handler:    _Bbooking_DeleteBooking_Handler,
		},
		{
			MethodName: "ListBooking",
			Handler:    _Bbooking_ListBooking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking.proto",
}