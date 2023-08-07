package firebase

import (
	"context"
	"fmt"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
)

func (repo *scopeRepo) Delete(
	ctx context.Context,
	opts hillchartsapi.ScopeDeleteOpts,
) error {

	err := repo.deleteRecord(ctx, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repo scopeRepo) deleteRecord(
	ctx context.Context,
	opts hillchartsapi.ScopeDeleteOpts,
) error {
	scope, sqErr := repo.findScopeQuery(string(opts.ID))
	if sqErr != nil {
		return sqErr
	}
	path := fmt.Sprintf("organizations/%s/squads/%s", opts.OrganizationID, scope.SquadID)

	_, err := repo.db.Doc(path).Collection("scopes").
		Doc(string(opts.ID)).
		Delete(ctx)
	if err != nil {
		return hcerrors.Wrap("repo.db.Doc(path).Collection(\"scopes\") Delete", err)
	}

	return nil
}
