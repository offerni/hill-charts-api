package squad

import hillchartsapi "github.com/offerni/hill-charts-api"

type Service struct {
	organizationRepo hillchartsapi.OrganizationRepository
	scopeRepo        hillchartsapi.ScopeRepository
	squadRepo        hillchartsapi.SquadRepository
}

type NewServiceOpts struct {
	OrganizationRepository hillchartsapi.OrganizationRepository
	ScopeRepository        hillchartsapi.ScopeRepository
	SquadRepository        hillchartsapi.SquadRepository
}

func (opts NewServiceOpts) Validate() error {
	if opts.OrganizationRepository == nil {
		return ErrNoOrganizationRepository
	}

	if opts.SquadRepository == nil {
		return ErrNoScopeRepository
	}

	if opts.SquadRepository == nil {
		return ErrNoSquadRepository
	}

	return nil
}

func NewService(opts NewServiceOpts) (*Service, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &Service{
		organizationRepo: opts.OrganizationRepository,
		scopeRepo:        opts.ScopeRepository,
		squadRepo:        opts.SquadRepository,
	}, nil
}
