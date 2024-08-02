package userentity

type User struct {
	Id           string `gorm:"primaryKey"`
	Name         string
	Username     string
	Password     string
	Role         string
	Organization string
}

func (User) TableName() string {
	return "user_management.user"
}
