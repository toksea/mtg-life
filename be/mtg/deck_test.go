package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAddCard(t *testing.T) {

	island := Card{
		name: "Island",
		img:  "/img/island.jpg",
	}

	myDeck := Deck{
		name: "my islands deck",
	}

	myDeck.addCard(&island, 60)

	log.Printf("my deck mainboard: %v", myDeck.mainboard)

	// 长度应该为 1
	assert.Equal(t, 1, len(myDeck.mainboard))

	// 获取第一张牌
	firstCard := myDeck.mainboard[0]

	// 应该是 DeckCard
	assert.IsType(t, DeckCard{}, firstCard)

	// 应该是 60 张 island
	assert.Equal(t, "Island", firstCard.card.name)
	assert.Equal(t, 60, firstCard.count)

}

func TestAddCardShouldMerge(t *testing.T) {

	island := Card{
		name: "Island",
		img:  "/img/island.jpg",
	}

	myDeck := Deck{
		name: "my islands deck",
	}

	myDeck.addCard(&island, 1)
	myDeck.addCard(&island, 1)

	// 长度应该为 1
	assert.Equal(t, 1, len(myDeck.mainboard))

	// 获取第一张牌
	firstCard := myDeck.mainboard[0]

	// 应该是 DeckCard
	assert.IsType(t, DeckCard{}, firstCard)

	// 应该是 2 张 island
	assert.Equal(t, "Island", firstCard.card.name)
	assert.Equal(t, 2, firstCard.count)

}

func TestAddFromDecklist(t *testing.T) {

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

	// 一共有 27 种不同的牌
	assert.Equal(t, 27, len(twinDeck.mainboard))

}
