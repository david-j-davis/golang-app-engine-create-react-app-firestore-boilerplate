package main

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

const firestoreAccountFile = "firebase.json"
const firestoreProjectID = "golang-gae-firestore-template"

type formData struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func main() {
	// Gin init
	r := gin.Default()
	// Serve from static build directory
	r.StaticFS("/", http.Dir("./build"))
	// routes
	r.POST("/api/user", userHandler)
	// run application on port 8080
	r.Run(":8080")
}

func writeLogIfError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func getNewFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	return firestore.NewClient(ctx, firestoreProjectID, option.WithServiceAccountFile(firestoreAccountFile))
}

// User enters the form fields to create a new user, which is then populated in our Firestore db.
func userHandler(c *gin.Context) {
	ctx := context.Background()

	client, err := getNewFirestoreClient(ctx)
	writeLogIfError(c, err)
	defer client.Close()
	// Get form data
	var form formData
	c.BindJSON(&form)

	// [START add user to firestore]
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"name":  form.Name,
		"email": form.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// [END add user to firestore]
	c.JSON(http.StatusOK, gin.H{"status": "user added to db"})
}
