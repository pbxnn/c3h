package service

import (
	"c3h/internal/biz"
	"context"

	pb "c3h/api/product_net"

	"github.com/go-kratos/kratos/v2/log"
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
	return &pb.SwitchReply{}, nil
}
func (s *ProductNetService) GetControlVars(ctx context.Context, req *pb.ControlVarRequest) (*pb.ControlVarReply, error) {
	return &pb.ControlVarReply{}, nil
}
func (s *ProductNetService) GetControlledVars(ctx context.Context, req *pb.ControlledVarRequest) (*pb.ControlledVarReply, error) {
	return &pb.ControlledVarReply{}, nil
}
func (s *ProductNetService) GetConfoundingVars(ctx context.Context, req *pb.ConfoundingVarRequest) (*pb.ConfoundingVarReply, error) {
	return &pb.ConfoundingVarReply{}, nil
}
func (s *ProductNetService) GetCatalyst(ctx context.Context, req *pb.CatalystRequest) (*pb.CatalystReply, error) {
	return &pb.CatalystReply{}, nil
}
