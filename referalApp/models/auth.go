package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sahilwahi7/referalApp/connection"
	"github.com/sahilwahi7/referalApp/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoggedInUser struct {
	UserName string `bson:"userName"`
	Password string `bson:"password"`
}

type ConcreteUser struct {
}

type UpdatedUser struct {
	UserName       string `json:"userName"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	LinkedProfile  string `json:"linkedinProfile`
	CurrentCompany string `json:"company"`
	Name           string `json:"name"`
	IsRefree       bool   `json:"isRefree`
}

type User interface {
	CheckUser(username string, password string) bool
	Authenticate(name string, username string, password string, id string, isrefree bool) bool
	FetchAll() *[]database.User
	CheckDetails(username string, linkedProfile string, currentCompany string, description string, title string) bool
}

func checkCollection(client *mongo.Client, collectionName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collectionNames, err := client.Database("User").ListCollectionNames(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range collectionNames {

		if name == collectionName {
			collectionExists := true

			return collectionExists

		}
	}

	return false
}

func createCollection(ctx context.Context, db *mongo.Database, collectionName string) error {
	opts := options.CreateCollection()
	err := db.CreateCollection(ctx, collectionName, opts)
	if err != nil {
		return err
	}
	return nil
}

func (u *ConcreteUser) CheckDetails(username string, linkedProfile string, currentCompany string, description string, title string) bool {
	c := &connection.ConcreteConnection{}
	client := c.Connect()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Printf("Bye from Referal DB")
	}()
	var result *database.User
	collection := client.Database("User").Collection("Refree")
	filter := bson.D{{"username", username}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No matching document found, please singup")
			return false
		} else {
			panic(err)
		}
	} else {
		collectionName := "User"
		collectionExists := checkCollection(client, collectionName)

		if collectionExists != true {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			err := createCollection(ctx, client.Database("User"), "User")
			if err != nil {
				log.Fatal(err)
				return false
			}
			fmt.Println("Collection 'User' created.")

		}

		collection := client.Database("User").Collection("User")
		updatedUser := &UpdatedUser{
			UserName:       username,
			Title:          title,
			Description:    description,
			LinkedProfile:  linkedProfile,
			CurrentCompany: currentCompany,
			Name:           result.Name,
			IsRefree:       result.IsRefree,
		}
		insertResult, err := collection.InsertOne(context.TODO(), updatedUser)
		if err != nil {
			panic(err)

		}
		insertedID := insertResult.InsertedID
		fmt.Println("Inserted document ID:", insertedID)
		return true
	}

}

func (u *ConcreteUser) CheckUser(username string, password string) bool {
	c := &connection.ConcreteConnection{}
	client := c.Connect()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Printf("Bye from Referal DB")
	}()
	//mongo connection
	collection := client.Database("User").Collection("Refree")
	var result LoggedInUser
	//what to find
	filter := bson.D{{"username", username}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No matching document found, please singup")
			return false
		} else {
			panic(err)
		}
	} else {
		if result.Password != password {
			fmt.Printf("Password is wrong for user:")
			return false
		}
	}
	return true
}

func (u *ConcreteUser) Authenticate(name string, username string, password string, id string, isrefree bool) bool {
	c := &connection.ConcreteConnection{}
	client := c.Connect()
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Printf("Bye from Referal DB")
	}()

	collection := client.Database("User").Collection("Refree")
	var result LoggedInUser
	//what to find
	filter := bson.D{{"username", username}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err == nil {
		return false
	}
	// if u.CheckUser(username, password) == true {
	// 	fmt.Printf("This is an already existing user...")
	// 	return false
	// }

	if len(password) < 8 {
		return false
	}
	newUser := &database.User{
		ID:       id,
		Password: password,
		Name:     name,
		UserName: username,
		IsRefree: isrefree,
	}
	insertResult, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		panic(err)

	}
	insertedID := insertResult.InsertedID
	fmt.Println("Inserted document ID:", insertedID)
	return true
}
