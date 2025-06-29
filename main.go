package main

import (
	"embed"
	"github.com/razshare/frizzante/web"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = web.NewServer()

func main() {
	server.Efs = efs
	server.AddRoute(web.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(web.Route{Pattern: "GET /chat", Handler: handlers.Chat})
	server.AddRoute(web.Route{Pattern: "GET /username", Handler: handlers.Username})
	server.AddRoute(web.Route{Pattern: "GET /chat/messages/stream", Handler: handlers.ChatMessagesStream})
	server.AddRoute(web.Route{Pattern: "POST /chat/messages/add", Handler: handlers.ChatMessagesAdd})
	server.AddRoute(web.Route{Pattern: "POST /chat/username/set", Handler: handlers.ChatUsernameSet})
	server.Start()
}
