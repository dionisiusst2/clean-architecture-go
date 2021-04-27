package domain

import (
	"time"

	"github.com/dionisiusst2/bakery-id/utils/validator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name      string         `json:"name" form:"name"`
	Email     string         `json:"email" form:"email"`
	Password  string         `json:"password,omitempty" form:"password"`
	Role      string         `json:"role" form:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (user *User) Validate() []string {
	err := make([]string, 0)

	if e := validator.ValidateNotEmpty("name", user.Name); e != "" {
		err = append(err, e)
	} else if e := validator.ValidateLengthBetween("name", user.Name, 5, 35); e != "" {
		err = append(err, e)
	}

	if e := validator.ValidateNotEmpty("password", user.Password); e != "" {
		err = append(err, e)
	} else if e := validator.ValidateLengthBetween("password", user.Password, 5, 32); e != "" {
		err = append(err, e)
	}

	if e := validator.ValidateNotEmpty("email", user.Email); e != "" {
		err = append(err, e)
	} else if e := validator.ValidateEmailFormat(user.Email); e != "" {
		err = append(err, e)
	}

	return err
}
