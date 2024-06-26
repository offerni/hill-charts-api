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
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	if err := svc.ValidateUserAndOrganization(ctx, ValidateUserAndOrganizationOpts{
		AccountID:      opts.AccountID,
		OrganizationID: opts.OrganizationID,
		UserID:         opts.UserID,
	}); err != nil {
		return nil, err
	}

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
		ModifiedAt:       squad.ModifiedAt,
		Name:             squad.Name,
		OrganizationID:   squad.OrganizationID,
		Scopes:           []*scope.FetchResponse{},
	}, nil
}

type CreateOpts struct {
	AccountID        hillchartsapi.AccountID
	CurrentCycleName *string
	Name             string
	OrganizationID   hillchartsapi.OrganizationID
	UserID           hillchartsapi.UserID
}

func (opts CreateOpts) Validate() error {
	if opts.AccountID == "" {
		return ErrNoAccountID
	}

	if opts.OrganizationID == "" {
		return ErrNoName
	}

	if opts.Name == "" {
		return ErrNoName
	}

	return nil
}
