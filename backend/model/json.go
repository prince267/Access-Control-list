package model

import (
	"encoding/json"
)

func Stringify(obj interface{}) string {
	byts, _ := json.Marshal(obj)
	return string(byts)
}
