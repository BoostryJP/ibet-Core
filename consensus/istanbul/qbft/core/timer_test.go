package core

import (
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	qbfttypes "github.com/ethereum/go-ethereum/consensus/istanbul/qbft/types"
	"github.com/ethereum/go-ethereum/consensus/istanbul/validator"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// manualTimer deliberately allows firing after Stop: a real AfterFunc callback
// may already be running when it is cancelled.
type manualTimer struct {
	mu      sync.Mutex
	stopped bool
	delay   time.Duration
	fire    func()
}

func (m *manualTimer) Stop() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	wasActive := !m.stopped
	m.stopped = true
	return wasActive
}

type timerTestClock struct {
	created chan *manualTimer
}

func (clock *timerTestClock) afterFunc(d time.Duration, f func()) timer {
	m := &manualTimer{delay: d, fire: f}
	clock.created <- m
	return m
}

// Only methods used by these consensus scenarios are implemented. An unexpected
// backend interaction fails the test rather than silently succeeding.
type timerTestBackend struct {
	istanbul.Backend
	address    common.Address
	validators istanbul.ValidatorSet
	mux        event.TypeMux
	last       *types.Block
	verify     func(istanbul.Proposal) (time.Duration, error)
	broadcasts chan uint64
}

func (b *timerTestBackend) Address() common.Address  { return b.address }
func (b *timerTestBackend) EventMux() *event.TypeMux { return &b.mux }
func (b *timerTestBackend) Validators(istanbul.Proposal) istanbul.ValidatorSet {
	return b.validators
}
func (b *timerTestBackend) LastProposal() (istanbul.Proposal, common.Address) {
	return b.last, common.Address{}
}
func (b *timerTestBackend) HasBadProposal(common.Hash) bool { return false }
func (b *timerTestBackend) Sign([]byte) ([]byte, error)     { return make([]byte, 65), nil }
func (b *timerTestBackend) Broadcast(_ istanbul.ValidatorSet, code uint64, _ []byte) error {
	b.broadcasts <- code
	return nil
}
func (b *timerTestBackend) Gossip(istanbul.ValidatorSet, uint64, []byte) error { return nil }
func (b *timerTestBackend) Verify(p istanbul.Proposal) (time.Duration, error) {
	if b.verify != nil {
		return b.verify(p)
	}
	return 0, nil
}

func timerBlock(number int64, withTx bool) *types.Block {
	b := types.NewBlockWithHeader(&types.Header{Number: big.NewInt(number), Time: uint64(number), Difficulty: big.NewInt(1)})
	if withTx {
		b = b.WithBody([]*types.Transaction{types.NewTransaction(0, common.Address{}, big.NewInt(1), 21000, big.NewInt(0), nil)}, nil)
	}
	return b
}

func newTimerTestCore(t *testing.T) (*core, *timerTestBackend, *timerTestClock) {
	t.Helper()
	cfg := *istanbul.DefaultConfig
	cfg.EmptyBlockPeriod = 10
	cfg.ProposerPolicy = istanbul.NewRoundRobinProposerPolicy()
	addresses := []common.Address{{1}, {2}, {3}, {4}}
	b := &timerTestBackend{
		address: addresses[0], validators: validator.NewSet(addresses, cfg.ProposerPolicy),
		last: timerBlock(588, false), broadcasts: make(chan uint64, 100),
	}
	c := New(b, &cfg).(*core)
	clock := &timerTestClock{created: make(chan *manualTimer, 100)}
	c.afterFunc = clock.afterFunc
	c.timerEvents = make(chan timerEvent)
	c.timerDone = make(chan struct{})
	c.startNewRound(common.Big0)
	t.Cleanup(func() {
		c.stopOnce.Do(func() { close(c.timerDone) })
		c.stopTimer()
		b.mux.Stop()
	})
	return c, b, clock
}

func receiveTimer(t *testing.T, clock *timerTestClock) *manualTimer {
	t.Helper()
	select {
	case m := <-clock.created:
		return m
	case <-time.After(5 * time.Second):
		t.Fatal("timer was not registered")
		return nil
	}
}

func captureTimerEvent(t *testing.T, c *core, m *manualTimer) timerEvent {
	t.Helper()
	done := make(chan struct{})
	go func() { m.fire(); close(done) }()
	select {
	case ev := <-c.timerEvents:
		<-done
		return ev
	case <-time.After(5 * time.Second):
		t.Fatal("timer callback did not deliver an event")
		return timerEvent{}
	}
}

func acceptTimerProposal(t *testing.T, c *core, p istanbul.Proposal) {
	t.Helper()
	msg := qbfttypes.NewPreprepare(c.current.Sequence(), c.current.Round(), p)
	msg.SetSource(c.valSet.GetProposer().Address())
	if err := c.handlePreprepareMsg(msg); err != nil {
		t.Fatal(err)
	}
}

func TestTimerEmptyDelayAndPreprepare(t *testing.T) {
	for _, tc := range []struct {
		name       string
		delayFirst bool
		proposer   bool
	}{
		{"proposer/preprepare-first", false, true},
		{"proposer/delay-first", true, true},
		{"validator/preprepare-first", false, false},
		{"validator/delay-first", true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c, b, clock := newTimerTestCore(t)
			if !tc.proposer {
				b.address = common.Address{2}
				c.address = b.address
			}
			r := &Request{Proposal: timerBlock(589, false)}
			if err := c.handleRequest(r); err != nil {
				t.Fatal(err)
			}
			delay := receiveTimer(t, clock)
			if delay.delay != 9*time.Second {
				t.Fatalf("delay = %v", delay.delay)
			}
			ev := captureTimerEvent(t, c, delay)
			var obsolete *manualTimer
			if tc.delayFirst {
				c.handleTimerEvent(ev)
				obsolete = receiveTimer(t, clock)
			}
			acceptTimerProposal(t, c, r.Proposal)
			current := receiveTimer(t, clock)
			if !tc.delayFirst {
				c.handleTimerEvent(ev)
			}
			if c.timers[emptyBlockTimer].timer != nil || c.timers[roundChangeTimer].timer != current {
				t.Fatal("PRE-PREPARE did not leave exactly the current round timer")
			}
			if obsolete != nil {
				c.handleTimerEvent(captureTimerEvent(t, c, obsolete))
				if c.current.Round().Sign() != 0 {
					t.Fatal("obsolete timer changed round")
				}
			}
			valid := captureTimerEvent(t, c, current)
			c.handleTimerEvent(valid)
			c.handleTimerEvent(valid)
			if c.current.Round().Uint64() != 1 {
				t.Fatal("valid timeout must change round exactly once")
			}
		})
	}
}

func TestTimerStaleTimeoutAcrossBlocks(t *testing.T) {
	c, b, clock := newTimerTestCore(t)
	r := &Request{Proposal: timerBlock(589, false)}
	if err := c.handleRequest(r); err != nil {
		t.Fatal(err)
	}
	c.handleTimerEvent(captureTimerEvent(t, c, receiveTimer(t, clock)))
	old := receiveTimer(t, clock)
	acceptTimerProposal(t, c, r.Proposal)
	receiveTimer(t, clock)
	b.last = r.Proposal.(*types.Block)
	if err := c.handleFinalCommitted(); err != nil {
		t.Fatal(err)
	}
	r = &Request{Proposal: timerBlock(590, true)}
	if err := c.handleRequest(r); err != nil {
		t.Fatal(err)
	}
	receiveTimer(t, clock)
	acceptTimerProposal(t, c, r.Proposal)
	receiveTimer(t, clock)
	b.last = r.Proposal.(*types.Block)
	if err := c.handleFinalCommitted(); err != nil {
		t.Fatal(err)
	}
	if err := c.handleRequest(&Request{Proposal: timerBlock(591, false)}); err != nil {
		t.Fatal(err)
	}
	delay := receiveTimer(t, clock)
	c.handleTimerEvent(captureTimerEvent(t, c, old))
	if c.current.Sequence().Int64() != 591 || c.current.Round().Sign() != 0 || c.timers[emptyBlockTimer].timer != delay {
		t.Fatal("old block's timeout affected the current empty block")
	}
}

func TestTimerDelayedProposalInvalidation(t *testing.T) {
	for _, action := range []string{"replacement", "preprepare", "committed-state", "final-committed", "round-change"} {
		t.Run(action, func(t *testing.T) {
			c, b, clock := newTimerTestCore(t)
			r := &Request{Proposal: timerBlock(589, false)}
			if err := c.handleRequest(r); err != nil {
				t.Fatal(err)
			}
			old := captureTimerEvent(t, c, receiveTimer(t, clock))
			switch action {
			case "replacement":
				if err := c.handleRequest(&Request{Proposal: timerBlock(589, true)}); err != nil {
					t.Fatal(err)
				}
			case "preprepare":
				acceptTimerProposal(t, c, r.Proposal)
			case "committed-state":
				c.setState(StateCommitted)
			case "final-committed":
				b.last = r.Proposal.(*types.Block)
				if err := c.handleFinalCommitted(); err != nil {
					t.Fatal(err)
				}
			case "round-change":
				c.startNewRound(common.Big1)
			}
			broadcasts, timers := len(b.broadcasts), len(clock.created)
			c.handleTimerEvent(old)
			if len(b.broadcasts) != broadcasts || len(clock.created) != timers {
				t.Fatal("invalidated delay broadcast a proposal or reset a timer")
			}
		})
	}
}

func TestTimerRequestPeriods(t *testing.T) {
	for _, tc := range []struct {
		name  string
		empty uint64
		block uint64
		tx    bool
		delay time.Duration
		kind  timerKind
	}{
		{"disabled", 0, 1, false, 10 * time.Second, roundChangeTimer},
		{"equal", 1, 1, false, 10 * time.Second, roundChangeTimer},
		{"shorter", 1, 2, false, 10 * time.Second, roundChangeTimer},
		{"empty", 10, 1, false, 9 * time.Second, emptyBlockTimer},
		{"transaction", 10, 1, true, 10 * time.Second, roundChangeTimer},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c, b, clock := newTimerTestCore(t)
			c.config.EmptyBlockPeriod = tc.empty
			c.config.BlockPeriod = tc.block
			r := &Request{Proposal: timerBlock(589, tc.tx)}
			originalTime := r.Proposal.(*types.Block).Time()
			if err := c.handleRequest(r); err != nil {
				t.Fatal(err)
			}
			m := receiveTimer(t, clock)
			if m.delay != tc.delay || c.timers[tc.kind].timer != m {
				t.Fatal("unexpected request timer")
			}
			if tc.kind == emptyBlockTimer {
				if len(b.broadcasts) != 0 || r.Proposal.(*types.Block).Time() != originalTime+9 {
					t.Fatal("incorrect empty block timing")
				}
				c.handleTimerEvent(captureTimerEvent(t, c, m))
			}
			if len(b.broadcasts) != 1 || <-b.broadcasts != qbfttypes.PreprepareCode {
				t.Fatal("expected a proposal broadcast")
			}
		})
	}
}

func TestTimerRoundTimeoutBackoff(t *testing.T) {
	for _, tc := range []struct {
		round int64
		cap   uint64
		want  time.Duration
	}{
		{0, 0, 10 * time.Second}, {1, 0, 20 * time.Second}, {2, 0, 40 * time.Second},
		{2, 30, 30 * time.Second}, {100, 30, 30 * time.Second},
	} {
		c, _, clock := newTimerTestCore(t)
		c.current.SetRound(big.NewInt(tc.round))
		c.config.MaxRequestTimeoutSeconds = tc.cap
		c.newRoundChangeTimer()
		if got := receiveTimer(t, clock).delay; got != tc.want {
			t.Fatalf("round %d: %v, want %v", tc.round, got, tc.want)
		}
	}
}

func TestTimerFuturePreprepare(t *testing.T) {
	for _, stale := range []bool{false, true} {
		c, b, clock := newTimerTestCore(t)
		calls := 0
		b.verify = func(istanbul.Proposal) (time.Duration, error) {
			calls++
			if calls == 1 {
				return time.Second, consensus.ErrFutureBlock
			}
			return 0, nil
		}
		p := qbfttypes.NewPreprepare(c.current.Sequence(), c.current.Round(), timerBlock(589, false))
		p.SetSource(c.valSet.GetProposer().Address())
		if err := c.handlePreprepareMsg(p); err != consensus.ErrFutureBlock {
			t.Fatalf("error = %v", err)
		}
		m := receiveTimer(t, clock)
		if stale {
			c.startNewRound(common.Big1)
		}
		c.handleTimerEvent(captureTimerEvent(t, c, m))
		if stale {
			if calls != 1 {
				t.Fatal("retried stale PRE-PREPARE")
			}
		} else if calls != 2 || c.state != StatePreprepared {
			t.Fatal("valid future PRE-PREPARE was not accepted")
		}
	}
}

func TestTimerStopRestart(t *testing.T) {
	c, b, clock := newTimerTestCore(t)
	if err := c.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { c.Stop() })
	if err := b.mux.Post(istanbul.RequestEvent{Proposal: timerBlock(589, false)}); err != nil {
		t.Fatal(err)
	}
	old := receiveTimer(t, clock)
	// Race delivery against Stop. Both a blocked sender and an already consumed
	// event must allow shutdown to finish.
	delivered := make(chan struct{})
	go func() { old.fire(); close(delivered) }()
	stopped := make(chan struct{})
	go func() { c.Stop(); close(stopped) }()
	for _, done := range []chan struct{}{delivered, stopped} {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			t.Fatal("shutdown blocked")
		}
	}
	for len(clock.created) > 0 {
		<-clock.created
	}
	for len(b.broadcasts) > 0 {
		<-b.broadcasts
	}
	if err := c.Start(); err != nil {
		t.Fatal(err)
	}
	// This callback still targets the old channels and must return immediately.
	delivered = make(chan struct{})
	go func() { old.fire(); close(delivered) }()
	select {
	case <-delivered:
	case <-time.After(5 * time.Second):
		t.Fatal("old callback entered restarted core")
	}
	if err := b.mux.Post(istanbul.RequestEvent{Proposal: timerBlock(589, false)}); err != nil {
		t.Fatal(err)
	}
	current := receiveTimer(t, clock)
	current.fire()
	select {
	case code := <-b.broadcasts:
		if code != qbfttypes.PreprepareCode {
			t.Fatalf("unexpected message: %d", code)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("restarted core failed to propose")
	}
	if err := c.Stop(); err != nil {
		t.Fatal(err)
	}
}

func TestTimerFuturePreprepareReplacement(t *testing.T) {
	c, b, clock := newTimerTestCore(t)
	calls := 0
	b.verify = func(istanbul.Proposal) (time.Duration, error) {
		calls++
		return time.Second, consensus.ErrFutureBlock
	}
	p := qbfttypes.NewPreprepare(c.current.Sequence(), c.current.Round(), timerBlock(589, false))
	p.SetSource(c.valSet.GetProposer().Address())
	for i := 0; i < 2; i++ {
		if err := c.handlePreprepareMsg(p); err != consensus.ErrFutureBlock {
			t.Fatalf("error = %v", err)
		}
	}
	old, current := receiveTimer(t, clock), receiveTimer(t, clock)
	c.handleTimerEvent(captureTimerEvent(t, c, old))
	if calls != 2 || c.timers[futurePreprepareTimer].timer != current {
		t.Fatal("superseded retry affected the current timer")
	}
	b.verify = nil
	c.handleTimerEvent(captureTimerEvent(t, c, current))
	if c.state != StatePreprepared {
		t.Fatal("replacement retry was not accepted")
	}
}

// Exercise the production AfterFunc adapter and the real event loop as well as
// the deterministic, manually delivered timer scenarios above.
func TestTimerRealTimeout(t *testing.T) {
	c, b, _ := newTimerTestCore(t)
	c.afterFunc = func(d time.Duration, f func()) timer { return time.AfterFunc(d, f) }
	c.config.EmptyBlockPeriod = 0
	c.config.RequestTimeout = 1
	if err := c.Start(); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { c.Stop() })
	if err := b.mux.Post(istanbul.RequestEvent{Proposal: timerBlock(589, true)}); err != nil {
		t.Fatal(err)
	}
	for _, want := range []uint64{qbfttypes.PreprepareCode, qbfttypes.RoundChangeCode} {
		select {
		case got := <-b.broadcasts:
			if got != want {
				t.Fatalf("message = %d, want %d", got, want)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("real timer failed to trigger a round change")
		}
	}
	if err := c.Stop(); err != nil {
		t.Fatal(err)
	}
	for _, slot := range c.timers {
		if slot.timer != nil {
			t.Fatal("timer remained active after shutdown")
		}
	}
}
