package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"cloud.google.com/go/firestore"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form["name"]
	email := r.Form["email"]

	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"name": name,
		"email": email,
	})
	if err != nil {
		log.Fatalf("Failed adding new user: %v", err)
	}
}

func main() {
	ctx := context.Background()

	// [START fs_initialize]
	// Sets your Google Cloud Platform project ID.
	projectID := "golang-gae-boilerplate"

	// Get a Firestore client.
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done.
	defer client.Close()
	// [END fs_initialize]

	http.HandleFunc("/api/user", userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}