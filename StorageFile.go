package UniStorage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dereking/utils/StrUtil"
)

type StorageFile struct {
	DataDir string
}

func NewStorageFile(dataDir string) *StorageFile {
	return &StorageFile{DataDir: dataDir}
}

func (a *StorageFile) SaveURL(url, ext string) (filefullpath string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	m := StrUtil.MD5(url)
	filefullpath = fmt.Sprintf("%s%c%s%c%s%c%s%c%s%s",
		a.DataDir, os.PathSeparator,
		m[0:1], os.PathSeparator,
		m[1:2], os.PathSeparator,
		m[2:3], os.PathSeparator,
		m, ext)

	os.MkdirAll(filepath.Dir(filefullpath), 0755)

	err = ioutil.WriteFile(filefullpath, data, 0755)
	os.Chmod(filefullpath, 0755)
	return

}

func (a *StorageFile) SaveObject(filename, ext string, obj interface{}) (filefullpath string, err error) {

	data, _ := json.Marshal(obj)

	m := StrUtil.MD5(filename)
	filefullpath = fmt.Sprintf("%s%c%s%c%s%c%s%c%s%s",
		a.DataDir, os.PathSeparator,
		m[0:1], os.PathSeparator,
		m[1:2], os.PathSeparator,
		m[2:3], os.PathSeparator,
		m, ext)
	os.MkdirAll(filepath.Dir(filefullpath), 0755)

	err = ioutil.WriteFile(filefullpath, data, 0755)
	os.Chmod(filefullpath, 0755)

	return
}
func (a *StorageFile) ReadObject(name, ext string) (data []byte, bExist bool, err error) {

	bExist, filefullpath := a.ExistObject(name, ext)
	if bExist {
		data, err = ioutil.ReadFile(filefullpath)
	}
	return
}

func (a *StorageFile) ExistObject(name, ext string) (bExist bool, fullpath string) {

	fullpath = IdToPath(name, a.DataDir, ext)

	return PathExist(fullpath), fullpath
}
