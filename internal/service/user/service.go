package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/artemmarkaryan/fisha-facade/pkg/database"
)

type Service struct{}

func (Service) Get(ctx context.Context, id int64) (us User, err error) {
	db, c, err := database.Get(ctx)()
	defer c()

	q, a, err := sq.
		Select("*").
		From(`"user"`).
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return
	}

	err = db.GetContext(ctx, &us, q, a...)

	return
}

func (Service) Login(ctx context.Context, user int64) (isNew bool, err error) {
	db, c, err := database.Get(ctx)()
	defer c()

	q, a, err := sq.
		Select("*").
		From(`"user"`).
		Where(sq.Eq{"id": user}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	var users []User
	if err = db.SelectContext(ctx, &users, q, a...); err != nil {
		return
	}

	if len(users) > 0 {
		return false, nil
	}

	q, a, err = sq.
		Insert(`"user"`).
		Columns("id").
		Values(user).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	_, err = db.ExecContext(ctx, q, a...)

	return true, err
}

func (Service) Forget(ctx context.Context, user int64) error {
	db, c, err := database.Get(ctx)()
	defer c()

	q, a, err := sq.
		Delete(`"user"`).
		Where(sq.Eq{"id": user}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	_, err = db.ExecContext(ctx, q, a...)

	return err
}
