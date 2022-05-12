package user_interest

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha/facade/pkg/database"
)

type Service struct{}

func (Service) ByUserIds(ctx context.Context, uId []int64) (ui []UserInterest, err error) {
	db, c, err := database.Get(ctx)()
	defer c()

	q, a, err := sq.
		Select("*").
		From("user_interest").
		Where(sq.Eq{"user_id": uId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	return ui, db.SelectContext(ctx, &ui, q, a...)
}
