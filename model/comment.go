package model

import (
    "cvwoapi/database"
    "gorm.io/gorm"
)


type Comment struct {
    gorm.Model
    Content string `gorm:"type:text" json:"content"`
    UserID  uint `gorm:"not null" json:"user_id"`
	PostRefer uint `gorm:"not null" json:"post_id"`
	User User `json:"user"`
}

func (comment *Comment) Save() (*Comment, error) {
    err := database.Database.Create(&comment).Error
    if err != nil {
        return &Comment{}, err
    }
	if comment.ID != 0 {
		err = database.Database.Model(&User{}).Where("id = ?", comment.UserID).Take(&comment.User).Error
		if err != nil {
			return &Comment{}, err
		}
	}
    return comment, nil
}

func (c *Comment) GetComments(post_id uint) (*[]Comment, error) {
    comments := []Comment{}
	err := database.Database.Where("post_refer = ?", post_id).Order("created_at desc").Find(&comments).Error
    if err != nil {
        return &[]Comment{}, err
    }
	if len(comments) > 0 {
		for i, _ := range comments {
			err := database.Database.Model(&User{}).Where("id = ?", comments[i].UserID).Take(&comments[i].User).Error
			if err != nil {
				return &[]Comment{}, err
			}
		}
	}
    return &comments, nil
}
