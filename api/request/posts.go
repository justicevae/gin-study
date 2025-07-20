package request

type PostList struct {
	KeyValue string `form:"key_value"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

type PostDetail struct {
	Id int `form:"id"`
}

type CreatePost struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePost struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type DeletePost struct {
	Id int `json:"id"`
}
