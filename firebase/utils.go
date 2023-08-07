package firebase

import (
	"fmt"

	"github.com/google/uuid"
)

func newPrefixedUUID(prefix string) string {
	uuid := uuid.New()

	return fmt.Sprintf("%s_%s", prefix, uuid)
}
