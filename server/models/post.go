package models

import (
	"time"
)

type Post struct {
	Id              int       `json:"postId" gorm:"primary_key;column:post_id"`
	UserId          int       `json:"userId" gorm:"column:user_id"`
	TargetCompanyId int       `json:"targetCompanyId" gorm:"column:target_company_id"`
	TargetPosition  string    `json:"targetPosition" gorm:"column:target_position"`
	Message         string    `json:"message"`
	Resume          string    `json:"resume" gorm:"type:string"`
	JobLink         string    `json:"jobLink" gorm:"column:job_link"`
	CreatedAt       time.Time `json:"createdAt"`
	Company         Company   `json:"company,omitempty" gorm:"foreignKey:target_company_id;references:Id"`
	User         	User   	  `json:"user,omitempty" gorm:"foreignKey:user_id;references:Id"`
}
