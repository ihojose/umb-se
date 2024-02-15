package database

import (
	_mysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func Connect() *gorm.DB {
	cfg := _mysql.Config{
		User:                 "root",
		Passwd:               "mysql",
		Addr:                 "localhost:3306",
		DBName:               "airlanese",
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       cfg.FormatDSN(),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	log.Println("Database connected...")

	return db
}
