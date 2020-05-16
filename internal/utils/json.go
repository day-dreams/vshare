package utils

import "encoding/json"

// ToPrettyJSON 转json string，
func ToPrettyJSON(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}
