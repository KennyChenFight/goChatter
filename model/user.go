package model

type User struct {
	Id               string `xorm:"pk" update:"fixed"`
	Email            string `binding:"required,email" update:"email"`
	PasswordDigest   string
	Name             string `binding:"required"`
	SelfIntroduction *string
	Picture          *string `binding:"omitempty,base64" update:"base64"`
}

func (u *User) TableName() string {
	return "users"
}
