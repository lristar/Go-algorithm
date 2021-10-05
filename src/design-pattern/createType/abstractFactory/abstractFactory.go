package abstractFactory

type AbstractFactory interface {
	CreateDirectory() IDirectory
	Get()
}

type IDirectory interface {
	Start()
	Stop()
}





