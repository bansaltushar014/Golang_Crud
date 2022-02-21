package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	model "github.com/bansaltushar014/GoLang_CRUD/model"
)

var client2 *mongo.Client

func CreateConnection() {
	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	client2 = client
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection Successful!")
	// defer client.Disconnect(ctx)

	/*
	   List databases
	*/
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}

func InsertData(docs []interface{}) {
	collection := client2.Database("goLangDB").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, insertErr := collection.InsertMany(ctx, docs)
	if insertErr != nil {
		log.Println("error")
		log.Fatal(insertErr)
	}
	fmt.Println(res)
	/*
		Iterate a cursor and print it
	*/
	cur, currErr := collection.Find(ctx, bson.D{})
	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)
	var posts []model.Post
	if err := cur.All(ctx, &posts); err != nil {
		panic(err)
	}
	fmt.Println(posts)

}

func FindData(docs interface{}) (model.Post, error) {
	result := model.Post{}
	collection := client2.Database("goLangDB").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	findErr := collection.FindOne(ctx, docs).Decode(&result)
	if findErr != nil {
		return result, findErr
	}
	log.Println(result)
	return result, nil
}

func UpdateData(findDoc interface{}, updateDoc interface{}) error {
	collection := client2.Database("goLangDB").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, updateErr := collection.UpdateOne(ctx, findDoc, updateDoc)
	if updateErr != nil {
		fmt.Println(updateErr)
		return updateErr
	}
	return nil
}
