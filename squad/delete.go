package squad

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
)

func (svc Service) Delete(
	ctx context.Context,
	opts DeleteOpts,
) error {
	// TODO: VALIDATION HERE
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
}
