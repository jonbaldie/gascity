package main

import (
	"github.com/jonbaldie/gascity/internal/cityinit"
	"github.com/jonbaldie/gascity/internal/fsys"
)

var _ cityinit.ScaffoldFS = fsys.OSScaffoldFS{}
