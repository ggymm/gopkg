package utils

import (
	"github.com/goccy/go-json"
)

func JsonEncode(data any) []byte {
	str, _ := JsonEncodeE(data)
	return str
}

func JsonEncodeE(data any) ([]byte, error) {
	str, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return str, nil
}

func JsonEncodes(data any) string {
	str, _ := JsonEncodesE(data)
	return str
}

func JsonEncodesE(data any) (string, error) {
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func JsonDecode(data string, v any) {
	_ = JsonDecodeE(data, v)
}

func JsonDecodeE(data string, v any) error {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return err
	}
	return nil
}

func JsonDecodes(data []byte, v any) {
	_ = JsonDecodesE(data, v)
}

func JsonDecodesE(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}
