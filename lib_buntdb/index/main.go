package main

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
	"strconv"
)

func main() {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// 索引：索引名，索引的匹配模式，索引的排序方式
	err = db.CreateIndex("f-c_id", "*-*", buntdb.IndexInt)
	if err != nil {
		log.Fatalln("create index err:", err)
	}

	err = db.Update(func(tx *buntdb.Tx) error {
		for fId := 1; fId <= 5; fId++ {
			for cId := 1; cId <= 10; cId++ {
				key := strconv.Itoa(fId) + "-" + strconv.Itoa(cId)
				value := "content of " + key
				tx.Set(key, value, nil)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalln("set data err:", err)
	}

	err = db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("f-c_id", func(key, value string) bool {
			fmt.Printf("key:%s, value:%s\n", key, value)
			return true
		})
		return nil
	})
	if err != nil {
		log.Fatalln("get data err:", err)
	}

}
