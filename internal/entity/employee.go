package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID          uuid.UUID `json:"id"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Dateofbirth time.Time `json:"dateofbirth"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Status      int8      `json:"status"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	e.CreatedAt = time.Now()
	return
}

func (e *Employee) BeforeUpdate(tx *gorm.DB) (err error) {
	e.UpdatedAt = time.Now()
	return
}
