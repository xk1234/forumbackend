package model

import (
    "cvwoapi/database"
    "gorm.io/gorm"
	// "fmt"
)


type Comment struct {
    gorm.Model
    Content string `gorm:"type:text" json:"content"`
    UserID  uint `gorm:"not null" json:"user_id"`
	PostRefer uint `gorm:"not null" json:"post_refer"`
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

func (comment *Comment) UpdateComment() (*Comment, error) {
    err := database.Database.Model(comment).Updates(comment).Error
    if err != nil {
        return nil, err
    }
    return comment, nil
}

func (comment *Comment) DeletePostComments(pid uint64) (int64, error) {
    var comments []Comment
	database.Database = database.Database.Debug()
    result := database.Database.Where("post_refer = ?", pid).Order("created_at desc").Find(&comments)
    if result.Error != nil {
        return 0, result.Error
    }
    result = database.Database.Delete(&comments)
    if result.Error != nil {
        return 0, result.Error
    }
    return result.RowsAffected, nil
}


func (comment *Comment) DeleteComment() (int64, error) {
	result := database.Database.Delete(comment)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}