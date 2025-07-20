package mysql

import "gorm.io/gorm"

type Users struct {
	Id int `json:"id"`
}

func (m *Users) TableName() string {
	return "users"
}
func (m *Users) DB() *gorm.DB {
	return GetDB()
}

func (m *Users) CreateUser(data map[string]interface{}) error {
	return m.DB().Table("users").Create(&data).Error
}

type UserInfo struct {
	Users
	Username string `json:"username"`
	Password string `json:"password"`
}

func (i *UserInfo) GetUserInfoByName(name string) error {
	return i.DB().Where("username = ?", name).Take(i).Error
}
