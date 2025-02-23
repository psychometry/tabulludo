package core

type GameLifecycle interface {
	Start()
	Pause()
	Stop()
	Save() ([]byte, error)
	Load(data []byte) error
}
