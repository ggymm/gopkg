package utils

import "testing"

func TestExtractZip(t *testing.T) {
	type args struct {
		filename string
		dstDir   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test1", args{"test1.zip", "test1"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractZip(tt.args.filename, tt.args.dstDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractZip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExtractZip() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompressZip(t *testing.T) {
	type args struct {
		dir  string
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"test1", "test1.zip"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CompressZip(tt.args.dir, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("CompressZip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
