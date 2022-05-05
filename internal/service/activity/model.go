package activity

import "time"

type Activity struct {
	Id        int64      `db:"id"`
	Name      string     `db:"name"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	Address   string     `db:"address"`
	Lon       float32    `db:"lon"`
	Lat       float32    `db:"lat"`
}
