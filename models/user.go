package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	FullName  string    `json:"fullname" gorm:"type: varchar(255)"`
	Students  Student   `json:"student"`
	StudentID int       `json:"student_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersProfileResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
