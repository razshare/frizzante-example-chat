package main

import (
	"embed"
	"main/lib/core/route"
	"main/lib/core/server"
	"main/lib/core/svelte/ssr"
	"main/lib/routes/handlers/chat"
	"main/lib/routes/handlers/fallback"
	"main/lib/routes/handlers/messages"
	"main/lib/routes/handlers/username"
	"os"
)

//go:embed app/dist
var efs embed.FS
var srv = server.New()
var dev = os.Getenv("DEV") == "1"
var render = ssr.New(ssr.Config{Efs: efs, Disk: dev})

func main() {
	defer server.Start(srv)
	srv.Efs = efs
	srv.Render = render
	srv.Routes = []route.Route{
		{Pattern: "GET /", Handler: fallback.View},
		{Pattern: "GET /chat", Handler: chat.View},
		{Pattern: "GET /username", Handler: username.View},
		{Pattern: "POST /username", Handler: username.Action},
		{Pattern: "GET /messages", Handler: messages.Stream},
		{Pattern: "POST /messages", Handler: messages.Action},
	}
}
