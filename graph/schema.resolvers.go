package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"os"

	hillchartsapi "github.com/offerni/hill-charts-api"
	"github.com/offerni/hill-charts-api/graph/generated"
	"github.com/offerni/hill-charts-api/graph/model"
	"github.com/offerni/hill-charts-api/squad"
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
	squads, err := r.squadSvc.List(ctx, squad.ListOpts{
		AccountID:      hillchartsapi.AccountID(1),
		OrganizationID: hillchartsapi.OrganizationID(os.Getenv("DEFAULT_ORG_ID")),
	})
	if err != nil {
		return nil, err
	}

	var squadsResp = make([]*model.Squad, len(squads.Data))

	for i, squad := range squads.Data {
		squadsResp[i].CurrentCycleName = squad.CurrentCycleName
		squadsResp[i].ID = string(squad.ID)
		squadsResp[i].Name = squad.Name
	}

	return &model.SquadList{
		Data:       []*model.Squad{},
		HasMore:    squads.HasMore,
		TotalCount: squads.TotalCount,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
