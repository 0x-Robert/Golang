package main

import (
	"log"
	"net/http"
	"time"

	chatgpt "github.com/chatgp/chatgpt-go"
)

func main() {
	// new chatgpt client
	token := `copy-from-cookies`
	cfValue := "copy-from-cookies"

	cookies := []*http.Cookie{
		{
			Name:  "__Secure-next-auth.session-token",
			Value: token,
		},
		{
			Name:  "cf_clearance",
			Value: cfValue,
		},
	}

	cli := chatgpt.NewClient(
		chatgpt.WithDebug(true),
		chatgpt.WithTimeout(60*time.Second),
		chatgpt.WithCookies(cookies),
	)

	message := "say hello to me"
	stream, err := cli.GetChatStream(message)
	if err != nil {
		log.Fatalf("get chat stream failed: %v\n", err)
	}

	var answer string
	for text := range stream.Stream {
		log.Printf("stream text: %s\n", text.Content)
		answer = text.Content
	}

	if stream.Err != nil {
		log.Fatalf("stream closed with error: %v\n", stream.Err)
	}

	log.Printf("q: %s, a: %s\n", message, answer)
}
