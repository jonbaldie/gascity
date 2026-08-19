#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat >&2 <<'USAGE'
Usage: install-bd-archive.sh VERSION [--cache]

Builds bd from jonbaldie/beads at VERSION (a git tag or commit) and installs it.
This fork publishes no GitHub release assets, so the install path is source
build / go install rather than an upstream release tarball.

Use --cache on self-hosted runners to install under RUNNER_TOOL_CACHE/HOME
and add that bin directory to GITHUB_PATH.
USAGE
}

version="${1:-}"
if [[ -z "$version" ]]; then
  usage
  exit 2
fi
shift || true

use_cache=false
while (($#)); do
  case "$1" in
    --cache) use_cache=true ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "Unknown argument: $1" >&2
      usage
      exit 2
      ;;
  esac
  shift
done

case "$(uname -s)" in
  Darwin) os=darwin ;;
  Linux) os=linux ;;
  *)
    echo "Unsupported OS: $(uname -s)" >&2
    exit 1
    ;;
esac

case "$(uname -m)" in
  arm64|aarch64) arch=arm64 ;;
  x86_64|amd64) arch=amd64 ;;
  *)
    echo "Unsupported architecture: $(uname -m)" >&2
    exit 1
    ;;
esac

platform_tuple="${os}_${arch}"
repo="${BD_REPO:-jonbaldie/beads}"

install_binary() {
  local src="$1"
  local dst="$2"
  mkdir -p "$(dirname "$dst")"
  install -m 0755 "$src" "$dst"
}

install_binary_with_sudo_fallback() {
  local src="$1"
  local dst="$2"
  local dst_dir
  dst_dir="$(dirname "$dst")"
  mkdir -p "$dst_dir"
  if [[ -w "$dst_dir" ]]; then
    install_binary "$src" "$dst"
  elif command -v sudo >/dev/null 2>&1; then
    sudo install -m 0755 "$src" "$dst"
  else
    echo "Cannot write $dst and sudo is unavailable" >&2
    exit 1
  fi
}

if $use_cache; then
  cache_root="${RUNNER_TOOL_CACHE:-$HOME/.local}"
  bin_dir="${cache_root}/gascity-bd/${version}/${platform_tuple}/bin"
else
  bin_dir="${BD_INSTALL_BIN_DIR:-/usr/local/bin}"
fi

target="${bin_dir}/bd"
if [[ -x "$target" ]]; then
  echo "Reusing cached bd ${version} at ${target}"
else
  if ! command -v git >/dev/null 2>&1; then
    echo "git is required to install bd from ${repo}" >&2
    exit 1
  fi
  if ! command -v go >/dev/null 2>&1; then
    echo "go is required to install bd from ${repo}" >&2
    exit 1
  fi
  tmp="$(mktemp -d)"
  trap 'rm -rf "$tmp"' EXIT
  src="${tmp}/beads-src"
  GIT_TERMINAL_PROMPT=0 git clone --depth 1 --filter=blob:none --branch "$version" \
    "https://github.com/${repo}" "$src"
  echo "Building bd ${version} from ${repo} with go build -tags gms_pure_go"
  go -C "$src" build -tags gms_pure_go -o "${tmp}/bd" ./cmd/bd
  if $use_cache; then
    install_binary "${tmp}/bd" "$target"
  else
    install_binary_with_sudo_fallback "${tmp}/bd" "$target"
  fi
fi

if $use_cache && [[ -n "${GITHUB_PATH:-}" ]]; then
  echo "$bin_dir" >> "$GITHUB_PATH"
fi

"$target" version
