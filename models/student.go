package models

type Student struct {
	Student_id       uint64 `json:"student_id" gorm:"unique;not null;PRIMARY_KEY;AUTO_INCREMENT"`
	Student_name     string `json:"student_name" binding:"required" gorm:"not null"`
	Student_age      uint64 `json:"student_age" binding:"required" gorm:"not null"`
	Student_address  string `json:"student_address" binding:"required" gorm:"not null"`
	Student_phone_no string `json:"student_phone_no" binding:"required" gorm:"not null"`
}
