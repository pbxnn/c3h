// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.19.1
// source: r401s/r401s.proto

package r401s

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

const OperationR401SConfirmReactorPerf = "/api.r401s.R401S/ConfirmReactorPerf"
const OperationR401SGetAPCControl = "/api.r401s.R401S/GetAPCControl"
const OperationR401SGetAll = "/api.r401s.R401S/GetAll"
const OperationR401SGetConfoundingVars = "/api.r401s.R401S/GetConfoundingVars"
const OperationR401SGetOperationVars = "/api.r401s.R401S/GetOperationVars"
const OperationR401SGetReactorPerformance = "/api.r401s.R401S/GetReactorPerformance"
const OperationR401SGetStatusVars = "/api.r401s.R401S/GetStatusVars"
const OperationR401SReset = "/api.r401s.R401S/Reset"
const OperationR401SSetAPCControl = "/api.r401s.R401S/SetAPCControl"
const OperationR401SSetControlSwitch = "/api.r401s.R401S/SetControlSwitch"

type R401SHTTPServer interface {
	ConfirmReactorPerf(context.Context, *ConfirmReactorPerfRequest) (*VarListReply, error)
	GetAPCControl(context.Context, *GetAPCControlRequest) (*VarListReply, error)
	GetAll(context.Context, *R401SGetAllRequest) (*R401SWsMessage, error)
	GetConfoundingVars(context.Context, *GetConfoundingVarsRequest) (*VarListReply, error)
	GetOperationVars(context.Context, *GetOperationVarsRequest) (*VarListReply, error)
	GetReactorPerformance(context.Context, *GetReactorPerformanceRequest) (*VarListReply, error)
	GetStatusVars(context.Context, *GetStatusVarsRequest) (*VarListReply, error)
	Reset(context.Context, *ResetRequest) (*VarReply, error)
	SetAPCControl(context.Context, *SetAPCControlRequest) (*VarReply, error)
	SetControlSwitch(context.Context, *SetControlSwitchRequest) (*VarReply, error)
}

func RegisterR401SHTTPServer(s *http.Server, srv R401SHTTPServer) {
	r := s.Route("/")
	r.GET("/c3h/r401s/apc-control", _R401S_GetAPCControl0_HTTP_Handler(srv))
	r.POST("/c3h/r401s/apc-control", _R401S_SetAPCControl0_HTTP_Handler(srv))
	r.POST("/c3h/r401s/control-switch", _R401S_SetControlSwitch0_HTTP_Handler(srv))
	r.POST("/c3h/r401s/reset", _R401S_Reset0_HTTP_Handler(srv))
	r.GET("/c3h/r401s/operation-vars", _R401S_GetOperationVars0_HTTP_Handler(srv))
	r.GET("/c3h/r401s/status-vars", _R401S_GetStatusVars0_HTTP_Handler(srv))
	r.GET("/c3h/r401s/confounding-vars", _R401S_GetConfoundingVars1_HTTP_Handler(srv))
	r.GET("/c3h/r401s/reactor-performance", _R401S_GetReactorPerformance0_HTTP_Handler(srv))
	r.POST("/c3h/r401s/reactor-perf", _R401S_ConfirmReactorPerf0_HTTP_Handler(srv))
	r.GET("/c3h/r401s/all", _R401S_GetAll0_HTTP_Handler(srv))
}

func _R401S_GetAPCControl0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAPCControlRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SGetAPCControl)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAPCControl(ctx, req.(*GetAPCControlRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarListReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_SetAPCControl0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetAPCControlRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SSetAPCControl)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetAPCControl(ctx, req.(*SetAPCControlRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_SetControlSwitch0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetControlSwitchRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SSetControlSwitch)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetControlSwitch(ctx, req.(*SetControlSwitchRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_Reset0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResetRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SReset)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Reset(ctx, req.(*ResetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_GetOperationVars0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOperationVarsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SGetOperationVars)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOperationVars(ctx, req.(*GetOperationVarsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarListReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_GetStatusVars0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetStatusVarsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SGetStatusVars)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetStatusVars(ctx, req.(*GetStatusVarsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarListReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_GetConfoundingVars1_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetConfoundingVarsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SGetConfoundingVars)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetConfoundingVars(ctx, req.(*GetConfoundingVarsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarListReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_GetReactorPerformance0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetReactorPerformanceRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SGetReactorPerformance)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetReactorPerformance(ctx, req.(*GetReactorPerformanceRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarListReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_ConfirmReactorPerf0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConfirmReactorPerfRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SConfirmReactorPerf)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ConfirmReactorPerf(ctx, req.(*ConfirmReactorPerfRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*VarListReply)
		return ctx.Result(200, reply)
	}
}

func _R401S_GetAll0_HTTP_Handler(srv R401SHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in R401SGetAllRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationR401SGetAll)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAll(ctx, req.(*R401SGetAllRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*R401SWsMessage)
		return ctx.Result(200, reply)
	}
}

type R401SHTTPClient interface {
	ConfirmReactorPerf(ctx context.Context, req *ConfirmReactorPerfRequest, opts ...http.CallOption) (rsp *VarListReply, err error)
	GetAPCControl(ctx context.Context, req *GetAPCControlRequest, opts ...http.CallOption) (rsp *VarListReply, err error)
	GetAll(ctx context.Context, req *R401SGetAllRequest, opts ...http.CallOption) (rsp *R401SWsMessage, err error)
	GetConfoundingVars(ctx context.Context, req *GetConfoundingVarsRequest, opts ...http.CallOption) (rsp *VarListReply, err error)
	GetOperationVars(ctx context.Context, req *GetOperationVarsRequest, opts ...http.CallOption) (rsp *VarListReply, err error)
	GetReactorPerformance(ctx context.Context, req *GetReactorPerformanceRequest, opts ...http.CallOption) (rsp *VarListReply, err error)
	GetStatusVars(ctx context.Context, req *GetStatusVarsRequest, opts ...http.CallOption) (rsp *VarListReply, err error)
	Reset(ctx context.Context, req *ResetRequest, opts ...http.CallOption) (rsp *VarReply, err error)
	SetAPCControl(ctx context.Context, req *SetAPCControlRequest, opts ...http.CallOption) (rsp *VarReply, err error)
	SetControlSwitch(ctx context.Context, req *SetControlSwitchRequest, opts ...http.CallOption) (rsp *VarReply, err error)
}

type R401SHTTPClientImpl struct {
	cc *http.Client
}

func NewR401SHTTPClient(client *http.Client) R401SHTTPClient {
	return &R401SHTTPClientImpl{client}
}

func (c *R401SHTTPClientImpl) ConfirmReactorPerf(ctx context.Context, in *ConfirmReactorPerfRequest, opts ...http.CallOption) (*VarListReply, error) {
	var out VarListReply
	pattern := "/c3h/r401s/reactor-perf"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationR401SConfirmReactorPerf))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) GetAPCControl(ctx context.Context, in *GetAPCControlRequest, opts ...http.CallOption) (*VarListReply, error) {
	var out VarListReply
	pattern := "/c3h/r401s/apc-control"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationR401SGetAPCControl))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) GetAll(ctx context.Context, in *R401SGetAllRequest, opts ...http.CallOption) (*R401SWsMessage, error) {
	var out R401SWsMessage
	pattern := "/c3h/r401s/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationR401SGetAll))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) GetConfoundingVars(ctx context.Context, in *GetConfoundingVarsRequest, opts ...http.CallOption) (*VarListReply, error) {
	var out VarListReply
	pattern := "/c3h/r401s/confounding-vars"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationR401SGetConfoundingVars))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) GetOperationVars(ctx context.Context, in *GetOperationVarsRequest, opts ...http.CallOption) (*VarListReply, error) {
	var out VarListReply
	pattern := "/c3h/r401s/operation-vars"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationR401SGetOperationVars))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) GetReactorPerformance(ctx context.Context, in *GetReactorPerformanceRequest, opts ...http.CallOption) (*VarListReply, error) {
	var out VarListReply
	pattern := "/c3h/r401s/reactor-performance"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationR401SGetReactorPerformance))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) GetStatusVars(ctx context.Context, in *GetStatusVarsRequest, opts ...http.CallOption) (*VarListReply, error) {
	var out VarListReply
	pattern := "/c3h/r401s/status-vars"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationR401SGetStatusVars))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) Reset(ctx context.Context, in *ResetRequest, opts ...http.CallOption) (*VarReply, error) {
	var out VarReply
	pattern := "/c3h/r401s/reset"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationR401SReset))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) SetAPCControl(ctx context.Context, in *SetAPCControlRequest, opts ...http.CallOption) (*VarReply, error) {
	var out VarReply
	pattern := "/c3h/r401s/apc-control"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationR401SSetAPCControl))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *R401SHTTPClientImpl) SetControlSwitch(ctx context.Context, in *SetControlSwitchRequest, opts ...http.CallOption) (*VarReply, error) {
	var out VarReply
	pattern := "/c3h/r401s/control-switch"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationR401SSetControlSwitch))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
