package beadmail

import (
	"testing"

	"github.com/jonbaldie/gascity/internal/beads"
	"github.com/jonbaldie/gascity/internal/mail"
	"github.com/jonbaldie/gascity/internal/mail/mailtest"
)

func TestBeadmailConformance(t *testing.T) {
	mailtest.RunProviderTests(t, func(_ *testing.T) mail.Provider {
		return New(beads.NewMemStore())
	})
}
