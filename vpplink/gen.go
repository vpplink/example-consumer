package vpplink

import (
	_ "github.com/vpplink/vpplink-gen"
)

//go:generate go run github.com/vpplink/vpplink-gen/cmd --binapi-package "go.fd.io/govpp/binapi" --output-dir ./impl
