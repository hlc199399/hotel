package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"hotel/internal/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/dao",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	gormDb, _ := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/hotel?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"))
	g.UseDB(gormDb) // reuse your gorm db
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(model.User{}, model.Room{}, model.Order{}, model.RoomCondition{})
	// Generate the code
	g.Execute()
}
