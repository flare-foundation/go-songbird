package rpcchainvm

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/flare-foundation/flare/api/proto/vmproto"
)

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
