package main

import (
	"fmt"
	"log"
)

func main() {
	store, err := newPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", store)

	server.Run()

	fmt.Println("Yea Buddy!")
}
