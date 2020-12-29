package main

import (
	"fmt"
	"milli-uni-odekakebiyori/model"
)

func main() {
	fmt.Println("Hello,world")

	game := model.NewGame()

	fmt.Println(game.Draw(10))
	fmt.Println(game.PlayerJoin(10))
	fmt.Println(game.Draw(10))
	fmt.Println(game.SetCurrentPlayer(10))
	fmt.Println(game.Draw(10))
	fmt.Println(game.Stage(10, 1, false))
	fmt.Println(game.Stage(10, 0, false))

}
