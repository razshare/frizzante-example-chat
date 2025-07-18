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
	servers.AddRoute(server, routes.Route{Pattern: "GET /", Handler: handlers.Default})
	servers.AddRoute(server, routes.Route{Pattern: "GET /chat", Handler: handlers.Chat})
	servers.AddRoute(server, routes.Route{Pattern: "GET /username", Handler: handlers.Username})
	servers.AddRoute(server, routes.Route{Pattern: "GET /chat/messages/stream", Handler: handlers.ChatMessagesStream})
	servers.AddRoute(server, routes.Route{Pattern: "POST /chat/messages/add", Handler: handlers.ChatMessagesAdd})
	servers.AddRoute(server, routes.Route{Pattern: "POST /chat/username/set", Handler: handlers.ChatUsernameSet})
	servers.Start(server)
}
