package chat

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
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
