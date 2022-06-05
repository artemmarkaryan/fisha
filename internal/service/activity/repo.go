package activity

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha-facade/pkg/database"
)

type repo struct{}

func (repo) getNear(
	ctx context.Context,
	lon, lat float64,
	distanceMeters float64,
	limit uint64,
	exclude []int64,
) (as []Activity, err error) {
	db, c, err := database.Get(ctx)()
	defer c()

	q, a, err := sq.
		Select("*").
		From("activity a").
		Where(sq.NotEq{"id": exclude}).
		Where(sq.Expr("earth_distance(ll_to_earth(a.lon, a.lat), ll_to_earth(?, ?)) < ?", lon, lat, distanceMeters)).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	return as, db.SelectContext(ctx, &as, q, a...)
}

func (repo) get(ctx context.Context, id int64) (a Activity, err error) {
	db, c, err := database.Get(ctx)()
	defer c()

	q, args, err := sq.
		Select("*").
		From("activity").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	return a, db.GetContext(ctx, &a, q, args...)
}
