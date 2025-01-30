// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.3
// - protoc             v5.29.3
// source: v1/service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserServiceCreateAddresses = "/user.service.v1.UserService/CreateAddresses"
const OperationUserServiceDeleteAddresses = "/user.service.v1.UserService/DeleteAddresses"
const OperationUserServiceGetAddresses = "/user.service.v1.UserService/GetAddresses"
const OperationUserServiceGetUserInfo = "/user.service.v1.UserService/GetUserInfo"
const OperationUserServiceSignin = "/user.service.v1.UserService/Signin"
const OperationUserServiceUpdateAddresses = "/user.service.v1.UserService/UpdateAddresses"

type UserServiceHTTPServer interface {
	CreateAddresses(context.Context, *Address) (*Address, error)
	DeleteAddresses(context.Context, *DeleteAddressesRequest) (*DeleteAddressesReply, error)
	GetAddresses(context.Context, *GetAddressesRequest) (*GetAddressesReply, error)
	GetUserInfo(context.Context, *GetUserInfoRequest) (*GetUserInfoResponse, error)
	Signin(context.Context, *SigninRequest) (*SigninReply, error)
	UpdateAddresses(context.Context, *Address) (*Address, error)
}

func RegisterUserServiceHTTPServer(s *http.Server, srv UserServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user", _UserService_Signin0_HTTP_Handler(srv))
	r.GET("/v1/user/profile", _UserService_GetUserInfo0_HTTP_Handler(srv))
	r.POST("/v1/user/address", _UserService_CreateAddresses0_HTTP_Handler(srv))
	r.PATCH("/v1/user/address", _UserService_UpdateAddresses0_HTTP_Handler(srv))
	r.DELETE("/v1/user/address", _UserService_DeleteAddresses0_HTTP_Handler(srv))
	r.GET("/v1/user/address", _UserService_GetAddresses0_HTTP_Handler(srv))
}

func _UserService_Signin0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SigninRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceSignin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Signin(ctx, req.(*SigninRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SigninReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetUserInfo0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetUserInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserInfo(ctx, req.(*GetUserInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserInfoResponse)
		return ctx.Result(200, reply)
	}
}

func _UserService_CreateAddresses0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Address
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceCreateAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAddresses(ctx, req.(*Address))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Address)
		return ctx.Result(200, reply)
	}
}

func _UserService_UpdateAddresses0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Address
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceUpdateAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateAddresses(ctx, req.(*Address))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Address)
		return ctx.Result(200, reply)
	}
}

func _UserService_DeleteAddresses0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteAddressesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceDeleteAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAddresses(ctx, req.(*DeleteAddressesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteAddressesReply)
		return ctx.Result(200, reply)
	}
}

func _UserService_GetAddresses0_HTTP_Handler(srv UserServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAddressesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserServiceGetAddresses)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAddresses(ctx, req.(*GetAddressesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAddressesReply)
		return ctx.Result(200, reply)
	}
}

type UserServiceHTTPClient interface {
	CreateAddresses(ctx context.Context, req *Address, opts ...http.CallOption) (rsp *Address, err error)
	DeleteAddresses(ctx context.Context, req *DeleteAddressesRequest, opts ...http.CallOption) (rsp *DeleteAddressesReply, err error)
	GetAddresses(ctx context.Context, req *GetAddressesRequest, opts ...http.CallOption) (rsp *GetAddressesReply, err error)
	GetUserInfo(ctx context.Context, req *GetUserInfoRequest, opts ...http.CallOption) (rsp *GetUserInfoResponse, err error)
	Signin(ctx context.Context, req *SigninRequest, opts ...http.CallOption) (rsp *SigninReply, err error)
	UpdateAddresses(ctx context.Context, req *Address, opts ...http.CallOption) (rsp *Address, err error)
}

type UserServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewUserServiceHTTPClient(client *http.Client) UserServiceHTTPClient {
	return &UserServiceHTTPClientImpl{client}
}

func (c *UserServiceHTTPClientImpl) CreateAddresses(ctx context.Context, in *Address, opts ...http.CallOption) (*Address, error) {
	var out Address
	pattern := "/v1/user/address"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceCreateAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) DeleteAddresses(ctx context.Context, in *DeleteAddressesRequest, opts ...http.CallOption) (*DeleteAddressesReply, error) {
	var out DeleteAddressesReply
	pattern := "/v1/user/address"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceDeleteAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) GetAddresses(ctx context.Context, in *GetAddressesRequest, opts ...http.CallOption) (*GetAddressesReply, error) {
	var out GetAddressesReply
	pattern := "/v1/user/address"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...http.CallOption) (*GetUserInfoResponse, error) {
	var out GetUserInfoResponse
	pattern := "/v1/user/profile"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserServiceGetUserInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) Signin(ctx context.Context, in *SigninRequest, opts ...http.CallOption) (*SigninReply, error) {
	var out SigninReply
	pattern := "/v1/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceSignin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserServiceHTTPClientImpl) UpdateAddresses(ctx context.Context, in *Address, opts ...http.CallOption) (*Address, error) {
	var out Address
	pattern := "/v1/user/address"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserServiceUpdateAddresses))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
