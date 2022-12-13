package dtos

type AccountDetailsResponseDto struct {
	ID       uint        `gorm:"primary_key; autoIncrement" json:"id"`
	Name     string      `gorm:"not null" json:"name"`
	Surname  string      `gorm:"not null" json:"surname"`
	Accounts interface{} `json:"accounts"`
}
