package service

import (
	"context"

	pb "c3h/api/r401s"
	"c3h/internal/biz"
	"c3h/internal/data/dao"
	"c3h/third_party/websocket"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	VarTypeSwitch = iota + 1
	VarTypeDouble
	VarTypeControlSwitch
)

// 模块
const (
	R401SSwitchModuleKey          = "R401S_switch_info"
	R401SOperationVarsModuleKey   = "R401S_operation_vars"
	R401SConfoundingVarsModuleKey = "R401S_confounding_vars"
	R401SStatusVarsModuleKey      = "R401S_status_vars"
	R401SReactorModuleKey         = "R401S_reactor_performance"
)

// 顶部开关
const (
	R401SMAPDControlMode = "R401S_MAPD_control_mode"
	R401SAPCMainSwitch   = "R401S_APC_main_switch"
	R401SAPCWriteSwitch  = "R401S_APC_write_switch"
	R401SCommunication   = "R401S_communication"
)

// 操作
const (
	R401SCycleMaterialsFlowKey   = "R401S_cycle_materials_flow"
	R401SHydrogenFlowKey         = "R401S_hydrogen_ratio"
	R401SIngressTemperatureKey   = "R401S_ingress_temperature"
	R401SHydrogenFlow            = "R401S_hydrogen_flow"
	R401SReactorPressure         = "R401S_reactor_pressure"
	R401SCycleMaterialsFlowCSKey = "R401S_cycle_materials_flow_control_switch"
	R401SHydrogenFlowCSKey       = "R401S_hydrogen_ratio_control_switch"
	R401SIngressTemperatureCSKey = "R401S_ingress_temperature_control_switch"
)

var editAbleVarMap = map[string]struct{}{
	R401SMAPDControlMode: {},
	R401SAPCMainSwitch:   {},
	R401SAPCWriteSwitch:  {},
}

var resetAbleVarMap = map[string]struct{}{
	R401SCycleMaterialsFlowKey: {},
	R401SHydrogenFlowKey:       {},
	R401SIngressTemperatureKey: {},
}

var controlSwitchMap = map[string]string{
	R401SCycleMaterialsFlowKey: R401SCycleMaterialsFlowCSKey,
	R401SHydrogenFlowKey:       R401SHydrogenFlowCSKey,
	R401SIngressTemperatureKey: R401SIngressTemperatureCSKey,
}

type R401SService struct {
	pb.UnimplementedR401SServer

	logger *log.Helper
	uc     *biz.R401SUsecase
	ws     *websocket.Server
}

func NewR401SService(uc *biz.R401SUsecase, logger log.Logger) *R401SService {
	return &R401SService{
		uc:     uc,
		logger: log.NewHelper(log.With(logger, "module", "R401S-service")),
	}
}

func (s *R401SService) GetAPCControl(ctx context.Context, req *pb.GetAPCControlRequest) (*pb.VarListReply, error) {
	list, err := s.uc.GetVarsByModule(ctx, R401SSwitchModuleKey)
	if err != nil {
		return nil, err
	}

	resp := &pb.VarListReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: s.formatVarDetailList(list),
		},
	}

	return resp, nil

	//resp := &pb.VarListReply{Code: 20000, Data: &pb.VarReplyData{}}
	//for _, item := range list {
	//	varDetail := &pb.VarDetail{
	//		Name:      item.Name,
	//		Key:       item.Key,
	//		Desc:      item.Desc,
	//		Type:      int32(item.Type),
	//		RealValue: item.RealValue,
	//		SetValue:  item.SetValue,
	//		Extra:     &pb.VarDetailExtra{},
	//	}
	//
	//	if _, ok := editAbleVarMap[item.Key]; ok {
	//		varDetail.Extra.EditAble = true
	//	}
	//
	//	resp.Data.List = append(resp.Data.List, varDetail)
	//}
	//
	//return resp, nil
}

func (s *R401SService) SetAPCControl(ctx context.Context, req *pb.SetAPCControlRequest) (*pb.VarReply, error) {

	item, err := s.uc.UpdateSwitchStatus(ctx, req.Key, req.SetValue)
	if err != nil {
		return nil, err
	}

	resp := &pb.VarReply{
		Code: 20000,
		Data: s.formatVarDetail(item, map[string]*dao.DataInfo{}),
	}

	return resp, nil
}

func (s *R401SService) SetControlSwitch(ctx context.Context, req *pb.SetControlSwitchRequest) (*pb.VarReply, error) {
	controlKey := controlSwitchMap[req.Key]
	list, err := s.uc.UpdateControlSwitchStatus(ctx, req.Key, controlKey, req.Status)
	if err != nil {
		return nil, err
	}

	resp := &pb.VarReply{
		Code: 20000,
		Data: s.formatVarDetailList(list)[0],
	}

	return resp, nil
}

func (s *R401SService) Reset(ctx context.Context, req *pb.ResetRequest) (*pb.VarReply, error) {
	item, err := s.uc.ResetVar(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	resp := &pb.VarReply{
		Code: 20000,
		Data: s.formatVarDetail(item, map[string]*dao.DataInfo{}),
	}

	return resp, nil
}

func (s *R401SService) GetOperationVars(ctx context.Context, req *pb.GetOperationVarsRequest) (*pb.VarListReply, error) {
	list, err := s.uc.GetVarsByModule(ctx, R401SOperationVarsModuleKey)
	if err != nil {
		return nil, err
	}

	resp := &pb.VarListReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: s.formatVarDetailList(list),
		},
	}

	return resp, nil

	//controlSwitchInfoMap := map[string]*dao.DataInfo{}
	//for idx, item := range list {
	//	if item.Type == VarTypeControlSwitch {
	//		controlSwitchInfoMap[item.Key] = list[idx]
	//		continue
	//	}
	//}
	//
	//for idx, item := range list {
	//	if item.Type == VarTypeControlSwitch {
	//		controlSwitchInfoMap[item.Key] = list[idx]
	//		continue
	//	}
	//
	//	varDetail := &pb.VarDetail{
	//		Name:      item.Name,
	//		Key:       item.Key,
	//		Desc:      item.Desc,
	//		Type:      int32(item.Type),
	//		RealValue: item.RealValue,
	//		SetValue:  item.SetValue,
	//		CalcValue: item.CalcValue,
	//		LowLimit:  item.LowLimit,
	//		HighLimit: item.HighLimit,
	//		Unit:      item.Unit,
	//		Extra:     &pb.VarDetailExtra{},
	//	}
	//
	//	if _, ok := resetAbleVarMap[item.Key]; ok {
	//		varDetail.Extra.ResetAble = true
	//	}
	//
	//	if k, ok := controlSwitchMap[item.Key]; ok {
	//		if v, ok := controlSwitchInfoMap[k]; ok {
	//			varDetail.Extra.HasControlSwitch = true
	//			varDetail.Extra.ControlSwitchStatus = int32(v.RealValue)
	//		}
	//	}
	//
	//	resp.Data.List = append(resp.Data.List, varDetail)
	//}
	//
	//return resp, nil
}

func (s *R401SService) GetConfoundingVars(ctx context.Context, req *pb.GetConfoundingVarsRequest) (*pb.VarListReply, error) {
	list, err := s.uc.GetVarsByModule(ctx, R401SConfoundingVarsModuleKey)
	if err != nil {
		return nil, err
	}

	resp := &pb.VarListReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: s.formatVarDetailList(list),
		},
	}

	return resp, nil

	//resp := &pb.VarListReply{Code: 20000, Data: &pb.VarReplyData{}}
	//for _, item := range list {
	//	varDetail := &pb.VarDetail{
	//		Name:      item.Name,
	//		Key:       item.Key,
	//		Desc:      item.Desc,
	//		Type:      int32(item.Type),
	//		RealValue: item.RealValue,
	//		SetValue:  item.SetValue,
	//		LowLimit:  item.LowLimit,
	//		HighLimit: item.HighLimit,
	//		Unit:      item.Unit,
	//	}
	//
	//	resp.Data.List = append(resp.Data.List, varDetail)
	//}
	//
	//return resp, nil
}

func (s *R401SService) GetStatusVars(ctx context.Context, req *pb.GetStatusVarsRequest) (*pb.VarListReply, error) {
	list, err := s.uc.GetVarsByModule(ctx, R401SStatusVarsModuleKey)
	if err != nil {
		return nil, err
	}

	resp := &pb.VarListReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: s.formatVarDetailList(list),
		},
	}

	return resp, nil

	//resp := &pb.VarListReply{Code: 20000, Data: &pb.VarReplyData{}}
	//for _, item := range list {
	//	varDetail := &pb.VarDetail{
	//		Name:      item.Name,
	//		Key:       item.Key,
	//		Desc:      item.Desc,
	//		Type:      int32(item.Type),
	//		RealValue: item.RealValue,
	//		SetValue:  item.SetValue,
	//		LowLimit:  item.LowLimit,
	//		HighLimit: item.HighLimit,
	//		Unit:      item.Unit,
	//	}
	//
	//	resp.Data.List = append(resp.Data.List, varDetail)
	//}
	//
	//return resp, nil
}

func (s *R401SService) GetReactorPerformance(ctx context.Context, req *pb.GetReactorPerformanceRequest) (*pb.VarListReply, error) {
	list, err := s.uc.GetVarsByModule(ctx, R401SReactorModuleKey)
	if err != nil {
		return nil, err
	}

	resp := &pb.VarListReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: s.formatVarDetailList(list),
		},
	}

	return resp, nil

	//resp := &pb.VarListReply{Code: 20000, Data: &pb.VarReplyData{}}
	//for _, item := range list {
	//	varDetail := &pb.VarDetail{
	//		Name:      item.Name,
	//		Key:       item.Key,
	//		Desc:      item.Desc,
	//		Type:      int32(item.Type),
	//		RealValue: item.RealValue,
	//		SetValue:  item.SetValue,
	//		LowLimit:  item.LowLimit,
	//		HighLimit: item.HighLimit,
	//		Unit:      item.Unit,
	//	}
	//
	//	resp.Data.List = append(resp.Data.List, varDetail)
	//}
	//
	//return resp, nil
}

func (s *R401SService) ConfirmReactorPerf(ctx context.Context, req *pb.ConfirmReactorPerfRequest) (*pb.VarListReply, error) {

	return &pb.VarListReply{
		Code: 20000,
	}, nil
}

func (s *R401SService) formatVarDetail(item *dao.DataInfo, controlSwitchInfoMap map[string]*dao.DataInfo) *pb.VarDetail {

	varDetail := &pb.VarDetail{
		Name:      item.Name,
		Key:       item.Key,
		Desc:      item.Desc,
		Type:      int32(item.Type),
		RealValue: item.RealValue,
		SetValue:  item.SetValue,
		CalcValue: item.CalcValue,
		LowLimit:  item.LowLimit,
		HighLimit: item.HighLimit,
		Unit:      item.Unit,
		Extra:     &pb.VarDetailExtra{},
	}

	if _, ok := resetAbleVarMap[item.Key]; ok {
		varDetail.Extra.ResetAble = true
	}

	if k, ok := controlSwitchMap[item.Key]; ok {
		if v, ok := controlSwitchInfoMap[k]; ok {
			varDetail.Extra.HasControlSwitch = true
			varDetail.Extra.ControlSwitchStatus = int32(v.RealValue)
		}
	}

	return varDetail
}

func (s *R401SService) formatVarDetailList(list []*dao.DataInfo) []*pb.VarDetail {

	var res []*pb.VarDetail
	controlSwitchInfoMap := map[string]*dao.DataInfo{}
	for idx, item := range list {
		if item.Type == VarTypeControlSwitch {
			controlSwitchInfoMap[item.Key] = list[idx]
			continue
		}
	}

	for idx, item := range list {
		if item.Type == VarTypeControlSwitch {
			continue
		}

		varDetail := s.formatVarDetail(list[idx], controlSwitchInfoMap)
		res = append(res, varDetail)
	}
	return res
}

func (s *R401SService) GetAll(ctx context.Context, req *pb.R401SGetAllRequest) (*pb.R401SWsMessage, error) {
	return s.getWsMessage(ctx)
}
