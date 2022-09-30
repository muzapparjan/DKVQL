package main

import (
	"dkvql"
	"fmt"
)

func main() {
	src := "insert \"username\" value \"admin\" timeout 30"

	sentence, err := dkvql.Parse(src)
	if err != nil {
		panic(err)
	}

	if insert, ok := sentence.(dkvql.Insert); ok {
		fmt.Printf("\nInsert: key={%v} value={%v} timeout={%v}", insert.Key.Value, insert.Value.Value, insert.Timeout.Value)
	}
}
