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

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context) (*model.ProjectList, error) {
	panic(fmt.Errorf("not implemented: Projects - projects"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Pages(ctx context.Context) (*model.ProjectList, error) {
	panic(fmt.Errorf("not implemented: Pages - pages"))
}
