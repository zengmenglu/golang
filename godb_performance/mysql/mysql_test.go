package mysql

import (
	"fmt"
	"testing"
)

func TestNewGorm(t *testing.T) {
	db:=NewGorm()
	defer db.Close()
	db.Exec("delete from kv_tbl")

	db.Create(&KVTable{[]byte("Jane"),[]byte("F")})

	var data KVTable
	db.First(&data,"k = ?", []byte("Jane"))
	fmt.Println(data)
}

func TestBatchSave(t *testing.T) {
	db:=NewGorm()
	defer db.Close()
	db.Exec("delete from kv_tbl")
	BatchSave(db,[]KVTable{{[]byte("A"),[]byte("A")}, {[]byte{'B'},[]byte{'B'}}})

	res:=[]KVTable{	}
	db.Find(&res)
	fmt.Println(res)
}