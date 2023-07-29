package utils

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	type args[T any] struct {
		boolExpression   bool
		trueReturnValue  T
		falseReturnValue T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[string]{
		{
			name: "测试是",
			args: struct {
				boolExpression   bool
				trueReturnValue  string
				falseReturnValue string
			}{boolExpression: true, trueReturnValue: "是", falseReturnValue: "否"},
			want: "是",
		},
		{
			name: "测试否",
			args: struct {
				boolExpression   bool
				trueReturnValue  string
				falseReturnValue string
			}{boolExpression: false, trueReturnValue: "是", falseReturnValue: "否"},
			want: "否",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := If(tt.args.boolExpression, tt.args.trueReturnValue, tt.args.falseReturnValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("If() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIfByFunc(t *testing.T) {
	type args[T any] struct {
		boolExpression          bool
		trueFuncForReturnValue  func() T
		falseFuncForReturnValue func() T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[string]{
		{
			name: "测试是",
			args: args[string]{
				boolExpression: true,
				trueFuncForReturnValue: func() string {
					return "是"
				},
				falseFuncForReturnValue: func() string {
					return "否"
				},
			},
			want: "是",
		},
		{
			name: "测试否",
			args: args[string]{
				boolExpression: false,
				trueFuncForReturnValue: func() string {
					return "是"
				},
				falseFuncForReturnValue: func() string {
					return "否"
				},
			},
			want: "否",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IfByFunc(tt.args.boolExpression, tt.args.trueFuncForReturnValue, tt.args.falseFuncForReturnValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IfByFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
