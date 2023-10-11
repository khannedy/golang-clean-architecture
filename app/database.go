package app

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewDatabase(viper *viper.Viper) (*gorm.DB, error) {
	username := viper.Get("database.username").(string)
	password := viper.Get("database.password").(string)
	host := viper.Get("database.host").(string)
	port := int(viper.Get("database.port").(float64))
	database := viper.Get("database.name").(string)
	idleConnection := viper.Get("database.pool.idle").(float64)
	maxConnection := viper.Get("database.pool.max").(float64)
	maxLifeTimeConnection := viper.Get("database.pool.lifetime").(float64)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	connection, err := db.DB()
	if err != nil {
		return nil, err
	}

	connection.SetMaxIdleConns(int(idleConnection))
	connection.SetMaxOpenConns(int(maxConnection))
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db, nil
}
