package core

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	qbfttypes "github.com/ethereum/go-ethereum/consensus/istanbul/qbft/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type timer interface {
	Stop() bool
}

type timerKind uint8

const (
	roundChangeTimer timerKind = iota
	emptyBlockTimer
	futurePreprepareTimer
	timerKindCount
)

func (kind timerKind) String() string {
	switch kind {
	case roundChangeTimer:
		return "round-change"
	case emptyBlockTimer:
		return "empty-block"
	case futurePreprepareTimer:
		return "future-preprepare"
	default:
		return "unknown"
	}
}

type scheduledTimer struct {
	timer      timer
	generation uint64
}

type timerEvent struct {
	kind         timerKind
	view         istanbul.View
	generation   uint64
	proposalHash common.Hash
	preprepare   *qbfttypes.Preprepare
}

// cancelTimer also invalidates callbacks which have already fired. All timer
// registration, cancellation and event handling runs on the consensus loop.
func (c *core) cancelTimer(kind timerKind) {
	slot := &c.timers[kind]
	if slot.timer != nil {
		slot.timer.Stop()
		slot.timer = nil
		c.currentLogger(true, nil).Trace("QBFT: stop timer", "kind", kind, "generation", slot.generation)
	}
	slot.generation++
}

func (c *core) scheduleTimer(kind timerKind, delay time.Duration, ev timerEvent) {
	c.cancelTimer(kind)
	slot := &c.timers[kind]
	ev.kind = kind
	ev.view = *c.currentView() // Copy big.Int values; do not capture mutable round state.
	ev.generation = slot.generation
	// Capture this execution's channels, so an old callback cannot enter a
	// restarted core. Callbacks never read or mutate consensus or timer state.
	events, done := c.timerEvents, c.timerDone
	slot.timer = c.afterFunc(delay, func() {
		select {
		case events <- ev:
		case <-done:
		}
	})
	c.currentLogger(true, nil).Trace("QBFT: start timer", "kind", kind, "generation", ev.generation, "delay", delay)
}

func (c *core) handleTimerEvent(ev timerEvent) {
	select {
	case <-c.timerDone:
		return
	default:
	}
	if ev.kind >= timerKindCount {
		return
	}
	slot := &c.timers[ev.kind]
	logger := c.currentLogger(true, nil).New("kind", ev.kind, "generation", ev.generation,
		"timer.sequence", ev.view.Sequence, "timer.round", ev.view.Round)
	if slot.timer == nil || ev.generation != slot.generation || c.current == nil || ev.view.Cmp(c.currentView()) != 0 {
		logger.Trace("QBFT: discard timer event", "reason", "inactive timer or stale generation/view")
		return
	}
	c.cancelTimer(ev.kind) // Consume once, even if the callback is delivered twice.
	logger.Trace("QBFT: handle timer event")
	switch ev.kind {
	case roundChangeTimer:
		c.handleTimeoutMsg()
	case emptyBlockTimer:
		request := c.current.pendingRequest
		if c.state != StateAcceptRequest || request == nil || request.Proposal.Hash() != ev.proposalHash {
			logger.Trace("QBFT: discard timer event", "reason", "proposal or state changed")
			return
		}
		c.newRoundChangeTimer()
		c.sendPreprepareMsg(request)
	case futurePreprepareTimer:
		// Apply the normal view/state checks again before retrying verification.
		if err := c.handleDecodedMessage(ev.preprepare); err != nil {
			return
		}
		data, err := rlp.EncodeToBytes(ev.preprepare)
		if err != nil {
			logger.Error("QBFT: can not encode deferred PRE-PREPARE", "err", err)
			return
		}
		c.backend.Gossip(c.valSet, ev.preprepare.Code(), data)
	}
}
