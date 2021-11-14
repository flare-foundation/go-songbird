// (c) 2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ipcs

import (
	"time"

	"github.com/flare-foundation/flare/api"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/rpc"
)

type Client struct {
	requester rpc.EndpointRequester
}

// NewClient returns a Client for interacting with the IPCS endpoint
func NewClient(uri string, requestTimeout time.Duration) Client {
	return &client{
		requester: rpc.NewEndpointRequester(uri, "/ext/ipcs", "ipcs", requestTimeout),
	}
}

func (c *client) PublishBlockchain(blockchainID string) (*PublishBlockchainReply, error) {
	res := &PublishBlockchainReply{}
	err := c.requester.SendRequest("publishBlockchain", &PublishBlockchainArgs{
		BlockchainID: blockchainID,
	}, res)
	return res, err
}

func (c *client) UnpublishBlockchain(blockchainID string) (bool, error) {
	res := &api.SuccessResponse{}
	err := c.requester.SendRequest("unpublishBlockchain", &UnpublishBlockchainArgs{
		BlockchainID: blockchainID,
	}, res)
	return res.Success, err
}

func (c *client) GetPublishedBlockchains() ([]ids.ID, error) {
	res := &GetPublishedBlockchainsReply{}
	err := c.requester.SendRequest("getPublishedBlockchains", nil, res)
	return res.Chains, err
}
