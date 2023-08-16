// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.19.1
// source: control_net/control_net.proto

package control_net

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

const OperationControlNetConfirmReactorPerf = "/api.control_net.ControlNet/ConfirmReactorPerf"
const OperationControlNetGetOperationVars = "/api.control_net.ControlNet/GetOperationVars"
const OperationControlNetGetR401APC = "/api.control_net.ControlNet/GetR401APC"
const OperationControlNetGetReactorPerf = "/api.control_net.ControlNet/GetReactorPerf"
const OperationControlNetGetReactorPerformance = "/api.control_net.ControlNet/GetReactorPerformance"
const OperationControlNetSetOperationVars = "/api.control_net.ControlNet/SetOperationVars"
const OperationControlNetSetR401APC = "/api.control_net.ControlNet/SetR401APC"

type ControlNetHTTPServer interface {
	ConfirmReactorPerf(context.Context, *ConfirmReactorPerfRequest) (*ConfirmReactorPerfReply, error)
	GetOperationVars(context.Context, *GetOperationVarsRequest) (*GetOperationVarsReply, error)
	GetR401APC(context.Context, *GetR401Request) (*GetR401Reply, error)
	GetReactorPerf(context.Context, *GetReactorPerfRequest) (*GetReactorPerfReply, error)
	GetReactorPerformance(context.Context, *GetReactorPerformanceRequest) (*GetReactorPerformanceReply, error)
	SetOperationVars(context.Context, *SetOperationVarsRequest) (*SetOperationVarsReply, error)
	SetR401APC(context.Context, *SetR401APCRequest) (*SetR401APCReply, error)
}

func RegisterControlNetHTTPServer(s *http.Server, srv ControlNetHTTPServer) {
	r := s.Route("/")
	r.GET("/c3h/control-net/r401-apc", _ControlNet_GetR401APC0_HTTP_Handler(srv))
	r.GET("/c3h/control-net/operation-vars", _ControlNet_GetOperationVars0_HTTP_Handler(srv))
	r.GET("/c3h/control-net/reactor-performance", _ControlNet_GetReactorPerformance0_HTTP_Handler(srv))
	r.POST("/c3h/control-net/r401-apc", _ControlNet_SetR401APC0_HTTP_Handler(srv))
	r.POST("/c3h/control-net/operation-vars", _ControlNet_SetOperationVars0_HTTP_Handler(srv))
	r.POST("/c3h/control-net/reactor-perf", _ControlNet_ConfirmReactorPerf0_HTTP_Handler(srv))
	r.GET("/c3h/control-net/reactor-perf", _ControlNet_GetReactorPerf0_HTTP_Handler(srv))
}

func _ControlNet_GetR401APC0_HTTP_Handler(srv ControlNetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetR401Request
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlNetGetR401APC)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetR401APC(ctx, req.(*GetR401Request))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetR401Reply)
		return ctx.Result(200, reply)
	}
}

func _ControlNet_GetOperationVars0_HTTP_Handler(srv ControlNetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOperationVarsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlNetGetOperationVars)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOperationVars(ctx, req.(*GetOperationVarsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOperationVarsReply)
		return ctx.Result(200, reply)
	}
}

func _ControlNet_GetReactorPerformance0_HTTP_Handler(srv ControlNetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetReactorPerformanceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlNetGetReactorPerformance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetReactorPerformance(ctx, req.(*GetReactorPerformanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetReactorPerformanceReply)
		return ctx.Result(200, reply)
	}
}

func _ControlNet_SetR401APC0_HTTP_Handler(srv ControlNetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetR401APCRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlNetSetR401APC)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetR401APC(ctx, req.(*SetR401APCRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetR401APCReply)
		return ctx.Result(200, reply)
	}
}

func _ControlNet_SetOperationVars0_HTTP_Handler(srv ControlNetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetOperationVarsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlNetSetOperationVars)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetOperationVars(ctx, req.(*SetOperationVarsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetOperationVarsReply)
		return ctx.Result(200, reply)
	}
}

func _ControlNet_ConfirmReactorPerf0_HTTP_Handler(srv ControlNetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConfirmReactorPerfRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlNetConfirmReactorPerf)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ConfirmReactorPerf(ctx, req.(*ConfirmReactorPerfRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ConfirmReactorPerfReply)
		return ctx.Result(200, reply)
	}
}

func _ControlNet_GetReactorPerf0_HTTP_Handler(srv ControlNetHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetReactorPerfRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationControlNetGetReactorPerf)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetReactorPerf(ctx, req.(*GetReactorPerfRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetReactorPerfReply)
		return ctx.Result(200, reply)
	}
}

type ControlNetHTTPClient interface {
	ConfirmReactorPerf(ctx context.Context, req *ConfirmReactorPerfRequest, opts ...http.CallOption) (rsp *ConfirmReactorPerfReply, err error)
	GetOperationVars(ctx context.Context, req *GetOperationVarsRequest, opts ...http.CallOption) (rsp *GetOperationVarsReply, err error)
	GetR401APC(ctx context.Context, req *GetR401Request, opts ...http.CallOption) (rsp *GetR401Reply, err error)
	GetReactorPerf(ctx context.Context, req *GetReactorPerfRequest, opts ...http.CallOption) (rsp *GetReactorPerfReply, err error)
	GetReactorPerformance(ctx context.Context, req *GetReactorPerformanceRequest, opts ...http.CallOption) (rsp *GetReactorPerformanceReply, err error)
	SetOperationVars(ctx context.Context, req *SetOperationVarsRequest, opts ...http.CallOption) (rsp *SetOperationVarsReply, err error)
	SetR401APC(ctx context.Context, req *SetR401APCRequest, opts ...http.CallOption) (rsp *SetR401APCReply, err error)
}

type ControlNetHTTPClientImpl struct {
	cc *http.Client
}

func NewControlNetHTTPClient(client *http.Client) ControlNetHTTPClient {
	return &ControlNetHTTPClientImpl{client}
}

func (c *ControlNetHTTPClientImpl) ConfirmReactorPerf(ctx context.Context, in *ConfirmReactorPerfRequest, opts ...http.CallOption) (*ConfirmReactorPerfReply, error) {
	var out ConfirmReactorPerfReply
	pattern := "/c3h/control-net/reactor-perf"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationControlNetConfirmReactorPerf))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ControlNetHTTPClientImpl) GetOperationVars(ctx context.Context, in *GetOperationVarsRequest, opts ...http.CallOption) (*GetOperationVarsReply, error) {
	var out GetOperationVarsReply
	pattern := "/c3h/control-net/operation-vars"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlNetGetOperationVars))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ControlNetHTTPClientImpl) GetR401APC(ctx context.Context, in *GetR401Request, opts ...http.CallOption) (*GetR401Reply, error) {
	var out GetR401Reply
	pattern := "/c3h/control-net/r401-apc"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlNetGetR401APC))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ControlNetHTTPClientImpl) GetReactorPerf(ctx context.Context, in *GetReactorPerfRequest, opts ...http.CallOption) (*GetReactorPerfReply, error) {
	var out GetReactorPerfReply
	pattern := "/c3h/control-net/reactor-perf"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlNetGetReactorPerf))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ControlNetHTTPClientImpl) GetReactorPerformance(ctx context.Context, in *GetReactorPerformanceRequest, opts ...http.CallOption) (*GetReactorPerformanceReply, error) {
	var out GetReactorPerformanceReply
	pattern := "/c3h/control-net/reactor-performance"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationControlNetGetReactorPerformance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ControlNetHTTPClientImpl) SetOperationVars(ctx context.Context, in *SetOperationVarsRequest, opts ...http.CallOption) (*SetOperationVarsReply, error) {
	var out SetOperationVarsReply
	pattern := "/c3h/control-net/operation-vars"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationControlNetSetOperationVars))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ControlNetHTTPClientImpl) SetR401APC(ctx context.Context, in *SetR401APCRequest, opts ...http.CallOption) (*SetR401APCReply, error) {
	var out SetR401APCReply
	pattern := "/c3h/control-net/r401-apc"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationControlNetSetR401APC))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}