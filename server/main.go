package main

import (
	"context"
	"fmt"
	"io"
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
	emulatorHost := "localhost:4000"
	fireStoreEndpoint := fmt.Sprintf("http://%s", emulatorHost)
	/*
		fmt.Sprintf() returns a string after formatting(similar to Fprintf())
	*/

	conf := &firebase.Config{
		ProjectID:     "testing-todos",
		DatabaseURL:   fireStoreEndpoint,
		StorageBucket: "",
	}
	app, err := firebase.NewApp(ctx, conf, option.WithoutAuthentication())

	if err != nil {
		fmt.Println("Error occured while creating a new app:", err)
		return
	}

	client, client_err := app.Firestore(ctx)
	if client_err != nil {
		fmt.Println("Error while creating a client:", client_err)
		return
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
		fmt.Println("Error while listening and serving =", err)
		return
	}
}
