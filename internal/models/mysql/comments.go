package mysql

import "gorm.io/gorm"

type Comments struct {
	Id int `json:"id"`
}

func (m *Comments) TableName() string {
	return "comments"
}
func (m *Comments) DB() *gorm.DB {
	return GetDB()
}

func (m *Comments) CreateComment(data map[string]interface{}) error {
	return m.DB().Create(data).Error
}

type CommentInfo struct {
	Comments
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
	PostId  int64  `json:"post_id"`
}

type CommentList []CommentInfo

func (l *CommentList) GetCommentList(postId int) (err error) {
	return GetDB().Where("post_id = ?", postId).Find(l).Error
}
