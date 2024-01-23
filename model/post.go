package model

import (
	"cvwoapi/database"
	"gorm.io/gorm"
	// "fmt"
)

type Post struct {
    gorm.Model
	Title string `gorm:"type:text" json:"title"`
    Content string `gorm:"type:text" json:"content"`
	Topic string `gorm:"type:text" json:"topic"`
    UserID  uint `gorm:"not null" json:"user_id"`
	Comments []Comment `gorm:"foreignKey:PostRefer"`
	User User `json:"user"`
}

func (post *Post) GetPost(pid uint64) (*Post, error) {
	var err error
	err = database.Database.Model(&Post{}).Where("id = ?", pid).Take(&post).Error
	if err != nil {
		return &Post{}, err
	}
	if post.ID != 0 {
		err = database.Database.Model(&User{}).Where("id = ?", post.UserID).Take(&post.User).Error
		if err != nil {
			return &Post{}, err
		}
	}
	
	return post, nil
}

func (post *Post) Save() (*Post, error) {
    err := database.Database.Create(&post).Error
    if err != nil {
        return &Post{}, err
    }
    return post, nil
}


func (p *Post) All(sort string) (*[]Post, error) {
	posts := []Post{}
    var err error
	query := database.Database.Model(&Post{})
    if sort == "oldest" {
        query = query.Order("created_at asc")
    } else if sort == "mostComments" {
        query = query.Select("posts.*, COUNT(comments.id) as comment_count").
                     Joins("left join comments on comments.post_refer = posts.id").
                     Group("posts.id").
                     Order("comment_count desc")
    } else {
        query = query.Order("created_at desc")
    }
    err = query.Limit(100).Find(&posts).Error
    if err != nil {
        return &[]Post{}, err
    }
    for i := range posts {
        err := database.Database.Model(&User{}).Where("id = ?", posts[i].UserID).Take(&posts[i].User).Error
        if err != nil {
            return &[]Post{}, err
        }
    }

    return &posts, nil
}


func (p *Post) UpdatePost() (*Post, error) {
	var err error
	err = database.Database.Model(&Post{}).Where("id = ?", p.ID).Updates(Post{Title: p.Title, Content: p.Content, Topic: p.Topic}).Error
	if err != nil {
		return &Post{}, err
	}

	return p, nil
}

func (p *Post) DeletePost() (int64, error) {
	result := database.Database.Delete(p)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}