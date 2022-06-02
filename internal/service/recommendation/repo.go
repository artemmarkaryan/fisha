package recommendation

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha-facade/pkg/database"
)

type repo struct{}

func (repo) Upsert(ctx context.Context, recs []R12n) error {
	db, c, err := database.Get(ctx)()
	if err != nil {
		return err
	}

	defer c()

	for _, rec := range recs {
		b := sq.
			Insert("recommendations").
			Values(rec.UserId, rec.ActivityId, rec.Rank, rec.Shown).
			Suffix("on conflict (user_id, activity_id) do update set rank = ?, shown = false", rec.Rank).
			PlaceholderFormat(sq.Dollar)
		q, a, err := b.ToSql()
		if err != nil {
			return err
		}

		if _, err = db.ExecContext(ctx, q, a...); err != nil {
			return err
		}
	}

	return nil
}

func (repo) Get(ctx context.Context, user int64) (activity int64, err error) {
	db, c, err := database.Get(ctx)()
	if err != nil {
		return
	}

	defer c()

	db.GetContext(ctx, &activity, sq.
		Select("activity_id").
		From("recommendations").
		Where(sq.Eq{"user_id": user}).
		PlaceholderFormat(sq.Dollar)
	)
}
