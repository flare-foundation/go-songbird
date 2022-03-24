// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validation

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"

	"github.com/flare-foundation/flare/ids"
)

const (
	validatorSetsCacheSize = 64
)

type Retriever interface {
	GetValidators(blockID ids.ID) (Set, error)
}

type CachingRetriever struct {
	retriever Retriever
	cache     *lru.Cache
}

func NewCachingRetriever(retriever Retriever) *CachingRetriever {

	cache, _ := lru.New(validatorSetsCacheSize)
	c := CachingRetriever{
		retriever: retriever,
		cache:     cache,
	}

	return &c
}

func (c *CachingRetriever) GetValidators(blockID ids.ID) (Set, error) {
	entry, ok := c.cache.Get(blockID)
	if ok {
		return entry.(Set), nil
	}
	set, err := c.retriever.GetValidators(blockID)
	if err != nil {
		return nil, fmt.Errorf("could not get validators for caching: %w", err)
	}
	c.cache.Add(blockID, set)
	return set, nil
}
