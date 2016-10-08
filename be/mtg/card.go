package main

import (
	"log"
)

type Card struct {
	name string
	img  string
}

func main() {

	card := Card{
		name: "Island",
		img:  "/img/island.jpg",
	}

	log.Printf("card: %v", card)

}
