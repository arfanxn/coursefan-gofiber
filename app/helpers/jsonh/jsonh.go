package jsonh

import "github.com/clarketm/json"

// MustMarshal marshals json or panics when fails
func MustMarshal(value any) []byte {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return bytes
}
