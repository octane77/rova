package dtos

type GetUserDto struct {
	ID       uint        `json:"id"`
	Name     string      `json:"name"`
	Surname  string      `json:"surname"`
	Accounts interface{} `json:"accounts"`
}
