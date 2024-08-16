package clients

import (
	"users_service/configs"
	pb "users_service/genproto/orders_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IClientsMeneger interface {
	BranchesService() pb.BranchesServiceClient
}

type grpcClients struct {
	branchesService pb.BranchesServiceClient
}

func NewClients(cfg *configs.Config) (IClientsMeneger, error) {

	connOrderService, err := grpc.NewClient(cfg.OrdersServiceGrpcHost + cfg.OrdersServiceGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		branchesService: pb.NewBranchesServiceClient(connOrderService),
	}, nil
}

func (g *grpcClients) BranchesService() pb.BranchesServiceClient {
	return g.branchesService
}
