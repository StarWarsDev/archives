package resolvers

import (
	"context"
	"errors"
	"fmt"
	"github.com/StarWarsDev/archives/internal/data"
	"github.com/StarWarsDev/archives/internal/gql"
	"github.com/StarWarsDev/archives/internal/gql/models"
	"strconv"
	"strings"
)

// Resolver wraps up all the GraphQL resolvers
type Resolver struct{}

// Query handles all query requests
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Command(ctx context.Context, id string) (*models.Command, error) {
	cards, err := data.CommandCards()
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		if card.ID == id {
			return card, nil
		}
	}

	return nil, fmt.Errorf("could not find command with ID %s", id)
}
func (r *queryResolver) Commands(ctx context.Context, query *string) ([]*models.Command, error) {
	if query == nil {
		return data.CommandCards()
	} else {
		field, term, err := parseQuery(query)
		if err != nil {
			return nil, err
		}
		if field != "" {
			if term == "" {
				return nil, errors.New("malformed query: term cannot be blank")
			}

			commandCards, err := data.CommandCards()
			if err != nil {
				return nil, err
			}

			var filteredCards []*models.Command
			for _, commandCard := range commandCards {
				switch field {
				case "id":
					if commandCard.ID == term {
						filteredCards = append(filteredCards, commandCard)
					}
				case "commander":
					if *commandCard.Commander == term {
						filteredCards = append(filteredCards, commandCard)
					}
				case "name":
					if commandCard.Name == term {
						filteredCards = append(filteredCards, commandCard)
					}
				case "pips":
					pips, err := strconv.Atoi(term)
					if err != nil {
						return nil, err
					}
					if commandCard.Pips == pips {
						filteredCards = append(filteredCards, commandCard)
					}
				case "faction":
					if *commandCard.Faction == term {
						filteredCards = append(filteredCards, commandCard)
					}
				default:
					return nil, fmt.Errorf("bad query: field [%s] is not searchable", field)
				}
			}

			return filteredCards, nil
		}
		return nil, errors.New("malformed query: query cannot be blank")
	}
}

func (r *queryResolver) Keyword(ctx context.Context, name string) (*models.Keyword, error) {
	keywords, err := data.Keywords()
	if err != nil {
		return nil, err
	}

	for _, keyword := range keywords {
		if keyword.Name == name {
			return keyword, nil
		}
	}

	return nil, fmt.Errorf("could not find keyword with name \"%s\"", name)
}
func (r *queryResolver) Keywords(ctx context.Context, query *string) ([]*models.Keyword, error) {
	if query == nil {
		// get everything
		return data.Keywords()
	} else {
		panic("query not implemented")
	}
}
func (r *queryResolver) Unit(ctx context.Context, id string) (*models.Unit, error) {
	cards, err := data.UnitCards()
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		if card.ID == id {
			query := "commander: " + card.Name
			commandCards, err := r.Commands(ctx, &query)
			if err != nil {
				return nil, err
			}
			card.CommandCards = commandCards

			return card, nil
		}
	}
	return nil, fmt.Errorf("could not find unit with ID %s", id)
}
func (r *queryResolver) Units(ctx context.Context, query *string) ([]*models.Unit, error) {
	if query == nil {
		// get everything
		units, err := data.UnitCards()
		if err != nil {
			return nil, err
		}

		for _, unit := range units {
			query := "commander: " + unit.Name
			commandCards, err := r.Commands(ctx, &query)
			if err != nil {
				return nil, err
			}
			unit.CommandCards = commandCards
		}

		return units, nil
	} else {
		panic("query not implemented")
	}
}
func (r *queryResolver) Upgrade(ctx context.Context, id string) (*models.Upgrade, error) {
	cards, err := data.UpgradeCards()
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		if card.ID == id {
			return card, nil
		}
	}
	return nil, fmt.Errorf("could not find upgrade with ID %s", id)
}
func (r *queryResolver) Upgrades(ctx context.Context, query *string) ([]*models.Upgrade, error) {
	if query == nil {
		// get everything
		upgrades, err := data.UpgradeCards()
		return upgrades, err
	} else {
		panic("query not implemented")
	}
}

func parseQuery(query *string) (field, term string, err error) {
	q := strings.Split(*query, ":")
	if len(q) < 1 || len(q) < 2 {
		err = errors.New("malformed query")
		return
	}
	field = strings.TrimSpace(q[0])
	term = strings.TrimSpace(q[1])
	return
}
