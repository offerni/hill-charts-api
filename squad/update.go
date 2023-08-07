package squad

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
	"github.com/offerni/hill-charts-api/scope"
)

func (svc Service) Update(
	ctx context.Context,
	opts UpdateOpts,
) (*FetchResponse, error) {
	// TODO: VALIDATION HERE

	err := svc.squadRepo.Update(ctx, hillchartsapi.SquadUpdateOpts{
		AccountID:        opts.AccountID,
		CurrentCycleName: opts.CurrentCycleName,
		ID:               opts.ID,
		Name:             opts.Name,
		OrganizationID:   opts.OrganizationID,
	})
	if err != nil {
		return nil, err
	}

	squad, err := svc.squadRepo.Find(ctx, hillchartsapi.SquadFindOpts{
		AccountID:      opts.AccountID,
		ID:             opts.ID,
		OrganizationID: opts.OrganizationID,
	})

	return &FetchResponse{
		CreatedAt:        squad.CreatedAt,
		CurrentCycleName: squad.CurrentCycleName,
		ID:               squad.ID,
		ModifiedAt:       squad.ModifiedAt,
		Name:             squad.Name,
		OrganizationID:   squad.OrganizationID,
		Scopes:           []*scope.FetchResponse{},
	}, nil
}

type UpdateOpts struct {
	AccountID        hillchartsapi.AccountID
	CurrentCycleName *string
	ID               hillchartsapi.SquadID
	Name             *string
	OrganizationID   hillchartsapi.OrganizationID
}
