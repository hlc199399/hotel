package svc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hotel/internal/config"
	"hotel/internal/dao"
	"hotel/internal/model"
)

type ServiceContext struct {
	Config config.Config
	Db     *dao.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, _ := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	db.AutoMigrate(&model.User{}, &model.Room{}, &model.Order{}, &model.RoomCondition{})
	return &ServiceContext{
		Config: c,
		Db:     dao.Use(db),
	}
}
