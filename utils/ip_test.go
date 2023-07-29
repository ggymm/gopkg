package utils

import "testing"

func TestGetLocalIp(t *testing.T) {
	t.Log(GetLocalIp())
}

func TestGetLocalIpList(t *testing.T) {
	t.Log(GetLocalIpList(true))
}
