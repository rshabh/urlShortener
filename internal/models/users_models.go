package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Uname     string    `json:"uname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
