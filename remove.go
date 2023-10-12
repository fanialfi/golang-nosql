package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Remove() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	selector := bson.M{"name": "John Wick"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("remove success")
}
