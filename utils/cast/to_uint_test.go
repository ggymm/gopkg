package cast

import (
	"github.com/ggymm/gopkg/types"
	"testing"
)

func TestToUnsignedE(t *testing.T) {
	type args struct {
		value   any
		options []Options
	}
	type testCase[T types.Unsigned] struct {
		name    string
		args    args
		want    T
		wantErr bool
	}
	tests := []testCase[uint]{
		{name: "int", args: args{value: 123}, want: 123, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToUnsignedE[uint](tt.args.value, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUnsignedE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToUnsignedE() got = %v, want %v", got, tt.want)
			}
		})
	}
}
