package activity

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha/facade/pkg/database"
)

type repo struct{}

func (repo) getNear(ctx context.Context, lon, lat float64, distanceMeters int, limit uint64) (as []Activity, err error) {
	db, c, err := database.Get(ctx)()
	defer c()

	q, a, err := sq.
		Select("*").
		From("activity a").
		Where(sq.Expr("earth_distance(ll_to_earth(a.lon, a.lat), ll_to_earth(?, ?)) < ?", lon, lat, distanceMeters)).
		Limit(limit).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	return as, db.SelectContext(ctx, &as, q, a...)
}
