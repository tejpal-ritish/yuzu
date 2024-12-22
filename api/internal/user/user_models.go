package user

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Created_at string `json:"created_at"`
}
