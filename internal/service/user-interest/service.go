package user_interest

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha-facade/pkg/database"
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

func (Service) Upsert(ctx context.Context, interests []UserInterest) (err error) {
	db, c, err := database.Get(ctx)()
	defer c()

	var q string
	var a []interface{}
	for _, ui := range interests {
		q, a, err = sq.
			Insert("user_interest").
			Columns("user_id", "interest_id", "rank", "created_at", "updated_at").
			Values(ui.UserId, ui.InterestId, ui.Rank, ui.CreatedAt, ui.UpdatedAt).
			Suffix("on conflict (user_id, interest_id) do update set rank = ?, updated_at = now()", ui.Rank).
			PlaceholderFormat(sq.Dollar).
			ToSql()
		if err != nil {
			return err
		}

		if _, err = db.ExecContext(ctx, q, a...); err != nil {
			return err
		}
	}

	return nil
}
