package testutil

import (
	"encoding/json"
)

func MustJsonify(data interface{}) string {
	d, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic(err)
	}
	return string(d)
}
