package model

// Card だよ
type Card struct {
	cardID int
	idolID int
}

// CardList はデッキに入る全カードのリストを返します
func CardList() []*Card {
	// TODO: 多分いつかいい感じなのになる
	// 0-51のidol_idのcardを2枚ずつ作る
	list := make([]*Card, 0, 104)
	for i := 0; i < 104; i++ {
		card := &Card{
			cardID: i,
			idolID: i % 52,
		}
		list = append(list, card)
	}
	return list
}

// Cards は複数カードの操作を扱います
type Cards []*Card

// FindByIdolID はidolIDが一致するカードを1枚返します
func (c Cards) FindByIdolID(idolID int) *Card {
	for _, card := range c {
		if card.idolID == idolID {
			return card
		}
	}
	return nil
}
