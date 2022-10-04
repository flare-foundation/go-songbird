package validators

import (
	"fmt"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/utils/constants"
)

const (
	songbirdValidatorWeight = 50_000
)

func loadSongbirdValidators() Set {
	weight := uint64(songbirdValidatorWeight)
	nodeIDs := []string{
		"NodeID-3M9KVT6ixi4gVMisbm5TnPXYXgFN5LHuv",
		"NodeID-NnX4fajAmyvpL9RLfheNdc47FKKDuQW8i",
		"NodeID-AzdF8JNU468uwZYGquHt7bhDrsggZpK67",
		"NodeID-FqeGcnLAXbDTthd382aP9uyu1i47paRRh",
		"NodeID-B9HuZ5hDkRodyRRsiMEHWgMmmMF7xSKbj",
		"NodeID-Jx3E1F7mfkseZmqnFgDUFV3eusMxVdT6Z",
		"NodeID-FnvWuwvJGezs4uaBLujkfeM8U3gmAUY3Z",
		"NodeID-LhVs6hzHjBcEkzA1Eu8Qxb9nEQAk1Qbgf",
		"NodeID-9SqDo3MxpvEDN4bE4rLTyM7HkkKAw4h96",
		"NodeID-4tStYRTi3KDxFmv1YHTZAQxbzeyMA7z52",
		"NodeID-8XnMh17zo6pB8Pa2zptRBi9TbbMZgij2t",
		"NodeID-Cn9P5wgg7d9RNLqm4dFLCUV2diCxpkj7f",
		"NodeID-PEDdah7g7Efiii1xw8ex2dH58oMfByzjb",
		"NodeID-QCt9AxMPt5nn445CQGoA3yktqkChnKmPY",
		"NodeID-9bWz6J61B8WbQtzeSyA1jsXosyVbuUJd1",
		"NodeID-DLMnewsEwtSH8Qk7p9RGzUVyZAaZVMKsk",
		"NodeID-7meEpyjmGbL577th58dm4nvvtVZiJusFp",
		"NodeID-JeYnnrUkuArAAe2Sjo47Z3X5yfeF7cw43",
		"NodeID-Fdwp9Wtjh5rxzuTCF9z4zrSM31y7ZzBQS",
		"NodeID-JdEBRLS98PansyFKQUzFKqk4xqrVZ41nC",
	}
	set := NewSet()
	for _, nodeID := range nodeIDs {
		shortID, err := ids.ShortFromPrefixedString(nodeID, constants.NodeIDPrefix)
		if err != nil {
			panic(fmt.Sprintf("invalid songbird validator node ID: %s", nodeID))
		}
		err = set.AddWeight(shortID, weight)
		if err != nil {
			panic(fmt.Sprintf("could not add weight for validator (node: %s, weight: %d): %s", nodeID, weight, err))
		}
	}
	return set
}
