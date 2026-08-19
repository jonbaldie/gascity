package main

import (
	"github.com/jonbaldie/gascity/internal/pathutil"
)

func normalizePathForCompare(path string) string {
	return pathutil.NormalizePathForCompare(path)
}

func samePath(a, b string) bool {
	return pathutil.SamePath(a, b)
}
