package network

import (
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/nspcc-dev/neo-go/pkg/network/capability"
)

const (
	maxPoolSize = 10000
	connRetries = 3
)

// Discoverer is an interface that is responsible for maintaining
// a healthy connection pool.
type Discoverer interface {
	BackFill(...string)
	GetFanOut() int
	NetworkSize() int
	PoolCount() int
	RequestRemote(int)
	RegisterBadAddr(string)
	RegisterGoodAddr(string, capability.Capabilities)
	RegisterConnectedAddr(string)
	UnregisterConnectedAddr(string)
	UnconnectedPeers() []string
	BadPeers() []string
	GoodPeers() []AddressWithCapabilities
}

// AddressWithCapabilities represents a node address with its capabilities.
type AddressWithCapabilities struct {
	Address      string
	Capabilities capability.Capabilities
}

// DefaultDiscovery default implementation of the Discoverer interface.
type DefaultDiscovery struct {
	seeds            []string
	transport        Transporter
	lock             sync.RWMutex
	dialTimeout      time.Duration
	badAddrs         map[string]bool
	connectedAddrs   map[string]bool
	goodAddrs        map[string]capability.Capabilities
	unconnectedAddrs map[string]int
	attempted        map[string]bool
	optimalFanOut    int32
	networkSize      int32
	requestCh        chan int
}

// NewDefaultDiscovery returns a new DefaultDiscovery.
func NewDefaultDiscovery(addrs []string, dt time.Duration, ts Transporter) *DefaultDiscovery {
	d := &DefaultDiscovery{
		seeds:            addrs,
		transport:        ts,
		dialTimeout:      dt,
		badAddrs:         make(map[string]bool),
		connectedAddrs:   make(map[string]bool),
		goodAddrs:        make(map[string]capability.Capabilities),
		unconnectedAddrs: make(map[string]int),
		attempted:        make(map[string]bool),
		requestCh:        make(chan int),
	}
	return d
}

func newDefaultDiscovery(addrs []string, dt time.Duration, ts Transporter) Discoverer {
	return NewDefaultDiscovery(addrs, dt, ts)
}

// BackFill implements the Discoverer interface and will backfill
// the pool with the given addresses.
func (d *DefaultDiscovery) BackFill(addrs ...string) {
	d.lock.Lock()
	d.backfill(addrs...)
	d.lock.Unlock()
}

func (d *DefaultDiscovery) backfill(addrs ...string) {
	for _, addr := range addrs {
		if d.badAddrs[addr] || d.connectedAddrs[addr] ||
			d.unconnectedAddrs[addr] > 0 {
			continue
		}
		d.pushToPoolOrDrop(addr)
	}
	d.updateNetSize()
}

// PoolCount returns the number of the available node addresses.
func (d *DefaultDiscovery) PoolCount() int {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return d.poolCount()
}

func (d *DefaultDiscovery) poolCount() int {
	return len(d.unconnectedAddrs)
}

// pushToPoolOrDrop tries to push the address given into the pool, but if the pool
// is already full, it just drops it.
func (d *DefaultDiscovery) pushToPoolOrDrop(addr string) {
	if len(d.unconnectedAddrs) < maxPoolSize {
		d.unconnectedAddrs[addr] = connRetries
	}
}

// RequestRemote tries to establish a connection with n nodes.
func (d *DefaultDiscovery) RequestRemote(requested int) {
	for ; requested > 0; requested-- {
		var nextAddr string
		d.lock.Lock()
		for addr := range d.unconnectedAddrs {
			if !d.connectedAddrs[addr] && !d.attempted[addr] {
				nextAddr = addr
				break
			}
		}

		if nextAddr == "" {
			// Empty pool, try seeds.
			for _, addr := range d.seeds {
				if !d.connectedAddrs[addr] && !d.attempted[addr] {
					nextAddr = addr
					break
				}
			}
		}
		if nextAddr == "" {
			d.lock.Unlock()
			// The pool is empty, but all seed nodes are already connected (or attempted),
			// we can end up in an infinite loop here, so drop the request.
			break
		}
		d.attempted[nextAddr] = true
		d.lock.Unlock()
		go d.tryAddress(nextAddr)
	}
}

// RegisterBadAddr registers the given address as a bad address.
func (d *DefaultDiscovery) RegisterBadAddr(addr string) {
	var isSeed bool
	d.lock.Lock()
	for _, seed := range d.seeds {
		if addr == seed {
			isSeed = true
			break
		}
	}
	if !isSeed {
		d.unconnectedAddrs[addr]--
		if d.unconnectedAddrs[addr] <= 0 {
			d.badAddrs[addr] = true
			delete(d.unconnectedAddrs, addr)
			delete(d.goodAddrs, addr)
		}
	}
	d.updateNetSize()
	d.lock.Unlock()
}

// UnconnectedPeers returns all addresses of unconnected addrs.
func (d *DefaultDiscovery) UnconnectedPeers() []string {
	d.lock.RLock()
	addrs := make([]string, 0, len(d.unconnectedAddrs))
	for addr := range d.unconnectedAddrs {
		addrs = append(addrs, addr)
	}
	d.lock.RUnlock()
	return addrs
}

// BadPeers returns all addresses of bad addrs.
func (d *DefaultDiscovery) BadPeers() []string {
	d.lock.RLock()
	addrs := make([]string, 0, len(d.badAddrs))
	for addr := range d.badAddrs {
		addrs = append(addrs, addr)
	}
	d.lock.RUnlock()
	return addrs
}

// GoodPeers returns all addresses of known good peers (that at least once
// succeeded handshaking with us).
func (d *DefaultDiscovery) GoodPeers() []AddressWithCapabilities {
	d.lock.RLock()
	addrs := make([]AddressWithCapabilities, 0, len(d.goodAddrs))
	for addr, cap := range d.goodAddrs {
		addrs = append(addrs, AddressWithCapabilities{
			Address:      addr,
			Capabilities: cap,
		})
	}
	d.lock.RUnlock()
	return addrs
}

// RegisterGoodAddr registers a known good connected address that has passed
// handshake successfully.
func (d *DefaultDiscovery) RegisterGoodAddr(s string, c capability.Capabilities) {
	d.lock.Lock()
	d.goodAddrs[s] = c
	delete(d.badAddrs, s)
	d.lock.Unlock()
}

// UnregisterConnectedAddr tells the discoverer that this address is no longer
// connected, but it is still considered a good one.
func (d *DefaultDiscovery) UnregisterConnectedAddr(s string) {
	d.lock.Lock()
	delete(d.connectedAddrs, s)
	d.backfill(s)
	d.lock.Unlock()
}

// RegisterConnectedAddr tells discoverer that the given address is now connected.
func (d *DefaultDiscovery) RegisterConnectedAddr(addr string) {
	d.lock.Lock()
	delete(d.unconnectedAddrs, addr)
	d.connectedAddrs[addr] = true
	d.updateNetSize()
	d.lock.Unlock()
}

// GetFanOut returns the optimal number of nodes to broadcast packets to.
func (d *DefaultDiscovery) GetFanOut() int {
	return int(atomic.LoadInt32(&d.optimalFanOut))
}

// NetworkSize returns the estimated network size.
func (d *DefaultDiscovery) NetworkSize() int {
	return int(atomic.LoadInt32(&d.networkSize))
}

// updateNetSize updates network size estimation metric. Must be called under read lock.
func (d *DefaultDiscovery) updateNetSize() {
	var netsize = len(d.connectedAddrs) + len(d.unconnectedAddrs) + 1 // 1 for the node itself.
	var fanOut = 2.5 * math.Log(float64(netsize-1))                   // -1 for the number of potential peers.
	if netsize == 2 {                                                 // log(1) == 0.
		fanOut = 1 // But we still want to push messages to the peer.
	}

	atomic.StoreInt32(&d.optimalFanOut, int32(fanOut+0.5)) // Truncating conversion, hence +0.5.
	atomic.StoreInt32(&d.networkSize, int32(netsize))
	updateNetworkSizeMetric(netsize)
	updatePoolCountMetric(d.poolCount())
}

func (d *DefaultDiscovery) tryAddress(addr string) {
	err := d.transport.Dial(addr, d.dialTimeout)
	d.lock.Lock()
	delete(d.attempted, addr)
	d.lock.Unlock()
	if err != nil {
		d.RegisterBadAddr(addr)
		time.Sleep(d.dialTimeout)
		d.RequestRemote(1)
	}
}
