package firebase

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/firebase/models"
	"google.golang.org/api/iterator"
)

func (repo *squadRepo) FindAll(
	ctx context.Context,
	opts hillchartsapi.SquadFindAllOpts,
) (*hillchartsapi.SquadList, error) {

	resp, err := repo.getSquadsCollection(ctx, opts)
	if err != nil {
		return nil, err
	}

	squads := []*hillchartsapi.Squad{}
	for _, mSquad := range resp {
		squads = append(squads, models.ConvertSquadModelToDomain(mSquad))
	}

	return &hillchartsapi.SquadList{
		Squads: squads,
		PaginatedList: &hillchartsapi.PaginatedList{
			HasMore:    false,
			TotalCount: len(squads),
		},
	}, nil
}

func (repo *squadRepo) getSquadsCollection(
	ctx context.Context,
	opts hillchartsapi.SquadFindAllOpts,
) ([]*models.Squad, error) {
	path := fmt.Sprintf("organizations/%s", opts.OrganizationID)

	iter := repo.db.Doc(path).Collection("squads").Where("organization_id", "==", opts.OrganizationID).Documents(ctx)

	var squadModel *models.Squad
	var squadModelMultiple = []*models.Squad{}
	for {
		squadRef, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, hcerrors.Wrap("iter.Next()", err)
		}

		if err := squadRef.DataTo(&squadModel); err != nil {
			return nil, hcerrors.Wrap("squadRef.DataTo(&squadData)", err)
		}

		spew.Dump(squadRef.Data())

		squadModelMultiple = append(squadModelMultiple, squadModel)
	}

	return squadModelMultiple, nil
}
