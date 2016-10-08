package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard(t *testing.T) {

	cases := []struct {
		name string
		img  string
	}{
		{
			"Island",
			"/img/island.jpg",
		},
	}

	for _, c := range cases {

		// construct
		card := Card{
			name: c.name,
			img:  c.img,
		}

		assert.Equal(t, card.name, c.name, "name not same")
		assert.Equal(t, card.img, c.img, "img not same")

	}

}
