package main

import (
	"fmt"
	"milli-uni-odekakebiyori/model"
)

func main() {
	fmt.Println("Hello,world")

	deck := model.NewDeck()

	fmt.Println(deck.List())
	fmt.Println(deck.Draw())
}
