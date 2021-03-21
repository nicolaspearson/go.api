package db

type User struct {
	Model
	Email     string `gorm:"column:email;index" json:"email"`
	Enabled   bool   `gorm:"column:first_name" json:"enabled"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	Address   string `gorm:"column:address" json:"address"`
	Password  string `gorm:"column:password" json:"password"`
}
