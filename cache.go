package models

type Cache interface {
	Sync(objs interface{}) error
}

type CacheManager interface {
	SyncCache() error
}
