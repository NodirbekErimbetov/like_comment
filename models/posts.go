package models

type PostPrimaryKey struct {
	Id string `json:"id"`
}

type CreatePost struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserId    string `json:"user_id"`
}

type Post struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserId    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdatePost struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GetListPostRequest struct {
	Page   int64  `json:"page"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListPostResponse struct {
	Count int64   `json:"count"`
	Posts  []*Post `json:"posts"`
}
