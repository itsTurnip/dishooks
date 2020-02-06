package dishooks

import (
	"testing"
)

var testURL = "https://discordapp.com/api/webhooks/671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX"

func TestWebhookFromURLString(t *testing.T) {
	webhook := &Webhook{
		ID:    "671422873239289884",
		Token: "G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX",
	}
	test, err := WebhookFromURL(testURL)
	if err != nil {
		t.Error("Got unexpected error ", err)
	}
	switch {
	case test.Token != webhook.Token:
		t.Errorf("Expected %s, got %s", webhook.Token, test.Token)
	case test.ID != webhook.ID:
		t.Errorf("Expected %s, got %s", webhook.ID, test.ID)
	}

}
func TestWebhookFromURLStringWithWrongURL(t *testing.T) {
	testData := map[string]string{
		"Another website":      "https://google.com/api/webhooks/671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX",
		"Discord broken url 1": "https://discordapp.com/api/webhooks/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX",
		"Discord broken url 2": "https://discordapp.com/api//671422873239289884/G0ArWEr3hgJ1I4THBIiwxkIbnGkHTGawikf3Z7be3afsZD-hCH-hNwWxU0rQAe3U7nkX",
	}
	for test, link := range testData {
		t.Run(test, func(t *testing.T) {
			webhook, err := WebhookFromURL(link)
			if webhook != nil || err == nil {
				t.Error("Wanted nothing, got webhook")
			}
		})
	}
}
func TestWebhookURL(t *testing.T) {
	webhook, _ := WebhookFromURL(testURL)
	if webhook.URL() != testURL {
		t.Fail()
	}
}
