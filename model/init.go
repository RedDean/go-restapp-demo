package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var DB *Database

// 初始化数据库连接
func (db *Database) Init() {
	DB = &Database{
		GetSelfDB(),
		GetDockerDB(),
	}
}

// 关闭数据库连接
func (db *Database) Close() {
	db.Self.Close()
	db.Docker.Close()
}

// 建立数据库连接
func openDB(username, password, addr, dbname string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		dbname,
		true,
		"Local",
	)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", dbname)
	}

	setupDB(db)

	return db
}

// 设置日志、连接数等数据库配置
func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitDockerDB() *gorm.DB {
	return openDB(viper.GetString("docker_db.username"),
		viper.GetString("docker_db.password"),
		viper.GetString("docker_db.addr"),
		viper.GetString("docker_db.name"))
}

func GetDockerDB() *gorm.DB {
	return InitDockerDB()
}
