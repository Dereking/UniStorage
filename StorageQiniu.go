package UniStorage

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

type StorageQiniu struct {
	Bucket string
	AK     string
	SK     string
}

func NewStorageQiniu(ak string, sk string, bucket string) *StorageQiniu {
	return &StorageQiniu{AK: ak, SK: sk, Bucket: bucket}
}

// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func (a *StorageQiniu) SaveURL(url, ext string) (filefullpath string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	dataLen := int64(len(data))

	path := IdToPath(url, "/", ext)
	filefullpath = path

	putPolicy := storage.PutPolicy{
		Scope: a.Bucket,
	}

	mac := qbox.NewMac(a.AK, a.SK)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": path,
		},
	}

	err = formUploader.Put(context.Background(), &ret, upToken, path,
		bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	return

}

func (a *StorageQiniu) SaveObject(filename, ext string, obj interface{}) (filefullpath string, err error) {

	data, _ := json.Marshal(obj)
	//data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	log.Println(string(data))

	putPolicy := storage.PutPolicy{
		Scope: a.Bucket,
	}

	mac := qbox.NewMac(a.AK, a.SK)
	upToken := putPolicy.UploadToken(mac)
	log.Println(upToken)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": filename + ext,
		},
	}

	err = formUploader.Put(context.Background(), &ret, upToken, filename,
		bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	return
}

func (a *StorageQiniu) ReadObject(name, ext string) (data []byte, bExist bool) {
	return
}

func (a *StorageQiniu) ExistObject(name, ext string) (bExist bool, fullpath string) {
	return
}
