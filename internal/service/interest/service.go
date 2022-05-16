package interest

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha-facade/pkg/database"
)

type Service struct{}

func (Service) List(ctx context.Context) (i []Interest, err error) {
	db, c, err := database.Get(ctx)()
	if err != nil {
		return nil, err
	}

	defer c()

	q, a, err := sq.Select("*").From("interest").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	return i, db.SelectContext(ctx, &i, q, a...)
}
