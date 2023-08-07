package firebase

import (
	"context"
	"fmt"
	"time"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/firebase/models"
)

func (repo *squadRepo) Create(
	ctx context.Context,
	opts hillchartsapi.SquadCreateOpts,
) (*hillchartsapi.Squad, error) {

	resp, err := repo.addDocWithoutID(ctx, opts)
	if err != nil {
		return nil, err
	}

	return models.ConvertSquadModelToDomain(resp), nil
}

func (repo squadRepo) addDocWithoutID(
	ctx context.Context,
	opts hillchartsapi.SquadCreateOpts,
) (*models.Squad, error) {
	path := fmt.Sprintf("organizations/%s", opts.OrganizationID)
	id := newPrefixedUUID(hillchartsapi.SquadIDPrefix)

	currentTime := time.Now()
	model := &models.Squad{
		AccountID:        string(opts.AccountID),
		CreatedAt:        &currentTime,
		CurrentCycleName: opts.CurrentCycleName,
		ID:               id,
		ModifiedAt:       &currentTime,
		Name:             opts.Name,
		OrganizationID:   string(opts.OrganizationID),
	}

	_, err := repo.db.Doc(path).Collection("squads").Doc(id).Set(ctx, model)
	if err != nil {
		return nil, hcerrors.Wrap("repo.db.Doc(path).Collection(\"squads\")", err)
	}

	return model, nil
}
