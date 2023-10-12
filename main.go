package main

import (
	"context"
)

func main() {
	// Insert()
	// Find()
	// Update()
	Remove()
}

type Student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

var ctx = context.Background()
