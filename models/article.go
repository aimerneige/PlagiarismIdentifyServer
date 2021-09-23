package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserID  uint
	Title   string
	Content string
	Likes   uint
}

type ArticleDto struct {
	ID       uint   `json:"id"`
	CreateAt string `json:"create"`
	UpdateAt string `json:"update"`
	UserID   uint   `json:"userid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Likes    uint   `json:"likes"`
}

func (a Article) ToDto() (dto ArticleDto) {
	timeTemplate := "2006-01-02 15:04"

	dto.ID = a.ID
	dto.CreateAt = a.CreatedAt.Local().Format(timeTemplate)
	dto.UpdateAt = a.UpdatedAt.Local().Format(timeTemplate)
	dto.UserID = a.UserID
	dto.Title = a.Title
	dto.Content = a.Content
	dto.Likes = a.Likes

	return
}
