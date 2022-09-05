package main

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
)

func main() {
	// 调用buntdb.Open()方法需要传入一个文件名的参数，指定数据保存的文件路径。
	// 如果传入特殊字符串:memory:，则buntdb不会将数据保存到磁盘。
	filePath := "./db"

	db, err := buntdb.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// 写事务：db.Update()
	// 读事务：db.View()
	// 同一时间只能存在一个写事务，
	// 但是可以同时存在多个并发的读事务。
	err = db.Update(func(tx *buntdb.Tx) error {
		// 如果之前的值被覆盖，则会返回true
		oldValue, replaced, err := tx.Set("testKey", "testValue", nil)
		if err != nil {
			return err
		}

		fmt.Printf("old value[%s] replaced[%t]\n", oldValue, replaced)
		return nil
	})
	if err != nil {
		log.Fatalln("set value err:", err)
	}

	err = db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get("testKey")
		if err != nil {
			return err
		}

		fmt.Printf("value is %s\n", val)
		return nil
	})
	if err != nil {
		log.Fatalln("get value err:", err)
	}
}
