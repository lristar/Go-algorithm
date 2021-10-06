package main

import (
	"Go-algorithm/src/design-pattern/createType/protoType"
	"Go-algorithm/src/design-pattern/createType/protoType/proto"
	"fmt"
)

func main() {
	maps:=protoType.CacheList{CacheMap: make(map[string]interface{})}
	r1:=proto.Rectangle{}
	r1.SetId(1)
	r1.Settypes("rectangle1")
	maps.SetCache(r1.GetType(),r1.Clone())
	r2:=proto.Rectangle{}
	r2.SetId(2)
	r2.Settypes("rectangle2")
	maps.SetCache(r2.GetType(),r2.Clone())
	c,err:=maps.GetCache("rectangle2")
	if err!=nil{
		return
	}
	r:=c.(*proto.Rectangle)
	fmt.Println(r.GetType())
}
