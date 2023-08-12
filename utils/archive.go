package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ExtractZip 解压 zip 文件
// name: 压缩文件名
// dstDir: 目标目录
func ExtractZip(name string, dstDir string) (string, error) {
	var (
		err     error
		zipFile *zip.ReadCloser
	)

	if dstDir == "" {
		// 创建临时目录
		dstDir, err = os.MkdirTemp("", "temp-zip-*")
		if err != nil {
			return dstDir, err
		}
	}

	// 读取压缩文件
	zipFile, err = zip.OpenReader(name)
	if err != nil {
		return dstDir, err
	}

	for _, file := range zipFile.Reader.File {
		// 忽略目录
		if file.FileInfo().IsDir() {
			continue
		}

		var (
			dstPath = filepath.Join(dstDir, file.Name)

			dstFile *os.File
			srcFile io.ReadCloser
		)

		// 如果文件目录不存在，则创建
		dstDir := filepath.Dir(dstPath)
		if _, err = os.Stat(dstDir); err != nil {
			if err = os.MkdirAll(dstDir, os.ModePerm); err != nil {
				return dstDir, err
			}
		}

		// 创建目标文件
		dstFile, err = os.OpenFile(dstPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return dstDir, err
		}

		// 打开源文件
		srcFile, err = file.Open()
		if err != nil {
			return dstDir, err
		}

		// 拷贝文件
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return dstDir, err
		}

		// 关闭文件
		_ = srcFile.Close()
		_ = dstFile.Close()
	}

	return dstDir, nil
}

// CompressZip 压缩文件夹
// dir: 目录名
// name: 压缩文件名
func CompressZip(dir string, name string) (err error) {
	var (
		dstFile   *os.File
		dstWriter *zip.Writer
	)

	// 创建压缩文件
	dstFile, err = os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		_ = dstFile.Close()
	}()

	// 创建压缩器
	dstWriter = zip.NewWriter(dstFile)

	// 遍历目录
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var (
			dst io.Writer

			relPath string
			relFile *os.File
		)

		// 获取相对路径
		relPath, err = filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		// 忽略当前目录
		if relPath == "." {
			return nil
		}

		if info.IsDir() {
			// 创建目录
			_, err = dstWriter.Create(relPath + "/")
			if err != nil {
				return err
			}
		} else {
			// 创建文件
			relFile, err = os.Open(path)
			if err != nil {
				return err
			}

			dst, err = dstWriter.Create(relPath)
			if err != nil {
				return err
			}

			// 拷贝文件
			_, err = io.Copy(dst, relFile)
			if err != nil {
				return err
			}

			_ = relFile.Close()
		}

		return nil
	})
	if err != nil {
		return err
	}

	// 关闭压缩器
	// 必须最后关闭
	err = dstWriter.Close()
	if err != nil {
		return err
	}

	return nil
}
