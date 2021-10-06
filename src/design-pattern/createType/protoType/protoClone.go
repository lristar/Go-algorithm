package protoType

import (
	"errors"
	"sync"
)

type Cloneable interface {
	Clone() Shape
}

type Shape interface {
	GetType() string
	GetId() int
	SetId(id int)
	Settypes(types string)
	Cloneable
}

type CacheList struct {
	sync.Mutex
	CacheMap map[string]interface{}
}

func (c *CacheList) SetCache(str string, i interface{}) {
	c.Lock()
	c.CacheMap[str] = i
	c.Unlock()
}

func (c *CacheList) GetCache(str string) (interface{}, error) {
	v, ok := c.CacheMap[str]
	if !ok {
		return nil, errors.New("没有这个key")
	}
	return v, nil
}


