package main

import (
	"github.com/jonbaldie/gascity/internal/config"
	"github.com/jonbaldie/gascity/internal/featureflags"
)

// applyFeatureFlags propagates daemon-level feature flags to the formula and
// molecule packages. Must be called after config.LoadWithIncludes and before
// any formula compilation or molecule instantiation.
func applyFeatureFlags(cfg *config.City) {
	featureflags.Apply(featureflags.FromConfig(cfg))
}
