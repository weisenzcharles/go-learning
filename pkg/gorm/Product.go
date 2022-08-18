package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//db, err := gorm.Open("mysql", "root:123456@(192.168.0.86)/test?charset=utf8mb4&parseTime=True&loc=Local")
	//defer db.Close()

	db, err := gorm.Open(mysql.Open("root:123456@(192.168.0.86)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "P0001", Price: 100})

	// Read
	var product Product
	// 根据整型主键查找
	db.First(&product, 2)
	// 查找 code 字段值为 P0001 的记录
	db.First(&product, "code = ?", "P0001")

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: "P0002"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "P0002"})

	// Delete - 删除 product
	db.Delete(&product, 1)
}
