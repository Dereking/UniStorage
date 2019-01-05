package UniStorage

type IStorage interface {
	SaveObject(name, ext string, obj interface{}) (filefullpath string, err error)
	SaveURL(url, ext string) (filefullpath string, err error)
}
