package conv

import (
	"reflect"
	"testing"
)

func Test_ToBytes(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"string", args{"hello"}, []byte("hello")},
		{"int16", args{int16(1)}, []byte{0, 1}},
		{"int32", args{int32(1)}, []byte{0, 0, 0, 1}},
		{"int64", args{int64(1)}, []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{"uint16", args{uint16(1)}, []byte{0, 1}},
		{"uint32", args{uint32(1)}, []byte{0, 0, 0, 1}},
		{"uint64", args{uint64(1)}, []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{"float32", args{float32(1)}, []byte{63, 128, 0, 0}},
		{"float64", args{float64(1)}, []byte{63, 240, 0, 0, 0, 0, 0, 0}},
		{"bool", args{true}, []byte{1}},
		{"int", args{123}, []byte{0, 0, 0, 123}},
		{"int", args{256}, []byte{0, 0, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBytes(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Log("got:", got, "want:", tt.want)
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
