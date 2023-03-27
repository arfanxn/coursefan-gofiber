package resources

import "encoding/json"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Bytes returns Response as bytes
func (response Response) Bytes() []byte {
	bytes, err := json.Marshal(response)
	if err != nil {
		return []byte{}
	}
	return bytes
}
