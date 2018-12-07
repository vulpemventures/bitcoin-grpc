package server

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/btcsuite/btcd/rpcclient"

	pb "github.com/vulpemventures/bitcoin-grpc/proto"
)

type BitcoinServer struct {
	rpcClient *rpcclient.Client
	mu        sync.Mutex // protects routeNotes
}

// GetBlockchainInfo returns the feature at the given point.
func (b *BitcoinServer) GetBlockchainInfo(ctx context.Context, point *pb.Request) (*pb.Reply, error) {
	jsonResponse, err := b.rpcClient.GetBlockChainInfo()
	if err != nil {
		return nil, err
	}

	blob, errorOnMarshaling := json.Marshal(jsonResponse)
	if errorOnMarshaling != nil {
		return nil, errorOnMarshaling
	}

	// No feature was found, return an unnamed feature
	return &pb.Reply{Blob: string(blob)}, nil
}

func (s *BitcoinServer) NewServer(host string, user string, password string) error {

	// Configure regtest client to connect the daemon
	connConfig := &rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         host,
		User:         user,
		Pass:         password,
	}

	client, err := rpcclient.New(connConfig, nil)
	if err != nil {
		return err
	}

	s.rpcClient = client

	return nil
}
