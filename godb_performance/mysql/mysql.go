package mysql

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type KVTable struct {
	K []byte `gorm:"k"`
	V []byte `gorm:"v"`
}

func (KVTable) TableName() string {
	return "kv_tbl"
}

func NewGorm() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@/gorm_example?charset=utf8&parseTime=True&loc=Local") // "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	if err != nil {
		panic(fmt.Sprintf("fail to connect db, err:%s", err))
	}
	return db
}

func BatchSave(db *gorm.DB,datas []KVTable)error{
	var buffer bytes.Buffer
	sql := "insert into `kv_tbl` (`k`,`v`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, e := range datas {
		if i == len(datas)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s');", e.K, e.V))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s'),", e.K, e.V))
		}
	}
	return db.Exec(buffer.String()).Error
}