package event_entity

import "time"

// gorm.Model definition
type BaseModel struct {
	ID        int `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}