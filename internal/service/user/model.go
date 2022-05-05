package user

import "time"

type User struct {
	Id              int64      `db:"id"`
	CreatedAt       *time.Time `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
	LastLocationLon float32    `db:"last_location_lon"`
	LastLocationLat float32    `db:"last_location_lat"`
}
