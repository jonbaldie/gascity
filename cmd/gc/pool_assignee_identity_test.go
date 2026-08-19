package main

import (
	"strings"
	"testing"
)

// TestSetTemplateEnvIdentityAlignsBeadsActorWithRuntimeAlias is the #5048
// write-side repro.
//
// Pool/cron workers claim via BEADS_ACTOR. Before this fix, setTemplateEnvIdentity
// stamped GC_ALIAS/GC_AGENT with the stable pool identity (bd.dog-1) while
// leaving BEADS_ACTOR on the sanitized tmux session name seeded by
// resolveTemplate (bd__dog-<beadID>). Claims therefore landed under a form
// later respawns no longer export in $GC_ALIAS, stranding healthy idle dogs.
func TestSetTemplateEnvIdentityAlignsBeadsActorWithRuntimeAlias(t *testing.T) {
	const (
		sanitizedSessionName = "bd__dog-ma-wisp-dw8d64"
		runtimeIdentity      = "bd.dog-1"
	)
	if !strings.Contains(sanitizedSessionName, "__") {
		t.Fatal("test premise broken: sanitized form must differ from dotted alias")
	}
	if sanitizedSessionName == runtimeIdentity {
		t.Fatal("test premise broken: sanitized session name and runtime identity must differ")
	}

	tp := &TemplateParams{
		Env: map[string]string{
			"GC_SESSION_NAME": sanitizedSessionName,
			"GC_AGENT":        sanitizedSessionName,
			"GC_ALIAS":        "bd.dog",
			"BEADS_ACTOR":     sanitizedSessionName,
		},
	}
	setTemplateEnvIdentity(tp, runtimeIdentity)

	if got := tp.Env["GC_ALIAS"]; got != runtimeIdentity {
		t.Fatalf("GC_ALIAS = %q, want runtime identity %q", got, runtimeIdentity)
	}
	if got := tp.Env["GC_AGENT"]; got != runtimeIdentity {
		t.Fatalf("GC_AGENT = %q, want runtime identity %q", got, runtimeIdentity)
	}
	if got := tp.Env["BEADS_ACTOR"]; got != runtimeIdentity {
		t.Fatalf("BEADS_ACTOR = %q, want runtime identity %q (claim assignee must match $GC_ALIAS)", got, runtimeIdentity)
	}
	if got := tp.Env["GC_SESSION_NAME"]; got != sanitizedSessionName {
		t.Fatalf("GC_SESSION_NAME = %q, want preserved sanitized handle %q", got, sanitizedSessionName)
	}
	if !tp.EnvIdentityStamped {
		t.Fatal("EnvIdentityStamped = false, want true")
	}
}
