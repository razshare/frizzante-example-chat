package handlers

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/sessions"
	"github.com/razshare/frizzante/views"
	"main/lib"
)

var messages = []string{}
var listOfConnections = map[string]*connections.Connection{}

func Chat(con *connections.Connection) {
	session := sessions.StartEmpty[lib.State](con)
	if session.State.Username == "" {
		connections.SendNavigate(con, "/username")
		return
	}
	connections.SendView(con, views.View{Name: "Chat", Data: map[string]any{
		"username": session.State.Username,
		"messages": messages,
	}})
}

func ChatMessagesAdd(con *connections.Connection) {
	session := sessions.StartEmpty[lib.State](con)
	if session.State.Username == "" {
		connections.SendNavigate(con, "/username")
		return
	}
	message := fmt.Sprintf("%s: %s", session.State.Username, connections.ReceiveForm(con).Get("message"))
	messages = append(messages, message)
	for _, conLocal := range listOfConnections {
		connections.SendMessage(conLocal, message)
	}
}

func ChatMessagesStream(con *connections.Connection) {
	session := sessions.StartEmpty[lib.State](con)
	if session.State.Username == "" {
		connections.SendNavigate(con, "/username")
		return
	}

	idObj, idError := uuid.NewV4()
	if idError != nil {
		connections.SendView(con, views.View{Name: "Chat", Data: map[string]any{
			"error": idError.Error(),
		}})
		return
	}

	connections.SendSseUpgrade(con)

	id := idObj.String()
	listOfConnections[id] = con
	<-connections.ReceiveCancellation(con)
	delete(listOfConnections, id)
}

func ChatUsernameSet(con *connections.Connection) {
	session := sessions.StartEmpty[lib.State](con)
	defer sessions.Save(session)
	session.State.Username = connections.ReceiveForm(con).Get("username")
	connections.SendNavigate(con, "/")
}
