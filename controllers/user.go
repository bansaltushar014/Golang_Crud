package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	model "github.com/bansaltushar014/GoLang_CRUD/model"
	database "github.com/bansaltushar014/GoLang_CRUD/services"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	log.Println(q["Title"][0])
	var findDoc interface{} = bson.D{{"title", q["Title"][0]}}
	findResult, FindErr := database.FindData(findDoc)
	if FindErr != nil {
		log.Fatal((FindErr))
	}

	result, jsonErr := json.Marshal(findResult)
	if jsonErr != nil {
		log.Fatal((jsonErr))
	}

	log.Println(len(findResult.Title))
	if len(findResult.Title) == 0 {
		io.WriteString(w, "Record does not exist!")
	} else {
		io.WriteString(w, string(result))
	}
}

func HandleWriteOrder(w http.ResponseWriter, r *http.Request) {
	// input request body
	decoder := json.NewDecoder(r.Body)
	var t model.Post
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	docs := []interface{}{
		bson.D{{"title", t.Title}, {"body", t.Body}},
	}

	database.InsertData(docs)
	io.WriteString(w, "Record has been updated!")
}

func HandleUpdateOrder(w http.ResponseWriter, r *http.Request) {
	// input request body
	decoder := json.NewDecoder(r.Body)
	var t model.Post
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	var findDoc interface{} = bson.D{{"title", t.Title}}
	updateDoc := bson.D{{"$set", bson.D{{"body", t.Body}}}}
	findResult, err := database.FindData(findDoc)
	log.Println(len(findResult.Title))

	if len(findResult.Title) == 0 {
		io.WriteString(w, "Record does not exist!")
	} else {
		updateErr := database.UpdateData(findDoc, updateDoc)
		if updateErr != nil {
			io.WriteString(w, "Record could not be updated!")
		} else {
			io.WriteString(w, "Record has been updated!")
		}
	}

}
