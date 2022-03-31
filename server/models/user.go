package models

type User struct {
	ID                  int    `json:"userId" gorm:"primary_key;not null"`
	FirstName           string `json:"firstName" gorm:"not null;default:null" validate:"required"`
	LastName            string `json:"lastName" gorm:"not null;default:null" validate:"required"`
	Location    		string `json:"currentLocation" gorm:"not null;default:null" validate:"required"`
	CompanyID    		int   `json:"currentCompanyId" gorm:"default:null"`
	CompanyName  		string `json:"currentCompanyName" gorm:"not null;default:null" validate:"required"`
	Position     		string `json:"currentPosition" gorm:"not null;default:null" validate:"required"`
	School              string `json:"school" gorm:"not null;default:null" validate:"required"`
	YearsOfExperienceID int   `json:"yearsOfExperienceId" gorm:"not null;default:null" validate:"required"`
	Email               string `json:"email" gorm:"uniqueIndex:uq_users_email;not null;default:null" validate:"required,email"`
	Password            string `json:"password" gorm:"not null;default:null" validate:"required"`
	Verified            *bool   `json:"verified" gorm:"not null;default:false"`
	Posts               []Post `json:"posts,omitempty"`
	YearsOfExperience   YearsOfExperience `json:"yearsOfExperience,omitempty"`
}

type UserProfile struct {
	FirstName           string `json:"firstName" validate:"required"`
	LastName            string `json:"lastName" validate:"required"`
	Location     		string `json:"currentLocation" validate:"required"`
	CompanyName  		string `json:"currentCompanyName" validate:"required"`
	Position     		string `json:"currentPosition" validate:"required"`
	School              string `json:"school" validate:"required"`
	YearsOfExperienceId int    `json:"yearsOfExperienceId" validate:"required"`	
}

type UserEmail struct {
	Email               string `json:"email" validate:"required,email"`
}

type UserPassword struct {
	CurrentPassword     string `json:"currentPassword" validate:"required"`
	NewPassword         string `json:"newPassword" validate:"required"`
}