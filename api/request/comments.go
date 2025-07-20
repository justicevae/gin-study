package request

type CreateComment struct {
	PostID  int    `json:"post_id"`
	Content string `json:"content"`
}

type GetCommentList struct {
	PostID int `form:"post_id"`
}
