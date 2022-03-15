package rpcchainvm

import (
	"context"
	"time"

	"github.com/flare-foundation/flare/database/manager"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/snow/consensus/snowman"
	"github.com/flare-foundation/flare/snow/engine/common"
	"github.com/flare-foundation/flare/snow/validators"
	"github.com/flare-foundation/flare/version"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/flare-foundation/flare/api/proto/vmproto"
)

type ChainVMMock struct {
	GetValidatorsFunc func(blockID ids.ID) (validators.Set, error)

	AppRequestFunc           func(nodeID ids.ShortID, requestID uint32, deadline time.Time, request []byte) error
	AppRequestFailedFunc     func(nodeID ids.ShortID, requestID uint32) error
	AppResponseFunc          func(nodeID ids.ShortID, requestID uint32, response []byte) error
	AppGossipFunc            func(nodeID ids.ShortID, msg []byte) error
	HealthCheckFunc          func() (interface{}, error)
	ConnectedFunc            func(id ids.ShortID, nodeVersion version.Application) error
	DisconnectedFunc         func(id ids.ShortID) error
	InitializeFunc           func(ctx *snow.Context, dbManager manager.Manager, genesisBytes []byte, upgradeBytes []byte, configBytes []byte, toEngine chan<- common.Message, fxs []*common.Fx, appSender common.AppSender) error
	SetStateFunc             func(state snow.State) error
	ShutdownFunc             func() error
	VersionFunc              func() (string, error)
	CreateStaticHandlersFunc func() (map[string]*common.HTTPHandler, error)
	CreateHandlersFunc       func() (map[string]*common.HTTPHandler, error)
	GetBlockFunc             func(ids.ID) (snowman.Block, error)
	ParseBlockFunc           func([]byte) (snowman.Block, error)
	BuildBlockFunc           func() (snowman.Block, error)
	SetPreferenceFunc        func(ids.ID) error
	LastAcceptedFunc         func() (ids.ID, error)
}

func (c ChainVMMock) GetValidators(blockID ids.ID) (validators.Set, error) {
	return c.GetValidatorsFunc(blockID)
}

func (c ChainVMMock) AppRequest(nodeID ids.ShortID, requestID uint32, deadline time.Time, request []byte) error {
	return c.AppRequestFunc(nodeID, requestID, deadline, request)
}

func (c ChainVMMock) AppRequestFailed(nodeID ids.ShortID, requestID uint32) error {
	return c.AppRequestFailedFunc(nodeID, requestID)
}

func (c ChainVMMock) AppResponse(nodeID ids.ShortID, requestID uint32, response []byte) error {
	return c.AppResponseFunc(nodeID, requestID, response)
}

func (c ChainVMMock) AppGossip(nodeID ids.ShortID, msg []byte) error {
	return c.AppGossipFunc(nodeID, msg)
}

func (c ChainVMMock) HealthCheck() (interface{}, error) {
	return c.HealthCheckFunc()
}

func (c ChainVMMock) Connected(id ids.ShortID, nodeVersion version.Application) error {
	return c.ConnectedFunc(id, nodeVersion)
}

func (c ChainVMMock) Disconnected(id ids.ShortID) error {
	return c.DisconnectedFunc(id)
}

func (c ChainVMMock) Initialize(ctx *snow.Context, dbManager manager.Manager, genesisBytes []byte, upgradeBytes []byte, configBytes []byte, toEngine chan<- common.Message, fxs []*common.Fx, appSender common.AppSender) error {
	return c.InitializeFunc(ctx, dbManager, genesisBytes, upgradeBytes, configBytes, toEngine, fxs, appSender)
}

func (c ChainVMMock) SetState(state snow.State) error {
	return c.SetStateFunc(state)
}

func (c ChainVMMock) Shutdown() error {
	return c.ShutdownFunc()
}

func (c ChainVMMock) Version() (string, error) {
	return c.VersionFunc()
}

func (c ChainVMMock) CreateStaticHandlers() (map[string]*common.HTTPHandler, error) {
	return c.CreateStaticHandlersFunc()
}

func (c ChainVMMock) CreateHandlers() (map[string]*common.HTTPHandler, error) {
	return c.CreateHandlersFunc()
}

func (c ChainVMMock) GetBlock(id ids.ID) (snowman.Block, error) {
	return c.GetBlockFunc(id)
}

func (c ChainVMMock) ParseBlock(bytes []byte) (snowman.Block, error) {
	return c.ParseBlockFunc(bytes)
}

func (c ChainVMMock) BuildBlock() (snowman.Block, error) {
	return c.BuildBlockFunc()
}

func (c ChainVMMock) SetPreference(id ids.ID) error {
	return c.SetPreferenceFunc(id)
}

func (c ChainVMMock) LastAccepted() (ids.ID, error) {
	return c.LastAcceptedFunc()
}

type VMClientMock struct {
	InitializeFunc           func(ctx context.Context, in *vmproto.InitializeRequest, opts ...grpc.CallOption) (*vmproto.InitializeResponse, error)
	SetStateFunc             func(ctx context.Context, in *vmproto.SetStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ShutdownFunc             func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateHandlersFunc       func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.CreateHandlersResponse, error)
	CreateStaticHandlersFunc func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.CreateStaticHandlersResponse, error)
	ConnectedFunc            func(ctx context.Context, in *vmproto.ConnectedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DisconnectedFunc         func(ctx context.Context, in *vmproto.DisconnectedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BuildBlockFunc           func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.BuildBlockResponse, error)
	ParseBlockFunc           func(ctx context.Context, in *vmproto.ParseBlockRequest, opts ...grpc.CallOption) (*vmproto.ParseBlockResponse, error)
	GetBlockFunc             func(ctx context.Context, in *vmproto.GetBlockRequest, opts ...grpc.CallOption) (*vmproto.GetBlockResponse, error)
	SetPreferenceFunc        func(ctx context.Context, in *vmproto.SetPreferenceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HealthFunc               func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.HealthResponse, error)
	VersionFunc              func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.VersionResponse, error)
	AppRequestFunc           func(ctx context.Context, in *vmproto.AppRequestMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AppRequestFailedFunc     func(ctx context.Context, in *vmproto.AppRequestFailedMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AppResponseFunc          func(ctx context.Context, in *vmproto.AppResponseMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AppGossipFunc            func(ctx context.Context, in *vmproto.AppGossipMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GatherFunc               func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.GatherResponse, error)
	BlockVerifyFunc          func(ctx context.Context, in *vmproto.BlockVerifyRequest, opts ...grpc.CallOption) (*vmproto.BlockVerifyResponse, error)
	BlockAcceptFunc          func(ctx context.Context, in *vmproto.BlockAcceptRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BlockRejectFunc          func(ctx context.Context, in *vmproto.BlockRejectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAncestorsFunc         func(ctx context.Context, in *vmproto.GetAncestorsRequest, opts ...grpc.CallOption) (*vmproto.GetAncestorsResponse, error)
	BatchedParseBlockFunc    func(ctx context.Context, in *vmproto.BatchedParseBlockRequest, opts ...grpc.CallOption) (*vmproto.BatchedParseBlockResponse, error)
	VerifyHeightIndexFunc    func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.VerifyHeightIndexResponse, error)
	GetBlockIDAtHeightFunc   func(ctx context.Context, in *vmproto.GetBlockIDAtHeightRequest, opts ...grpc.CallOption) (*vmproto.GetBlockIDAtHeightResponse, error)
	FetchValidatorsFunc      func(ctx context.Context, in *vmproto.FetchValidatorsRequest, opts ...grpc.CallOption) (*vmproto.FetchValidatorsResponse, error)
}

func (v VMClientMock) Initialize(ctx context.Context, in *vmproto.InitializeRequest, opts ...grpc.CallOption) (*vmproto.InitializeResponse, error) {
	return v.InitializeFunc(ctx, in, opts...)
}

func (v VMClientMock) SetState(ctx context.Context, in *vmproto.SetStateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.SetStateFunc(ctx, in, opts...)
}

func (v VMClientMock) Shutdown(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.ShutdownFunc(ctx, in, opts...)
}

func (v VMClientMock) CreateHandlers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.CreateHandlersResponse, error) {
	return v.CreateHandlersFunc(ctx, in, opts...)
}

func (v VMClientMock) CreateStaticHandlers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.CreateStaticHandlersResponse, error) {
	return v.CreateStaticHandlersFunc(ctx, in, opts...)
}

func (v VMClientMock) Connected(ctx context.Context, in *vmproto.ConnectedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.ConnectedFunc(ctx, in, opts...)
}

func (v VMClientMock) Disconnected(ctx context.Context, in *vmproto.DisconnectedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.DisconnectedFunc(ctx, in, opts...)
}

func (v VMClientMock) BuildBlock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.BuildBlockResponse, error) {
	return v.BuildBlockFunc(ctx, in, opts...)
}

func (v VMClientMock) ParseBlock(ctx context.Context, in *vmproto.ParseBlockRequest, opts ...grpc.CallOption) (*vmproto.ParseBlockResponse, error) {
	return v.ParseBlockFunc(ctx, in, opts...)
}

func (v VMClientMock) GetBlock(ctx context.Context, in *vmproto.GetBlockRequest, opts ...grpc.CallOption) (*vmproto.GetBlockResponse, error) {
	return v.GetBlockFunc(ctx, in, opts...)
}

func (v VMClientMock) SetPreference(ctx context.Context, in *vmproto.SetPreferenceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.SetPreferenceFunc(ctx, in, opts...)
}

func (v VMClientMock) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.HealthResponse, error) {
	return v.HealthFunc(ctx, in, opts...)
}

func (v VMClientMock) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.VersionResponse, error) {
	return v.VersionFunc(ctx, in, opts...)
}

func (v VMClientMock) AppRequest(ctx context.Context, in *vmproto.AppRequestMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.AppRequestFunc(ctx, in, opts...)
}

func (v VMClientMock) AppRequestFailed(ctx context.Context, in *vmproto.AppRequestFailedMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.AppRequestFailedFunc(ctx, in, opts...)
}

func (v VMClientMock) AppResponse(ctx context.Context, in *vmproto.AppResponseMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.AppResponseFunc(ctx, in, opts...)
}

func (v VMClientMock) AppGossip(ctx context.Context, in *vmproto.AppGossipMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.AppGossipFunc(ctx, in, opts...)
}

func (v VMClientMock) Gather(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.GatherResponse, error) {
	return v.GatherFunc(ctx, in, opts...)
}

func (v VMClientMock) BlockVerify(ctx context.Context, in *vmproto.BlockVerifyRequest, opts ...grpc.CallOption) (*vmproto.BlockVerifyResponse, error) {
	return v.BlockVerifyFunc(ctx, in, opts...)
}

func (v VMClientMock) BlockAccept(ctx context.Context, in *vmproto.BlockAcceptRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.BlockAcceptFunc(ctx, in, opts...)
}

func (v VMClientMock) BlockReject(ctx context.Context, in *vmproto.BlockRejectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return v.BlockRejectFunc(ctx, in, opts...)
}

func (v VMClientMock) GetAncestors(ctx context.Context, in *vmproto.GetAncestorsRequest, opts ...grpc.CallOption) (*vmproto.GetAncestorsResponse, error) {
	return v.GetAncestorsFunc(ctx, in, opts...)
}

func (v VMClientMock) BatchedParseBlock(ctx context.Context, in *vmproto.BatchedParseBlockRequest, opts ...grpc.CallOption) (*vmproto.BatchedParseBlockResponse, error) {
	return v.BatchedParseBlockFunc(ctx, in, opts...)
}

func (v VMClientMock) VerifyHeightIndex(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vmproto.VerifyHeightIndexResponse, error) {
	return v.VerifyHeightIndexFunc(ctx, in, opts...)
}

func (v VMClientMock) GetBlockIDAtHeight(ctx context.Context, in *vmproto.GetBlockIDAtHeightRequest, opts ...grpc.CallOption) (*vmproto.GetBlockIDAtHeightResponse, error) {
	return v.GetBlockIDAtHeightFunc(ctx, in, opts...)
}

func (v VMClientMock) FetchValidators(ctx context.Context, in *vmproto.FetchValidatorsRequest, opts ...grpc.CallOption) (*vmproto.FetchValidatorsResponse, error) {
	return v.FetchValidatorsFunc(ctx, in, opts...)
}
