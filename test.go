package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/yefriddavid/pushNotification/config"
	//"pushNotification/config"
    //"context"
    "path/filepath"
    "google.golang.org/api/option"

)

func main() {

	//app, _, _ := config.SetupFirebase()
	app, _, _ := SetupFirebase()
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
			Body:  "Perro hp!!",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}






//

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {

    ctx := context.Background()

    //serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
    // serviceAccountKeyFilePath, err := filepath.Abs("./google-services.json")
    serviceAccountKeyFilePath, err := filepath.Abs("./notificationsapp-25689-firebase-adminsdk-c2a6h-a861f1eeca.json")

    if err != nil {
        panic("Unable to load serviceAccountKeys.json file")
    }

    opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
fmt.Println("test")

    //Firebase admin SDK initialization
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        panic("Firebase load error")
    }

    //Messaging client
    client, e := app.Messaging(ctx)

fmt.Println(serviceAccountKeyFilePath)
fmt.Println(e)
fmt.Println("test3")

    return app, ctx, client
}
