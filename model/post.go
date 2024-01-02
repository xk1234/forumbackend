package model

import (
    "cvwoapi/database"
    "gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Content string `gorm:"type:text" json:"content"`
    UserID  uint `gorm:"not null" json:"user_id"`
	Comments []Comment `gorm:"foreignKey:PostRefer"`
	User User `json:"user"`
}

func (post *Post) Save() (*Post, error) {
    err := database.Database.Create(&post).Error
    if err != nil {
        return &Post{}, err
    }
    return post, nil
}


func (p *Post) All() (*[]Post, error) {
	var err error
	posts := []Post{}
	err = database.Database.Model(&Post{}).Limit(100).Order("created_at desc").Find(&posts).Error
	if err != nil {
		return &[]Post{}, err
	}
	if len(posts) > 0 {
		for i, _ := range posts {
			err := database.Database.Model(&User{}).Where("id = ?", posts[i].UserID).Take(&posts[i].User).Error
			if err != nil {
				return &[]Post{}, err
			}
		}
	}
	return &posts, nil
}