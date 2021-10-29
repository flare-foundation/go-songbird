package network

import (
	"context"
	"crypto"
	"net"
	"testing"
	"time"

	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/network/message"
	"github.com/flare-foundation/flare/network/throttling"
	"github.com/flare-foundation/flare/snow/networking/benchlist"
	"github.com/flare-foundation/flare/snow/validators"
	"github.com/flare-foundation/flare/utils"
	"github.com/flare-foundation/flare/utils/hashing"
	"github.com/flare-foundation/flare/utils/logging"
	"github.com/flare-foundation/flare/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

type TestMsg struct {
	op    message.Op
	bytes []byte
}

func newTestMsg(op message.Op, bits []byte) *TestMsg {
	return &TestMsg{op: op, bytes: bits}
}

func (m *TestMsg) Op() message.Op {
	return m.op
}

func (*TestMsg) Get(message.Field) interface{} {
	return nil
}

func (m *TestMsg) Bytes() []byte {
	return m.bytes
}

func (m *TestMsg) BytesSavedCompression() int {
	return 0
}

func TestPeer_Close(t *testing.T) {
	initCerts(t)

	log := logging.NoLog{}
	ip := utils.NewDynamicIPDesc(
		net.IPv6loopback,
		0,
	)
	id := ids.ShortID(hashing.ComputeHash160Array([]byte(ip.IP().String())))
	networkID := uint32(0)
	appVersion := version.NewDefaultApplication("app", 0, 1, 0)
	versionParser := version.NewDefaultApplicationParser()

	listener := &testListener{
		addr: &net.TCPAddr{
			IP:   net.IPv6loopback,
			Port: 0,
		},
		inbound: make(chan net.Conn, 1<<10),
		closed:  make(chan struct{}),
	}
	caller := &testDialer{
		addr: &net.TCPAddr{
			IP:   net.IPv6loopback,
			Port: 0,
		},
		outbounds: make(map[string]*testListener),
	}
	serverUpgrader0 := NewTLSServerUpgrader(tlsConfig0)
	clientUpgrader0 := NewTLSClientUpgrader(tlsConfig0)

	vdrs := validators.NewSet()
	handler := &testHandler{}

	versionManager := version.NewCompatibility(
		appVersion,
		appVersion,
		time.Now(),
		appVersion,
		appVersion,
		time.Now(),
		appVersion,
	)

	netwrk, err := NewDefaultNetwork(
		"",
		prometheus.NewRegistry(),
		log,
		id,
		ip,
		networkID,
		versionManager,
		versionParser,
		listener,
		caller,
		serverUpgrader0,
		clientUpgrader0,
		vdrs,
		vdrs,
		handler,
		throttling.InboundConnThrottlerConfig{},
		HealthConfig{},
		benchlist.NewManager(&benchlist.Config{}),
		defaultAliasTimeout,
		cert0.PrivateKey.(crypto.Signer),
		defaultPeerListSize,
		defaultGossipPeerListTo,
		defaultGossipPeerListFreq,
		defaultGossipAcceptedFrontierSize,
		defaultGossipOnAcceptSize,
		true,
		defaultInboundMsgThrottler,
		defaultOutboundMsgThrottler,
		ids.Set{},
	)
	assert.NoError(t, err)
	assert.NotNil(t, netwrk)

	ip1 := utils.NewDynamicIPDesc(
		net.IPv6loopback,
		1,
	)
	caller.outbounds[ip1.IP().String()] = listener
	conn, err := caller.Dial(context.Background(), ip1.IP())
	assert.NoError(t, err)

	basenetwork := netwrk.(*network)

	newmsgbytes := []byte("hello")

	// fake a peer, and write a message
	peer := newPeer(basenetwork, conn, ip1.IP())
	peer.sendQueue = [][]byte{}
	testMsg := newTestMsg(message.GetVersion, newmsgbytes)
	peer.Send(testMsg, true)

	go func() {
		err := netwrk.Close()
		assert.NoError(t, err)
	}()

	peer.Close()
}
