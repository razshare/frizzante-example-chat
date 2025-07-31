package handlers

import (
	"github.com/razshare/frizzante/connections"
	"github.com/razshare/frizzante/views"
)

func Username(con *connections.Connection) {
	con.SendView(views.View{Name: "Username"})
}
