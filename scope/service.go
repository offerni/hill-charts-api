package scope

import hillchartsapi "github.com/offerni/hill-charts-api"

type Service struct {
	organizationRepo hillchartsapi.OrganizationRepository
	scopeRepo        hillchartsapi.ScopeRepository
	userRepo         hillchartsapi.UserRepository
}

type NewServiceOpts struct {
	OrganizationRepository hillchartsapi.OrganizationRepository
	ScopeRepository        hillchartsapi.ScopeRepository
	UserRepository         hillchartsapi.UserRepository
}

func (opts NewServiceOpts) Validate() error {
	if opts.ScopeRepository == nil {
		return ErrNoScopeRepository
	}

	if opts.OrganizationRepository == nil {
		return ErrNoOrganizationRepository
	}

	if opts.UserRepository == nil {
		return ErrNoUserRepository
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
		userRepo:         opts.UserRepository,
	}, nil
}
