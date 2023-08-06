package firebase

import "errors"

var (
	ErrNoDB          error = errors.New("No DB")
	ErrNoRecordFound error = errors.New("No Record Found")
)
