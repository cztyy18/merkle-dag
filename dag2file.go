package merkledag

import (
	"encoding/json"
	"fmt"
	"strings"
)



func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	
	var treeObj Object
	treeObj = *getObjectByHash(store, hash)
	obj := getNode(store, path, hp, treeObj)
	var data []byte
	if obj.Links != nil {
		data = obj.Data
	} else {
		getDfsData(store, obj, data)
	}
	return data
}
func getDfsData(store KVStore, object Object, data []byte) []byte {
	obj := &Object{}
	for i := 0; i < len(object.Links); i++ {
		obj = getObjectByHash(store, object.Links[i].Hash)
		if obj.Links != nil {
			getDfsData(store, *obj, data)
		} else {
			data = append(data, obj.Data...)
		}
	}
	return data
}
func getNode(store KVStore, path string, hp HashPool, object Object) Object {
	for _, part := range splitPath(path) {
		var dirHash []byte
		for i := 0; i < len(object.Links); i++ {
			if object.Links[i].Name == part {
				dirHash = object.Links[i].Hash
				break
			}
		}

		jsonObj, _ := store.Get(dirHash)
		var obj Object
		err := json.Unmarshal(jsonObj, &obj)
		if err != nil {
			fmt.Println("解析字符串错误")
		} else {
			object = obj
		}
	}

	return object
}

func splitPath(path string) []string {

	return strings.Split(path, "/")
}

func getObjectByHash(store KVStore, hash []byte) *Object {
	obj := &Object{}
	jsonTreeObj, _ := store.Get(hash)
	err := json.Unmarshal(jsonTreeObj, &obj)
	if err != nil {
		fmt.Println("解析字符串错误")
	}
	return objest the Hash2File function
}
	return nil

