package firebase

import (
	"context"
	"fmt"
	"time"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/firebase/models"
)

func (repo *scopeRepo) Create(
	ctx context.Context,
	opts hillchartsapi.ScopeCreateOpts,
) (*hillchartsapi.Scope, error) {

	resp, err := repo.addDocWithoutID(ctx, opts)
	if err != nil {
		return nil, err
	}

	return models.ConvertScopeModelToDomain(resp), nil
}

func (repo scopeRepo) addDocWithoutID(
	ctx context.Context,
	opts hillchartsapi.ScopeCreateOpts,
) (*models.Scope, error) {
	path := fmt.Sprintf("organizations/%s/squads/%s", opts.OrganizationID, opts.SquadID)
	id := newPrefixedUUID(hillchartsapi.ScopeIDPrefix)

	currentTime := time.Now()
	model := &models.Scope{
		AccountID:  string(opts.AccountID),
		Colour:     opts.Colour,
		CreatedAt:  &currentTime,
		ID:         id,
		ModifiedAt: &currentTime,
		Name:       opts.Name,
		Progress:   opts.Progress,
		SquadID:    string(opts.SquadID),
	}

	_, err := repo.db.Doc(path).Collection("scopes").Doc(id).Set(ctx, model)
	if err != nil {
		return nil, hcerrors.Wrap("repo.db.Doc(path).Collection(\"scopes\")", err)
	}

	return model, nil
}
