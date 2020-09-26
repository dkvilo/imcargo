package core

import "encoding/json"

// Response unmarshal json and prepare
func Response(res interface{}) []byte {
	data, _ := json.Marshal(res)
	return data 
}
