package dao

import (
	"context"
	"github.com/Amore14rn/888Starz/internal/dal/postgres"
	"github.com/Amore14rn/888Starz/internal/domain/products/model"
	"github.com/Amore14rn/888Starz/pkg/errors"
	psql "github.com/Amore14rn/888Starz/pkg/postgresql"
	"github.com/Amore14rn/888Starz/pkg/tracing"
	sq "github.com/Masterminds/squirrel"
	"strconv"
)

type ProductDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewProductDAO(client psql.Client) *ProductDAO {
	return &ProductDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *ProductDAO) All(ctx context.Context) ([]model.Products, error) {
	all, err := repo.findBy(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]model.Products, len(all))
	for i, e := range all {
		resp[i] = e.ToDomain()
	}

	return resp, nil
}

func (repo *ProductDAO) Create(ctx context.Context, req model.CreateProducts) (model.Products, error) {
	sql, args, err := repo.qb.
		Insert(postgres.ProductTable).
		Columns(
			"id",
			"description",
			"quantity",
			"created_at",
		).
		Values(
			req.ID,
			req.Description,
			req.Quantity,
			req.CreatedAt,
		).ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return model.Products{}, err
	}

	tracing.SpanEvent(ctx, "Insert Product query")
	tracing.TraceVal(ctx, "sql", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := repo.client.Exec(ctx, sql, args...)
	if execErr != nil {
		execErr = psql.ErrDoQuery(execErr)
		tracing.Error(ctx, execErr)

		return model.Products{}, execErr
	}

	if cmd.RowsAffected() == 0 {
		return model.Products{}, errors.New("nothing inserted")
	}

	return model.NewProduct(
		req.ID,
		req.Description,
		req.Quantity,
		req.Tags,
		req.CreatedAt,
		nil), nil
}

func (repo *ProductDAO) GetByID(ctx context.Context, id string) (model.Products, error) {
	statement := repo.qb.
		Select(
			"id",
			"description",
			"quantity",
			"created_at",
		).
		From(postgres.ProductTable).
		Where(sq.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return model.Products{}, err
	}

	tracing.SpanEvent(ctx, "Select Product by ID")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		tracing.Error(ctx, err)

		return model.Products{}, err
	}

	defer rows.Close()

	var e ProductStorage
	if rows.Next() {
		if err = rows.Scan(
			&e.ID,
			&e.Description,
			&e.Quantity,
			&e.CreatedAt,
		); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			tracing.Error(ctx, err)

			return model.Products{}, err
		}
	}

	return e.ToDomain(), nil
}

func (repo *ProductDAO) findBy(ctx context.Context) ([]ProductStorage, error) {
	statement := repo.qb.
		Select(
			"id",
			"description",
			"quantity",
			"created_at",
		).
		From(postgres.ProductTable)

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return nil, err
	}

	tracing.SpanEvent(ctx, "Select Product")
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

	entities := make([]ProductStorage, rows.CommandTag().RowsAffected())

	for rows.Next() {
		var e ProductStorage
		if err = rows.Scan(
			&e.ID,
			&e.Description,
			&e.Quantity,
			&e.CreatedAt,
		); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			tracing.Error(ctx, err)

			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (repo *ProductDAO) GetProduct(ctx context.Context, id string) (model.Products, error) {
	statement := repo.qb.
		Select(
			"id",
			"description",
			"quantity",
			"created_at",
		).
		From(postgres.ProductTable).
		Where(sq.Eq{"id": id})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return model.Products{}, err
	}

	tracing.SpanEvent(ctx, "Select Product by ID")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	rows, err := repo.client.Query(ctx, query, args...)
	if err != nil {
		err = psql.ErrDoQuery(err)
		tracing.Error(ctx, err)

		return model.Products{}, err
	}

	defer rows.Close()

	var e ProductStorage
	if rows.Next() {
		if err = rows.Scan(
			&e.ID,
			&e.Description,
			&e.Quantity,
			&e.CreatedAt,
		); err != nil {
			err = psql.ErrScan(psql.ParsePgError(err))
			tracing.Error(ctx, err)

			return model.Products{}, err
		}
	}

	return e.ToDomain(), nil
}

func (repo *ProductDAO) Update(ctx context.Context, req model.UpdateProducts) error {
	statement := repo.qb.
		Update(postgres.ProductTable).
		Set("description", req.Description).
		Set("quantity", req.Quantity).
		Set("updated_at", req.UpdatedAt).
		Where(sq.Eq{"id": req.ID})

	query, args, err := statement.ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}

	tracing.SpanEvent(ctx, "Update Product")
	tracing.TraceVal(ctx, "SQL", query)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := repo.client.Exec(ctx, query, args...)
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

func (repo *ProductDAO) Delete(ctx context.Context, id string) error {
	sql, args, err := repo.qb.
		Delete(postgres.ProductTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		err = psql.ErrCreateQuery(err)
		tracing.Error(ctx, err)

		return err
	}

	tracing.SpanEvent(ctx, "Delete Product")
	tracing.TraceVal(ctx, "SQL", sql)
	for i, arg := range args {
		tracing.TraceIVal(ctx, "arg-"+strconv.Itoa(i), arg)
	}

	cmd, execErr := repo.client.Exec(ctx, sql, args...)
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
