package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func handler(w http.ResponseWriter, _ *http.Request) {
	_, err := io.WriteString(w, "Welcome to the basic Http server")
	if err != nil {
		fmt.Println("error", err)
		return
	}
}

func main() {
	fmt.Println("Hello World this is a web server in go")

	ctx := context.Background()

	// setting up local firestore emulator

	conf := &firebase.Config{ProjectID: "testing-todos-bd80f"}
	opts := []option.ClientOption{
		option.WithEndpoint("localhost:8080"),
		option.WithoutAuthentication(),
	}
	app, err := firebase.NewApp(ctx, conf, opts...)

	if err != nil {
		fmt.Println("Error while creating a new firebase app:", err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	_, _, addErr := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Gourab",
		"last":  "Das",
		"born":  2004,
	})

	if addErr != nil {
		fmt.Println("Error occured:", addErr)
	}

	http.HandleFunc("/", handler)
	fmt.Println("Listening to port 5000")
	er := http.ListenAndServe(":5000", nil)
	if er != nil {
		fmt.Println("Error while listening and serving =", er)
		return
	}
}
