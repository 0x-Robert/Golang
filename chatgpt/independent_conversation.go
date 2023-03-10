package main

import (
	"log"
	"net/http"
	"time"

	chatgpt "github.com/chatgp/chatgpt-go"
)

func main() {
	token := `copy-from-cookies`
	cfValue := "copy-from-cookies"
	puid := "copy-from-cookies"

	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: token,
		},
		{
			Name:  "cf_clearance",
			Value: cfValue,
		},
		{
			Name:  "_puid",
			Value: puid,
		},
	}

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
	)

	// chat in independent conversation
	message := "say hello to me"
	text, err := cli.GetChatText(message)
	if err != nil {
		log.Fatalf("get chat text failed: %v", err)
	}

	log.Printf("q: %s, a: %s\n", message, text.Content)
}
