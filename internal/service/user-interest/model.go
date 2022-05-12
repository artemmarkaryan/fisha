package user_interest

import "time"

type UserInterest struct {
	UserId     int64      `db:"user_id"`
	InterestId int64      `db:"interest_id"`
	Rank       float64    `db:"rank"`
	CreatedAt  *time.Time `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}
