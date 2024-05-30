package dao

import (
	initialization "Concurrency-Backend/init"
	"sync"

	"gorm.io/gorm"
)

// dao层通过db指针访问数据库接口
var (
	db     *gorm.DB
	dbOnce sync.Once
)

func dataBaseInitialization() {
	dbOnce.Do(func() {
		db = initialization.GetDB()
	})
}
