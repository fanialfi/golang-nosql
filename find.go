package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Find() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "fani alfirdaus"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]Student, 0)
	for csr.Next(ctx) {
		var row Student

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Printf("Name:\t%s\n", result[0].Name)
		fmt.Printf("Grade:\t%d\n", result[0].Grade)
	}
}
