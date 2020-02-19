package resolvers

import (
	"context"
	"github.com/StarWarsDev/archives/internal/transform"
	legionhq "github.com/StarWarsDev/go-legion-hq"

	"github.com/StarWarsDev/archives/internal/gql"
	"github.com/StarWarsDev/archives/internal/gql/models"
)

// Resolver wraps up all the GraphQL resolvers
type Resolver struct{}

// Query handles all query requests
func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Command(ctx context.Context, id string) (*models.Command, error) {
	panic("not implemented")
}
func (r *queryResolver) Commands(ctx context.Context, query string) ([]*models.Command, error) {
	if query == "" {
		data, err := legionhq.GetData()
		if err != nil {
			return nil, err
		}

		var commands []*models.Command

		for _, card := range data.CommandCards() {
			command := transform.CardToCommand(card)
			commands = append(commands, &command)
		}

		return commands, nil
	} else {
		panic("query not implemented")
	}
}
func (r *queryResolver) Keyword(ctx context.Context, name string) (*models.Keyword, error) {
	panic("not implemented")
}
func (r *queryResolver) Keywords(ctx context.Context, query string) ([]*models.Keyword, error) {
	if query == "" {
		// get everything
		panic("not implemented")
	} else {
		panic("query not implemented")
	}
}
func (r *queryResolver) Unit(ctx context.Context, id string) (*models.Unit, error) {
	panic("not implemented")
}
func (r *queryResolver) Units(ctx context.Context, query string) ([]*models.Unit, error) {
	if query == "" {
		// get everything
		panic("not implemented")
	} else {
		panic("query not implemented")
	}
}
func (r *queryResolver) Upgrade(ctx context.Context, id string) (*models.Upgrade, error) {
	panic("not implemented")
}
func (r *queryResolver) Upgrades(ctx context.Context, query string) ([]*models.Upgrade, error) {
	if query == "" {
		// get everything
		panic("not implemented")
	} else {
		panic("query not implemented")
	}
}
