package data

import (
	"fmt"
	"github.com/StarWarsDev/go-legion-data"
	"strings"
	"time"
)

var (
	extLastFetch time.Time
	extData      legiondata.Data
)

func getExtData() (legiondata.Data, error) {
	elapsed := time.Since(extLastFetch)
	if elapsed.Minutes() > 5 {
		data, err := legiondata.GetData()
		if err != nil {
			return data, err
		}
		extData = data
		extLastFetch = time.Now()
	}

	return extData, nil
}

func ExtUnit(name, rank string) (*legiondata.Unit, error) {
	data, err := getExtData()
	if err != nil {
		return nil, err
	}

	for _, unit := range data.Units {
		if strings.ToLower(unit.Name) == strings.ToLower(name) &&
			strings.ToLower(unit.Rank) == strings.ToLower(rank) {
			return &unit, nil
		}
	}

	return nil, fmt.Errorf("cound not find unit named %s", name)
}

func ExtUpgrade(name string) (*legiondata.Upgrade, error) {
	data, err := getExtData()
	if err != nil {
		return nil, err
	}

	for _, upgrade := range data.Upgrades {
		if upgrade.Name == name {
			return &upgrade, nil
		}
	}

	return nil, fmt.Errorf("cound not find upgrade named %s", name)
}

func ExtCommandCard(name string) (*legiondata.CommandCard, error) {
	data, err := getExtData()
	if err != nil {
		return nil, err
	}

	for _, card := range data.CommandCards {
		if card.Name == name {
			return &card, nil
		}
	}

	return nil, fmt.Errorf("cound not find command card named %s", name)
}
