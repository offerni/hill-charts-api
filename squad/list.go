package squad

import (
	"context"

	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
	"github.com/offerni/hill-charts-api/scope"
)

func (svc *Service) List(ctx context.Context, opts ListOpts) (*ListResponse, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	org, err := svc.organizationRepo.Find(ctx, hillchartsapi.OrganizationFindOpts{
		AccountID: opts.AccountID,
		ID:        opts.OrganizationID,
	})
	if err != nil {
		return nil, ErrOganizationNotFound
	}

	user, err := svc.userRepo.Find(ctx, hillchartsapi.UserFindOpts{
		AccountID:      opts.AccountID,
		ID:             opts.UserID,
		OrganizationID: opts.OrganizationID,
	})
	if err != nil {
		return nil, ErrUserNotFound
	}

	if user.OrganizationID != org.ID {
		return nil, ErrUserDoesNotBelongToOrganization
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
			CreatedAt:        squad.CreatedAt,
			CurrentCycleName: squad.CurrentCycleName,
			ID:               hillchartsapi.SquadID(squad.ID),
			ModifiedAt:       squad.ModifiedAt,
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
	for _, s := range scopes {
		resp = append(resp, &scope.FetchResponse{
			Colour:     s.Colour,
			CreatedAt:  s.CreatedAt,
			ID:         s.ID,
			ModifiedAt: s.ModifiedAt,
			Name:       s.Name,
			Progress:   s.Progress,
			SquadID:    s.SquadID,
		})
	}

	return resp
}

type ListOpts struct {
	AccountID      hillchartsapi.AccountID
	OrganizationID hillchartsapi.OrganizationID
	SquadID        hillchartsapi.SquadID
	UserID         hillchartsapi.UserID
}

type ListResponse struct {
	Data       []*FetchResponse
	HasMore    bool
	TotalCount int
}

func (opts ListOpts) Validate() error {
	if opts.AccountID == "" {
		return hillchartsapi.ErrNoAccountID
	}

	if opts.OrganizationID == "" {
		return hillchartsapi.ErrNoOrganizationID
	}

	if opts.SquadID == "" {
		return hillchartsapi.ErrNoSquadID
	}

	if opts.UserID == "" {
		return hillchartsapi.ErrNoUserID
	}

	return nil
}
