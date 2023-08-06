package firebase

import (
	"context"
	"fmt"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/firebase/models"
	"google.golang.org/api/iterator"
)

func (repo *scopeRepo) FindAll(
	ctx context.Context,
	opts hillchartsapi.ScopeFindAllOpts,
) (*hillchartsapi.ScopeList, error) {

	resp, err := repo.getScopesCollection(ctx, opts)
	if err != nil {
		return nil, err
	}

	scopes := []*hillchartsapi.Scope{}
	for _, mScope := range resp {
		scopes = append(scopes, models.ConvertScopeModelToDomain(mScope))
	}

	return &hillchartsapi.ScopeList{
		Scopes: scopes,
		PaginatedList: &hillchartsapi.PaginatedList{
			HasMore:    false,
			TotalCount: len(scopes),
		},
	}, nil
}

func (repo *scopeRepo) getScopesCollection(
	ctx context.Context,
	opts hillchartsapi.ScopeFindAllOpts,
) ([]*models.Scope, error) {
	path := fmt.Sprintf("organizations/%s/squads/%s", opts.OrganizationID, opts.SquadID)

	iter := repo.db.Doc(path).Collection("scopes").Where("squad_id", "==", opts.SquadID).Documents(ctx)

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

	scopeModels := []*models.Scope{}
	for _, scope := range scopeModelMultiple {
		scopeModels = append(scopeModels, &models.Scope{
			AccountID: scope.AccountID,
			Colour:    scope.Colour,
			ID:        scope.ID,
			Name:      scope.Name,
			Progress:  scope.Progress,
			SquadID:   scope.SquadID,
		})

	}

	return scopeModels, nil
}
