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
		AddRoute(frz.Route{Pattern: "GET /chat", Handler: handlers.Chat}).
		AddRoute(frz.Route{Pattern: "GET /username", Handler: handlers.Username}).
		AddRoute(frz.Route{Pattern: "GET /chat/messages/stream", Handler: handlers.ChatMessagesStream}).
		AddRoute(frz.Route{Pattern: "POST /chat/messages/add", Handler: handlers.ChatMessagesAdd}).
		AddRoute(frz.Route{Pattern: "POST /chat/username/set", Handler: handlers.ChatUsernameSet}).
		Start()
}
