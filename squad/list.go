package squad

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	hillchartsapi "github.com/offerni/hill-charts-api"
	hcerrors "github.com/offerni/hill-charts-api/errors"
)

func (svc *Service) List(ctx context.Context, opts ListOpts) (*ListResponse, error) {
	org, err := svc.organizationRepo.Find(ctx, hillchartsapi.OrganizationFindOpts{
		AccountID: opts.AccountID,
		ID:        opts.OrganizationID,
	})
	if err != nil {
		return nil, hcerrors.Wrap("svc.organizationRepo.Find", err)
	}

	spew.Dump(org)

	return &ListResponse{
		HasMore:    false,
		TotalCount: 10, // revisit all this
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
