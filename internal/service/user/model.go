package user

import "time"

type User struct {
	Id              int64      `db:"id"`
	CreatedAt       *time.Time `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
	LastLocationLon *float64   `db:"last_location_lon"`
	LastLocationLat *float64   `db:"last_location_lat"`
}

func (u User) ValidLocation() bool {
	return u.LastLocationLon != nil && *u.LastLocationLon > 0 && u.LastLocationLat != nil && *u.LastLocationLat > 0
}
