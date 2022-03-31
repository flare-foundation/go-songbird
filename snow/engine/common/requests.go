// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"fmt"
	"strings"

	"github.com/flare-foundation/flare/ids"
)

const (
	minRequestsSize = 32
)

type req struct {
	validator ids.ShortID
	id        uint32
}

// Requests tracks pending container messages from a peer.
type Requests struct {
	reqsToID map[ids.ShortID]map[uint32]ids.ID
	idToReq  map[ids.ID]req
}

// Add a request. Assumes that requestIDs are unique. Assumes that containerIDs
// are only in one request at a time.
func (r *Requests) Add(validator ids.ShortID, requestID uint32, containerID ids.ID) {
	if r.reqsToID == nil {
		r.reqsToID = make(map[ids.ShortID]map[uint32]ids.ID, minRequestsSize)
	}
	vdrReqs, ok := r.reqsToID[validator]
	if !ok {
		vdrReqs = make(map[uint32]ids.ID)
		r.reqsToID[validator] = vdrReqs
	}
	vdrReqs[requestID] = containerID

	if r.idToReq == nil {
		r.idToReq = make(map[ids.ID]req, minRequestsSize)
	}
	r.idToReq[containerID] = req{
		validator: validator,
		id:        requestID,
	}
}

// Remove attempts to abandon a requestID sent to a validator. If the request is
// currently outstanding, the requested ID will be returned along with true. If
// the request isn't currently outstanding, false will be returned.
func (r *Requests) Remove(validator ids.ShortID, requestID uint32) (ids.ID, bool) {
	vdrReqs, ok := r.reqsToID[validator]
	if !ok {
		return ids.ID{}, false
	}
	containerID, ok := vdrReqs[requestID]
	if !ok {
		return ids.ID{}, false
	}

	if len(vdrReqs) == 1 {
		delete(r.reqsToID, validator)
	} else {
		delete(vdrReqs, requestID)
	}

	delete(r.idToReq, containerID)
	return containerID, true
}

// RemoveAny outstanding requests for the container ID. True is returned if the
// container ID had an outstanding request.
func (r *Requests) RemoveAny(containerID ids.ID) bool {
	req, ok := r.idToReq[containerID]
	if !ok {
		return false
	}

	r.Remove(req.validator, req.id)
	return true
}

// Len returns the total number of outstanding requests.
func (r *Requests) Len() int { return len(r.idToReq) }

// Contains returns true if there is an outstanding request for the container
// ID.
func (r *Requests) Contains(containerID ids.ID) bool {
	_, ok := r.idToReq[containerID]
	return ok
}

func (r Requests) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Requests: (Num Validators = %d)", len(r.reqsToID)))
	for validator, reqs := range r.reqsToID {
		sb.WriteString(fmt.Sprintf("\n  VDR[%s]: (Outstanding Requests %d)", validator, len(reqs)))
		for reqID, containerID := range reqs {
			sb.WriteString(fmt.Sprintf("\n    Request[%d]: %s", reqID, containerID))
		}
	}
	return sb.String()
}
