package studentsdto

type StudentResponse struct {
	ID     int    `json:"id"`
	UserID int    `json:"-"`
	Avatar string `json:"avatar_url" form:"image" gorm:"type: varchar(255)"`
	Name   string `json:"name" gorm:"type: varchar(255)"`
	Gender string `json:"gender" gorm:"type: varchar(255)"`
	Dob    string `json:"dob" gorm:"type: varchar(255)"`
}

func (StudentResponse) TableName() string {
	return "students"
}
