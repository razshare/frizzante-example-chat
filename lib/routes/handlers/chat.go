package handlers

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
)

var messages = make([]string, 0)
var listOfConnections = map[string]*connections.Connection{}

func Chat(con *connections.Connection) {
	session := sessions.New(con, lib.State{}).Start()
	if session.State.Username == "" {
		con.SendNavigate("/username")
		return
	}
	con.SendView(views.View{Name: "Chat", Data: map[string]any{
		"username": session.State.Username,
		"messages": messages,
	}})
}

func ChatMessagesAdd(con *connections.Connection) {
	session := sessions.New(con, lib.State{}).Start()
	if session.State.Username == "" {
		con.SendNavigate("/username")
		return
	}
	message := fmt.Sprintf("%s: %s", session.State.Username, con.ReceiveForm().Get("message"))
	messages = append(messages, message)
	for _, conLocal := range listOfConnections {
		conLocal.SendMessage(message)
	}
}

func ChatMessagesStream(con *connections.Connection) {
	session := sessions.New(con, lib.State{}).Start()
	if session.State.Username == "" {
		con.SendNavigate("/username")
		return
	}

	idObj, idError := uuid.NewV4()
	if idError != nil {
		con.SendView(views.View{Name: "Chat", Data: map[string]any{
			"error": idError.Error(),
		}})
		return
	}

	con.SendSseUpgrade()

	id := idObj.String()
	listOfConnections[id] = con
	<-con.ReceiveCancellation()
	delete(listOfConnections, id)
}

func ChatUsernameSet(con *connections.Connection) {
	session := sessions.New(con, lib.State{}).Start()
	defer session.Save()
	session.State.Username = con.ReceiveForm().Get("username")
	con.SendNavigate("/")
}
