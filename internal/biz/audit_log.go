package biz

import (
	"c3h/api/platform"
	"c3h/api/r401s"
	"c3h/internal/data/dao"
	"context"
	"encoding/json"
	"io"

	"go.opentelemetry.io/otel/trace"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// http请求相关操作类型（1-1000）
const (
	_ = iota

	OpR401SConfirmReactorPerformance
	OpR401SResetOperationVar
	OpR401SSetAPCControl
	OpR401SSetOperationControlSwitch
)

// cron job相关操作类型（1001-2000）
const (
	OperationCollectData = iota + 1001
)

// http请求相关资源类型（1-1000）
const (
	_ = iota

	ResourceR401SReactorPerformance
	ResourceR401SResetOperationVar
	ResourceR401SSetOperationControlSwitch
	ResourceR401SSetAPCControl
)

// cron job相关资源类型（1001-2000）
const (
	ResourceCollectData = iota + 1001
)

var operationMap = map[string]uint{

	r401s.OperationR401SConfirmReactorPerf: OpR401SConfirmReactorPerformance,
	r401s.OperationR401SReset:              OpR401SResetOperationVar,
	r401s.OperationR401SSetControlSwitch:   OpR401SSetOperationControlSwitch,
	r401s.OperationR401SSetAPCControl:      OpR401SSetAPCControl,

	platform.OperationPlatformCollectData: OperationCollectData,
}

var resourceMap = map[string]uint{

	r401s.OperationR401SConfirmReactorPerf: ResourceR401SReactorPerformance,
	r401s.OperationR401SReset:              ResourceR401SResetOperationVar,
	r401s.OperationR401SSetControlSwitch:   OpR401SSetOperationControlSwitch,
	r401s.OperationR401SSetAPCControl:      OpR401SSetAPCControl,

	platform.OperationPlatformCollectData: ResourceCollectData,
}

type AuditLogRepo interface {
	AddRecord(ctx context.Context, d *dao.AuditLog) error
	ListRecord(ctx context.Context, conds map[string]interface{}) ([]*dao.AuditLog, error)
}

type AuditLogUsecase struct {
	logger *log.Helper
	alr    AuditLogRepo
}

type logDetails struct {
	Request  json.RawMessage `json:"req,omitempty"`
	Response json.RawMessage `json:"resp,omitempty"`
	CronData interface{}     `json:"cronData,omitempty"`
	Error    error           `json:"error"`
	TraceID  string          `json:"trace_id"`
}

func NewAuditLogUsecase(alr AuditLogRepo, logger log.Logger) *AuditLogUsecase {

	return &AuditLogUsecase{
		alr:    alr,
		logger: log.NewHelper(log.With(logger, "module", "audit-log-usecase")),
	}
}

func (alu *AuditLogUsecase) AddRequestRecord(ctx context.Context, req, reply interface{}, replyErr error) {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return
	}

	info, ok := tr.(*http.Transport)
	if !ok {
		return
	}

	op := info.Operation()
	operation, ok := operationMap[op]
	if !ok {
		return
	}

	r := info.Request()
	reqBody, _ := io.ReadAll(r.Body)
	codec, _ := http.CodecForRequest(r, "Accept")
	respBody, _ := codec.Marshal(reply)

	data := logDetails{
		Request:  reqBody,
		Response: respBody,
		Error:    replyErr,
		TraceID:  trace.SpanContextFromContext(ctx).TraceID().String(),
	}

	details, err := codec.Marshal(data)
	if err != nil {
		alu.logger.Warn("marshal log detail err:%s", err.Error())
	}

	d := &dao.AuditLog{
		Operation: operation,
		Resource:  resourceMap[op],
		Details:   string(details),
		IsSuccess: 1,
	}

	if replyErr != nil {
		d.IsSuccess = 0
	}

	err = alu.alr.AddRecord(ctx, d)
	if err != nil {
		alu.logger.Warn("add record err:%s", err.Error())
	}
}

func (alu *AuditLogUsecase) AddCronRecord(ctx context.Context, op string, cronData interface{}, cronErr error) {
	operation, ok := operationMap[op]
	if !ok {
		return
	}

	resource, ok := resourceMap[op]
	if !ok {
		return
	}

	data := logDetails{
		CronData: cronData,
		Error:    cronErr,
		TraceID:  trace.SpanContextFromContext(ctx).TraceID().String(),
	}

	details, err := json.Marshal(data)
	if err != nil {
		alu.logger.Warn("marshal log detail err:%s", err.Error())
	}

	d := &dao.AuditLog{
		Operation: operation,
		Resource:  resource,
		Details:   string(details),
		IsSuccess: 1,
	}

	if cronErr != nil {
		d.IsSuccess = 0
	}

	err = alu.alr.AddRecord(ctx, d)
	if err != nil {
		alu.logger.Warn("add record err:%s", err.Error())
	}
}
