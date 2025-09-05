package fallback

import (
	"main/lib/core/client"
	"main/lib/core/send"
	"main/lib/routes/handlers/chat"
)

func View(c *client.Client) {
	send.FileOrElse(c, func() { chat.View(c) })
}
