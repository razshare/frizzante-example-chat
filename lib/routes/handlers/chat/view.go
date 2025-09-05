package chat

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/view"
	"main/lib/routes/handlers/messages"
	"main/lib/session"
)

func View(c *client.Client) {
	s := session.Start(receive.SessionId(c))
	if s.Username == "" {
		send.Navigate(c, "/username")
		return
	}
	send.View(c, view.View{Name: "Chat", Props: map[string]any{
		"username": s.Username,
		"messages": messages.Messages,
	}})
}
