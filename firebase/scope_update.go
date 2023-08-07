package firebase

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
)

func (repo *scopeRepo) Update(
	ctx context.Context,
	opts hillchartsapi.ScopeUpdateOpts,
) error {

	err := repo.updateRecords(ctx, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repo scopeRepo) updateRecords(
	ctx context.Context,
	opts hillchartsapi.ScopeUpdateOpts,
) error {
	scope, sqErr := repo.findScopeQuery(string(opts.ID))
	if sqErr != nil {
		return sqErr
	}
	path := fmt.Sprintf("organizations/%s/squads/%s", opts.OrganizationID, scope.SquadID)

	updateFields := repo.buildUpdateFields(opts)

	_, err := repo.db.Doc(path).Collection("scopes").
		Doc(string(opts.ID)).
		Update(ctx, updateFields)
	if err != nil {
		return hcerrors.Wrap("repo.db.Doc(path).Collection(\"scopes\") Update", err)
	}

	return nil
}

func (repo scopeRepo) buildUpdateFields(opts hillchartsapi.ScopeUpdateOpts) []firestore.Update {
	currentTime := time.Now()
	updateFields := []firestore.Update{
		// TODO: create consts in the model for the fields
		{
			Path:  "modified_at",
			Value: currentTime,
		},
	}

	if opts.Colour != nil {
		updateFields = append(updateFields, firestore.Update{
			Path:  "colour",
			Value: opts.Colour,
		})
	}

	if opts.Name != nil {
		updateFields = append(updateFields, firestore.Update{
			Path:  "name",
			Value: opts.Name,
		})
	}

	if opts.Progress != nil {
		updateFields = append(updateFields, firestore.Update{
			Path:  "progress",
			Value: opts.Progress,
		})
	}

	return updateFields
}
