package main

import (
	"embed"
	"github.com/joho/godotenv"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	"main/lib/handlers"
	"os"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	if err := godotenv.Load(".env"); err != nil {
		server.Notifier.SendError(err)
	} else {
		server.Address = os.Getenv("server.address")
		server.SecureAddress = os.Getenv("server.secure_address")
		server.Key = os.Getenv("server.key")
		server.Certificate = os.Getenv("server.certificate")
		server.PublicRoot = os.Getenv("server.public_root")
		server.AppRoot = os.Getenv("server.app_root")
		server.ServerJs = os.Getenv("server.server_js")
		server.IndexHtml = os.Getenv("server.index_html")
	}

	server.Efs = efs
	server.AddRoute(routes.Route{Pattern: "GET /", Handler: handlers.Default})
	server.AddRoute(routes.Route{Pattern: "GET /chat", Handler: handlers.Chat})
	server.AddRoute(routes.Route{Pattern: "GET /username", Handler: handlers.Username})
	server.AddRoute(routes.Route{Pattern: "GET /chat/messages/stream", Handler: handlers.ChatMessagesStream})
	server.AddRoute(routes.Route{Pattern: "POST /chat/messages/add", Handler: handlers.ChatMessagesAdd})
	server.AddRoute(routes.Route{Pattern: "POST /chat/username/set", Handler: handlers.ChatUsernameSet})
	server.Start()
}
