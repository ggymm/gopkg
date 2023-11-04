package files

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string) (err error) {
	err = Mkdir(filepath.Dir(dst))
	if err != nil {
		return err
	}

	var (
		srcFile *os.File
		dstFile *os.File

		srcFileInfo os.FileInfo
	)

	// 读取源文件
	srcFile, err = os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		_ = srcFile.Close()
	}()

	// 创建目标文件
	dstFile, err = os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = dstFile.Close()
	}()

	// 复制文件
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// 获取源文件信息
	srcFileInfo, err = os.Stat(src)
	if err != nil {
		return err
	}

	// 修改目标文件权限
	return os.Chmod(dst, srcFileInfo.Mode())
}
