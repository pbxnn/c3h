package service

import (
	"c3h/internal/biz"
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "c3h/api/control_net"
)

type ControlNetService struct {
	pb.UnimplementedControlNetServer

	cuc    *biz.ControlNetUsecase
	logger *log.Helper
}

func NewControlNetService(cuc *biz.ControlNetUsecase, logger log.Logger) *ControlNetService {
	return &ControlNetService{
		cuc:    cuc,
		logger: log.NewHelper(log.With(logger, "module", "product-network-service")),
	}
}

func (s *ControlNetService) SetR401APC(ctx context.Context, req *pb.SetR401APCRequest) (*pb.SetR401APCReply, error) {
	return &pb.SetR401APCReply{}, nil
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
