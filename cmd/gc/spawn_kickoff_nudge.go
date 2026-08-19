package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/jonbaldie/gascity/internal/beads"
	sessionpkg "github.com/jonbaldie/gascity/internal/session"
)

const (
	spawnKickoffNudgeSource         = "spawn"
	spawnKickoffNudgeDefaultMessage = "check for assigned work"
)

// spawnKickoffNudge is the first-work prompt for a session the reconciler just
// spawned. Mechanical: a non-empty trigger bead means the session exists to
// run that work, so the first prompt is the assignment. No trigger bead means
// nothing to enqueue.
func spawnKickoffNudge(info sessionpkg.Info, sessionName, message string, now time.Time) (queuedNudge, bool) {
	triggerID := strings.TrimSpace(info.TriggerBeadID)
	if triggerID == "" {
		return queuedNudge{}, false
	}
	sessionName = strings.TrimSpace(sessionName)
	if sessionName == "" {
		return queuedNudge{}, false
	}
	message = strings.TrimSpace(message)
	if message == "" {
		message = spawnKickoffNudgeDefaultMessage
	}
	return newQueuedNudgeWithOptions(sessionName, message, spawnKickoffNudgeSource, now, queuedNudgeOptions{
		ID:        spawnKickoffNudgeID(info.ID, triggerID),
		SessionID: info.ID,
		Reference: &nudgeReference{Kind: "bead", ID: triggerID},
	}), true
}

func spawnKickoffNudgeID(sessionID, triggerID string) string {
	return "kickoff-" + strings.TrimSpace(sessionID) + "-" + strings.TrimSpace(triggerID)
}

// enqueueSpawnKickoffNudge queues the first-work prompt after a successful
// reconciler spawn. cityPath empty skips (hand-built commit fixtures); a
// duplicate kickoff id is a no-op.
func enqueueSpawnKickoffNudge(cityPath string, store beads.Store, result startResult, now time.Time, stderr io.Writer) {
	if cityPath == "" || result.err != nil {
		return
	}
	sessionName := strings.TrimSpace(result.prepared.candidate.name())
	if sessionName == "" {
		sessionName = strings.TrimSpace(result.prepared.candidate.tp.SessionName)
	}
	message := strings.TrimSpace(result.prepared.candidate.tp.Hints.Nudge)
	item, ok := spawnKickoffNudge(result.prepared.candidate.info, sessionName, message, now)
	if !ok {
		return
	}
	if err := enqueueQueuedNudgeWithStore(cityPath, beads.NudgesStore{Store: store}, item); err != nil {
		fmt.Fprintf(stderr, "session reconciler: enqueueing spawn kickoff nudge for %s: %v\n", sessionName, err) //nolint:errcheck
	}
}

func sessionFrontBeads(sessFront *sessionpkg.Store) beads.Store {
	if sessFront == nil {
		return nil
	}
	return sessFront.Store().Store
}
