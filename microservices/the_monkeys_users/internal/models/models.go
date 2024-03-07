package models

import (
	"database/sql"
)

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

type UserProfileRes struct {
	ProfileId     string         `protobuf:"bytes,1,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	Username      string         `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName     string         `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string         `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	DateOfBirth   sql.NullTime   `protobuf:"bytes,5,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	RoleId        int64          `protobuf:"varint,6,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	Bio           sql.NullString `protobuf:"bytes,7,opt,name=bio,proto3" json:"bio,omitempty"`
	AvatarUrl     sql.NullString `protobuf:"bytes,8,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	CreatedAt     sql.NullTime   `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     sql.NullTime   `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Address       sql.NullString `protobuf:"bytes,11,opt,name=address,proto3" json:"address,omitempty"`
	ContactNumber sql.NullInt64  `protobuf:"varint,12,opt,name=contact_number,json=contactNumber,proto3" json:"contact_number,omitempty"`
	UserStatus    string         `protobuf:"bytes,13,opt,name=user_status,json=userStatus,proto3" json:"user_status,omitempty"`
}
