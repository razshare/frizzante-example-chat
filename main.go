package main

import (
	"embed"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	server.Efs = efs
	server.AddRoute(routes.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(routes.Route{Pattern: "GET /chat", Handler: handlers.Chat})
	server.AddRoute(routes.Route{Pattern: "GET /username", Handler: handlers.Username})
	server.AddRoute(routes.Route{Pattern: "GET /chat/messages/stream", Handler: handlers.ChatMessagesStream})
	server.AddRoute(routes.Route{Pattern: "POST /chat/messages/add", Handler: handlers.ChatMessagesAdd})
	server.AddRoute(routes.Route{Pattern: "POST /chat/username/set", Handler: handlers.ChatUsernameSet})
	server.Start()
}
