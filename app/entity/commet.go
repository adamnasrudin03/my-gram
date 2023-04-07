package entity

type Comment struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	UserID  uint64 `json:"user_id"`
	User    User   `json:"user,omitempty"`
	PhotoID uint64 `json:"photo_id"`
	Photo   Photo  `json:"photo,omitempty"`
	Message string `gorm:"not null" json:"message"`
	GORMModel
}
