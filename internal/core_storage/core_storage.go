package core_storage

import (
	"github.com/balerter/balerter/internal/alert/alert"
)

type CoreStorageKV interface {
	Put(string, string) error
	Get(string) (string, error)
	Upsert(string, string) error
	Delete(string) error
	All() (map[string]string, error)
}

type CoreStorageAlert interface {
	GetOrNew(string) (*alert.Alert, error)
	All() ([]*alert.Alert, error)
	Release(a *alert.Alert) error
	Get(string) (*alert.Alert, error)
}

type CoreStorage interface {
	Name() string
	KV() CoreStorageKV
	Alert() CoreStorageAlert
	Stop() error
}
