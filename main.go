package main

import (
	"embed"
	"github.com/razshare/frizzante/libsrv"
	"main/lib/handlers"
)

//go:embed app/dist
var efs embed.FS
var server = libsrv.NewServer()

func main() {
	server.Efs = efs
	server.AddRoute(libsrv.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(libsrv.Route{Pattern: "GET /chat", Handler: handlers.Chat})
	server.AddRoute(libsrv.Route{Pattern: "GET /username", Handler: handlers.Username})
	server.AddRoute(libsrv.Route{Pattern: "GET /chat/messages/stream", Handler: handlers.ChatMessagesStream})
	server.AddRoute(libsrv.Route{Pattern: "POST /chat/messages/add", Handler: handlers.ChatMessagesAdd})
	server.AddRoute(libsrv.Route{Pattern: "POST /chat/username/set", Handler: handlers.ChatUsernameSet})
	server.Start()
}
