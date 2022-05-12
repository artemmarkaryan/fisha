package activity_interest

import "time"

type ActivityInterest struct {
	ActivityId int64      `db:"activity_id"`
	InterestId int64      `db:"interest_id"`
	Rank       float64    `db:"rank"`
	CreatedAt  *time.Time `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
}
