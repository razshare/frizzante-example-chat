package messages

import (
	"main/lib/core/client"
	"main/lib/core/receive"
	"main/lib/core/send"
	"main/lib/core/view"
	"main/lib/session"

	uuid "github.com/nu7hatch/gouuid"
)

func Stream(c *client.Client) {
	s := session.Start(receive.SessionId(c))

	if s.Username == "" {
		send.Navigate(c, "/username")
		return
	}

	id, err := uuid.NewV4()
	if err != nil {
		send.View(c, view.View{Name: "Chat", Props: map[string]any{
			"error": err.Error(),
		}})
		return
	}

	send.SseUpgrade(c)

	Clients[id.String()] = c

	<-receive.Cancellation(c)

	delete(Clients, id.String())
}
