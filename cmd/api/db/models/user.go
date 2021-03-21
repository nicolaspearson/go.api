package db

type User struct {
	Model
	Email     string `gorm:"column:email;index" json:"email"`
	Enabled   bool   `gorm:"column:enabled" json:"enabled"`
	FirstName string `gorm:"column:firstName" json:"firstName"`
	LastName  string `gorm:"column:lastName" json:"lastName"`
	Password  string `gorm:"column:password" json:"password"`
}
