package messages

import (
	"fmt"
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"main/lib/session"
)

func Action(c *client.Client) {
	s := session.Start(receive.SessionId(c))
	if s.Username == "" {
		send.Navigate(c, "/username")
		return
	}

	msg := fmt.Sprintf("%s: %s", s.Username, receive.Form(c).Get("message"))
	Messages = append(Messages, msg)

	for _, cloc := range Clients {
		send.Message(cloc, msg)
	}
}
