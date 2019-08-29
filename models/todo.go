package models

import "time"

type Todo struct {
	ID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Text string
	Done bool
	UserID int
	User User

}
