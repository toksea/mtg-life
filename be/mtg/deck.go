package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type DeckCard struct {
	card  *Card
	count int
}

type Deck struct {
	name      string
	mainboard []DeckCard // slice
}

// interfaces
func (d *Deck) addFromDecklist(decklist string) {
	for _, line := range strings.Split(decklist, "\n") {

		log.Printf("decklist line %v", line)

		if line == "" {
			continue
		}

		deckcardParts := strings.SplitN(line, " ", 2)

		log.Printf("decklist line parts %v", deckcardParts)

		// @TODO handle err
		count, _ := strconv.Atoi(deckcardParts[0]) // x is an int
		cardname := deckcardParts[1]

		log.Printf("decklist line detail %v, %v", cardname, count)

		// @TODO find card
		card := Card{
			name: cardname,
			// img:  "/img/island.jpg",
		}

		d.addCard(&card, count)
	}
}

func (d *Deck) lookUpCardByName(name string) (deckCard *DeckCard) {
	return
}

func (d *Deck) addCard(card *Card, count int) {

	// 搜索是否已有此牌，
	// 没有则新加，有则
	// 加数量

	deckCard := DeckCard{
		card:  card,
		count: count,
	}

	d.mainboard = append(d.mainboard, deckCard)

}

func (d *Deck) getDecklist() (decklist string) {

	decklist = d.name + "\n"

	for i := range d.mainboard {

		card := d.mainboard[i].card
		count := d.mainboard[i].count

		log.Printf("deck list: %v, %v",
			card, count)

		decklist += fmt.Sprintf("%v %v\n", count, card.name)

	}

	return
}
