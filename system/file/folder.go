package file

import (
	"os"
	"path"
	"runtime"
)

func CurrentPath() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		absPath = path.Dir(filename)
	}

	return absPath
}

// Mkdir create dir if not exists
func Mkdir(dir string) (err error) {
	if _, err = os.Stat(dir); err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
