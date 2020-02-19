package transform

import (
	"github.com/StarWarsDev/archives/internal/gql/models"
	legionhq "github.com/StarWarsDev/go-legion-hq"
	"strconv"
)

func ImagePathToURL(path string) string {
	return "https://raw.githubusercontent.com/NicholasCBrown/legion-HQ-2.0/master/public" + path
}

func CardToCommand(card legionhq.Card) models.Command {
	pips, err := strconv.Atoi(card.CardSubType)
	if err != nil {
		pips = 0
	}
	return models.Command{
		ID:           card.ID,
		CardType:     card.CardType,
		CardSubType:  card.CardSubType,
		Name:         card.CardName,
		Requirements: card.Requirements,
		Icon:         ImagePathToURL(card.IconLocation),
		Image:        ImagePathToURL(card.ImageLocation),
		Commander:    &card.Commander,
		Faction:      &card.Faction,
		Keywords:     card.Keywords,
		Pips:         pips,
	}
}
