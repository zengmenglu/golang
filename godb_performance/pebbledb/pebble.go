package pebbledb

import (
	"fmt"
	"github.com/cockroachdb/pebble"
)


func NewPebbleDB(path string) *pebble.DB{
	pebbleDB, err := pebble.Open(path, &pebble.Options{})
	if err != nil {
		panic(fmt.Sprintf("pebbleDB init fail:%s",err))
	}
	return pebbleDB
}


