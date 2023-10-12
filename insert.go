package main

import (
	"fmt"
	"log"
)

func Insert() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, Student{Name: "fani alfirdaus", Grade: 17})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, Student{Name: "Ethan Nut", Grade: 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("insert success")
}
