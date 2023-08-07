package squad

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
)

func (svc Service) Delete(
	ctx context.Context,
	opts DeleteOpts,
) error {
	if err := opts.Validate(); err != nil {
		return err
	}

	if err := svc.ValidateUserAndOrganization(ctx, ValidateUserAndOrganizationOpts{
		AccountID:      opts.AccountID,
		OrganizationID: opts.OrganizationID,
		UserID:         opts.UserID,
	}); err != nil {
		return err
	}

	err := svc.squadRepo.Delete(ctx, hillchartsapi.SquadDeleteOpts{
		AccountID:      opts.AccountID,
		ID:             opts.ID,
		OrganizationID: opts.OrganizationID,
	})
	if err != nil {
		return err
	}

	return nil
}

type DeleteOpts struct {
	AccountID      hillchartsapi.AccountID
	ID             hillchartsapi.SquadID
	OrganizationID hillchartsapi.OrganizationID
	UserID         hillchartsapi.UserID
}

func (opts DeleteOpts) Validate() error {
	if opts.AccountID == "" {
		return ErrNoAccountID
	}

	if opts.OrganizationID == "" {
		return ErrNoName
	}

	if opts.ID == "" {
		return ErrNoID
	}

	return nil
}
