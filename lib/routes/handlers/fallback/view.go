package fallback

import (
	"github.com/razshare/frizzante/client"
	"github.com/razshare/frizzante/send"
	"main/lib/routes/handlers/chat"
)

func View(c *client.Client) {
	send.FileOrElse(c, func() { chat.View(c) })
}
