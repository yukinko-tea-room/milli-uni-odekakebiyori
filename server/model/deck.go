package model

import "errors"

var (
	// ErrDeckEmpty はデッキが空なことを示すエラーです
	ErrDeckEmpty = errors.New("error: deck empty")
)

type Deck struct {
	deck []*Card
}

// NewDeck は新しいDeckを作成して返します
func NewDeck() *Deck {
	// TODO: シャッフルしてデッキセット
	return &Deck{
		deck: CardList(),
	}
}

// List
func (d *Deck) List() []*Card {
	return d.deck
}

// Draw
func (d *Deck) Draw() (*Card, int, error) {
	if len(d.deck) <= 0 {
		return nil, 0, ErrDeckEmpty
	}
	drawCard := d.deck[0]
	d.deck = d.deck[1:]
	return drawCard, len(d.deck), nil
}
