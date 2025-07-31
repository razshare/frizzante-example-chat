package main

import (
	"embed"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	handlers2 "main/lib/routes/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	server.Efs = efs

	server.Routes = []routes.Route{
		{Pattern: "GET /", Handler: handlers2.Default},
		{Pattern: "GET /chat", Handler: handlers2.Chat},
		{Pattern: "GET /username", Handler: handlers2.Username},
		{Pattern: "GET /chat/messages/stream", Handler: handlers2.ChatMessagesStream},
		{Pattern: "POST /chat/messages/add", Handler: handlers2.ChatMessagesAdd},
		{Pattern: "POST /chat/username/set", Handler: handlers2.ChatUsernameSet},
	}

	server.Start()
}
