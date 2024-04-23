package convert

import "encoding/json"

func FromJSON[T any](data []byte) (T, error) {
	var t T
	err := json.Unmarshal(data, &t)
	return t, err
}

func ToJSON[T any](t T) ([]byte, error) {
	return json.Marshal(t)
}
