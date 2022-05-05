package user_interest

import "time"

type UserInterest struct {
	UserId     int64      `json:"user_id"`
	InterestId int64      `json:"interest_id"`
	Rank       float32    `json:"rank"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
