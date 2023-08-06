package firebase

import (
	"cloud.google.com/go/firestore"
	hillchartsapi "github.com/offerni/hill-charts-api"
)

type scopeRepo struct {
	db *firestore.Client
}

func NewScopeRepository(db *firestore.Client) (hillchartsapi.ScopeRepository, error) {
	if db == nil {
		return nil, ErrNoDB
	}

	return &scopeRepo{db}, nil
}
