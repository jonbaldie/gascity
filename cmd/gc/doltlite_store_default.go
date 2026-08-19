//go:build !gascity_native_beads

package main

import "github.com/jonbaldie/gascity/internal/beads"

func openOptimizedDoltliteStore(_ string, _ *beads.BdStore) (beads.Store, bool) {
	return nil, false
}
