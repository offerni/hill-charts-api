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

	if err := svc.ValidateUserAndOrganization(ctx, ValidateUserAndOrganizationOpts{
		AccountID:      opts.AccountID,
		OrganizationID: opts.OrganizationID,
		UserID:         opts.UserID,
	}); err != nil {
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
	UserID         hillchartsapi.UserID
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

type ValidateUserAndOrganizationOpts struct {
	AccountID      hillchartsapi.AccountID
	OrganizationID hillchartsapi.OrganizationID
	UserID         hillchartsapi.UserID
}

func (svc *Service) ValidateUserAndOrganization(
	ctx context.Context,
	opts ValidateUserAndOrganizationOpts,
) error {

	org, err := svc.organizationRepo.Find(ctx, hillchartsapi.OrganizationFindOpts{
		AccountID: opts.AccountID,
		ID:        opts.OrganizationID,
	})
	if err != nil {
		return ErrOganizationNotFound
	}

	user, err := svc.userRepo.Find(ctx, hillchartsapi.UserFindOpts{
		AccountID:      opts.AccountID,
		ID:             opts.UserID,
		OrganizationID: opts.OrganizationID,
	})
	if err != nil {
		return ErrUserNotFound
	}

	if user.OrganizationID != org.ID {
		return ErrUserDoesNotBelongToOrganization
	}

	return nil
}
