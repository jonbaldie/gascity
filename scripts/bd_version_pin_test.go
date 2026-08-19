package scripts_test

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/jonbaldie/gascity/internal/deps"
)

// TestBDVersionPins keeps every independently-edited bd version anchor in
// lockstep, the same way TestDoltVersionPins does for Dolt. Before this test the
// bd floors drifted apart: deps.env BD_VERSION, the init hard-dependency floor
// bdMinVersion, the ready-projection feature floor bdReadyProjectionMinVersion,
// and the bd_compatibility config enum were all hand-edited with no cross-check,
// so a regression like #3135 (a 1.0.5 flag emitted ahead of the pinned 1.0.4
// floor) could merge green. This test makes deps.env the single source of truth
// and fails loudly the moment an anchor moves without the others.
func TestBDVersionPins(t *testing.T) {
	root := repoRoot(t)
	env := readDotenv(t, filepath.Join(root, "deps.env"))

	bdRepo := env["BD_REPO"]               // GitHub owner/name this fork installs bd from
	bdVersion := env["BD_VERSION"]         // installable default (v-prefixed release tag)
	bdPrev := env["BD_PREV_VERSION"]       // min-supported matrix cell (downloadable)
	bdCurrent := env["BD_CURRENT_VERSION"] // bleeding-edge matrix cell (built from source)
	bdCurrentRef := env["BD_CURRENT_REF"]  // beads commit the current cell builds from

	if bdRepo != "jonbaldie/beads" {
		t.Fatalf("deps.env BD_REPO = %q, want jonbaldie/beads", bdRepo)
	}

	if bdVersion == "" {
		t.Fatal("deps.env missing BD_VERSION")
	}
	if bdPrev == "" {
		t.Fatal("deps.env missing BD_PREV_VERSION (the minimum-supported contract-matrix cell)")
	}
	if bdCurrent == "" {
		t.Fatal("deps.env missing BD_CURRENT_VERSION (the bleeding-edge contract-matrix cell)")
	}

	// The current cell has no release tarball, so it is built from a pinned beads
	// commit. A non-deterministic ref (branch name, short SHA) would make the cell
	// irreproducible; require a full 40-char commit SHA.
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(bdCurrentRef) {
		t.Fatalf("deps.env BD_CURRENT_REF = %q, want a full 40-char jonbaldie/beads commit SHA", bdCurrentRef)
	}
	if !regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[0-9A-Za-z.-]+)?$`).MatchString(bdCurrent) {
		t.Fatalf("deps.env BD_CURRENT_VERSION = %q, want a semver token", bdCurrent)
	}
	// The native Go store, the bleeding-edge contract-matrix cell, and the
	// source-built agent image must all use the same upstream commit. A drift
	// here can pair one schema catalog with another version's write behavior.
	goMod := readFile(t, root, "go.mod")
	goModMatch := regexp.MustCompile(`(?m)^\s*github\.com/steveyegge/beads\s+v\S+-([0-9a-f]{12})\s*$`).FindStringSubmatch(goMod)
	if goModMatch == nil {
		t.Fatal("go.mod missing a pseudo-version pin for github.com/steveyegge/beads")
	}
	if got, want := goModMatch[1], bdCurrentRef[:12]; got != want {
		t.Fatalf("go.mod beads pseudo-version commit = %q, want BD_CURRENT_REF prefix %q", got, want)
	}
	dockerfile := readFile(t, root, "contrib/k8s/Dockerfile.agent")
	if !strings.Contains(dockerfile, "ARG BD_SOURCE_REF="+bdCurrentRef) {
		t.Fatalf("contrib/k8s/Dockerfile.agent BD_SOURCE_REF must equal deps.env BD_CURRENT_REF (%s)", bdCurrentRef)
	}
	if !strings.Contains(dockerfile, "ARG BD_BUILD="+bdCurrentRef[:10]) {
		t.Fatalf("contrib/k8s/Dockerfile.agent BD_BUILD must equal the first 10 characters of BD_CURRENT_REF (%s)", bdCurrentRef[:10])
	}

	// Anchor roles, kept as distinct contracts so a promotion cannot quietly
	// collapse them:
	//   BD_PREV_VERSION -- the minimum-supported bd (the matrix floor cell).
	//   BD_VERSION      -- the installable default; must be >= the floor.
	// The init hard-dependency floor (bdMinVersion) is the minimum-supported
	// version restated as a Go constant, so it must track BD_PREV_VERSION, not
	// BD_VERSION. Tying it to BD_VERSION would drag the hard floor up the moment
	// BD_VERSION is promoted (e.g. -> v1.0.5) and drop support for the
	// min-supported matrix cell these contract tests exist to keep green.
	bdMin := extractGoStringConst(t, root, "cmd/gc/init_provider_readiness.go", "bdMinVersion")
	if bdMin != strings.TrimPrefix(bdPrev, "v") {
		t.Fatalf("bdMinVersion = %q but deps.env BD_PREV_VERSION = %q (want %q); the init hard floor is the minimum-supported bd and must track BD_PREV_VERSION, not BD_VERSION",
			bdMin, bdPrev, strings.TrimPrefix(bdPrev, "v"))
	}
	// The installable default may move ahead of the floor but never behind it.
	if deps.CompareVersions(bdVersion, bdPrev) < 0 {
		t.Fatalf("deps.env BD_VERSION = %q is older than BD_PREV_VERSION = %q; the installable default must be at least the minimum-supported version",
			bdVersion, bdPrev)
	}

	// The ready-projection feature floor (#3135's regressing surface) must exist
	// and be strictly newer than the init floor, otherwise the gated path is dead
	// for every supported bd. Compare semantically -- the same way the runtime
	// gate in bdstore_ready_projection.go does (deps.CompareVersions) -- so a
	// floor that is merely different from the init floor, including an older one,
	// cannot pass.
	readyFloor := extractGoStringConst(t, root, "internal/beads/bdstore_ready_projection.go", "bdReadyProjectionMinVersion")
	if readyFloor == "" {
		t.Fatal("internal/beads/bdstore_ready_projection.go missing bdReadyProjectionMinVersion const")
	}
	if deps.CompareVersions(readyFloor, bdMin) <= 0 {
		t.Fatalf("bdReadyProjectionMinVersion (%q) must be strictly newer than bdMinVersion (%q); a feature floor at or below the init floor gates nothing", readyFloor, bdMin)
	}

	// The bd_compatibility config enum is the operator-facing mirror of the two
	// floors; both floor values must appear as enum members so they cannot diverge.
	cfg := readFile(t, root, "internal/config/config.go")
	for _, member := range []string{"enum=bd-" + bdMin, "enum=bd-" + readyFloor} {
		if !strings.Contains(cfg, member) {
			t.Fatalf("internal/config/config.go bd_compatibility enum missing %q (floors: init=%s ready=%s)", member, bdMin, readyFloor)
		}
	}

	// This fork publishes no beads GitHub release assets, so CI must install bd
	// from jonbaldie/beads source (`go install` / `go build`), never from an
	// upstream release tarball URL. BD_PREV_VERSION and BD_VERSION remain the
	// git tags the script checks out.
	install := readFile(t, root, ".github/scripts/install-bd-archive.sh")
	if strings.Contains(install, "github.com/gastownhall/beads/releases") ||
		strings.Contains(install, "github.com/steveyegge/beads/releases") {
		t.Fatal(".github/scripts/install-bd-archive.sh still downloads bd from an upstream beads release URL")
	}
	if !strings.Contains(install, "jonbaldie/beads") {
		t.Fatal(".github/scripts/install-bd-archive.sh must install from jonbaldie/beads")
	}
	if !strings.Contains(install, "go build") && !strings.Contains(install, "go install") {
		t.Fatal(".github/scripts/install-bd-archive.sh must build or go-install bd from the fork module")
	}

	// Every workflow that pins BD_VERSION must pin the same value as deps.env, so a
	// bump in one place cannot leave a stale matrix cell behind. Validate every
	// assignment in both .yml and .yaml workflows: a file-level presence check
	// would let a stale pin ride along beside a correct one.
	assertWorkflowPins(t, root, "BD_VERSION", bdVersion)

	// The devcontainer README restates the installed version in prose, which
	// makes it an anchor like any other -- and it was the only one no test read,
	// so it sat at v1.0.4 through the promotion to v1.1.0. A doc anchor nothing
	// asserts is how the next bump goes half-applied.
	assertDocPinAnchor(t, root, ".devcontainer/README.md", "BD_VERSION", bdVersion)
}

// assertDocPinAnchor fails when a doc restates a deps.env pin as
// "`KEY` from `deps.env` (currently VALUE)" and VALUE has drifted. The phrasing
// is the contract: prose that names the variable without restating the value is
// not an anchor and is not matched, so a doc can always opt out by dropping the
// parenthetical rather than by going stale.
func assertDocPinAnchor(t *testing.T, root, rel, key, want string) {
	t.Helper()
	re := regexp.MustCompile("`" + regexp.QuoteMeta(key) + "` from `deps\\.env` \\(currently ([^)]+)\\)")
	m := re.FindStringSubmatch(readFile(t, root, rel))
	if m == nil {
		t.Fatalf("%s no longer restates %s as \"`%s` from `deps.env` (currently <value>)\"; either restore that phrasing or drop this assertion with the anchor", rel, key, key)
	}
	if got := strings.TrimSpace(m[1]); got != want {
		t.Errorf("%s says %s is currently %q, want %q (deps.env)", rel, key, got, want)
	}
}

// TestScanPinAssignments proves the workflow pin scanner catches the partial
// drift a file-level presence check missed: a stale BD_VERSION sharing a file
// with a correct one is still reported with its line, while a
// `${{ env.BD_VERSION }}` reference is not treated as an assignment.
func TestScanPinAssignments(t *testing.T) {
	const fixture = `env:
  BD_VERSION: "v1.0.4"
  DOLT_VERSION: "2.1.7"
jobs:
  stale:
    env:
      BD_VERSION: "v1.0.3"
    steps:
      - with:
          bd-version: ${{ env.BD_VERSION }}
`
	got := scanPinAssignments("BD_VERSION", fixture)
	want := []pinAssignment{
		{line: 2, value: "v1.0.4"},
		{line: 7, value: "v1.0.3"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("scanPinAssignments(BD_VERSION) = %+v, want %+v", got, want)
	}
}

// readDotenv parses simple KEY=VALUE lines, ignoring comments and blanks.
func readDotenv(t *testing.T, path string) map[string]string {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	out := map[string]string{}
	for _, line := range strings.Split(string(content), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		k, v, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		out[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
	return out
}

func readFile(t *testing.T, root, rel string) string {
	t.Helper()
	content, err := os.ReadFile(filepath.Join(root, rel))
	if err != nil {
		t.Fatalf("read %s: %v", rel, err)
	}
	return string(content)
}

// extractGoStringConst returns the value of a `name = "..."` Go string constant,
// or "" if the file does not declare it. The pattern is anchored to a real
// declaration form -- the identifier must start a line, optionally preceded by
// indentation and the `const` keyword -- so a comment or prose example naming the
// same identifier above the real const cannot be matched first.
func extractGoStringConst(t *testing.T, root, rel, name string) string {
	t.Helper()
	re := regexp.MustCompile(`(?m)^\s*(?:const\s+)?` + regexp.QuoteMeta(name) + `\s*=\s*"([^"]+)"`)
	m := re.FindStringSubmatch(readFile(t, root, rel))
	if m == nil {
		return ""
	}
	return m[1]
}

// pinAssignment is a single `KEY: value` mapping entry found in a workflow file,
// carrying its 1-based line number for diagnostics.
type pinAssignment struct {
	line  int
	value string
}

// scanPinAssignments returns every `key: value` assignment in content. It matches
// only a mapping key -- optional indentation, the exact key, then a colon -- so a
// reference such as `bd-version: ${{ env.BD_VERSION }}` is not mistaken for an
// assignment of BD_VERSION. Surrounding quotes and any trailing comment are
// stripped from the captured value.
func scanPinAssignments(key, content string) []pinAssignment {
	re := regexp.MustCompile(`^\s*` + regexp.QuoteMeta(key) + `:\s*["']?([^"'\s#]+)["']?`)
	var out []pinAssignment
	for i, line := range strings.Split(content, "\n") {
		if m := re.FindStringSubmatch(line); m != nil {
			out = append(out, pinAssignment{line: i + 1, value: m[1]})
		}
	}
	return out
}

// assertWorkflowPins fails for every workflow assignment of key whose value is not
// want, scanning both .yml and .yaml workflows and reporting each offending file
// and line. Validating every assignment -- not just file-level presence -- catches
// a file that mixes a correct pin with a stale one, and reporting via t.Errorf
// rather than t.Fatalf surfaces all stale pins in a single run.
func assertWorkflowPins(t *testing.T, root, key, want string) {
	t.Helper()
	dir := filepath.Join(root, ".github", "workflows")
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if ext := filepath.Ext(path); ext != ".yml" && ext != ".yaml" {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(root, path)
		for _, a := range scanPinAssignments(key, string(content)) {
			if a.value != want {
				t.Errorf("%s:%d pins %s to %q, want %q (deps.env)", rel, a.line, key, a.value, want)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk workflows: %v", err)
	}
}
