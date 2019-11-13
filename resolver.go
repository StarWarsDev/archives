package main

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Units(ctx context.Context) ([]*Unit, error) {
	units := []*Unit{
		{
			ID:      "stormtroopers",
			Name:    "Stormtroopers",
			Side:    SideDark,
			Points:  44,
			Rank:    RankCorps,
			Type:    UnitTypeTrooper,
			Minis:   4,
			Wounds:  1,
			Courage: 1,
			Defense: DefenseDiceRed,
			Surge: Surge{
				Attack: &AttackSurgeHit,
			},
		},
	}

	return units, nil
}
