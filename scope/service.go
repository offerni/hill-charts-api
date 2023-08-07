package scope

import hillchartsapi "github.com/offerni/hill-charts-api"

type Service struct {
	scopeRepo hillchartsapi.ScopeRepository
}

type NewServiceOpts struct {
	ScopeRepository hillchartsapi.ScopeRepository
}

func (opts NewServiceOpts) Validate() error {
	if opts.ScopeRepository == nil {
		return ErrNoScopeRepository
	}

	return nil
}

func NewService(opts NewServiceOpts) (*Service, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &Service{

		scopeRepo: opts.ScopeRepository,
	}, nil
}
