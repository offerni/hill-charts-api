package firebase

import (
	"cloud.google.com/go/firestore"
	hillchartsapi "github.com/offerni/hill-charts-api"
)

type squadRepo struct {
	db *firestore.Client
}

func NewSquadRepository(db *firestore.Client) (hillchartsapi.SquadRepository, error) {
	if db == nil {
		return nil, ErrNoDB
	}

	return &squadRepo{db}, nil
}
