package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/yefriddavid/pushNotification/config"
	//"pushNotification/config"
)

func main() {

	app, _, _ := config.SetupFirebase()
	sendToToken(app)
	}

func sendToToken(app *firebase.App) {

	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	//registrationToken := "d3adVu0tMDyBUTPkgc_l-0:APA91bHjb6-wWkT1ABGSasFqxrsOR3AdfcTjLc8b7f7yukWLt32GS4UA5XdIwZ8p98oOLp-CBcyuYaCYdEPRji_f2WSXO9JKb7XPjotm_3bdkk-7hJyxJS8JuUHt82xzGGJ6Aacy0QWb"
	registrationToken := "dHD7ZvywQX-rvrRWid33FH:APA91bHwypgYEwrnxK2IUpipe36pNVsmV3FD0yJ1U7XswDBoKXX6lWLEnP6LTynOH9W6cYoRHrot-Juck7JPwfpyfSFLFBpz7wxmDqjXXkH_7lGt_zm_cvYWj4DcF5j3whm2hKgK02yU"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Notification Test",
			Body:  "Hello React!!",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}


