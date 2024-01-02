package model

import (
    "cvwoapi/database"
    "gorm.io/gorm"
    "html"
    "strings"
)
type User struct {
    gorm.Model
    Username string `gorm:"size:255;not null;unique" json:"username"`
    Posts    []Post
	Comments []Comment
}

func (user *User) Save() (*User, error) {
    err := database.Database.Create(&user).Error
    if err != nil {
        return &User{}, err
    }
    return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
    user.Username = html.EscapeString(strings.TrimSpace(user.Username))
    return nil
}

func FindUserByUsername(username string) (User, error) {
    var user User
    err := database.Database.Where("username=?", username).Find(&user).Error
    if err != nil {
        return User{}, err
    }
    return user, nil
}

func FindUserById(id uint) (User, error) {
    var user User
    err := database.Database.Preload("Posts").Where("ID=?", id).Find(&user).Error
    if err != nil {
        return User{}, err
    }
    return user, nil
}
