package utils

import (
	"github.com/goccy/go-json"
)

func JsonEncode(data interface{}) []byte {
	str, _ := JsonEncodeE(data)
	return str
}

func JsonEncodeE(data interface{}) ([]byte, error) {
	str, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return str, nil
}

func JsonEncodes(data interface{}) string {
	str, _ := JsonEncodesE(data)
	return str
}

func JsonEncodesE(data interface{}) (string, error) {
	str, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(str), nil
}

func JsonDecode(data string, v interface{}) {
	_ = JsonDecodeE(data, &v)
}

func JsonDecodeE(data string, v interface{}) error {
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		return err
	}
	return nil
}

func JsonDecodes(data []byte, v interface{}) {
	_ = JsonDecodesE(data, &v)
}

func JsonDecodesE(data []byte, v interface{}) error {
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}
