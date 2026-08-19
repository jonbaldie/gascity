package tmux

import (
	"context"
	"errors"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"
)

// wakeKeystrokeExecutor is a tmux fake for the controller wake/keystroke path
// (SendKeysDebounced). It answers pane-mode, provider-env, and busy-indicator
// probes so a unit test can drive the submit-confirm seam without a live TUI.
type wakeKeystrokeExecutor struct {
	mu sync.Mutex

	calls    [][]string
	provider string
	// busyAfter is the number of Enter keystrokes after which capture-pane
	// reports a busy indicator. 0 means already busy before any Enter (mid-turn).
	// -1 means never busy.
	busyAfter int
	enters    int
}

func (e *wakeKeystrokeExecutor) execute(args []string) (string, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	cp := slices.Clone(args)
	e.calls = append(e.calls, cp)

	switch {
	case slices.Contains(args, "#{pane_in_mode}"):
		return "0", nil
	case slices.Contains(args, "#{session_attached}"):
		return "1", nil
	case slices.Contains(args, "show-environment"):
		if e.provider == "" {
			return "", errors.New("unknown variable: GC_PROVIDER")
		}
		return "GC_PROVIDER=" + e.provider, nil
	case slices.Contains(args, "capture-pane"):
		if e.busyAfter == 0 || (e.busyAfter > 0 && e.enters >= e.busyAfter) {
			return "esc to interrupt", nil
		}
		return "❯ ", nil
	case slices.Contains(args, "list-windows"):
		return "1000", nil
	case slices.Contains(args, "list-panes"):
		return "%0\tsh\t1", nil
	case isSubmitEnterCall(args):
		e.enters++
		return "", nil
	default:
		return "", nil
	}
}

func (e *wakeKeystrokeExecutor) executeCtx(_ context.Context, args []string) (string, error) {
	return e.execute(args)
}

func (e *wakeKeystrokeExecutor) snapshot() (calls [][]string, enters int) {
	e.mu.Lock()
	defer e.mu.Unlock()
	return slices.Clone(e.calls), e.enters
}

func isSubmitEnterCall(args []string) bool {
	return slices.Contains(args, "send-keys") &&
		slices.Contains(args, "Enter") &&
		!slices.Contains(args, "-l") &&
		!slices.Contains(args, "-X")
}

func literalPayloads(calls [][]string) []string {
	var out []string
	for _, c := range calls {
		if !slices.Contains(c, "send-keys") || !slices.Contains(c, "-l") {
			continue
		}
		out = append(out, c[len(c)-1])
	}
	return out
}

func newWakeTmux(exec executor) *Tmux {
	return &Tmux{
		cfg:   Config{SocketName: "x", NudgeLockTimeout: 30 * time.Second},
		exec:  exec,
		sleep: noSleep,
	}
}

// TestSendKeysDebouncedRefusesBusyPaneWithoutTyping is the mid-turn half of
// gastownhall/gascity#4935: a wake that lands on an already-busy pane must not
// type wake text (exactly-one-copy unsubmitted input) and then report success.
// Fail closed before the paste; the controller retries on a later tick.
func TestSendKeysDebouncedRefusesBusyPaneWithoutTyping(t *testing.T) {
	fe := &wakeKeystrokeExecutor{provider: "claude", busyAfter: 0}
	tm := newWakeTmux(fe)

	err := tm.SendKeysDebounced("sess", "# gc-wake", 0)
	if err == nil {
		t.Fatal("SendKeysDebounced err = nil, want ErrSubmitPaneBusy (mid-turn must not type)")
	}
	if !errors.Is(err, ErrSubmitPaneBusy) {
		t.Fatalf("err = %v, want errors.Is(err, ErrSubmitPaneBusy)", err)
	}

	calls, enters := fe.snapshot()
	if payloads := literalPayloads(calls); len(payloads) != 0 {
		t.Fatalf("typed %v into a busy pane; mid-turn must not leave unsubmitted wake text", payloads)
	}
	if enters != 0 {
		t.Fatalf("Enter sends = %d, want 0 (must not submit into a busy turn)", enters)
	}
}

// TestSendKeysDebouncedFailsClosedWhenEnterNeverConfirms is the keystroke-path
// twin of TestNudgeSessionReturnsUnconfirmedErrorWhenNeverBusyForClaude: handing
// Enter to tmux is not delivery. An eligible provider whose busy indicator never
// appears must not report success with wake text stranded in the input box.
func TestSendKeysDebouncedFailsClosedWhenEnterNeverConfirms(t *testing.T) {
	fe := &wakeKeystrokeExecutor{provider: "claude", busyAfter: -1}
	tm := newWakeTmux(fe)

	err := tm.SendKeysDebounced("sess", "# gc-wake", 0)
	if err == nil {
		t.Fatal("SendKeysDebounced err = nil, want ErrNudgeSubmitUnconfirmed (unconfirmed Enter is not success)")
	}
	if !errors.Is(err, ErrNudgeSubmitUnconfirmed) {
		t.Fatalf("err = %v, want errors.Is(err, ErrNudgeSubmitUnconfirmed)", err)
	}

	calls, _ := fe.snapshot()
	if payloads := literalPayloads(calls); len(payloads) != 1 || payloads[0] != "# gc-wake" {
		t.Fatalf("literals = %v, want exactly one copy of the wake text", payloads)
	}
}

// TestSendKeysDebouncedReEntersWhileIdle proves a dropped first Enter is
// retried while the pane stays idle, then stops once busy is observed — the
// same submitEnterAndConfirm contract NudgeSession uses, on the keystroke seam.
func TestSendKeysDebouncedReEntersWhileIdle(t *testing.T) {
	fe := &wakeKeystrokeExecutor{provider: "claude", busyAfter: 2}
	tm := newWakeTmux(fe)

	if err := tm.SendKeysDebounced("sess", "# gc-wake", 0); err != nil {
		t.Fatalf("SendKeysDebounced: %v", err)
	}
	_, enters := fe.snapshot()
	if enters != 2 {
		t.Fatalf("Enter sends = %d, want 2 (initial dropped + one re-send)", enters)
	}
}

// TestSendKeysDebouncedBestEffortWhenNotEligible is the control: a pane whose
// busy indicator we cannot read (shell, unknown provider) keeps the historical
// single-Enter, nil-on-tmux-accept contract. Verification is a promise about a
// readable indicator; without one, fail-closed would reject every shell poke.
func TestSendKeysDebouncedBestEffortWhenNotEligible(t *testing.T) {
	fe := &wakeKeystrokeExecutor{provider: "", busyAfter: -1}
	tm := newWakeTmux(fe)

	if err := tm.SendKeysDebounced("sess", "# gc-wake", 0); err != nil {
		t.Fatalf("SendKeysDebounced: %v (ineligible pane must stay best-effort)", err)
	}
	calls, enters := fe.snapshot()
	if enters != 1 {
		t.Fatalf("Enter sends = %d, want 1 (best-effort single submit)", enters)
	}
	if payloads := literalPayloads(calls); len(payloads) != 1 || payloads[0] != "# gc-wake" {
		t.Fatalf("literals = %v, want [# gc-wake]", payloads)
	}
}

// TestNudgeSessionRefusesBusyPaneWithoutTyping closes the same mid-turn hole
// on the production controller wake path (Nudge → NudgeSession). Confirming
// submit by "pane is busy" is a false positive when the pane was already busy
// before the paste — that is how a wake reports success with stranded text.
func TestNudgeSessionRefusesBusyPaneWithoutTyping(t *testing.T) {
	fe := &wakeKeystrokeExecutor{provider: "claude", busyAfter: 0}
	tm := newWakeTmux(fe)
	err := tm.NudgeSession("sess", "# gc-wake")
	if err == nil {
		t.Fatal("NudgeSession err = nil, want ErrSubmitPaneBusy")
	}
	if !errors.Is(err, ErrSubmitPaneBusy) {
		t.Fatalf("err = %v, want errors.Is(err, ErrSubmitPaneBusy)", err)
	}

	calls, enters := fe.snapshot()
	for _, payload := range literalPayloads(calls) {
		if strings.Contains(payload, "# gc-wake") {
			t.Fatalf("NudgeSession typed %q into a busy pane", payload)
		}
	}
	if enters != 0 {
		t.Fatalf("Enter sends = %d, want 0", enters)
	}
}
