package transform

import (
	"fmt"
	"github.com/StarWarsDev/archives/internal/gql/models"
	legiondata "github.com/StarWarsDev/go-legion-data"
	legionhq "github.com/StarWarsDev/go-legion-hq"
	"strconv"
)

func ImagePathToURL(cardType, imageName string) string {
	return fmt.Sprintf("https://d2b46bduclcqna.cloudfront.net/%sCards/%s", cardType, imageName)
}

func CardToCommand(card *legionhq.Card, extCard *legiondata.CommandCard) models.Command {
	pips, err := strconv.Atoi(card.CardSubType)
	if err != nil {
		pips = 0
	}

	commandCard := models.Command{
		ID:          card.ID,
		CardType:    card.CardType,
		CardSubType: card.CardSubType,
		Name:        card.CardName,
		Image:       ImagePathToURL(card.CardType, card.ImageLocation),
		Commander:   card.Commander,
		Faction:     card.Faction,
		Pips:        pips,
	}

	// loop through the requirements so we can filter out blanks
	for _, requirement := range card.Requirements {
		if requirement != "" {
			commandCard.Requirements = append(commandCard.Requirements, requirement)
		}
	}

	// loop through the keywords so we can filter out blanks
	for _, keyword := range card.Keywords {
		if keyword != "" {
			commandCard.Keywords = append(commandCard.Keywords, keyword)
		}
	}

	if extCard != nil {
		commandCard.Orders = extCard.Orders
		commandCard.Text = extCard.Description
		if extCard.Weapon != nil {
			commandCard.Weapon = &models.Weapon{
				Name: extCard.Weapon.Name,
				Range: &models.Range{
					From: extCard.Weapon.Range.From,
					To:   extCard.Weapon.Range.To,
				},
				Keywords: extCard.Weapon.Keywords,
				Dice: &models.Dice{
					Black: extCard.Weapon.Dice.Black,
					Red:   extCard.Weapon.Dice.Red,
					White: extCard.Weapon.Dice.White,
				},
			}

			if extCard.Weapon.Surge != nil {
				commandCard.Weapon.Surge = &models.Surge{
					Attack:  extCard.Weapon.Surge.Attack,
					Defense: extCard.Weapon.Surge.Defense,
				}
			}
		}
	}

	return commandCard
}

// CardToUnit converts a legionhq card into a Unit
func CardToUnit(card *legionhq.Card, extUnit *legiondata.Unit) models.Unit {
	unit := models.Unit{
		ID:          card.ID,
		Name:        card.CardName,
		CardType:    card.CardType,
		CardSubType: card.CardSubType,
		Image:       ImagePathToURL(card.CardType, card.ImageLocation),
		Unique:      card.IsUnique,
		Cost:        card.Cost,
		Rank:        card.Rank,
		Faction:     card.Faction,
	}

	// loop through the slots to filter out blanks
	for _, slot := range card.UpgradeBar {
		if slot != "" {
			unit.Slots = append(unit.Slots, slot)
		}
	}

	// loop through the requirements so we can filter out blanks
	for _, requirement := range card.Requirements {
		if requirement != "" {
			unit.Requirements = append(unit.Requirements, requirement)
		}
	}

	// loop through the keywords so we can filter out blanks
	for _, keyword := range card.Keywords {
		if keyword != "" {
			unit.Keywords = append(unit.Keywords, keyword)
		}
	}

	if extUnit != nil {
		// enrich unit data
		unit.Wounds = extUnit.Wounds
		unit.Courage = extUnit.Courage
		unit.Resilience = extUnit.Resilience
		unit.Surge = &models.Surge{
			Attack:  extUnit.Surge.Attack,
			Defense: extUnit.Surge.Defense,
		}
		unit.Entourage = extUnit.Entourage
		for _, weapon := range extUnit.Weapons {
			weap := models.Weapon{
				Name: weapon.Name,
				Range: &models.Range{
					From: weapon.Range.From,
					To:   weapon.Range.To,
				},
				Keywords: weapon.Keywords,
				Dice: &models.Dice{
					Black: weapon.Dice.Black,
					Red:   weapon.Dice.Red,
					White: weapon.Dice.White,
				},
			}

			if weapon.Surge != nil {
				weap.Surge = &models.Surge{
					Attack:  weapon.Surge.Attack,
					Defense: weapon.Surge.Defense,
				}
			}

			unit.Weapons = append(unit.Weapons, &weap)
		}
	}

	return unit
}

func CardToUpgrade(card *legionhq.Card, extUpgrade *legiondata.Upgrade) models.Upgrade {
	upgrade := models.Upgrade{
		ID:          card.ID,
		CardType:    card.CardType,
		CardSubType: card.CardSubType,
		Name:        card.CardName,
		Image:       ImagePathToURL(card.CardType, card.ImageLocation),
		Unique:      card.IsUnique,
		Cost:        card.Cost,
	}

	// loop through the requirements so we can filter out blanks
	for _, requirement := range card.Requirements {
		if requirement != "" {
			upgrade.Requirements = append(upgrade.Requirements, requirement)
		}
	}

	// loop through the keywords so we can filter out blanks
	for _, keyword := range card.Keywords {
		if keyword != "" {
			upgrade.Keywords = append(upgrade.Keywords, keyword)
		}
	}

	if extUpgrade != nil {
		upgrade.Exhaust = extUpgrade.Exhaust != nil && *extUpgrade.Exhaust
		upgrade.UnitTypeExclusions = extUpgrade.UnitTypeExclusions
		upgrade.Text = extUpgrade.Description
		if extUpgrade.Weapon != nil {
			upgrade.Weapon = &models.Weapon{
				Name: extUpgrade.Weapon.Name,
				Range: &models.Range{
					From: extUpgrade.Weapon.Range.From,
					To:   extUpgrade.Weapon.Range.To,
				},
				Keywords: extUpgrade.Weapon.Keywords,
				Dice: &models.Dice{
					Black: extUpgrade.Weapon.Dice.Black,
					Red:   extUpgrade.Weapon.Dice.Red,
					White: extUpgrade.Weapon.Dice.White,
				},
			}

			if extUpgrade.Weapon.Surge != nil {
				upgrade.Weapon.Surge = &models.Surge{
					Attack:  extUpgrade.Weapon.Surge.Attack,
					Defense: extUpgrade.Weapon.Surge.Defense,
				}
			}
		}
	}

	return upgrade
}
