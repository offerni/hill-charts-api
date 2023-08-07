package scope

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
)

func (svc Service) Delete(
	ctx context.Context,
	opts DeleteOpts,
) error {
	// TODO: VALIDATION HERE
	err := svc.scopeRepo.Delete(ctx, hillchartsapi.ScopeDeleteOpts{
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
	ID             hillchartsapi.ScopeID
	OrganizationID hillchartsapi.OrganizationID
}
