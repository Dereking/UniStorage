package UniStorage

type IReader interface {
	ReadObject(name, ext string) (data []byte, bExist bool)
	ExistObject(name, ext string) (bExist bool, fullpath string)
}
