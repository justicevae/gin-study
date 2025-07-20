package request

type Register struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
