package studentsdto

type CreateStudentRequest struct {
	UserID int    `json:"-"`
	Avatar string `json:"avatar_url" form:"image" gorm:"type: varchar(255)"`
	Name   string `json:"name" gorm:"type: varchar(255)"`
	Gender string `json:"gender" gorm:"type: varchar(255)"`
	Dob    string `json:"dob" gorm:"type: varchar(255)"`

}

type UpdateStudentRequest struct {
	Avatar string `json:"avatar_url" form:"image" gorm:"type: varchar(255)"`
	Name   string `json:"name" gorm:"type: varchar(255)" form:"name"`
	Gender string `json:"gender" gorm:"type: varchar(255)" form:"gender"`
	Dob    string `json:"dob" gorm:"type: varchar(255)" form:"dob"`
	UserID int `json:"user_id" form:"user_id"`
}
