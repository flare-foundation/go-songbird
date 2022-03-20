// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
	lru "github.com/hashicorp/golang-lru"
)

const (
	validatorSetsCacheSize = 64
)

type Retriever interface {
	GetValidators(blockID ids.ID) (Set, error)
}

type cachingRetriever struct {
	retriever Retriever
	cache     *lru.Cache
}

func NewCachingRetriever(retriever Retriever) Retriever {

	cache, _ := lru.New(validatorSetsCacheSize)
	c := cachingRetriever{
		retriever: retriever,
		cache:     cache,
	}

	return &c
}

func (c *cachingRetriever) GetValidators(blockID ids.ID) (Set, error) {
	entry, ok := c.cache.Get(blockID)
	if ok {
		return entry.(Set), nil
	}
	set, err := c.retriever.GetValidators(blockID)
	if err != nil {
		return nil, fmt.Errorf("could not get validator set (block: %s): %w", blockID.String(), err)
	}
	c.cache.Add(blockID, set)
	return set, nil
}
