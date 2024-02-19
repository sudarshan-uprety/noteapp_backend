package structure

import "time"

type RegisterInputStruct struct {
	Email            string `json:"email" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Confirm_Password string `json:"confirm_password" binding:"required"`
	Phone            string `json:"phone" binding:"required"`
	Full_Name        string `json:"full_name" binding:"required"`
}

type RegisterOutputStruct struct {
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Full_Name  string    `json:"full_name"`
	Created_at time.Time `json:"created_at"`
}

type LoginInputStruct struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshInputStruct struct {
	Refresh_Token string `json:"refresh_token" binding:"required"`
}
