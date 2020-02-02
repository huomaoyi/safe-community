/**
 * @Description: 
 * @Version: 1.0.0
 * @Author: liteng
 * @Date: 2020-02-02 14:26
 */

package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"safe-community/common"
	"safe-community/core/dao/models"
	"sync"
)

var gdb *gorm.DB
var storeInterface models.IStore
var storeOnce sync.Once

//Store ...
type Store struct {
	db *gorm.DB
}

//NewStore ...
func NewStore(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

//SingleStore ...
func SingleStore() models.IStore {
	storeOnce.Do(func() {
		args := common.GetConfig().GetValue("database", "mysql_url")
		log.Println("args: ", args)
		var err error
		gdb, err = gorm.Open("mysql", args)
		if err != nil {
			panic(err)
		}
		gdb.SingularTable(true)
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return defaultTableName
		}

		gdb.DB().SetMaxOpenConns(50)
		gdb.DB().SetMaxIdleConns(50)

		storeInterface = NewStore(gdb)
		gdb.LogMode(true)
	})
	return storeInterface
}

//BeginTx ...
func (s *Store) BeginTx() (models.IStore, error) {
	db := s.db.Begin()
	if db.Error != nil {
		return nil, db.Error
	}
	return NewStore(db), nil
}

//Rollback ...
func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

//CommitTx ...
func (s *Store) CommitTx() error {
	return s.db.Commit().Error
}