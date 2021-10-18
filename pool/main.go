package main

import (
	"encoding/json"
	"sync"
)

type Person struct {
	Name   string
	Remark string
}

// 初始化一个对象池
var personPool = sync.Pool{
	New: func() interface{} {
		return new(Person)
	},
}

func main() {
	person := personPool.Get().(*Person) // 从pool中获取对象
	data := []byte(`{Name: "A", Remark: "remark.."}`)
	json.Unmarshal(data, person)
	personPool.Put(person) // 将对象放回对象池
}
