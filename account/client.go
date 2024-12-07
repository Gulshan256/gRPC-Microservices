package account

import (
	"context"

	"github.com/Gulshan256/go-gRPC-Microservices/account/github.com/Gulshan256/go-gRPC-Microservices/pb"
	"google.golang.org/grpc"
)

type client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

func newClient(conn *grpc.ClientConn) *client {
	return &client{conn: conn, service: pb.NewAccountServiceClient(conn)}
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) PostAccount(ctx context.Context, name string) (*Account, error) {
	r, err := c.service.PostAccount(
		ctx,
		&pb.PostAccountRequest{Name: name},
	)
	if err != nil {
		return nil, err
	}
	return &Account{ID: r.Account.Id, Name: r.Account.Name}, nil
}

func (c *client) GetAccount(ctx context.Context, id string) (*Account, error) {
	r, err := c.service.GetAccount(
		ctx,
		&pb.GetAccountRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &Account{ID: r.Account.Id, Name: r.Account.Name}, nil
}


func (c *client) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]*Account, error) {
	r, err := c.service.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{Skip: skip, Take: take},
	)
	if err != nil {
		return nil, err
	}
	accounts := make([]*Account, 0, len(r.Accounts))
	for _, a := range r.Accounts {
		accounts = append(accounts, &Account{ID: a.Id, Name: a.Name})
	}
	return accounts, nil
}