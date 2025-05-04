package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type CreateUser struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Bio      string `json:"bio"`
	MediaUrl string `json:"media_url"`
	Password string `json:"password"`
}

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	Bio       string `json:"bio"`
	MediaUrl  string `json:"media_url"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUser struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	Bio       string `json:"bio"`
	MediaUrl  string `json:"media_url"`
	Password  string `json:"password"`
}

type GetListUserRequest struct {
	Page   int64  `json:"page"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
}

type GetListUserResponse struct {
	Count int64   `json:"count"`
	Users []*User `json:"users"`
}
