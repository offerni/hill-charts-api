package firebase

import (
	"cloud.google.com/go/firestore"
	hillchartsapi "github.com/offerni/hill-charts-api"
)

type organizationRepo struct {
	db *firestore.Client
}

func NewOrganizationRepository(db *firestore.Client) (hillchartsapi.OrganizationRepository, error) {
	if db == nil {
		return nil, ErrNoDB
	}

	return &organizationRepo{db}, nil
}
