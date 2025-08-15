package messages

import (
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/receive"
	"github.com/razshare/frizzante/send"
	"github.com/razshare/frizzante/view"
	"main/lib/session"
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
