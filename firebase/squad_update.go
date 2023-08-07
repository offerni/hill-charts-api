package firebase

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
)

func (repo *squadRepo) Update(
	ctx context.Context,
	opts hillchartsapi.SquadUpdateOpts,
) error {

	err := repo.updateRecords(ctx, opts)
	if err != nil {
		return err
	}

	return nil
}

func (repo squadRepo) updateRecords(
	ctx context.Context,
	opts hillchartsapi.SquadUpdateOpts,
) error {
	path := fmt.Sprintf("organizations/%s", opts.OrganizationID)

	updateFields := repo.buildUpdateFields(opts)

	_, err := repo.db.Doc(path).Collection("squads").
		Doc(string(opts.ID)).
		Update(ctx, updateFields)
	if err != nil {
		return hcerrors.Wrap("repo.db.Doc(path).Collection(\"squads\") Update", err)
	}

	return nil
}

func (repo squadRepo) buildUpdateFields(opts hillchartsapi.SquadUpdateOpts) []firestore.Update {
	currentTime := time.Now()
	updateFields := []firestore.Update{
		// TODO: create consts in the model for the fields
		{
			Path:  "modified_at",
			Value: currentTime,
		},
	}

	if opts.CurrentCycleName != nil {
		updateFields = append(updateFields, firestore.Update{
			Path:  "current_cycle_name",
			Value: opts.CurrentCycleName,
		})
	}

	if opts.Name != nil {
		updateFields = append(updateFields, firestore.Update{
			Path:  "name",
			Value: opts.Name,
		})
	}
	return updateFields
}
