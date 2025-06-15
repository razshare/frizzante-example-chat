package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
)

func ChatRoom(c *frz.Connection) {
	s, o := frz.Session(c, lib.State{})
	o.Save(s)

	c.SendView(frz.View{Name: "ChatRoom", Data: map[string]any{
		"username": s.Username,
	}})
}
