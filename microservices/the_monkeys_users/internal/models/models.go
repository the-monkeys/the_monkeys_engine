package models

import (
	"database/sql"
)

// TODO: Change as per schema user_account table
type UserAccount struct {
	Id            int64          `json:"id"`
	AccountId     string         `json:"account_id"`
	UserName      string         `json:"username"`
	FirstName     string         `json:"firstname"`
	LastName      string         `json:"lastname"`
	Email         string         `json:"email"`
	DateOfBirth   sql.NullTime   `json:"date_of_birth"`
	Bio           sql.NullString `json:"bio"`
	Address       sql.NullString `json:"address"`
	ContactNumber sql.NullInt64  `json:"contact_number"`
	UserStatus    string         `json:"user_status"`
}

type UserAuthInfo struct {
	Id                       int64          `json:"id"`
	UserId                   int64          `json:"user_id"`
	PasswordHash             sql.NullString `json:"password_hash"`
	CreatedAt                sql.NullTime   `json:"created_at"`
	PasswordRecoveryToken    sql.NullString `json:"password_recovery_token"`
	PasswordRecoveryTimeout  sql.NullTime   `json:"password_recovery_timeout"`
	PasswordUpdatedAt        sql.NullTime   `json:"password_updated_at"`
	EmailVerificationToken   string         `json:"email_verification_token"`
	EmailVerificationTimeout sql.NullTime   `json:"email_verification_timeout"`
	EmailValidationStatus    string         `json:"email_validation_satus"`
	EmailValidationTime      sql.NullTime   `json:"email_validation_time"`
	AuthProviderId           int64          `json:"auth_provider_id"`
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
