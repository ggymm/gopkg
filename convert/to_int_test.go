package convert

import (
	"github.com/ggymm/gopkg/types"
	"testing"
)

func Test_toSignedE_Int(t *testing.T) {
	type args struct {
		value   any
		options []Options
	}
	type testCase[T types.Signed] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[int]{
		{name: "int", args: args{value: 123}, want: 123, wantErr: false},
		{name: "int16", args: args{value: int16(123)}, want: 123, wantErr: false},
		{name: "int32", args: args{value: int32(123)}, want: 123, wantErr: false},
		{name: "int64", args: args{value: int64(123)}, want: 123, wantErr: false},
		{name: "uint", args: args{value: uint(123)}, want: 123, wantErr: false},
		{name: "uint16", args: args{value: uint16(123)}, want: 123, wantErr: false},
		{name: "uint32", args: args{value: uint32(123)}, want: 123, wantErr: false},
		{name: "uint64", args: args{value: uint64(123)}, want: 123, wantErr: false},
		{name: "float32", args: args{value: float32(123)}, want: 123, wantErr: false},
		{name: "float64", args: args{value: float64(123)}, want: 123, wantErr: false},
		{name: "string", args: args{value: "123"}, want: 123, wantErr: false},
		{name: "[]byte", args: args{value: []byte{0, 0, 0, 123}}, want: 123, wantErr: false},
		{name: "[]byte", args: args{value: []byte{0, 0, 1, 0}}, want: 256, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toSignedE[int](tt.args.value, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Log("got:", got, "want:", tt.want)
				t.Errorf("toSignedE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Log("got:", got, "want:", tt.want)
				t.Errorf("toSignedE() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toSignedE_Int32(t *testing.T) {
	type args struct {
		value   any
		options []Options
	}
	type testCase[T types.Signed] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[int32]{
		{name: "int", args: args{value: 123}, want: int32(123), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toSignedE[int32](tt.args.value, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Log("got:", got, "want:", tt.want)
				t.Errorf("toSignedE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Log("got:", got, "want:", tt.want)
				t.Errorf("toSignedE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
