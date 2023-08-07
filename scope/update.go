package scope

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
)

func (svc Service) Update(
	ctx context.Context,
	opts UpdateOpts,
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

	err := svc.scopeRepo.Update(ctx, hillchartsapi.ScopeUpdateOpts{
		AccountID:      opts.AccountID,
		Colour:         opts.Colour,
		ID:             opts.ID,
		Name:           opts.Name,
		OrganizationID: opts.OrganizationID,
		Progress:       opts.Progress,
	})
	if err != nil {
		return nil, err
	}

	scope, err := svc.scopeRepo.Find(ctx, hillchartsapi.ScopeFindOpts{
		AccountID:      opts.AccountID,
		ID:             opts.ID,
		OrganizationID: opts.OrganizationID,
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

type UpdateOpts struct {
	AccountID      hillchartsapi.AccountID
	Colour         *string
	ID             hillchartsapi.ScopeID
	Name           *string
	OrganizationID hillchartsapi.OrganizationID
	Progress       *float32
	UserID         hillchartsapi.UserID
}

func (opts UpdateOpts) Validate() error {
	if opts.AccountID == "" {
		return hillchartsapi.ErrNoAccountID
	}

	if opts.OrganizationID == "" {
		return ErrNoOrganizationID
	}

	if opts.ID == "" {
		return ErrNoID
	}

	return nil
}
