package userentity

func (userentity *Entity) TableName() string {
	return "user"
}

func (userentity Entity) ChangeEmail(email string) {
	userentity.Email = email
}
