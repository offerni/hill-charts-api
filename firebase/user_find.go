package firebase

import (
	"context"
	"fmt"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/firebase/models"
)

func (repo *userRepo) Find(
	ctx context.Context,
	opts hillchartsapi.UserFindOpts,
) (*hillchartsapi.User, error) {

	path := fmt.Sprintf("organizations/%s/users/%s", opts.OrganizationID, opts.ID)

	orgSnap, err := repo.db.Doc(path).Get(ctx)
	if err != nil {
		return nil, hcerrors.Wrap("repo.db.Doc(path).Get(ctx)", err)
	}

	var orgData *models.User
	if err := orgSnap.DataTo(&orgData); err != nil {
		return nil, hcerrors.Wrap("orgSnap.DataTo(&orgData)", err)
	}

	return models.ConvertUserModelToDomain(orgData), nil
}
