package handlers

import (
	"github.com/razshare/frizzante/libcon"
	"github.com/razshare/frizzante/libview"
)

func Username(con *libcon.Connection) {
	con.SendView(libview.View{Name: "Username"})
}
