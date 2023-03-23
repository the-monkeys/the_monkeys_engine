package auth

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func SetupFirebase() (*auth.Client, error) {

	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	fmt.Printf("opt: %v\n", opt)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	//Firebase Auth
	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return auth, nil
}

// AuthMiddleware : to verify all authorized operations
func FirebaseAuthChq(c *gin.Context) *auth.Token {
	firebaseAuth := c.MustGet("firebaseAuth").(*auth.Client)

	authorizationToken := c.GetHeader("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
		c.Abort()
		return nil
	}

	//verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token"})
		c.Abort()
		return nil
	}
	return token
}
