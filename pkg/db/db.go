package db

import (
	"fmt"
	"go-fiber/pkg/secret"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var G *gorm.DB

var Models = []interface{}{
	&User{},
	&Auth{},
	&Role{},
	&Permission{},
	&UserHasRole{},
	&RoleHashPermission{},
	&Contact{},
	&Address{},
}

func Setup() {
	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta`,
		secret.DB_HOST,
		secret.DB_USER,
		secret.DB_PASS,
		secret.DB_NAME,
		secret.DB_PORT,
		secret.DB_SSLMODE,
	)

	var gormLogger logger.Interface
	if secret.APP_ENV != "development" {
		gormLogger = logger.Default.LogMode(logger.Silent)
	} else {
		gormLogger = logger.Default.LogMode(logger.Info)
	}

	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: secret.DB_SCHEMA + ".",
		},
		SkipDefaultTransaction: true,
		Logger:                 gormLogger,
		NowFunc: func() time.Time {
			return time.Now().Local() // timestamps
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	G = connect
}
