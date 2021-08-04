package main

import (
	"awesomeProject/leveldb"
	"awesomeProject/mysql"
	"awesomeProject/pebbledb"
	"fmt"
	"github.com/cockroachdb/pebble"
	leveldb2 "github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"math/rand"
	"testing"
	"time"
)

var (
	levelPath  = "./data/level"
	pebblePath = "./data/pebble"
)
var (
	levelDBWo *opt.WriteOptions
)

func init() {
	rand.Seed(time.Now().UnixNano())

	levelDBWo = &opt.WriteOptions{Sync: true}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randByte() []byte {
	// random string

	//b := make([]byte, rand.Intn(1024)+1)
	//for i := range b {
	//	b[i] = letterBytes[rand.Intn(len(letterBytes))]
	//}
	//return b

	// random int
	n := rand.Int()
	return []byte(fmt.Sprintf("%d", n))
}

// level
func BenchmarkLevelSet(b *testing.B) {
	levelDB := leveldb.NewLevelDB(levelPath)
	defer levelDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := randByte()
		levelDB.Put(s, s, levelDBWo)
	}
}

func BenchmarkLevelBatchSet(b *testing.B) {
	levelDB := leveldb.NewLevelDB(levelPath)
	defer levelDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := randByte()
		s1 := randByte()
		batch := new(leveldb2.Batch)
		batch.Put(s, s)
		batch.Put(s1, s1)
		levelDB.Write(batch, levelDBWo)
	}
}
func BenchmarkLevelDBDel(b *testing.B) {
	db := leveldb.NewLevelDB(levelPath)
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := randByte()
		db.Delete(s, levelDBWo)
	}
}
func BenchmarkLevelGet(b *testing.B) {
	levelDB := leveldb.NewLevelDB(levelPath)
	defer levelDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := randByte()
		levelDB.Get(s, nil)
	}
}

// pebble
func BenchmarkPebbleSet(b *testing.B) {
	pebbleDB := pebbledb.NewPebbleDB(pebblePath)
	defer pebbleDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := randByte()
		pebbleDB.Set(s, s, pebble.Sync)
	}
}

func BenchmarkPebbleBatchSet(b *testing.B) {
	pebbleDB := pebbledb.NewPebbleDB(pebblePath)
	defer pebbleDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := randByte()
		s1 := randByte()
		batch := pebbleDB.NewBatch()
		batch.Set(s, s, nil)
		batch.Set(s1, s1, nil)
		pebbleDB.Apply(batch, pebble.Sync)
	}
}

func BenchmarkPebbleDel(b *testing.B) {
	db := pebbledb.NewPebbleDB(pebblePath)
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := randByte()
		db.Delete(s, pebble.Sync)
	}
}

func BenchmarkPebbleGet(b *testing.B) {
	pebbleDB := pebbledb.NewPebbleDB(pebblePath)
	defer pebbleDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := randByte()
		pebbleDB.Get(s)
	}
}

// mysql
func BenchmarkMySqlSet(b *testing.B) {
	db := mysql.NewGorm()
	//db.Exec("delete from kv_tbl")
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := randByte()
		db.Create(&mysql.KVTable{s, s})
	}
}

func BenchmarkMysqlBatchSet(b *testing.B) {
	db := mysql.NewGorm()
	//db.Exec("delete from kv_tbl")
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := randByte()
		s1 := randByte()
		data := []mysql.KVTable{{K: s, V: s}, {K: s1, V: s1}}
		mysql.BatchSave(db, data)
	}
}

func BenchmarkMysqlUpdate(b *testing.B) {
	db := mysql.NewGorm()
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := randByte()
		s1 := randByte()
		db.Model(&mysql.KVTable{}).Where("k=?", s).Update("v", s1)
	}
}

func BenchmarkMysqlGet(b *testing.B) {
	db := mysql.NewGorm()
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := randByte()
		db.Where("k=?", s).First(&mysql.KVTable{})
	}
}

func BenchmarkMysqlDel(b *testing.B) {
	db := mysql.NewGorm()
	defer db.Close()
	for i := 0; i < b.N; i++ {
		s := randByte()
		db.Where("k=?", s).Delete(&mysql.KVTable{})
	}
}
