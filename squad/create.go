package squad

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
	"github.com/offerni/hill-charts-api/scope"
)

func (svc Service) Create(
	ctx context.Context,
	opts CreateOpts,
) (*FetchResponse, error) {
	// TODO: VALIDATION HERE

	squad, err := svc.squadRepo.Create(ctx, hillchartsapi.SquadCreateOpts{
		AccountID:        opts.AccountID,
		CurrentCycleName: opts.CurrentCycleName,
		Name:             opts.Name,
		OrganizationID:   opts.OrganizationID,
	})

	if err != nil {
		return nil, err
	}

	return &FetchResponse{
		CreatedAt:        squad.CreatedAt,
		CurrentCycleName: squad.CurrentCycleName,
		ID:               squad.ID,
		ModifieddAt:      squad.ModifiedAt,
		Name:             squad.Name,
		OrganizationID:   squad.OrganizationID,
		Scopes:           []*scope.FetchResponse{},
	}, nil
}

type CreateOpts struct {
	AccountID        hillchartsapi.AccountID
	CurrentCycleName string
	Name             string
	OrganizationID   hillchartsapi.OrganizationID
}
