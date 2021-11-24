package cache

import "log"

type dataGetter interface {
	GetData(key string) string
}

type CachedDataGetter struct {
	dataGetter
	cache map[string]string
}

func NewCachedDataGetter(dg dataGetter) *CachedDataGetter {
	return &CachedDataGetter{dg, make(map[string]string)}
}

func (cd *CachedDataGetter) GetData(key string) string {
	if v, ok := cd.cache[key]; ok {
		log.Printf("Key [%s] found in cache", key)
		return v
	}

	log.Printf("Key [%s] not found in cache", key)
	v := cd.dataGetter.GetData(key)
	cd.cache[key] = v
	return v
}
