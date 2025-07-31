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

	server.Routes = []routes.Route{
		{Pattern: "GET /", Handler: handlers.Default},
		{Pattern: "GET /chat", Handler: handlers.Chat},
		{Pattern: "GET /username", Handler: handlers.Username},
		{Pattern: "GET /chat/messages/stream", Handler: handlers.ChatMessagesStream},
		{Pattern: "POST /chat/messages/add", Handler: handlers.ChatMessagesAdd},
		{Pattern: "POST /chat/username/set", Handler: handlers.ChatUsernameSet},
	}

	server.Start()
}
