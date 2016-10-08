package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Card struct {
	name string
	img  string
}

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

func (d *Deck) addCard(card *Card, count int) {
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

func main() {

	island := Card{
		name: "Island",
		img:  "/img/island.jpg",
	}

	myDeck := Deck{
		name: "my islands deck",
	}

	myDeck.addCard(&island, 60)

	log.Printf("my deck: %v", myDeck.getDecklist())

	decklist := `
3 Deceiver Exarch
2 Pestermite
4 Snapcaster Mage
2 Tasigur, the Golden Fang
1 Vendilion Clique
4 Island
1 Mountain
1 Swamp
1 Blood Crypt
2 Bloodstained Mire
1 Desolate Lighthouse
4 Polluted Delta
3 Scalding Tarn
2 Steam Vents
2 Sulfur Falls
1 Tectonic Edge
2 Watery Grave
3 Splinter Twin
2 Cryptic Command
1 Dismember
1 Dispel
1 Kolaghan's Command
4 Lightning Bolt
4 Remand
2 Spell Snare
2 Terminate
4 Serum Visions
`
	twinDeck := Deck{
		name: "Splinter Twin",
	}

	twinDeck.addFromDecklist(decklist)

	log.Printf("twin deck: %v", twinDeck.getDecklist())

}
