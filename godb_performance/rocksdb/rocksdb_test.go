package rocksdb

import (
	"github.com/tecbot/gorocksdb"
	"testing"
)



func TestNewRocksDB(t *testing.T) {
	db:=NewRocksDB("./data")
	defer db.Close()
}

func TestRockDBWR(t *testing.T){
	db:=NewRocksDB("./data")
	defer db.Close()

	wo:=gorocksdb.NewDefaultWriteOptions()
	ro:=gorocksdb.NewDefaultReadOptions()

	err:=db.Put(wo,[]byte("rockkey"),[]byte("rockval"))
	if err!=nil{
		t.Error("rocksdb put fail:",err)
	}

	val,err:=db.Get(ro,[]byte("rockkey"))
	if err != nil{
		t.Error("rocksdb get fail:",err)
	}
	t.Logf("val:%s\n",val.Data())
}