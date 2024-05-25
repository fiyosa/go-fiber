package db

import (
	"fmt"
	"go-fiber/pkg/secret"
	"log"
	"os"
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

	var setLogger logger.Interface
	if secret.APP_ENV != "development" {
		setLogger = logger.Default.LogMode(logger.Silent)
	} else {
		setLogger = gormLogger()
	}

	connect, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: secret.DB_SCHEMA + ".",
		},
		SkipDefaultTransaction: true,
		Logger:                 setLogger,
		NowFunc: func() time.Time {
			return time.Now().Local() // timestamps
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	G = connect
}

func gormLogger() logger.Interface {
	_, err := os.Stat("logs")
	if os.IsNotExist(err) {
		err := os.MkdirAll("logs", 0755)
		if err != nil {
			log.Fatalf("error creating directory logs: %v", err)
		}
	}

	file, err := os.OpenFile("logs/gorm.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             0,           // Disable slow threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
}
