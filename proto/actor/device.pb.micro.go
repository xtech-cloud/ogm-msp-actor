// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/actor/device.proto

package actor

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Device service

func NewDeviceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Device service

type DeviceService interface {
	// 列举
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*DomainListResponse, error)
}

type deviceService struct {
	c    client.Client
	name string
}

func NewDeviceService(name string, c client.Client) DeviceService {
	return &deviceService{
		c:    c,
		name: name,
	}
}

func (c *deviceService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*DomainListResponse, error) {
	req := c.c.NewRequest(c.name, "Device.List", in)
	out := new(DomainListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Device service

type DeviceHandler interface {
	// 列举
	List(context.Context, *ListRequest, *DomainListResponse) error
}

func RegisterDeviceHandler(s server.Server, hdlr DeviceHandler, opts ...server.HandlerOption) error {
	type device interface {
		List(ctx context.Context, in *ListRequest, out *DomainListResponse) error
	}
	type Device struct {
		device
	}
	h := &deviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Device{h}, opts...))
}

type deviceHandler struct {
	DeviceHandler
}

func (h *deviceHandler) List(ctx context.Context, in *ListRequest, out *DomainListResponse) error {
	return h.DeviceHandler.List(ctx, in, out)
}
