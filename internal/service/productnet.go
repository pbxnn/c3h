package service

import (
	"c3h/internal/biz"
	"context"

	pb "c3h/api/product_net"

	"github.com/go-kratos/kratos/v2/log"
)

const (
	SwitchTypeSuccess = "success"
	SwitchTypeDanger  = "danger"
	SwitchTypePrimary = "primary"
)

type ProductNetService struct {
	pb.UnimplementedProductNetServer

	puc    *biz.ProductNetUsecase
	logger *log.Helper
}

func NewProductNetService(puc *biz.ProductNetUsecase, logger log.Logger) *ProductNetService {
	return &ProductNetService{
		puc:    puc,
		logger: log.NewHelper(log.With(logger, "module", "product-network-service")),
	}
}

func (s *ProductNetService) GetSwitchInfo(ctx context.Context, req *pb.SwitchRequest) (*pb.SwitchReply, error) {
	return &pb.SwitchReply{
		Code: 20000,
		Data: &pb.SwitchReplyData{
			DC_402_Total:  &pb.SwitchDetail{Type: SwitchTypeSuccess, Text: "on"},
			DC_402_Manual: &pb.SwitchDetail{Type: SwitchTypeSuccess, Text: "on"},
			DC_402_AB:     &pb.SwitchDetail{Type: SwitchTypePrimary, Text: "A"},
			MAPDExport:    &pb.SwitchDetail{Type: SwitchTypePrimary, Text: "600.00"},
			Drier:         &pb.SwitchDetail{Type: SwitchTypeDanger, Text: "off"},
			Communication: &pb.SwitchDetail{Type: SwitchTypeDanger, Text: "off"},
		},
	}, nil
}
func (s *ProductNetService) GetControlVars(ctx context.Context, req *pb.ControlVarRequest) (*pb.ControlVarReply, error) {
	return &pb.ControlVarReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: []*pb.VarDetail{
				{
					Name:      "UPPER_FIC40013",
					Desc:      "上层H2流量",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
				{
					Name:      "UPPER_H_ALK",
					Desc:      "上层氢炔比",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "mol/mol",
				},
				{
					Name:      "LOWER_FIC40013",
					Desc:      "下层H2流量",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
				{
					Name:      "TOTAL_H2_FLOW",
					Desc:      "总H2流量",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
			},
		},
	}, nil
}

func (s *ProductNetService) GetControlledVars(ctx context.Context, req *pb.ControlledVarRequest) (*pb.ControlledVarReply, error) {
	return &pb.ControlledVarReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: []*pb.VarDetail{
				{
					Name:      "AI40001",
					Desc:      "入口MAPD浓度",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "mol%",
				},
				{
					Name:      "AI40002",
					Desc:      "出口MAPD浓度",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "mol%",
				},
				{
					Name:      "AI40003",
					Desc:      "入口C3流量",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
				{
					Name:      "AI40004",
					Desc:      "C3回流量",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
			},
		},
	}, nil
}
func (s *ProductNetService) GetConfoundingVars(ctx context.Context, req *pb.ConfoundingVarRequest) (*pb.ConfoundingVarReply, error) {
	return &pb.ConfoundingVarReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: []*pb.VarDetail{
				{
					Name:      "AI40001",
					Desc:      "A/B入口温度",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "°C",
				},
				{
					Name:      "AI40002",
					Desc:      "A/B出口温度",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "°C",
				},
				{
					Name:      "AI40003",
					Desc:      "FA409压力",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
			},
		},
	}, nil
}
func (s *ProductNetService) GetCatalyst(ctx context.Context, req *pb.CatalystRequest) (*pb.CatalystReply, error) {
	return &pb.CatalystReply{
		Code: 20000,
		Data: &pb.VarReplyData{
			List: []*pb.VarDetail{
				{
					Name:      "AI400011",
					Desc:      "丙烯出入口化验分析值",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
				{
					Name:      "AI400012",
					Desc:      "丙烷出入口化验分析值",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "mol/mol",
				},
				{
					Name:      "AI400013",
					Desc:      "丙二烯出入口化验分析值",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
				{
					Name:      "AI400014",
					Desc:      "甲基乙炔出入口化验分析值",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
				{
					Name:      "AI400015",
					Desc:      "C3流量人工录入值",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
				{
					Name:      "AI400016",
					Desc:      "甲基乙炔出入口化验分析值",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "Kg/h",
				},
				{
					Name:      "AI400017",
					Desc:      "选择性",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "%",
				},
				{
					Name:      "AI400018",
					Desc:      "转化率",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "%",
				},
				{
					Name:      "AI400019",
					Desc:      "丙烯增量",
					RealValue: 123.45,
					SetValue:  200.00,
					LowLimit:  100.00,
					HighLimit: 300.00,
					Unit:      "%",
				},
			},
		},
	}, nil
}
