package utils

import (
	"reflect"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	type args struct {
		data string
		v    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{"{\"name\": \"json\"}", &map[string]interface{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonDecode(tt.args.data, tt.args.v)
			t.Log(tt.args.v)
		})
	}
}

func TestJsonDecodes(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"test1", args{[]byte("{\"name\": \"json\"}"), &map[string]interface{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonDecodes(tt.args.data, tt.args.v)
			t.Log(tt.args.v)
		})
	}
}

func TestJsonEncode(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"test1", args{map[string]interface{}{"name": "json"}}, []byte("{\"name\":\"json\"}")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JsonEncode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonEncodes(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{map[string]interface{}{"name": "json"}}, "{\"name\":\"json\"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JsonEncodes(tt.args.data); got != tt.want {
				t.Errorf("JsonEncodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
