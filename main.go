package main

import (
	"embed"
	"github.com/razshare/frizzante/environments"
	"github.com/razshare/frizzante/routes"
	"github.com/razshare/frizzante/servers"
	"github.com/razshare/frizzante/traces"
	"main/lib/handlers"
	"os"
)

//go:embed app/dist
var efs embed.FS
var server = servers.New()

func main() {
	server.Efs = efs

	if err := environments.LoadDotenv(".env"); err != nil {
		traces.Trace(server.ErrorLog, err)
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
