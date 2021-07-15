package middleware

import (
	"context"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo"
	"google.golang.org/api/option"
)

// 認証用ミドルウェア
func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}
		// firebase SDKのセットアップ
		opt := option.WithCredentialsFile("koolhaas-api-firebase.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error initializing app: %v\n", err)
			return err
		}

		client, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("errors: %v\n", err)
			return err
		}

		auth := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(auth, "Bearer ", "", 1)

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Printf("error verifying ID token: %v\n", err)
			return err
		}

		log.Printf("Verified ID token %v\n", token)
		c.Set("token", token)
		return next(c)
	}
}

func Auth() echo.MiddlewareFunc {
	return auth
}
