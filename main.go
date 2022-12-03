package main

import (
	"log"

	_ "github.com/lib/pq"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {

	generateSchema()
	// main1()
}

func generateSchema() {
	err := entc.Generate("./ent/schema",
		&gen.Config{},
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
