package interest

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha/facade/pkg/database"
)

type Service struct{}

func (Service) Names(ctx context.Context) (i []string, err error) {
	dbp, err := database.Get(ctx)
	if err != nil {
		return nil, err
	}

	db, c, err := dbp()
	if err != nil {
		return nil, err
	}

	defer c()

	q, a, err := sq.Select("name").From("interest").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	return i, db.SelectContext(ctx, &i, q, a...)
}
