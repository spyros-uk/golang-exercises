package main

import (
	"./api"
	"./person"
	"fmt"
)

func main() {
		fmt.Println("Main is running...")
		Help()
		person.Walk()
		graphql.Moto()
}
