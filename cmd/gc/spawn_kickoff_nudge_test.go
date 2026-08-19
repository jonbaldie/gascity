package main

import (
	"encoding/json"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/jonbaldie/gascity/internal/agent"
	"github.com/jonbaldie/gascity/internal/beadmeta"
	"github.com/jonbaldie/gascity/internal/beads"
	"github.com/jonbaldie/gascity/internal/clock"
	"github.com/jonbaldie/gascity/internal/events"
	"github.com/jonbaldie/gascity/internal/nudgequeue"
	"github.com/jonbaldie/gascity/internal/session/sessiontest"
)

type routedWorkPair struct {
	BeadID   string
	RoutedTo string
}

func routedBeadPayload(t *testing.T, id, routedTo string) json.RawMessage {
	t.Helper()
	payload, err := json.Marshal(beads.Bead{
		ID:       id,
		Status:   "open",
		Type:     "task",
		Metadata: beads.StringMap{beadmeta.RoutedToMetadataKey: routedTo},
	})
	if err != nil {
		t.Fatalf("marshal routed bead payload: %v", err)
	}
	return payload
}

// routedWorkPairsFromEvents is the Go statement of the nudge-on-route selector.
// Formula steps stamp gc.routed_to at creation, so bead.updated-only misses them.
func routedWorkPairsFromEvents(evts []events.Event) []routedWorkPair {
	pairs := make([]routedWorkPair, 0, len(evts))
	seen := make(map[string]struct{}, len(evts))
	for _, event := range evts {
		switch event.Type {
		case events.BeadCreated, events.BeadUpdated:
		default:
			continue
		}
		bead, ok := beads.DecodeBeadEventPayload(event.Payload)
		if !ok {
			continue
		}
		routedTo := strings.TrimSpace(bead.Metadata[beadmeta.RoutedToMetadataKey])
		if bead.ID == "" || routedTo == "" {
			continue
		}
		key := bead.ID + "|" + routedTo
		if _, dup := seen[key]; dup {
			continue
		}
		seen[key] = struct{}{}
		pairs = append(pairs, routedWorkPair{BeadID: bead.ID, RoutedTo: routedTo})
	}
	return pairs
}

func TestRoutedWorkPairsFromEvents_CreatedAlreadyRoutedVsUpdatedOnly(t *testing.T) {
	created := events.Event{
		Type:    events.BeadCreated,
		Payload: routedBeadPayload(t, "gp-zeqis", "gws-decomposer"),
	}
	updated := events.Event{
		Type:    events.BeadUpdated,
		Payload: routedBeadPayload(t, "gp-qqr3a", "gws-operator"),
	}
	unroutedCreated := events.Event{
		Type:    events.BeadCreated,
		Payload: routedBeadPayload(t, "gp-bare", ""),
	}

	got := routedWorkPairsFromEvents([]events.Event{created, updated, unroutedCreated})
	if len(got) != 2 {
		t.Fatalf("pairs = %#v, want created-already-routed and updated-routed", got)
	}
	if got[0] != (routedWorkPair{BeadID: "gp-zeqis", RoutedTo: "gws-decomposer"}) {
		t.Errorf("first pair = %#v, want the creation-stamped route", got[0])
	}
	if got[1] != (routedWorkPair{BeadID: "gp-qqr3a", RoutedTo: "gws-operator"}) {
		t.Errorf("second pair = %#v, want the updated route", got[1])
	}

	updatedOnly := make([]events.Event, 0, 3)
	for _, event := range []events.Event{created, updated, unroutedCreated} {
		if event.Type == events.BeadUpdated {
			updatedOnly = append(updatedOnly, event)
		}
	}
	gotUpdatedOnly := routedWorkPairsFromEvents(updatedOnly)
	if len(gotUpdatedOnly) != 1 || gotUpdatedOnly[0].BeadID != "gp-qqr3a" {
		t.Fatalf("bead.updated-only pairs = %#v, want only the later update (the creation-stamped step is the #4382 gap)", gotUpdatedOnly)
	}
}

func TestSpawnKickoffNudge_TriggerBeadIsAssignment(t *testing.T) {
	now := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	info := sessiontest.SeedBead(t, beads.Bead{
		ID:     "gws-p0es",
		Title:  "decomposer",
		Type:   sessionBeadType,
		Labels: []string{sessionBeadLabel},
		Metadata: map[string]string{
			"session_name":                    "gws-p0es",
			beadmeta.TriggerBeadIDMetadataKey: "gp-zeqis",
		},
	})

	item, ok := spawnKickoffNudge(info, "gws-p0es", "gc hook --claim --json", now)
	if !ok {
		t.Fatal("spawnKickoffNudge returned false, want a first-work prompt for a trigger bead")
	}
	if item.Agent != "gws-p0es" {
		t.Errorf("Agent = %q, want gws-p0es", item.Agent)
	}
	if item.SessionID != "gws-p0es" {
		t.Errorf("SessionID = %q, want gws-p0es", item.SessionID)
	}
	if item.Source != spawnKickoffNudgeSource {
		t.Errorf("Source = %q, want %q", item.Source, spawnKickoffNudgeSource)
	}
	if item.Message != "gc hook --claim --json" {
		t.Errorf("Message = %q, want the configured claim nudge", item.Message)
	}
	if item.ID != "kickoff-gws-p0es-gp-zeqis" {
		t.Errorf("ID = %q, want stable kickoff-<session>-<trigger>", item.ID)
	}
	if item.Reference == nil || item.Reference.Kind != "bead" || item.Reference.ID != "gp-zeqis" {
		t.Errorf("Reference = %#v, want bead gp-zeqis", item.Reference)
	}
}

func TestSpawnKickoffNudge_NoTriggerBeadIsSilent(t *testing.T) {
	info := sessiontest.SeedBead(t, beads.Bead{
		ID:     "gws-idle",
		Title:  "idle",
		Type:   sessionBeadType,
		Labels: []string{sessionBeadLabel},
		Metadata: map[string]string{
			"session_name": "gws-idle",
		},
	})
	if _, ok := spawnKickoffNudge(info, "gws-idle", "check for assigned work", time.Now()); ok {
		t.Fatal("spawnKickoffNudge returned true without a trigger bead")
	}
}

func TestSpawnKickoffNudge_EmptyMessageUsesDefault(t *testing.T) {
	now := time.Date(2026, 8, 19, 12, 0, 0, 0, time.UTC)
	info := sessiontest.SeedBead(t, beads.Bead{
		ID:     "gws-p0es",
		Title:  "decomposer",
		Type:   sessionBeadType,
		Labels: []string{sessionBeadLabel},
		Metadata: map[string]string{
			"session_name":                    "gws-p0es",
			beadmeta.TriggerBeadIDMetadataKey: "gp-zeqis",
		},
	})
	item, ok := spawnKickoffNudge(info, "gws-p0es", "  ", now)
	if !ok {
		t.Fatal("spawnKickoffNudge returned false")
	}
	if item.Message != spawnKickoffNudgeDefaultMessage {
		t.Errorf("Message = %q, want %q", item.Message, spawnKickoffNudgeDefaultMessage)
	}
}

func TestEnqueueSpawnKickoffNudge_PrimesReconcilerSpawn(t *testing.T) {
	clearGCEnv(t)
	disableManagedDoltRecoveryForTest(t)
	t.Setenv("GC_BEADS", "file")
	cityPath := t.TempDir()
	store := beads.NewMemStore()
	session, err := store.Create(beads.Bead{
		Title:  "decomposer",
		Type:   sessionBeadType,
		Labels: []string{sessionBeadLabel},
		Metadata: map[string]string{
			"session_name":                    "gws-p0es",
			"state":                           "creating",
			beadmeta.TriggerBeadIDMetadataKey: "gp-zeqis",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	info := sessiontest.SeedBead(t, session)
	now := time.Date(2026, 8, 19, 12, 0, 1, 0, time.UTC)
	result := startResult{
		prepared: preparedStart{
			candidate: startCandidate{
				info: info,
				tp: TemplateParams{
					SessionName: "gws-p0es",
					Hints:       agent.StartupHints{Nudge: "gc hook --claim --json"},
				},
			},
			cityPath: cityPath,
			coreHash: "core",
			liveHash: "live",
		},
		outcome:  "success",
		started:  now.Add(-time.Second),
		finished: now,
	}

	enqueueSpawnKickoffNudge(cityPath, store, result, now, io.Discard)

	state, err := nudgequeue.LoadState(cityPath)
	if err != nil {
		t.Fatalf("LoadState: %v", err)
	}
	if len(state.Pending) != 1 {
		t.Fatalf("pending nudges = %d, want 1 first-work prompt", len(state.Pending))
	}
	got := state.Pending[0]
	if got.Source != spawnKickoffNudgeSource {
		t.Errorf("Source = %q, want %q", got.Source, spawnKickoffNudgeSource)
	}
	if got.Message != "gc hook --claim --json" {
		t.Errorf("Message = %q, want the configured claim nudge", got.Message)
	}
	if !strings.Contains(got.ID, "gp-zeqis") {
		t.Errorf("ID = %q, want it keyed on the trigger bead", got.ID)
	}
}

func TestCommitStartResult_EnqueuesKickoffForTriggerBead(t *testing.T) {
	clearGCEnv(t)
	disableManagedDoltRecoveryForTest(t)
	t.Setenv("GC_BEADS", "file")
	cityPath := t.TempDir()
	store := beads.NewMemStore()
	session, err := store.Create(beads.Bead{
		Title:  "decomposer",
		Type:   sessionBeadType,
		Labels: []string{sessionBeadLabel},
		Metadata: map[string]string{
			"session_name":                    "gws-p0es",
			"state":                           "creating",
			beadmeta.TriggerBeadIDMetadataKey: "gp-zeqis",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	now := time.Date(2026, 8, 19, 12, 0, 1, 0, time.UTC)
	result := startResult{
		prepared: preparedStart{
			candidate: startCandidate{
				info: sessiontest.SeedBead(t, session),
				tp: TemplateParams{
					SessionName:  "gws-p0es",
					TemplateName: "decomposer",
					Hints:        agent.StartupHints{Nudge: "check for assigned work"},
				},
			},
			cityPath: cityPath,
			coreHash: "core",
			liveHash: "live",
		},
		outcome:  "success",
		started:  now.Add(-time.Second),
		finished: now,
	}
	if !commitStartResult(result, sessionFrontDoor(store), &clock.Fake{Time: now}, events.NewFake(), 0, ioDiscard{}, ioDiscard{}) {
		t.Fatal("commitStartResult returned false")
	}
	state, err := nudgequeue.LoadState(cityPath)
	if err != nil {
		t.Fatalf("LoadState: %v", err)
	}
	if len(state.Pending) != 1 {
		t.Fatalf("pending after reconciler spawn = %d, want 1 (first prompt without a human nudge)", len(state.Pending))
	}
}
