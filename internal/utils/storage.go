package utils

type StorageProxy interface {
	// 简单的key-value
	Get(key string) interface{}
	Rm(key string)
	Set(key string, value interface{})
}

type storage struct {
	center map[string]interface{}
}

func (s *storage) Get(key string) interface{} {
	return s.center[key]
}

func (s *storage) Rm(key string) {
	delete(s.center, key)
}

func (s *storage) Set(key string, value interface{}) {
	s.center[key] = value
}

var (
	s storage
)

func Storage() StorageProxy {
	return &s
}
