package ui

type Widget interface {
	GetCore() *Core
	Update()
	Draw()
	Close()
}
