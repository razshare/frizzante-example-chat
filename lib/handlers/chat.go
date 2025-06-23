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
	state, _ := sessions.Start[lib.State](con)
	if state.Username == "" {
		con.SendNavigate("/username")
		return
	}
	con.SendView(views.View{Name: "Chat", Data: map[string]any{
		"username": state.Username,
		"messages": messages,
	}})
}

func ChatMessagesAdd(con *connections.Connection) {
	state, _ := sessions.Start[lib.State](con)
	if state.Username == "" {
		con.SendNavigate("/username")
		return
	}
	message := fmt.Sprintf("%s: %s", state.Username, con.ReceiveForm().Get("message"))
	messages = append(messages, message)
	for _, conLocal := range listOfConnections {
		conLocal.SendMessage(message)
	}
}

func ChatMessagesStream(con *connections.Connection) {
	state, _ := sessions.Start[lib.State](con)
	if state.Username == "" {
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
	state, operator := sessions.Start[lib.State](con)
	defer operator.Save(state)
	state.Username = con.ReceiveForm().Get("username")
	con.SendNavigate("/")
}
