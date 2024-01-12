package dao

import (
	"gorm.io/gorm"
	"sync"
)

// dao层通过db指针访问数据库接口
var (
	db     *gorm.DB
	dbOnce sync.Once
)

func dataBaseInitialization() {
	dbOnce.Do(func() {
		// db = initialization.GetDB()
	})
}
