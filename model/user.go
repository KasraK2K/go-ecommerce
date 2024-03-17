package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"id"                 bson:"id"                 gorm:"type:uint;primaryKey;<-:false"`
	Email         string         `json:"email"              bson:"email"              gorm:"type:string;unique;not null;"            validate:"required,email,min=6,max=32"`
	Password      string         `json:"password,omitempty" bson:"password,omitempty" gorm:"type:string;check:length(password) >= 8" validate:"required,min=8,max=32"`
	ContactNumber string         `json:"contact_number"     bson:"contact_number"     gorm:"type:string;"`
	FirstName     string         `json:"first_name"         bson:"first_name"         gorm:"type:string;"`
	Surname       string         `json:"surname"            bson:"surname"            gorm:"type:string;"`
	IsActive      bool           `json:"is_active"          bson:"is_active"          gorm:"default:true"`
	IsAdmin       bool           `json:"is_admin"           bson:"is_admin"           gorm:"default:false"`
	IsArchive     bool           `json:"is_archive"         bson:"is_archive"         gorm:"default:false"`
	Permission    string         `json:"permission"         bson:"permission"         gorm:"type:string"`
	CreatedAt     time.Time      `json:"created_at"         bson:"created_at"         gorm:"type:timestamptz;autoCreateTime;"`
	UpdatedAt     time.Time      `json:"updated_at"         bson:"updated_at"         gorm:"type:timestamptz;autoUpdateTime;"`
	ArchiveAt     gorm.DeletedAt `json:"archive_at"         bson:"archive_at"         gorm:"type:timestamptz;index"`
}

type UserFilter struct {
	ID        uint   `json:"id,omitempty"         bson:"id,omitempty"          validate:"omitempty"`
	Email     string `json:"email,omitempty"      bson:"email,omitempty"       validate:"omitempty,email,min=6,max=32"`
	IsActive  bool   `json:"is_active,omitempty"  bson:"is_active,omitempty"`
	IsAdmin   bool   `json:"is_admin,omitempty"   bson:"is_admin,omitempty"`
	IsArchive bool   `json:"is_archive,omitempty" bson:"is_archive,omitempty"`
}

type UserUpdate struct {
	ID            uint   `json:"id" bson:"id"`
	Email         string `json:"email"          bson:"email"`
	Password      string `json:"password"       bson:"password"`
	ContactNumber string `json:"contact_number" bson:"contact_number"`
	FirstName     string `json:"first_name"     bson:"first_name"`
	Surname       string `json:"surname"        bson:"surname"`
	IsActive      bool   `json:"is_active"      bson:"is_active"`
	IsAdmin       bool   `json:"is_admin"       bson:"is_admin"`
	Permission    string `json:"permission"     bson:"permission"`
}

type UserLoginPayload struct {
	Email    string `json:"email"    bson:"email"    validate:"required,email,min=6,max=32"`
	Password string `json:"password" bson:"password" validate:"required,min=8,max=32"`
	Platform int    `json:"platform" bson:"platform" validate:"required,oneof=1 2"`
}
