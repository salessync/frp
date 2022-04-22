package frpc

import (
	"embed"

	"github.com/salessync/frp/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
