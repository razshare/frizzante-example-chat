package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
)

func Username(con *connections.Connection) {
	connections.SendView(con, views.View{Name: "Username"})
}
