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

	squads := make([]*hillchartsapi.Squad, len(resp))

	for i, squad := range resp {
		squads[i].AccountID = hillchartsapi.AccountID(squad.AccountID)
		squads[i].CurrentCycleName = ""
		squads[i].ID = hillchartsapi.SquadID(squad.ID)
		squads[i].Name = squad.Name
		squads[i].OrganizationID = hillchartsapi.OrganizationID(squad.OrganizationID)
	}

	spew.Dump(resp)

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

	var squadData []*models.Squad

	for {
		squadRef, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, hcerrors.Wrap("iter.Next()", err)
		}
		fmt.Println(squadRef.Data())
	}
	return squadData, nil
}
