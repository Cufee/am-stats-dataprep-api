package types

import "time"

type UserSession struct {
	UserID     string    `json:"user_id" firestore:"user_id"`
	Expiration time.Time `json:"expiration" firestore:"expiration"`
}
