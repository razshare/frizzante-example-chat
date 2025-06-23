package handlers

import (
	"fmt"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libsession"
	"github.com/razshare/frizzante/libview"
	"main/lib"
)

var messages = []string{}
var connections = map[string]*libcon.Connection{}

func Chat(con *libcon.Connection) {
	state, _ := libsession.Session(con, lib.State{})
	if state.Username == "" {
		con.SendNavigate("/username")
		return
	}
	con.SendView(libview.View{Name: "Chat", Data: map[string]any{
		"username": state.Username,
		"messages": messages,
	}})
}

func ChatMessagesAdd(con *libcon.Connection) {
	state, _ := libsession.Session(con, lib.State{})
	if state.Username == "" {
		con.SendNavigate("/username")
		return
	}
	message := fmt.Sprintf("%s: %s", state.Username, con.ReceiveForm().Get("message"))
	messages = append(messages, message)
	for _, connection := range connections {
		connection.SendMessage(message)
	}
}

func ChatMessagesStream(con *libcon.Connection) {
	state, _ := libsession.Session(con, lib.State{})
	if state.Username == "" {
		con.SendNavigate("/username")
		return
	}

	con.SendSseUpgrade()

	idObject, idError := uuid.NewV4()
	if idError != nil {
		con.SendView(libview.View{Name: "Chat", Data: map[string]any{
			"error": idError.Error(),
		}})
		return
	}

	id := idObject.String()
	connections[id] = con
	<-con.ReceiveCancellation()
	delete(connections, id)
}

func ChatUsernameSet(con *libcon.Connection) {
	state, op := libsession.Session(con, lib.State{})
	defer op.Save(state)
	state.Username = con.ReceiveForm().Get("username")
	con.SendNavigate("/")
}
