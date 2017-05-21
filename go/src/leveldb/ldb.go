package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func getData(db *leveldb.DB, key string) string {
	data, err := db.Get([]byte(key), nil)
	if err == nil {
		return string(data)
	}
	fmt.Printf("data: nil, %s\n", err)
	return ""
}

func main() {

	db, _ := leveldb.OpenFile("/tmp/users.ldb", nil)
	defer db.Close()

	data := getData(db, "key")
	fmt.Println("data:" + data)
	_ = db.Put([]byte("key"), []byte("value"), nil)
	data = getData(db, "key")
	fmt.Println("data:" + data)
}
