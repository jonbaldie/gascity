package runtime_test

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/jonbaldie/gascity/internal/runtime"
	"github.com/jonbaldie/gascity/internal/runtime/runtimetest"
)

func TestFakeConformance(t *testing.T) {
	var counter int64

	runtimetest.RunProviderTests(t, func(_ *testing.T) (runtime.Provider, runtime.Config, string) {
		return runtime.NewFake(), runtime.Config{}, fmt.Sprintf("fake-conform-%d", atomic.AddInt64(&counter, 1))
	})
}
