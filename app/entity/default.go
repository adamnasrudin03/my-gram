package entity

import "time"

type GORMModel struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updated_at,omitempty"`
}
