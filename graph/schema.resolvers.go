package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/offerni/hill-charts-api/graph/generated"
	"github.com/offerni/hill-charts-api/graph/model"
)

// CreateProject is the resolver for the CreateProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, opts model.NewProject) (*model.Project, error) {
	panic(fmt.Errorf("not implemented: CreateProject - CreateProject"))
}

// CreateScope is the resolver for the CreateScope field.
func (r *mutationResolver) CreateScope(ctx context.Context, opts model.NewScope) (*model.Scope, error) {
	panic(fmt.Errorf("not implemented: CreateScope - CreateScope"))
}

// Pages is the resolver for the pages field.
func (r *queryResolver) Pages(ctx context.Context) (*model.ProjectList, error) {
	panic(fmt.Errorf("not implemented: Pages - pages"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
