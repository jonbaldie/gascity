package formula

import (
	"bytes"
	"testing"

	"github.com/BurntSushi/toml"
)

// FuzzParseTOMLValidateRoundTrip checks the public formula TOML boundary.
// Parsing can accept structurally incomplete drafts, so validation errors are
// expected input outcomes; every valid formula must survive canonical encoding
// and parsing again.
func FuzzParseTOMLValidateRoundTrip(f *testing.F) {
	f.Add([]byte("formula = \"mol-example\"\n\n[[steps]]\nid = \"work\"\ntitle = \"Do work\"\n"))
	f.Add([]byte("formula = \"exp-example\"\ntype = \"expansion\"\n\n[[template]]\nid = \"work\"\ntitle = \"Do {target}\"\n"))
	f.Add([]byte("formula = \"graph-example\"\ncontract = \"graph.v2\"\n\n[[steps]]\nid = \"work\"\ntitle = \"Do work\"\n"))

	f.Fuzz(func(t *testing.T, data []byte) {
		parsed, err := NewParser().ParseTOML(data)
		if err != nil {
			return
		}
		if err := parsed.Validate(); err != nil {
			return
		}

		var first bytes.Buffer
		if err := toml.NewEncoder(&first).Encode(parsed); err != nil {
			t.Fatalf("valid formula cannot be encoded: %v", err)
		}
		reparsed, err := NewParser().ParseTOML(first.Bytes())
		if err != nil {
			t.Fatalf("encoded formula cannot be parsed: %v\ncanonical:\n%s", err, first.Bytes())
		}
		if err := reparsed.Validate(); err != nil {
			t.Fatalf("encoded formula does not validate: %v\ncanonical:\n%s", err, first.Bytes())
		}
		var second bytes.Buffer
		if err := toml.NewEncoder(&second).Encode(reparsed); err != nil {
			t.Fatalf("reparsed formula cannot be encoded: %v", err)
		}
		if string(second.Bytes()) != string(first.Bytes()) {
			t.Fatalf("formula canonical form is not stable\nfirst:\n%s\nsecond:\n%s", first.Bytes(), second.Bytes())
		}
	})
}
