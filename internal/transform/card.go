package transform

import (
	"github.com/StarWarsDev/archives/internal/gql/models"
	legionhq "github.com/StarWarsDev/go-legion-hq"
	"strconv"
)

func ImagePathToURL(path string) string {
	return "https://raw.githubusercontent.com/NicholasCBrown/legion-HQ-2.0/master/public" + path
}

func CardToCommand(card *legionhq.Card) models.Command {
	pips, err := strconv.Atoi(card.CardSubType)
	if err != nil {
		pips = 0
	}

	commander := card.Commander
	faction := card.Faction

	return models.Command{
		ID:           card.ID,
		CardType:     card.CardType,
		CardSubType:  card.CardSubType,
		Name:         card.CardName,
		Requirements: card.Requirements,
		Icon:         ImagePathToURL(card.IconLocation),
		Image:        ImagePathToURL(card.ImageLocation),
		Commander:    &commander,
		Faction:      &faction,
		Keywords:     card.Keywords,
		Pips:         pips,
	}
}

// CardToUnit converts a legionhq card into a Unit
func CardToUnit(card *legionhq.Card) models.Unit {
	isUnique := card.IsUnique
	return models.Unit{
		ID:           card.ID,
		Name:         card.CardName,
		CardType:     card.CardType,
		CardSubType:  card.CardSubType,
		Icon:         ImagePathToURL(card.IconLocation),
		Image:        ImagePathToURL(card.ImageLocation),
		Requirements: card.Requirements,
		Unique:       &isUnique,
		Cost:         card.Cost,
		Rank:         card.Rank,
		Faction:      card.Faction,
		Slots:        card.UpgradeBar,
		Keywords:     card.Keywords,
	}
}

func CardToUpgrade(card *legionhq.Card) models.Upgrade {
	isUnique := card.IsUnique
	return models.Upgrade{
		ID:           card.ID,
		CardType:     card.CardType,
		CardSubType:  card.CardSubType,
		Name:         card.CardName,
		Requirements: card.Requirements,
		Icon:         ImagePathToURL(card.IconLocation),
		Image:        ImagePathToURL(card.ImageLocation),
		Unique:       &isUnique,
		Cost:         card.Cost,
		Keywords:     card.Keywords,
	}
}
