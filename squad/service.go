package squad

import hillchartsapi "github.com/offerni/hill-charts-api"

type Service struct {
	organizationRepo hillchartsapi.OrganizationRepository
}

type NewServiceOpts struct {
	OrganizationRepository hillchartsapi.OrganizationRepository
}

func (opts NewServiceOpts) Validate() error {
	if opts.OrganizationRepository == nil {
		return ErrNoOrganizationRepository
	}

	return nil
}

func NewService(opts NewServiceOpts) (*Service, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &Service{
		organizationRepo: opts.OrganizationRepository,
	}, nil
}
