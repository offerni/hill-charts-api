package firebase

import (
	"context"
	"fmt"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/firebase/models"
)

func (repo *squadRepo) Find(
	ctx context.Context,
	opts hillchartsapi.SquadFindOpts,
) (*hillchartsapi.Squad, error) {

	path := fmt.Sprintf("organizations/%s/squads/%s", opts.OrganizationID, opts.ID)

	orgSnap, err := repo.db.Doc(path).Get(ctx)
	if err != nil {
		return nil, hcerrors.Wrap("repo.db.Doc(path).Get(ctx)", err)
	}

	var orgData *models.Squad
	if err := orgSnap.DataTo(&orgData); err != nil {
		return nil, hcerrors.Wrap("orgSnap.DataTo(&orgData)", err)
	}

	return models.ConvertSquadModelToDomain(orgData), nil
}
