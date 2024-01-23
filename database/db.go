package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

var Database *gorm.DB

func Connect() {
    var err error
    dbstr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Africa/Lagos", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
    Database, err = gorm.Open(postgres.Open(dbstr), &gorm.Config{})

    if err != nil {
        panic(err)
    } else {
        fmt.Println("Successfully connected to the database")
    }
}