package config

import "testing"

// FuzzParseMarshalRoundTrip checks the public TOML configuration boundary.
// Invalid TOML and configurations rejected by validation are expected input
// outcomes; every accepted configuration must have a stable canonical form.
func FuzzParseMarshalRoundTrip(f *testing.F) {
	f.Add([]byte("[workspace]\nname = \"city\"\n"))
	f.Add([]byte("[workspace]\nname = \"city\"\n\n[[agent]]\nname = \"worker\"\n"))
	f.Add([]byte("[daemon]\nformula_v2 = true\n\n[workspace]\nname = \"city\"\n"))

	f.Fuzz(func(t *testing.T, data []byte) {
		cfg, err := Parse(data)
		if err != nil {
			return
		}

		canonical, err := cfg.Marshal()
		if err != nil {
			t.Fatalf("accepted config cannot be marshaled: %v", err)
		}
		reparsed, err := Parse(canonical)
		if err != nil {
			t.Fatalf("canonical config cannot be parsed: %v\ncanonical:\n%s", err, canonical)
		}
		recanonical, err := reparsed.Marshal()
		if err != nil {
			t.Fatalf("reparsed config cannot be marshaled: %v", err)
		}
		if string(recanonical) != string(canonical) {
			t.Fatalf("config canonical form is not stable\nfirst:\n%s\nsecond:\n%s", canonical, recanonical)
		}
	})
}
