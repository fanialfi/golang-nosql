package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Update() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "Ethan Nut"}
	var changes = Student{Name: "John Wick", Grade: 2}

	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("update success")
}
