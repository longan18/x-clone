package model

import "time"

type RefreshToken struct {
	ID         uint      `json:"id"`
	UserId     uint      `json:"user_id"`
	Token      string    `json:"token"`
	ExpiryDate time.Time `json:"expiry_date"`
}
