package util

import (
	"os"
	"path/filepath"
	"strings"
)

//CurrentDirectory 获取程序运行路径
func CurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	return OsPathConverter(dir)
}
func OsPathConverter(source string) string {
	return strings.Replace(strings.Replace(source, "\\", string(os.PathSeparator), -1), "/", string(os.PathSeparator), -1)
}
