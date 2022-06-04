package activity

type Activity struct {
	Id        int64   `db:"id"`
	Name      string  `db:"name"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
	Address   string  `db:"address"`
	Lon       float64 `db:"lon"`
	Lat       float64 `db:"lat"`
	Meta      string  `db:"meta"`
}
