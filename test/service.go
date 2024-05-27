package test

import (
	"fmt"
	"go-fiber/pkg/db"
	"go-fiber/pkg/hash"
	"go-fiber/pkg/secret"
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func connect(t *testing.T) *gorm.DB {
	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta`,
		secret.GetEnv("DB_HOST", "localhost"),
		secret.GetEnv("DB_USER", "postgres"),
		secret.GetEnv("DB_PASS", "\"\""),
		secret.GetEnv("DB_NAME", "go-fiber"),
		secret.GetEnv("DB_PORT", "5432"),
		secret.GetEnv("DB_SSLMODE", "disable"),
	)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: secret.GetEnv("DB_SCHEMA", "public") + ".",
		},
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().Local() // timestamps
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	return conn
}

func CreateUser(t *testing.T) db.User {
	pass, _ := hash.Create("Password")
	user := db.User{
		Username: "test",
		Name:     "test",
		Password: pass,
	}
	if err := connect(t).Table("users").Create(&user).Error; err != nil {
		t.Fatal(err.Error())
	}
	return user
}

func DeleteUser(t *testing.T, username string) {
	if err := connect(t).Table("users").Where("username = ?", username).Delete(&db.User{}).Error; err != nil {
		t.Fatal(err.Error())
	}
}
