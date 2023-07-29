package utils

import (
	"github.com/goccy/go-json"

	"github.com/ggymm/gopkg/constant"
	"github.com/ggymm/gopkg/log"
)

func JsonEncode(data interface{}) string {
	str, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg(constant.JsonEncodeError)
		return ""
	}
	return string(str)
}

func JsonDecode(data string, v interface{}) {
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		log.Error().Err(err).Msg(constant.JsonDecodeError)
		return
	}
}
