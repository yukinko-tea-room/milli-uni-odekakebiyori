package model

import (
	"errors"
	"fmt"
)

// PlayerAction はプレイヤーの1行動を表します
type PlayerAction struct {
	PlayerID   int
	CardID     int
	ActionType int
}

// StageCard はプレイヤーが場に出したカードを表します
type StageCard struct {
	card  *Card
	isTap bool
}

// Game は1ゲームが持つ情報を管理します
type Game struct {
	PlayerHand    map[int]Cards
	PlayerStage   map[int][]*StageCard
	Deck          *Deck
	Log           []*PlayerAction
	CurrentPlayer int // 現在のターンのPlayerID
}

// NewGame 月曜日が町にやってくる 桜舞う電車に飛び乗る
func NewGame() *Game {
	playerHand := make(map[int]Cards, 4)
	playerStage := make(map[int][]*StageCard, 4)
	return &Game{
		Deck:        NewDeck(),
		PlayerHand:  playerHand,
		PlayerStage: playerStage,
	}
}

// PlayerJoin はゲームにプレイヤーを登録します
func (g *Game) PlayerJoin(playerID int) error {
	if _, ok := g.PlayerHand[playerID]; ok {
		return fmt.Errorf("error: playerID:%d was already joined", playerID)
	}
	if _, ok := g.PlayerStage[playerID]; ok {
		return fmt.Errorf("error: playerID:%d was already joined", playerID)
	}

	g.PlayerHand[playerID] = Cards{}
	g.PlayerStage[playerID] = []*StageCard{}

	return nil
}

// Draw は現在のターンのプレイヤーがドローします
func (g *Game) Draw(playerID int) ([]*Card, error) {
	if playerID != g.CurrentPlayer {
		// ターンじゃない
		return []*Card{}, fmt.Errorf("error: playerID:%d is not current player", playerID)
	}
	card, _, err := g.Deck.Draw()
	if err != nil {
		// 流局
		return []*Card{}, errors.New("error: deck was empty")
	}
	g.PlayerHand[g.CurrentPlayer] = append(g.PlayerHand[g.CurrentPlayer], card)
	return g.PlayerHand[g.CurrentPlayer], nil
}

func (g *Game) SetCurrentPlayer(playerID int) error {
	if _, ok := g.PlayerHand[playerID]; !ok {
		return fmt.Errorf("error: playerID:%d not found in this game player stage", playerID)
	}
	if _, ok := g.PlayerStage[playerID]; !ok {
		return fmt.Errorf("error: playerID:%d not found in this game player stage", playerID)
	}

	g.CurrentPlayer = playerID
	return nil
}

// Stage は手札のアイドルを場に出します
func (g *Game) Stage(playerID, idolID int, isTap bool) ([]*StageCard, error) {
	hands, ok := g.PlayerHand[playerID]
	if !ok {
		// このゲームのプレイヤーちゃうで
		return []*StageCard{}, fmt.Errorf("error: playerID:%d not found in this game player hands", playerID)
	}

	stage, ok := g.PlayerStage[playerID]
	if !ok {
		// このゲームのプレイヤーちゃうで
		return []*StageCard{}, fmt.Errorf("error: playerID:%d not found in this game player stage", playerID)
	}

	if playerID != g.CurrentPlayer {
		// ターンじゃない
		return []*StageCard{}, fmt.Errorf("error: playerID:%d is not current player", playerID)
	}

	card := hands.FindByIdolID(idolID)
	if card == nil {
		// そのカード持ってないで
		return []*StageCard{}, fmt.Errorf("error: playerID:%d idolID:%d not found int hand", playerID, idolID)
	}

	g.PlayerStage[playerID] = append(stage, &StageCard{
		card:  card,
		isTap: isTap,
	})
	return g.PlayerStage[playerID], nil
}
