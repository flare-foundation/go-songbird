// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gkeystore

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/flare-foundation/flare/api/keystore"
	"github.com/flare-foundation/flare/api/proto/gkeystoreproto"
	"github.com/flare-foundation/flare/api/proto/rpcdbproto"
	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/database/rpcdb"
	"github.com/flare-foundation/flare/utils/math"
	"github.com/flare-foundation/flare/vms/rpcchainvm/grpcutils"
)

var _ gkeystoreproto.KeystoreServer = &Server{}

// Server is a snow.Keystore that is managed over RPC.
type Server struct {
	gkeystoreproto.UnimplementedKeystoreServer
	ks     keystore.BlockchainKeystore
	broker *plugin.GRPCBroker
}

// NewServer returns a keystore connected to a remote keystore
func NewServer(ks keystore.BlockchainKeystore, broker *plugin.GRPCBroker) *Server {
	return &Server{
		ks:     ks,
		broker: broker,
	}
}

func (s *Server) GetDatabase(
	_ context.Context,
	req *gkeystoreproto.GetDatabaseRequest,
) (*gkeystoreproto.GetDatabaseResponse, error) {
	db, err := s.ks.GetRawDatabase(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	closer := dbCloser{Database: db}

	// start the db server
	dbBrokerID := s.broker.NextId()
	go s.broker.AcceptAndServe(dbBrokerID, func(opts []grpc.ServerOption) *grpc.Server {
		opts = append(opts,
			grpc.MaxRecvMsgSize(math.MaxInt),
			grpc.MaxSendMsgSize(math.MaxInt),
		)
		server := grpc.NewServer(opts...)
		closer.closer.Add(server)
		db := rpcdb.NewServer(&closer)
		rpcdbproto.RegisterDatabaseServer(server, db)
		return server
	})
	return &gkeystoreproto.GetDatabaseResponse{DbServer: dbBrokerID}, nil
}

type dbCloser struct {
	database.Database
	closer grpcutils.ServerCloser
}

func (db *dbCloser) Close() error {
	err := db.Database.Close()
	db.closer.Stop()
	return err
}
