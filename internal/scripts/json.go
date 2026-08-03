package scripts

import (
	"encoding/json"
	"io"
)

func DecodeJSON[T any](r io.Reader) (T, error) {
	var data T

	decoder := json.NewDecoder(r)
	err := decoder.Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func EncodeJSON[T any](w io.Writer, data T) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}
