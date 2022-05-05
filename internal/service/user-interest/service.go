package user_interest

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha/facade/pkg/database"
)

type Service struct{}

func (Service) ByUserIds(ctx context.Context, uId []int64) (ui []UserInterest, err error) {
	dbp, err := database.Get(ctx)
	if err != nil {
		return
	}

	db, c, err := dbp()
	defer c()

	q, a, err := sq.
		Select("i.*").
		From("user_interest ui").
		InnerJoin("interest i on ui.interest_id = i.id").
		Where(sq.Eq{"ui.user_id": uId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	return ui, db.SelectContext(ctx, &ui, q, a...)
}
