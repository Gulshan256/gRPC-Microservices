package account

import (
	"context"
	"fmt"
	"net"

	"github.com/Gulshan256/go-gRPC-Microservices/account/github.com/Gulshan256/go-gRPC-Microservices/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{service: s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

// post account
func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	account, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.PostAccountResponse{Account: &pb.Account{
		Id:   account.ID,
		Name: account.Name,
	}}, nil

}

// get account
func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := s.service.GetAccount(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{Account: &pb.Account{
		Id:   account.ID,
		Name: account.Name,
	}}, nil
}

// get accounts
func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	accounts, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}

	pbAccounts := []*pb.Account{}
	for _, account := range accounts {
		pbAccounts = append(pbAccounts, &pb.Account{
			Id:   account.ID,
			Name: account.Name,
		})
	}

	return &pb.GetAccountsResponse{Accounts: pbAccounts}, nil
}
