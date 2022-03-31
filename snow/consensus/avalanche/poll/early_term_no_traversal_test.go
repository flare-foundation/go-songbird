// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poll

import (
	"testing"

	"github.com/flare-foundation/flare/ids"
)

func TestEarlyTermNoTraversalResults(t *testing.T) {
	alpha := 1

	vtxID := ids.ID{1}
	votes := []ids.ID{vtxID}

	validator1 := ids.ShortID{1} // k = 1

	validators := ids.ShortBag{}
	validators.Add(validator1)

	factory := NewEarlyTermNoTraversalFactory(alpha)
	poll := factory.New(validators)

	poll.Vote(validator1, votes)
	if !poll.Finished() {
		t.Fatalf("Poll did not terminate after receiving k votes")
	}

	result := poll.Result()
	if list := result.List(); len(list) != 1 {
		t.Fatalf("Wrong number of vertices returned")
	} else if retVtxID := list[0]; retVtxID != vtxID {
		t.Fatalf("Wrong vertex returned")
	} else if set := result.GetSet(vtxID); set.Len() != 1 {
		t.Fatalf("Wrong number of votes returned")
	}
}

func TestEarlyTermNoTraversalString(t *testing.T) {
	alpha := 2

	vtxID := ids.ID{1}
	votes := []ids.ID{vtxID}

	validator1 := ids.ShortID{1}
	validator2 := ids.ShortID{2} // k = 2

	validators := ids.ShortBag{}
	validators.Add(
		validator1,
		validator2,
	)

	factory := NewEarlyTermNoTraversalFactory(alpha)
	poll := factory.New(validators)

	poll.Vote(validator1, votes)

	expected := `waiting on Bag: (Size = 1)
    ID[BaMPFdqMUQ46BV8iRcwbVfsam55kMqcp]: Count = 1
received UniqueBag: (Size = 1)
    ID[SYXsAycDPUu4z2ZksJD5fh5nTDcH3vCFHnpcVye5XuJ2jArg]: Members = 0000000000000002`
	if result := poll.String(); expected != result {
		t.Fatalf("Poll should have returned:\n%s\nbut returned\n%s", expected, result)
	}
}

func TestEarlyTermNoTraversalDropsDuplicatedVotes(t *testing.T) {
	alpha := 2

	vtxID := ids.ID{1}
	votes := []ids.ID{vtxID}

	validator1 := ids.ShortID{1}
	validator2 := ids.ShortID{2} // k = 2

	validators := ids.ShortBag{}
	validators.Add(
		validator1,
		validator2,
	)

	factory := NewEarlyTermNoTraversalFactory(alpha)
	poll := factory.New(validators)

	poll.Vote(validator1, votes)
	if poll.Finished() {
		t.Fatalf("Poll finished after less than alpha votes")
	}
	poll.Vote(validator1, votes)
	if poll.Finished() {
		t.Fatalf("Poll finished after getting a duplicated vote")
	}
	poll.Vote(validator2, votes)
	if !poll.Finished() {
		t.Fatalf("Poll did not terminate after receiving k votes")
	}
}

func TestEarlyTermNoTraversalTerminatesEarly(t *testing.T) {
	alpha := 3

	vtxID := ids.ID{1}
	votes := []ids.ID{vtxID}

	validator1 := ids.ShortID{1}
	validator2 := ids.ShortID{2}
	validator3 := ids.ShortID{3}
	vdr4 := ids.ShortID{4}
	vdr5 := ids.ShortID{5} // k = 5

	validators := ids.ShortBag{}
	validators.Add(
		validator1,
		validator2,
		validator3,
		vdr4,
		vdr5,
	)

	factory := NewEarlyTermNoTraversalFactory(alpha)
	poll := factory.New(validators)

	poll.Vote(validator1, votes)
	if poll.Finished() {
		t.Fatalf("Poll finished after less than alpha votes")
	}
	poll.Vote(validator2, votes)
	if poll.Finished() {
		t.Fatalf("Poll finished after less than alpha votes")
	}
	poll.Vote(validator3, votes)
	if !poll.Finished() {
		t.Fatalf("Poll did not terminate early after receiving alpha votes for one vertex and none for other vertices")
	}
}

func TestEarlyTermNoTraversalForSharedAncestor(t *testing.T) {
	alpha := 4

	vtxA := ids.ID{1}
	vtxB := ids.ID{2}
	vtxC := ids.ID{3}
	vtxD := ids.ID{4}

	// If validators 1-3 vote for frontier vertices
	// B, C, and D respectively, which all share the common ancestor
	// A, then we cannot terminate early with alpha = k = 4
	// If the final vote is cast for any of A, B, C, or D, then
	// vertex A will have transitively received alpha = 4 votes
	validator1 := ids.ShortID{1}
	validator2 := ids.ShortID{2}
	validator3 := ids.ShortID{3}
	vdr4 := ids.ShortID{4}

	validators := ids.ShortBag{}
	validators.Add(validator1)
	validators.Add(validator2)
	validators.Add(validator3)
	validators.Add(vdr4)

	factory := NewEarlyTermNoTraversalFactory(alpha)
	poll := factory.New(validators)

	poll.Vote(validator1, []ids.ID{vtxB})
	if poll.Finished() {
		t.Fatalf("Poll finished early after receiving one vote")
	}
	poll.Vote(validator2, []ids.ID{vtxC})
	if poll.Finished() {
		t.Fatalf("Poll finished early after receiving two votes")
	}
	poll.Vote(validator3, []ids.ID{vtxD})
	if poll.Finished() {
		t.Fatalf("Poll terminated early, when a shared ancestor could have received alpha votes")
	}
	poll.Vote(vdr4, []ids.ID{vtxA})
	if !poll.Finished() {
		t.Fatalf("Poll did not terminate after receiving all outstanding votes")
	}
}

func TestEarlyTermNoTraversalWithFastDrops(t *testing.T) {
	alpha := 2

	validator1 := ids.ShortID{1}
	validator2 := ids.ShortID{2}
	validator3 := ids.ShortID{3} // k = 3

	validators := ids.ShortBag{}
	validators.Add(
		validator1,
		validator2,
		validator3,
	)

	factory := NewEarlyTermNoTraversalFactory(alpha)
	poll := factory.New(validators)

	poll.Vote(validator1, nil)
	if poll.Finished() {
		t.Fatalf("Poll finished early after dropping one vote")
	}
	poll.Vote(validator2, nil)
	if !poll.Finished() {
		t.Fatalf("Poll did not terminate after dropping two votes")
	}
}
