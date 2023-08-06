package squad

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/scope"
)

func (svc *Service) List(ctx context.Context, opts ListOpts) (*ListResponse, error) {
	_, err := svc.organizationRepo.Find(ctx, hillchartsapi.OrganizationFindOpts{
		AccountID: opts.AccountID,
		ID:        opts.OrganizationID,
	})
	if err != nil {
		return nil, hcerrors.Wrap("svc.organizationRepo.Find", err)
	}

	squadsList, err := svc.squadRepo.FindAll(ctx, hillchartsapi.SquadFindAllOpts{
		AccountID:      opts.AccountID,
		OrganizationID: opts.OrganizationID,
	})
	if err != nil {
		return nil, hcerrors.Wrap("svc.SquadRepo.FindAll", err)
	}

	squads := []*FetchResponse{}
	for _, squad := range squadsList.Squads {
		scopessList, err := svc.scopeRepo.FindAll(ctx, hillchartsapi.ScopeFindAllOpts{
			AccountID:      opts.AccountID,
			OrganizationID: opts.OrganizationID,
			SquadID:        squad.ID,
		})
		if err != nil {
			return nil, hcerrors.Wrap("svc.SquadRepo.FindAll", err)
		}

		squads = append(squads, &FetchResponse{
			ID:               hillchartsapi.SquadID(squad.ID),
			CurrentCycleName: squad.CurrentCycleName,
			Name:             squad.Name,
			OrganizationID:   squad.OrganizationID,
			Scopes:           svc.buildScopesFetchResponse(scopessList.Scopes),
		})
	}

	return &ListResponse{
		Data:       squads,
		HasMore:    squadsList.HasMore,
		TotalCount: squadsList.TotalCount,
	}, nil
}

func (svc Service) buildScopesFetchResponse(scopes []*hillchartsapi.Scope) []*scope.FetchResponse {
	resp := []*scope.FetchResponse{}
	for _, squad := range scopes {
		resp = append(resp, &scope.FetchResponse{
			Colour:   squad.Colour,
			ID:       squad.ID,
			Name:     squad.Name,
			Progress: squad.Progress,
			SquadID:  squad.SquadID,
		})
	}

	return resp
}

type ListOpts struct {
	AccountID      hillchartsapi.AccountID
	OrganizationID hillchartsapi.OrganizationID
	SquadID        hillchartsapi.SquadID
}

type ListResponse struct {
	Data       []*FetchResponse
	HasMore    bool
	TotalCount int
}
