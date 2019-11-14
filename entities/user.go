package entities

type UserId struct {
	Id  int64 `json:"id" form:"id"`
}

type UserInfo struct {
	Id  int64 `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	IsDelete  string `json:"is_delete" form:"is_delete"`
}