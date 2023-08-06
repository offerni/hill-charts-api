package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/offerni/hill-charts-api/graph/generated"
	"github.com/offerni/hill-charts-api/graph/model"
)

// CreateSquad is the resolver for the CreateSquad field.
func (r *mutationResolver) CreateSquad(ctx context.Context, opts model.NewSquad) (*model.Squad, error) {
	panic(fmt.Errorf("not implemented: CreateSquad - CreateSquad"))
}

// CreateScope is the resolver for the CreateScope field.
func (r *mutationResolver) CreateScope(ctx context.Context, opts model.NewScope) (*model.Scope, error) {
	panic(fmt.Errorf("not implemented: CreateScope - CreateScope"))
}

// Squads is the resolver for the Squads field.
func (r *queryResolver) Squads(ctx context.Context) (*model.SquadList, error) {
	panic(fmt.Errorf("not implemented: Squads - Squads"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
