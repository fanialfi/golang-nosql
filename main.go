package main

import (
	"context"
)

func main() {
	// Insert()
	// Find()
	// Update()
	// Remove()
	Aggregate()
}

type Student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

var ctx = context.Background()
