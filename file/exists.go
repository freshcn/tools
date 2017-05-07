package file

import "os"

// Exists 判断文件是否存在
// path为文件路径
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
