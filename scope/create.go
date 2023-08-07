package scope

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
)

func (svc Service) Create(
	ctx context.Context,
	opts CreateOpts,
) (*FetchResponse, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	scope, err := svc.scopeRepo.Create(ctx, hillchartsapi.ScopeCreateOpts{
		AccountID:      opts.AccountID,
		Colour:         opts.Colour,
		Name:           opts.Name,
		OrganizationID: opts.OrganizationID,
		Progress:       opts.Progress,
		SquadID:        opts.SquadID,
	})

	if err != nil {
		return nil, err
	}

	return &FetchResponse{
		Colour:     scope.Colour,
		CreatedAt:  scope.CreatedAt,
		ID:         scope.ID,
		ModifiedAt: scope.ModifiedAt,
		Name:       scope.Name,
		Progress:   scope.Progress,
		SquadID:    scope.SquadID,
	}, nil
}

type CreateOpts struct {
	AccountID      hillchartsapi.AccountID
	Colour         string
	Name           string
	OrganizationID hillchartsapi.OrganizationID
	Progress       float32
	SquadID        hillchartsapi.SquadID
}

func (opts CreateOpts) Validate() error {
	if opts.AccountID == "" {
		return hillchartsapi.ErrNoAccountID
	}

	if opts.Colour == "" {
		return ErrNoColour
	}

	if opts.Name == "" {
		return ErrNoName
	}

	if opts.OrganizationID == "" {
		return ErrNoOrganizationID
	}

	if opts.SquadID == "" {
		return ErrNoSquadID
	}

	return nil
}
