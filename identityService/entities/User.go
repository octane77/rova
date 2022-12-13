package entities

type User struct {
	ID       uint   `gorm:"primary_key; autoIncrement" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Surname  string `gorm:"not null" json:"surname"`
	Password string `gorm:"->;<-; not null" json:"-"`
}
