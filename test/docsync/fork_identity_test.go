package docsync

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	forkModule     = "github.com/jonbaldie/gascity"
	forkGoInstall  = "go install github.com/jonbaldie/gascity/cmd/gc@latest"
	forkCloneURL   = "https://github.com/jonbaldie/gascity.git"
	upstreamOrg    = "gastownhall"
	upstreamModule = "github.com/" + upstreamOrg + "/gascity"
)

func TestForkModulePath(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(repoRoot(), "go.mod"))
	if err != nil {
		t.Fatalf("read go.mod: %v", err)
	}
	first := strings.SplitN(string(data), "\n", 2)[0]
	want := "module " + forkModule
	if first != want {
		t.Fatalf("go.mod module line = %q, want %q", first, want)
	}
}

func TestForkInstallDocsRecommendGoInstall(t *testing.T) {
	files := []string{
		"README.md",
		"docs/getting-started/installation.md",
	}
	for _, rel := range files {
		body := mustReadRepoFile(t, rel)
		if !strings.Contains(body, forkGoInstall) {
			t.Errorf("%s does not recommend %q as the install path", rel, forkGoInstall)
		}
		if strings.Contains(body, "brew install gascity") {
			t.Errorf("%s still recommends `brew install gascity` as a way to install this fork", rel)
		}
		if strings.Contains(body, "This taps the `gastownhall/gascity`") ||
			strings.Contains(body, "brew install gastownhall/gascity") ||
			strings.Contains(body, "brew untap gastownhall/gascity") {
			t.Errorf("%s still recommends the upstream Homebrew tap gastownhall/gascity", rel)
		}
		if strings.Contains(body, "github.com/"+upstreamOrg+"/gascity/releases/download") {
			t.Errorf("%s still recommends upstream release tarballs as the way to install this fork", rel)
		}
		if strings.Contains(body, "gh attestation") && strings.Contains(body, "--repo gastownhall/gascity") {
			t.Errorf("%s still recommends attesting against gastownhall/gascity as the way to install this fork", rel)
		}
	}

	install := mustReadRepoFile(t, "docs/getting-started/installation.md")
	if !strings.Contains(install, "git clone "+forkCloneURL) {
		t.Errorf("installation.md clone URL is not %q", forkCloneURL)
	}
}

func TestCodingStandardsUseForkModulePath(t *testing.T) {
	body := mustReadRepoFile(t, "CODING_STANDARDS.md")
	if strings.Contains(body, "Keep the module path `"+upstreamModule+"`") {
		t.Fatal("CODING_STANDARDS.md still tells contributors to keep the upstream module path")
	}
	if !strings.Contains(body, "Keep the module path `"+forkModule+"`") {
		t.Fatalf("CODING_STANDARDS.md does not tell contributors to keep %s", forkModule)
	}
}

func TestUnreleasedChangelogCompareUsesFork(t *testing.T) {
	body := mustReadRepoFile(t, "CHANGELOG.md")
	want := "[Unreleased]: https://github.com/jonbaldie/gascity/compare/"
	if !strings.Contains(body, want) {
		t.Fatalf("CHANGELOG.md unreleased compare link does not point at this fork")
	}
}

func mustReadRepoFile(t *testing.T, rel string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(repoRoot(), rel))
	if err != nil {
		t.Fatalf("read %s: %v", rel, err)
	}
	return string(data)
}
