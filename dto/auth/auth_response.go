package authdto

type LoginResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	UserID   int    `json:"user_id"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Status   string `json:"status"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	Id       int    `gorm:"type: int" json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullname"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Status   string `gorm:"type: varchar(255)" json:"status" `
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
