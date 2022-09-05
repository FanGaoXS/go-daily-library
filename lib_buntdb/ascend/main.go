package main

import (
	"fmt"
	"github.com/tidwall/buntdb"
	"log"
)

func main() {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Update(func(tx *buntdb.Tx) error {
		data := map[string]string{
			"a": "Ava",
			"b": "Bella",
			"c": "Carol",
			"d": "Diana",
			"e": "Eileen",
		}
		for k, v := range data {
			_, _, _ = tx.Set(k, v, nil)
		}
		return nil
	})
	if err != nil {
		log.Fatalln("set data err:", err)
	}

	err = db.View(func(tx *buntdb.Tx) error {
		// 升序遍历
		_ = tx.Ascend("", func(key, value string) bool {
			fmt.Printf("key:%s, value:%s\n", key, value)
			return true
		})
		return nil
	})
	if err != nil {
		log.Fatalln("get data err:", err)
	}
}
