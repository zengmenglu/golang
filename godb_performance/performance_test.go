package main

import (
	"awesomeProject/leveldb"
	"awesomeProject/mysql"
	"awesomeProject/pebbledb"
	"fmt"
	"github.com/cockroachdb/pebble"
	leveldb2 "github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"testing"
)

var (
	levelPath  = "./data/level"
	pebblePath = "./data/pebble"
)

// pebble
func BenchmarkPebbleSet(b *testing.B) {
	pebbleDB := pebbledb.NewPebbleDB(pebblePath)
	defer pebbleDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		pebbleDB.Set(s, s, pebble.Sync)
	}
}

func BenchmarkPebbleBatchSet(b *testing.B) {
	pebbleDB := pebbledb.NewPebbleDB(pebblePath)
	defer pebbleDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		s1 := []byte(fmt.Sprintf("%d%d", i, i+b.N))
		batch := pebbleDB.NewBatch()
		batch.Set(s, s, nil)
		batch.Set(s1, s1, nil)
		pebbleDB.Apply(batch, &pebble.WriteOptions{Sync: true})
	}
}

func BenchmarkPebbleGet(b *testing.B) {
	pebbleDB := pebbledb.NewPebbleDB(pebblePath)
	defer pebbleDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		pebbleDB.Get(s)
	}
}

// level
func BenchmarkLevelSet(b *testing.B) {
	levelDB := leveldb.NewLevelDB(levelPath)
	defer levelDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		levelDB.Put(s, s, &opt.WriteOptions{Sync: true})
	}
}

func BenchmarkLevelBatchSet(b *testing.B) {
	levelDB := leveldb.NewLevelDB(levelPath)
	defer levelDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		s1 := []byte(fmt.Sprintf("%d%d", i, i+b.N))
		batch := new(leveldb2.Batch)
		batch.Put(s, s)
		batch.Put(s1, s1)
		levelDB.Write(batch, &opt.WriteOptions{Sync: true})
	}
}

func BenchmarkLevelGet(b *testing.B) {
	levelDB := leveldb.NewLevelDB(levelPath)
	defer levelDB.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		levelDB.Get(s, nil)
	}
}

// mysql
func BenchmarkMySqlSet(b *testing.B) {
	db := mysql.NewGorm()
	db.Exec("delete from kv_tbl")
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		db.Create(&mysql.KVTable{s, s})
	}
}

func BenchmarkMysqlBatchSet(b *testing.B) {
	db := mysql.NewGorm()
	db.Exec("delete from kv_tbl")
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		s1 := []byte(fmt.Sprintf("%d", i+b.N))
		data := []mysql.KVTable{{K: s, V: s}, {K: s1, V: s1}}
		mysql.BatchSave(db, data)
	}
}

func BenchmarkMysqlUpdate(b *testing.B) {
	db := mysql.NewGorm()
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		s1 := []byte(fmt.Sprintf("%d", i+b.N))
		db.Model(&mysql.KVTable{}).Where("k=?", s).Update("v", s1)
	}
}

func BenchmarkMysqlGet(b *testing.B) {
	db := mysql.NewGorm()
	defer db.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := []byte(fmt.Sprintf("%d", i))
		db.Where("k=?", s).First(&mysql.KVTable{})
	}
}
