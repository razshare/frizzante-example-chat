package username

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/session"
)

func Action(c *client.Client) {
	s := session.Start(receive.SessionId(c))

	s.Username = receive.Form(c).Get("username")

	send.Navigate(c, "/")
}
