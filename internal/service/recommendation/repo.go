package recommendation

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha/facade/pkg/database"
)

type repo struct{}

func (repo) Upsert(ctx context.Context, recs []R12n) error {
	dbp, err := database.Get(ctx)
	if err != nil {
		return err
	}

	db, c, err := dbp()
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
