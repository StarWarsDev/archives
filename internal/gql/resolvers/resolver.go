package resolvers

import (
	"context"
	"errors"
	"fmt"
	"github.com/StarWarsDev/archives/internal/data"
	"github.com/StarWarsDev/archives/internal/gql"
	"github.com/StarWarsDev/archives/internal/gql/models"
	"regexp"
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
	commandCards, err := data.CommandCards()
	if err != nil {
		return nil, err
	}

	if query == nil {
		return commandCards, nil
	} else {
		field, term, err := parseQuery(query)
		if err != nil {
			return nil, err
		}

		var filteredCards []*models.Command
		for _, commandCard := range commandCards {
			isMatch := false
			switch field {
			case "id":
				isMatch = isExactMatch(commandCard.ID, term)
			case "commander":
				match, err := isRegexpMatch(commandCard.Commander, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "name":
				match, err := isRegexpMatch(commandCard.Name, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "pips":
				pips, err := strconv.Atoi(term)
				if err != nil {
					return nil, err
				}
				isMatch = commandCard.Pips == pips
			case "faction":
				match, err := isRegexpMatch(commandCard.Faction, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			default:
				return nil, fmt.Errorf("bad query: field [%s] is not searchable", field)
			}

			if isMatch {
				filteredCards = append(filteredCards, commandCard)
			}
		}

		return filteredCards, nil
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
	keywords, err := data.Keywords()
	if err != nil {
		return nil, err
	}

	if query == nil {
		// get everything
		return keywords, nil
	} else {
		field, term, err := parseQuery(query)
		if err != nil {
			return nil, err
		}

		var filteredKeywords []*models.Keyword
		for _, keyword := range keywords {
			isMatch := false
			switch field {
			case "name":
				match, err := isRegexpMatch(keyword.Name, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "description":
				// do a regex match rather than an exact string match
				match, err := isRegexpMatch(keyword.Description, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			default:
				return nil, fmt.Errorf("bad query: field [%s] is not searchable", field)
			}

			if isMatch {
				filteredKeywords = append(filteredKeywords, keyword)
			}
		}

		return filteredKeywords, nil
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
	cards, err := data.UnitCards()
	if err != nil {
		return nil, err
	}

	if query == nil {
		return cards, nil
	} else {
		field, term, err := parseQuery(query)
		if err != nil {
			return nil, err
		}

		var filteredCards []*models.Unit
		for _, card := range cards {
			isMatch := false
			switch field {
			case "id":
				isMatch = isExactMatch(card.ID, term)
			case "name":
				match, err := isRegexpMatch(card.Name, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "cardType":
				match, err := isRegexpMatch(card.CardType, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "cardSubType":
				match, err := isRegexpMatch(card.CardSubType, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "unique":
				unique, err := strconv.ParseBool(term)
				if err != nil {
					return nil, err
				}
				isMatch = card.Unique == unique
			case "requirements":
				for _, requirement := range card.Requirements {
					if !isMatch {
						match, err := isRegexpMatch(requirement, term)
						if err != nil {
							return nil, err
						}
						isMatch = match
					}
				}
			case "rank":
				match, err := isRegexpMatch(card.Rank, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "faction":
				match, err := isRegexpMatch(card.Faction, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "slots":
				for _, slot := range card.Slots {
					if !isMatch {
						match, err := isRegexpMatch(slot, term)
						if err != nil {
							return nil, err
						}
						isMatch = match
					}
				}
			case "keywords":
				for _, keyword := range card.Keywords {
					if !isMatch {
						match, err := isRegexpMatch(keyword, term)
						if err != nil {
							return nil, err
						}
						isMatch = match
					}
				}
			default:
				return nil, fmt.Errorf("bad query: field [%s] is not searchable", field)
			}

			if isMatch {
				q := "commander: " + card.Name
				commandCards, err := r.Commands(ctx, &q)
				if err != nil {
					return nil, err
				}
				card.CommandCards = commandCards
				filteredCards = append(filteredCards, card)
			}
		}

		return filteredCards, nil
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
	cards, err := data.UpgradeCards()
	if err != nil {
		return nil, err
	}

	if query == nil {
		return cards, nil
	} else {
		field, term, err := parseQuery(query)
		if err != nil {
			return nil, err
		}

		var filteredCards []*models.Upgrade
		for _, card := range cards {
			isMatch := false
			switch field {
			case "id":
				isMatch = isExactMatch(card.ID, term)
			case "name":
				match, err := isRegexpMatch(card.Name, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "cardType":
				match, err := isRegexpMatch(card.CardType, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "cardSubType":
				match, err := isRegexpMatch(card.CardSubType, term)
				if err != nil {
					return nil, err
				}
				isMatch = match
			case "unique":
				unique, err := strconv.ParseBool(term)
				if err != nil {
					return nil, err
				}
				isMatch = card.Unique == unique
			case "requirements":
				for _, requirement := range card.Requirements {
					if !isMatch {
						match, err := isRegexpMatch(requirement, term)
						if err != nil {
							return nil, err
						}
						isMatch = match
					}
				}
			case "keywords":
				for _, keyword := range card.Keywords {
					if !isMatch {
						match, err := isRegexpMatch(keyword, term)
						if err != nil {
							return nil, err
						}
						isMatch = match
					}
				}
			default:
				return nil, fmt.Errorf("bad query: field [%s] is not searchable", field)
			}

			if isMatch {
				filteredCards = append(filteredCards, card)
			}
		}

		return filteredCards, nil
	}
}

func (r *queryResolver) CommunityLinks(ctx context.Context) ([]*models.LinkGroup, error) {
	return data.CommunityLinks()
}

func parseQuery(query *string) (field, term string, err error) {
	if *query == "" {
		return "", "", errors.New("malformed query: query cannot be blank")
	}

	q := strings.Split(*query, ":")
	if len(q) < 1 || len(q) < 2 {
		err = errors.New("malformed query")
		return
	}
	field = strings.TrimSpace(q[0])
	term = strings.TrimSpace(q[1])

	if field == "" {
		return "", "", errors.New("malformed query: field cannot be blank")
	}

	if term == "" {
		return "", "", errors.New("malformed query: field cannot be blank")
	}

	return
}

func isExactMatch(subject, term string) bool {
	return subject == term
}

func isRegexpMatch(subject, pattern string) (bool, error) {
	return regexp.MatchString(pattern, subject)
}
