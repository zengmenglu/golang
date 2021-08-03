package rocksdb

import (
	"fmt"
	"github.com/tecbot/gorocksdb"
)

func NewRocksDB(path string)*gorocksdb.DB{
	opts:=gorocksdb.NewDefaultOptions()
	db,err:=gorocksdb.OpenDb(opts,path)
	if err != nil{
		panic(fmt.Sprintf("new rocksdb fail:%s",err))
	}
	return db
}

