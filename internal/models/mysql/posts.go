package mysql

import (
	"gin-study/api/request"
	"gorm.io/gorm"
)

type Posts struct {
	Id int `json:"id"`
}

func (m *Posts) TableName() string {
	return "posts"
}
func (m *Posts) DB() *gorm.DB {
	return GetDB()
}

func (m *Posts) CreatePost(data map[string]interface{}) error {
	return m.DB().Create(&data).Error
}

func (m *Posts) UpdatePost(id int, data map[string]interface{}) error {
	return m.DB().Where("id = ?", id).Updates(&data).Error
}

type PostsInfo struct {
	Posts
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  int    `json:"user_id"`
}

func (i *PostsInfo) GetInfoById(id int) error {
	return i.DB().Where("id = ?", id).Take(i).Error
}

type PostList []PostsInfo

func (l *PostList) GetPostList(param request.PostList) (count int64, err error) {
	db := GetDB().Model(&Posts{})
	if param.KeyValue != "" {
		db = db.Where("title like ?", "%"+param.KeyValue+"%")
	}
	if err = db.Count(&count).Error; err != nil {
		return
	}
	if err = db.Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(l).Error; err != nil {
		return
	}
	return
}
