package models

type SignUp struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Bio      string `json:"bio"`
	MediaUrl string `json:"media_url"`
	Password string `json:"password"`
}

type SignIn struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type SignInResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	Bio       string `json:"bio"`
	MediaUrl  string `json:"media_url"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Follow struct {
	FollowerId string `json:"follower_id"`
	FollowedId string `json:"followed_id"`
	CreatedAt  string `json:"created_at"`
}

type FollowList struct {
	Count   int64     `json:"count"`
	Follows []*Follow `json:"follows"`
}
