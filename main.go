package main

import (
	"net/http"

	handlers "github.com/sgirones/fb-messenger-image-day/handlers"
	"github.com/paked/configure"
	"github.com/paked/messenger"
)

var (
	conf        = configure.New()
	verifyToken = conf.String("verify-token", "mad-skrilla", "The token used to verify facebook")
	verify      = conf.Bool("should-verify", false, "Whether or not the app should verify itself")
	pageToken   = conf.String("page-token", "not skrilla", "The token that is used to verify the page on facebook")
)


func handlersSetup(client *messenger.Messenger) {
	client.HandleMessage(handlers.MessageReceived)
	client.HandleDelivery(handlers.MessageDelivered)
}

func main() {
	conf.Use(configure.NewFlag())
	conf.Use(configure.NewEnvironment())
	conf.Use(configure.NewJSONFromFile("config.json"))
	conf.Parse()

	client := messenger.New(messenger.Options{
		Verify:      *verify,
		VerifyToken: *verifyToken,
		Token:       *pageToken,
	})

	handlersSetup(client)

	//start the server
	http.ListenAndServe(":5646", client.Handler())

}