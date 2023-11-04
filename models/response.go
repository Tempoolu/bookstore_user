package models

import "time"

type Book struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at"`
}

type BookResult struct {
	Result    string      `json:"result"`
	Book      []Book    `json:"book"`
}
