package nullh

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

func NullStringToNullUUID(nullStr null.String) (nullUUID uuid.NullUUID) {
	parsedUUID, err := uuid.Parse(nullStr.String)
	if (err != nil) || !nullStr.Valid {
		return
	}
	nullUUID.UUID = parsedUUID
	nullUUID.Valid = true
	return
}
