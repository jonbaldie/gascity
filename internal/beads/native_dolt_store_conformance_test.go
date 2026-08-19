package beads_test

import (
	"testing"

	"github.com/jonbaldie/gascity/internal/beads"
	"github.com/jonbaldie/gascity/internal/beads/beadstest"
)

func TestNativeDoltStoreConformance(t *testing.T) {
	beadstest.RunStoreTests(t, beads.NewNativeDoltStoreForConformance)
}
