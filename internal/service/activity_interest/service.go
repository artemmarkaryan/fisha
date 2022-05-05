package activity_interest

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha/facade/pkg/database"
)

type Service struct{}

func (Service) ByActivityIds(ctx context.Context, uId []int64) (ui []ActivityInterest, err error) {
	dbp, err := database.Get(ctx)
	if err != nil {
		return
	}

	db, c, err := dbp()
	defer c()

	q, a, err := sq.
		Select("i.*").
		From("activity_interest ai").
		InnerJoin("interest i on ai.interest_id = i.id").
		Where(sq.Eq{"ai.activity_id": uId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	return ui, db.SelectContext(ctx, &ui, q, a...)
}
