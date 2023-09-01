package convert

import "testing"

func TestToString(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{1}, "1"},
		{"test2", args{1.1}, "1.1"},
		{"test3", args{"1"}, "1"},
		{"test4", args{true}, "true"},
		{"test5", args{[]byte("1")}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.i); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
