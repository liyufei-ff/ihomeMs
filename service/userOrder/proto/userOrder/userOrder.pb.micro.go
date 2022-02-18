// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/userOrder/userOrder.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserOrder service

type UserOrderService interface {
	CreateOrder(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetOrderInfo(ctx context.Context, in *GetReq, opts ...client.CallOption) (*GetResp, error)
	UpdateStatus(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*UpdateResp, error)
	UpdateOrderComment(ctx context.Context, in *OrderCommentReq, opts ...client.CallOption) (*OrderCommentResp, error)
}

type userOrderService struct {
	c    client.Client
	name string
}

func NewUserOrderService(name string, c client.Client) UserOrderService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.userOrder"
	}
	return &userOrderService{
		c:    c,
		name: name,
	}
}

func (c *userOrderService) CreateOrder(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "UserOrder.CreateOrder", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) GetOrderInfo(ctx context.Context, in *GetReq, opts ...client.CallOption) (*GetResp, error) {
	req := c.c.NewRequest(c.name, "UserOrder.GetOrderInfo", in)
	out := new(GetResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) UpdateStatus(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*UpdateResp, error) {
	req := c.c.NewRequest(c.name, "UserOrder.UpdateStatus", in)
	out := new(UpdateResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userOrderService) UpdateOrderComment(ctx context.Context, in *OrderCommentReq, opts ...client.CallOption) (*OrderCommentResp, error) {
	req := c.c.NewRequest(c.name, "UserOrder.UpdateOrderComment", in)
	out := new(OrderCommentResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserOrder service

type UserOrderHandler interface {
	CreateOrder(context.Context, *Request, *Response) error
	GetOrderInfo(context.Context, *GetReq, *GetResp) error
	UpdateStatus(context.Context, *UpdateReq, *UpdateResp) error
	UpdateOrderComment(context.Context, *OrderCommentReq, *OrderCommentResp) error
}

func RegisterUserOrderHandler(s server.Server, hdlr UserOrderHandler, opts ...server.HandlerOption) error {
	type userOrder interface {
		CreateOrder(ctx context.Context, in *Request, out *Response) error
		GetOrderInfo(ctx context.Context, in *GetReq, out *GetResp) error
		UpdateStatus(ctx context.Context, in *UpdateReq, out *UpdateResp) error
		UpdateOrderComment(ctx context.Context, in *OrderCommentReq, out *OrderCommentResp) error
	}
	type UserOrder struct {
		userOrder
	}
	h := &userOrderHandler{hdlr}
	return s.Handle(s.NewHandler(&UserOrder{h}, opts...))
}

type userOrderHandler struct {
	UserOrderHandler
}

func (h *userOrderHandler) CreateOrder(ctx context.Context, in *Request, out *Response) error {
	return h.UserOrderHandler.CreateOrder(ctx, in, out)
}

func (h *userOrderHandler) GetOrderInfo(ctx context.Context, in *GetReq, out *GetResp) error {
	return h.UserOrderHandler.GetOrderInfo(ctx, in, out)
}

func (h *userOrderHandler) UpdateStatus(ctx context.Context, in *UpdateReq, out *UpdateResp) error {
	return h.UserOrderHandler.UpdateStatus(ctx, in, out)
}

func (h *userOrderHandler) UpdateOrderComment(ctx context.Context, in *OrderCommentReq, out *OrderCommentResp) error {
	return h.UserOrderHandler.UpdateOrderComment(ctx, in, out)
}
