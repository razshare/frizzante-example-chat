package handlers

import "github.com/razshare/frizzante/frz"

func Username(c *frz.Connection) {
	c.SendView(frz.View{Name: "Username"})
}
