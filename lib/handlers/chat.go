package handlers

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/frz"
	"main/lib"
)

var messages = []string{}
var connections = map[string]*frz.Connection{}

func Chat(c *frz.Connection) {
	s, _ := frz.Session(c, lib.State{})
	if s.Username == "" {
		c.SendNavigate("/username")
		return
	}
	c.SendView(frz.View{Name: "Chat", Data: map[string]any{
		"username": s.Username,
		"messages": messages,
	}})
}

func ChatMessagesAdd(c *frz.Connection) {
	s, _ := frz.Session(c, lib.State{})
	if s.Username == "" {
		c.SendNavigate("/username")
		return
	}
	message := fmt.Sprintf("%s: %s", s.Username, c.ReceiveForm().Get("message"))
	messages = append(messages, message)
	for _, connection := range connections {
		connection.SendMessage(message)
	}
}

func ChatMessagesStream(c *frz.Connection) {
	s, _ := frz.Session(c, lib.State{})
	if s.Username == "" {
		c.SendNavigate("/username")
		return
	}

	c.SendSseUpgrade()

	idObject, idError := uuid.NewV4()
	if idError != nil {
		c.SendView(frz.View{Name: "Chat", Data: map[string]any{
			"error": idError.Error(),
		}})
		return
	}

	id := idObject.String()
	connections[id] = c
	<-c.ReceiveCancellation()
	delete(connections, id)
}

func ChatUsernameSet(c *frz.Connection) {
	s, o := frz.Session(c, lib.State{})
	defer o.Save(s)
	s.Username = c.ReceiveForm().Get("username")
	c.SendNavigate("/")
}
