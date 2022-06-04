package user

import "time"

type User struct {
	Id        int64      `db:"id"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Lon       *float64   `db:"lon"`
	Lat       *float64   `db:"lat"`
}

func (u User) ValidLocation() bool {
	return u.Lon != nil && *u.Lon > 0 && u.Lat != nil && *u.Lat > 0
}
