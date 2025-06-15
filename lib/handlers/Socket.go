package handlers

import (
	"github.com/razshare/frizzante/frz"
	"main/lib"
)

var connections = map[string]*frz.Connection{}

func Socket(c *frz.Connection) {
	s, _ := frz.Session(c, lib.State{})

	if s.Username == "" {
		c.SendForbidden("username is empty")
		return
	}

	connections[s.Username] = c

	alive := c.IsAlive()
	c.SendWsUpgrade()

	// Read and broadcast.
	for *alive {
		message := c.ReceiveMessage()
		for _, connection := range connections {
			connection.SendMessage(message)
		}
	}
}
