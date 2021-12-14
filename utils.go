package UniStorage

import (
	"fmt"
	"github.com/Dereking/utils/StrUtil"
	"os"
)

// golang新版本的应该
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

//IdToPath 根据唯一id计算出存储路径。baseDir不包含末尾斜杠
func IdToPath(id, baseDir, ext string) (fullpath string) {

	m := StrUtil.MD5(id)
	fullpath = fmt.Sprintf("%s%c%s%c%s%c%s%c%s%s",
		baseDir, os.PathSeparator,
		m[0:1], os.PathSeparator,
		m[1:2], os.PathSeparator,
		m[2:3], os.PathSeparator,
		m, ext)

	return fullpath
}
