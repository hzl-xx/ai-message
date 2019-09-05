package model

import (
	"messageserver/utils/configure"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

var db *gorm.DB

/*
初始化数据库
*/

func Init() {

	var (
		err error
		dbType, dbName, user, password, host, tablePrefix, charset string
	)

	// yaml配置数据
	sec, err := configure.Cfg.GetSection("database")

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	charset = sec.Key("CHARSET").String()

	// 连接数据库
	conf := user+":"+password+"@tcp("+host+")/"+dbName+"?charset="+charset+"&parseTime=True&loc=Local"
	db,err = gorm.Open(dbType, conf)
	if err != nil {
		log.Fatal(err)
	}

	// 禁用复数表名
	db.SingularTable(true)
	// 设置表前缀
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName;
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
