package models

import "database/sql"

type MyProfile struct {
	Id            int64
	FirstName     string
	LastName      string
	Email         string
	CreateTime    string
	IsActive      bool
	CountryCode   string
	Mobile        string
	About         string
	Instagram     string
	Twitter       string
	EmailVerified bool
}

type TheMonkeysUser struct {
	Id                          int64          `json:"id"`
	ProfileId                   string         `json:"profile_id"`
	Username                    string         `json:"username"`
	FirstName                   string         `json:"first_name"`
	LastName                    string         `json:"last_name"`
	Email                       string         `json:"email"`
	Password                    string         `json:"password"`
	PasswordVerificationToken   sql.NullString `json:"password_verification_token"`
	PasswordVerificationTimeout sql.NullTime   `json:"password_verification_timeout"`
	EmailVerificationStatus     string         `json:"email_verified"`
	UserStatus                  string         `json:"is_active,omitempty"`
	EmailVerificationToken      string         `json:"email_verification_token"`
	EmailVerificationTimeout    sql.NullTime   `json:"email_verification_timeout"`
	MobileVerificationToken     string         `json:"mobile_verification_token"`
	MobileVerificationTimeout   sql.NullTime   `json:"mobile_verification_timeout"`
	LoginMethod                 string         `json:"login_method"`
}
