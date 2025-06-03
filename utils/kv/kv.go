package kv

type KeyMapInterface interface {
	Set(any, any)
	Get(any) any
}

func NewKeyMap() *KeyMap {
	return &KeyMap{m: make(map[any]any)}
}

type KeyMap struct {
	m map[any]any
}

func (k *KeyMap) Set(key, val any) {
	k.m[key] = val
}

func (k *KeyMap) Get(key any) any {
	return k.m[key]
}
