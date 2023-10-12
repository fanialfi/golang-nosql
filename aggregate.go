package main

import (
	"fmt"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

var pipeline = make([]bson.M, 0)

func Aggregate() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
[
	{ "$group" : {
		"_id": null,
		"Total": {"$sum": 1}
	}},
	{ "$project": {
		"Total": 1,
		"_id": 0
	}}
]
`)), true, &pipeline)

	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M

		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Printf("total\t: %d\n", result[0]["Total"])
	}
}
