package main

import (
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/cocoagaurav/httpHandler/model"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"strconv"
)

var (
	app    *firebase.App
	client *auth.Client
)

func FirebaseStartAuth() {
	var err error
	conf := &firebase.Config{ServiceAccountID: "firebase-adminsdk-6b9tl@testproject-fa267.iam.gserviceaccount.com"}

	opt := option.WithCredentialsFile("/Users/gaurav/Downloads/testproject-fa267-firebase-adminsdk-6b9tl-341dae30de.json")

	app, err = firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
}

func CreateFireBaseUser(user *model.User) *auth.UserRecord {
	var err error
	client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	params := (&auth.UserToCreate{}).
		DisplayName(user.Name).
		Password(strconv.Itoa(user.Id))
	u, err := client.CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %v\n", u)
	return u
}

func GenerateToken(uid string) string {
	client, _ = app.Auth(context.Background())
	token, _ := client.CustomToken(context.Background(), uid)

	return token

}

func VerifyToken(token string) *auth.Token {
	fmt.Printf("\n varifying token is:%v", token)

	//client, _ = app.Auth(context.Background())

	tok, err := client.VerifyIDToken(context.Background())
	if err != nil {
		fmt.Printf(" \n err is:%v", err)
		return nil
	}
	fmt.Printf("\n return token is:%T", tok)

	fmt.Printf("\n varified token is:%v", tok)

	fmt.Println("label 9")

	return tok

}

func GetUserCreds(authId string) *auth.UserRecord {
	user, err := client.GetUser(context.Background(), authId)
	if err != nil {
		fmt.Printf("error is:", err)
		return nil
	}
	return user
}
