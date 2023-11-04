package models

import "time"

type User struct {
	ID       int       `gorm:"primary_key" json:"id"`
	Name     string    `json:"title"`
    Books    []Book    `gorm:"-" json:"books"`
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at"`
}

func (u *User) Get(id string) {
	db.First(u, id)
}
