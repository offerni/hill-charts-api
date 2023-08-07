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
	"github.com/offerni/hill-charts-api/scope"
	"github.com/offerni/hill-charts-api/squad"
)

// CreateScope is the resolver for the CreateScope field.
func (r *mutationResolver) CreateScope(ctx context.Context, opts model.NewScope) (*model.Scope, error) {
	scope, err := r.scopeSvc.Create(ctx, scope.CreateOpts{
		AccountID:      hillchartsapi.AccountID("1"),
		Colour:         opts.Colour,
		Name:           opts.Name,
		OrganizationID: hillchartsapi.OrganizationID(os.Getenv("DEFAULT_ORG_ID")),
		SquadID:        hillchartsapi.SquadID(opts.SquadID),
	})

	if err != nil {
		return nil, err
	}

	return &model.Scope{
		Colour:     scope.Colour,
		CreatedAt:  timeToStr(scope.CreatedAt),
		ID:         string(scope.ID),
		ModifiedAt: timeToStr(scope.ModifiedAt),
		Name:       scope.Name,
		Progress:   float32ToStr(scope.Progress),
		SquadID:    string(scope.SquadID),
	}, nil
}

// CreateSquad is the resolver for the CreateSquad field.
func (r *mutationResolver) CreateSquad(ctx context.Context, opts model.NewSquad) (*model.Squad, error) {
	squad, err := r.squadSvc.Create(ctx, squad.CreateOpts{
		AccountID:        hillchartsapi.AccountID("1"),
		CurrentCycleName: *opts.CurrentCycleName,
		Name:             opts.Name,
		OrganizationID:   hillchartsapi.OrganizationID(os.Getenv("DEFAULT_ORG_ID")),
	})

	if err != nil {
		return nil, err
	}

	return &model.Squad{
		CreatedAt:        timeToStr(squad.CreatedAt),
		CurrentCycleName: squad.CurrentCycleName,
		ID:               string(squad.ID),
		ModifiedAt:       timeToStr(squad.ModifiedAt),
		Name:             squad.Name,
		Scope:            []*model.Scope{},
	}, nil
}

// DeleteScope is the resolver for the DeleteScope field.
func (r *mutationResolver) DeleteScope(ctx context.Context, id *string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteScope - DeleteScope"))
}

// DeleteSquad is the resolver for the DeleteSquad field.
func (r *mutationResolver) DeleteSquad(ctx context.Context, id *string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteSquad - DeleteSquad"))
}

// UpdateScope is the resolver for the UpdateScope field.
func (r *mutationResolver) UpdateScope(ctx context.Context, id string, opts *model.UpdateScope) (*model.Scope, error) {
	panic(fmt.Errorf("not implemented: UpdateScope - UpdateScope"))
}

// UpdateSquad is the resolver for the UpdateSquad field.
func (r *mutationResolver) UpdateSquad(ctx context.Context, id string, opts *model.UpdateSquad) (*model.Squad, error) {
	panic(fmt.Errorf("not implemented: UpdateSquad - UpdateSquad"))
}

// Squads is the resolver for the Squads field.
func (r *queryResolver) Squads(ctx context.Context) (*model.SquadList, error) {
	squads, err := r.squadSvc.List(ctx, squad.ListOpts{
		AccountID:      hillchartsapi.AccountID("1"),
		OrganizationID: hillchartsapi.OrganizationID(os.Getenv("DEFAULT_ORG_ID")),
		UserID:         hillchartsapi.UserID(os.Getenv("TEMP_DEFAULT_USER_ID")),
	})
	if err != nil {
		return nil, err
	}

	return &model.SquadList{
		Data:       buildSquadsResponseFromList(squads.Data),
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
