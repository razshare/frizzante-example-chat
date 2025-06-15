package main

import (
	"embed"
	"github.com/razshare/frizzante/frz"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS

func main() {
	frz.NewServer().
		WithEfs(efs).
		AddRoute(frz.Route{Pattern: "GET /", Handler: handlers.Default}).
		AddRoute(frz.Route{Pattern: "GET /socket", Handler: handlers.Socket}).
		Start()
}
