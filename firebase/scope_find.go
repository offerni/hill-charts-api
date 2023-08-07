package firebase

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/firebase/models"
	"google.golang.org/api/iterator"
)

func (repo *scopeRepo) Find(
	ctx context.Context,
	opts hillchartsapi.ScopeFindOpts,
) (*hillchartsapi.Scope, error) {

	scope, err := repo.findScopeQuery(string(opts.ID))
	if err != nil {
		return nil, err
	}

	return models.ConvertScopeModelToDomain(scope), nil
}

func (repo scopeRepo) findScopeQuery(id string) (*models.Scope, error) {
	ctx := context.Background()

	iter := repo.db.CollectionGroup("scopes").Where("id", "==", id).Documents(ctx)
	var scopeModel models.Scope
	var scopeModelMultiple []models.Scope

	for {
		scopeRef, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, hcerrors.Wrap("iter.Next()", err)
		}

		if err := scopeRef.DataTo(&scopeModel); err != nil {
			return nil, hcerrors.Wrap("scopeRef.DataTo(&scopeData)", err)
		}

		scopeModelMultiple = append(scopeModelMultiple, scopeModel)
	}

	return &scopeModel, nil
}
