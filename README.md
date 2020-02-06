# dishooks
[![GoDoc](https://godoc.org/github.com/itsTurnip/dishooks?status.svg)](https://godoc.org/github.com/itsTurnip/dishooks)

dishooks is a simple Discord webhook API wrapper.

Example:
```golang
package main 

import "github.com/itsTurnip/dishooks"

func main() {
    webhook, _ = dishooks.WebhookFromURL("https://discordapp.com/api/webhooks/671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX")
    _ = webhook.Get()
    _, _ = webhook.SendMessage(&dishooks.WebhookMessage{
        Content: "Hello world!",
    })
}
```
