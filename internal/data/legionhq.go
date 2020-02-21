package data

import (
	"github.com/StarWarsDev/archives/internal/gql/models"
	"github.com/StarWarsDev/archives/internal/transform"
	"github.com/StarWarsDev/go-legion-hq"
	"time"
)

var (
	lastFetch time.Time
	data      legionhq.Data
)

func getData() (legionhq.Data, error) {
	elapsed := time.Since(lastFetch)
	if elapsed.Minutes() >= 5 {
		d, err := legionhq.GetData()
		if err != nil {
			return d, err
		}
		data = d
		lastFetch = time.Now()
	}

	return data, nil
}

// CommandCards returns all data with cardType: command
func CommandCards() ([]*models.Command, error) {
	data, err := getData()
	if err != nil {
		return nil, err
	}

	var commands []*models.Command

	for _, card := range data.CommandCards() {
		command := transform.CardToCommand(&card)
		commands = append(commands, &command)
	}

	return commands, nil
}

// UnitCards returns all data with cardType: unit
func UnitCards() ([]*models.Unit, error) {
	data, err := getData()
	if err != nil {
		return nil, err
	}

	var units []*models.Unit

	for _, card := range data.UnitCards() {
		unit := transform.CardToUnit(&card)
		units = append(units, &unit)
	}

	return units, nil
}

// UpgradeCards returns all data with cardType: upgrade
func UpgradeCards() ([]*models.Upgrade, error) {
	data, err := getData()
	if err != nil {
		return nil, err
	}

	var upgrades []*models.Upgrade

	for _, card := range data.UpgradeCards() {
		upgrade := transform.CardToUpgrade(&card)
		upgrades = append(upgrades, &upgrade)
	}

	return upgrades, nil
}

// Keywords returns a slice of keywords
func Keywords() ([]*models.Keyword, error) {
	data, err := getData()
	if err != nil {
		return nil, err
	}

	var keywords []*models.Keyword

	for _, collection := range []map[string]string{data.KeywordDict, data.AdditionalKeywords} {
		for name, description := range collection {
			keyword := models.Keyword{
				Name:        name,
				Description: description,
			}
			keywords = append(keywords, &keyword)
		}
	}

	return keywords, nil
}

func CommunityLinks() ([]*models.LinkGroup, error) {
	data, err := getData()
	if err != nil {
		return nil, err
	}

	var linkGroups []*models.LinkGroup

	for name, linksIn := range data.CommunityLinks {
		var links []*models.Link

		for _, link := range linksIn {
			links = append(links, &models.Link{
				Name: link.Name,
				URL:  link.URL,
			})
		}

		linkGroups = append(linkGroups, &models.LinkGroup{
			Name:  name,
			Links: links,
		})
	}

	return linkGroups, nil
}
