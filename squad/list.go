package squad

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
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
		squads = append(squads, &FetchResponse{
			ID:               hillchartsapi.SquadID(squad.ID),
			CurrentCycleName: squad.CurrentCycleName,
			Name:             squad.Name,
			OrganizationID:   squad.OrganizationID,
			Scopes:           []*hillchartsapi.Scope{},
		})
	}

	return &ListResponse{
		Data:       squads,
		HasMore:    squadsList.HasMore,
		TotalCount: squadsList.TotalCount,
	}, nil
}

type ListOpts struct {
	AccountID      hillchartsapi.AccountID
	OrganizationID hillchartsapi.OrganizationID
}

type ListResponse struct {
	Data       []*FetchResponse
	HasMore    bool
	TotalCount int
}
