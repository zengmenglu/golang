package leveldb

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)
func NewLevelDB(path string) *leveldb.DB{
	db, err := leveldb.OpenFile(path, nil)
	if err != nil{
		panic(fmt.Sprintf("new levelDB fail:%s",err))
	}
	return db
}
