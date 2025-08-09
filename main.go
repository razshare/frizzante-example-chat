package main

import (
	"embed"
	"github.com/razshare/frizzante/route"
	"github.com/razshare/frizzante/server"
	"main/lib/routes/handlers/chat"
	"main/lib/routes/handlers/fallback"
	"main/lib/routes/handlers/messages"
	"main/lib/routes/handlers/username"
)

//go:embed app/dist
var efs embed.FS
var conf = server.Default()

func main() {
	defer server.Start(conf)
	conf.Container.Efs = efs
	conf.Routes = []route.Route{
		{Pattern: "GET /", Handler: fallback.View},
		{Pattern: "GET /chat", Handler: chat.View},
		{Pattern: "GET /username", Handler: username.View},
		{Pattern: "POST /username", Handler: username.Action},
		{Pattern: "GET /messages", Handler: messages.Stream},
		{Pattern: "POST /messages", Handler: messages.Action},
	}
}
