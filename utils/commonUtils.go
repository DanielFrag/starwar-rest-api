package utils

import (
	"bytes"
	"encoding/json"
)

//FormatJSON format the interface 'i' to json
func FormatJSON(i interface{}) []byte {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(i)
	return b.Bytes()
}
