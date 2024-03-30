package models

import (
	"time"
	"gorm.io/gorm"
)

type Book struct {
	Id          int64     `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Author      string    `gorm:"type:varchar(255)" json:"author"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return
}

func (b *Book) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now()
	return
}
