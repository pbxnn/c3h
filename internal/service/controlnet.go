package service

import (
	"context"

	pb "c3h/api/control_net"
)

type ControlNetService struct {
	pb.UnimplementedControlNetServer
}

func NewControlNetService() *ControlNetService {
	return &ControlNetService{}
}

func (s *ControlNetService) GetR401APC(ctx context.Context, req *pb.GetR401Request) (*pb.GetR401Reply, error) {
	return &pb.GetR401Reply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: []*pb.VarDetail{
				{
					Name:     "MAPD",
					Desc:     "MAPD控制",
					SetValue: 600.00,
					EditAble: true,
					IsSwitch: false,
				},
				{
					Name:         "APC switch",
					Desc:         "APC开关选择",
					EditAble:     true,
					IsSwitch:     true,
					SwitchStatus: true,
				},
				{
					Name:         "APC 写入开关",
					Desc:         "APC 写入开关",
					EditAble:     true,
					IsSwitch:     true,
					SwitchStatus: false,
				},
				{
					Name:         "通讯指示灯",
					Desc:         "展示通讯状况",
					SetValue:     1,
					EditAble:     false,
					IsSwitch:     true,
					SwitchStatus: true,
				},
			},
		},
	}, nil
}
func (s *ControlNetService) GetOperationVars(ctx context.Context, req *pb.GetOperationVarsRequest) (*pb.GetOperationVarsReply, error) {
	return &pb.GetOperationVarsReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: []*pb.VarDetail{
				{
					Name:      "循环物料流量数据输入",
					Desc:      "输入循环物料流量数据",
					RealValue: 100.00,
					SetValue:  200.00,
					HighLimit: 300.00,
					LowLimit:  50.00,
					Unit:      "%mol",
				},
				{
					Name:      "配氢比数据输入",
					Desc:      "输入配氢比数据",
					RealValue: 100.00,
					SetValue:  200.00,
					HighLimit: 300.00,
					LowLimit:  50.00,
					Unit:      "%",
				},
				{
					Name:      "入口温度数据输入",
					Desc:      "输入入口温度数据",
					RealValue: 100.00,
					SetValue:  200.00,
					HighLimit: 300.00,
					LowLimit:  50.00,
					Unit:      "°C",
				},
				{
					Name:      "氢气流量数据输入",
					Desc:      "输入氢气流量数据",
					RealValue: 100.00,
					SetValue:  200.00,
					HighLimit: 300.00,
					LowLimit:  50.00,
					Unit:      "Kg/h",
				},
			},
		},
	}, nil
}
func (s *ControlNetService) GetReactorPerformance(ctx context.Context, req *pb.GetReactorPerformanceRequest) (*pb.GetReactorPerformanceReply, error) {
	return &pb.GetReactorPerformanceReply{}, nil
}
func (s *ControlNetService) SetR401APC(ctx context.Context, req *pb.SetR401APCRequest) (*pb.SetR401APCReply, error) {
	return &pb.SetR401APCReply{
		Code: 20000,
	}, nil
}
func (s *ControlNetService) SetOperationVars(ctx context.Context, req *pb.SetOperationVarsRequest) (*pb.SetOperationVarsReply, error) {
	return &pb.SetOperationVarsReply{}, nil
}
func (s *ControlNetService) ConfirmReactorPerf(ctx context.Context, req *pb.ConfirmReactorPerfRequest) (*pb.ConfirmReactorPerfReply, error) {
	return &pb.ConfirmReactorPerfReply{}, nil
}
func (s *ControlNetService) GetReactorPerf(ctx context.Context, req *pb.GetReactorPerfRequest) (*pb.GetReactorPerfReply, error) {
	return &pb.GetReactorPerfReply{}, nil
}
