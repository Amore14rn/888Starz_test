package dao

import (
	"context"
	"encoding/json"
	"github.com/Amore14rn/888Starz/internal/dal/postgres"
	"github.com/Amore14rn/888Starz/internal/domain/user/model"
	"github.com/Amore14rn/888Starz/pkg/errors"
	psql "github.com/Amore14rn/888Starz/pkg/postgresql"
	"github.com/Amore14rn/888Starz/pkg/tracing"
	sq "github.com/Masterminds/squirrel"
	"strconv"
)

type UserDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewUserStorage(client psql.Client) *UserDAO {
	return &UserDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *UserDAO) All(ctx context.Context) ([]model.User, error) {
	all, err := repo.findBy(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]model.User, len(all))
	for i, e := range all {
		resp[i] = e.ToDomain()
	}

	return resp, nil
}

func (u *UserDAO) Create(ctx context.Context, req model.CreateUser) error {
	sql, args, err := u.qb.
		Insert(postgres.UserTable).
		Columns(
			"id",
			"first_name",
			"last_name",
			"full_name",
			"age",
			"is_married",
			"password",
		).
		Values(
			req.ID,
			req.FirstName,
			req.LastName,
			req.FullName,
			req.Age,
			req.IsMarried,
			req.Password,
		).ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}
	tracing.SpanEvent(ctx, "Insert Product query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := u.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("nothing inserted")
	}

	return nil
}

func (repo *UserDAO) findBy(ctx context.Context) ([]UserStorage, error) {
	statement := repo.qb.
		Select(
			"id",
			"first_name",
			"last_name",
			"age",
			"is_married",
			"password",
			"created_at",
			"updated_at",
		).
		From(postgres.UserTable + " u")

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Select User")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	defer rows.Close()

	entities := make([]UserStorage, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var e UserStorage
		if err = rows.Scan(
			&e.ID,
			&e.FirstName,
			&e.LastName,
			&e.FullName,
			&e.Age,
			&e.IsMarried,
			&e.Password,
			&e.CreatedAt,
			&e.Orders,
		); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			tracing.Error(ctx, err)

			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (u *UserDAO) GetUser(ctx context.Context, id string) (model.User, error) {
	user, err := u.findByID(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user.ToDomain(), nil
}

func (u *UserDAO) findByID(ctx context.Context, id string) (UserStorage, error) {
	statement := u.qb.
		Select(
			"id",
			"first_name",
			"last_name",
			"age",
			"is_married",
			"password",
			"created_at",
			"updated_at",
		).
		From(postgres.UserTable + " u").
		Where(sq.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return UserStorage{}, err
	}

	tracing.SpanEvent(ctx, "Select User")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	row := u.client.QueryRow(ctx, query, args...)

	var e UserStorage
	if err = row.Scan(
		&e.ID,
		&e.FirstName,
		&e.LastName,
		&e.FullName,
		&e.Age,
		&e.IsMarried,
		&e.Password,
		&e.CreatedAt,
		&e.Orders,
	); err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)

		return UserStorage{}, err
	}

	return e, nil
}

func (u *UserDAO) GetUserByName(ctx context.Context, name string) (model.User, error) {
	user, err := u.findByName(ctx, name)
	if err != nil {
		return model.User{}, err
	}

	return user.ToDomain(), nil
}

func (u *UserDAO) findByName(ctx context.Context, name string) (UserStorage, error) {
	statement := u.qb.
		Select(
			"id",
			"first_name",
			"last_name",
			"age",
			"is_married",
			"password",
			"created_at",
			"updated_at",
		).
		From(postgres.UserTable + " u").
		Where(sq.Eq{"first_name": name})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return UserStorage{}, err
	}

	tracing.SpanEvent(ctx, "Select User")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	row := u.client.QueryRow(ctx, query, args...)

	var e UserStorage
	if err = row.Scan(
		&e.ID,
		&e.FirstName,
		&e.LastName,
		&e.FullName,
		&e.Age,
		&e.IsMarried,
		&e.Password,
		&e.CreatedAt,
		&e.Orders,
	); err != nil {
		err = psql.ErrScan(psql.ParsePgError(err))
		tracing.Error(ctx, err)

		return UserStorage{}, err
	}

	return e, nil
}

func (u *UserDAO) Update(ctx context.Context, req model.UpdateUser) error {
	sql, args, err := u.qb.
		Update(postgres.UserTable).
		Set("first_name", req.FirstName).
		Set("last_name", req.LastName).
		Set("full_name", req.FullName).
		Set("age", req.Age).
		Set("is_married", req.IsMarried).
		Set("password", req.Password).
		Set("updated_at", req.UpdatedAt).
		Where(sq.Eq{"id": req.ID}).
		ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}
	tracing.SpanEvent(ctx, "Update User query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := u.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("nothing updated")
	}

	return nil
}

func (u *UserDAO) Delete(ctx context.Context, id string) error {
	sql, args, err := u.qb.
		Delete(postgres.UserTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}
	tracing.SpanEvent(ctx, "Delete User query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := u.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("nothing deleted")
	}

	return nil
}

func (u *UserDAO) CreateOrder(ctx context.Context, req model.CreateOrder) error {
	productsJSON, err := json.Marshal(req.Products)
	if err != nil {
		return err
	}

	sql, args, err := sq.
		Insert(postgres.OrderTable).
		Columns(
			"id",
			"user_id",
			"product_id",
			"products",
			"time_stamp",
		).
		Values(
			req.ID,
			req.UserID,
			req.ProductID,
			string(productsJSON),
			req.TimeStamp,
		).ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}
	tracing.SpanEvent(ctx, "Insert Order query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := u.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return execErr
	}

	if cmd.RowsAffected() == 0 {
		return errors.New("nothing inserted")
	}

	return nil
}

func (u *UserDAO) AreProductsAvailable(ctx context.Context, productID string, products []model.OrderProduct) bool {
	return true
}
