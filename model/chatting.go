package model

type Chatting struct {
	Id         string `xorm:"pk" update:"fixed"`
	SenderId   string `binding:"required"`
	ReceiverId string `binding:"required"`
	Content    string `binding:"required"`
}

func (c *Chatting) TableName() string {
	return "chatting"
}
