package resolvers

import (
	"context"

	"github.com/StarWarsDev/archives/internal/gql"
	"github.com/StarWarsDev/archives/internal/gql/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() gql.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Command(ctx context.Context, id string) (*models.Command, error) {
	panic("not implemented")
}
func (r *queryResolver) Commands(ctx context.Context, query string) ([]*models.Command, error) {
	panic("not implemented")
}
func (r *queryResolver) Keyword(ctx context.Context, name string) (*models.Keyword, error) {
	panic("not implemented")
}
func (r *queryResolver) Keywords(ctx context.Context, query string) ([]*models.Keyword, error) {
	panic("not implemented")
}
func (r *queryResolver) Unit(ctx context.Context, id string) (*models.Unit, error) {
	panic("not implemented")
}
func (r *queryResolver) Units(ctx context.Context, query string) ([]*models.Unit, error) {
	panic("not implemented")
}
func (r *queryResolver) Upgrade(ctx context.Context, id string) (*models.Upgrade, error) {
	panic("not implemented")
}
func (r *queryResolver) Upgrades(ctx context.Context, query string) ([]*models.Upgrade, error) {
	panic("not implemented")
}
