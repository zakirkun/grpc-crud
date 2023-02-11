package model

import "time"

type Movie struct {
	ID          string `gorm:"primarykey"`
	Title       string
	Genre       string
	Description string
	Thumbnail   string
	CreatedAt   time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime:false"`
}
