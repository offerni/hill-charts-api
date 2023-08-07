package firebase

import (
	"context"
	"fmt"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
)

func (repo *squadRepo) Delete(
	ctx context.Context,
	opts hillchartsapi.SquadDeleteOpts,
) error {

	err := repo.deleteRecord(ctx, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repo squadRepo) deleteRecord(
	ctx context.Context,
	opts hillchartsapi.SquadDeleteOpts,
) error {
	path := fmt.Sprintf("organizations/%s", opts.OrganizationID)
	_, err := repo.db.Doc(path).Collection("squads").
		Doc(string(opts.ID)).
		Delete(ctx)
	if err != nil {
		return hcerrors.Wrap("repo.db.Doc(path).Collection(\"squads\") Delete", err)
	}

	return nil
}
