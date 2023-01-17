package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"userservice/internal/dto/req"
	"userservice/internal/model"
	"userservice/internal/repository"
	"userservice/internal/repository/sqltype"
	"userservice/internal/util"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/exp/slog"
)

type userRepository struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewUserRepository(logger *slog.Logger, db *sql.DB) repository.IUserRepository {
	return &userRepository{
		logger: logger,
		db:     db,
	}
}

func (ur userRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	const scope = "userRepository#FindByID"
	user := &sqltype.User{}
	err := ur.db.QueryRow(
		`
			SELECT "id", "email", "password", "first_name", "last_name", "created_at", "updated_at", "deleted_at"
			FROM "users_tab"
			WHERE "id" = $1 AND "deleted_at" IS NULL;
		`,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		ur.logger.Error(
			"Failed to find a user by its ID",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", scope, repository.ErrDataNotFound.SetError(err))
		}
		pgError, ok := err.(*pgconn.PgError)
		if !ok {
			return nil, fmt.Errorf("%s: %w", scope, repository.ErrUnknown.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, repository.ErrUnknown.SetError(pgError))
	}
	ur.logger.Info(
		"Found a user by its ID",
		slog.String("request_id", ctx.Value(util.RequestID).(string)),
		slog.String("scope", scope),
	)
	return user.ToModel(), nil
}

func (ur userRepository) Create(ctx context.Context, userDto req.UserDto) (*model.User, error) {
	const scope = "userRepository#Create"
	user := &sqltype.User{}
	err := ur.db.QueryRow(
		`
			INSERT INTO "users_tab" ("email", "password", "first_name", "last_name", "created_at", "updated_at")
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING "id", "email", "password", "first_name", "last_name", "created_at", "updated_at", "deleted_at";
		`,
		userDto.Email,
		userDto.Password,
		userDto.FirstName,
		userDto.LastName,
		time.Now(),
		time.Now(),
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		ur.logger.Error(
			"Failed to create a user",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		pgError, ok := err.(*pgconn.PgError)
		if !ok {
			return nil, fmt.Errorf("%s: %w", scope, repository.ErrUnknown.SetError(err))
		}
		if pgError.Code == pgerrcode.UniqueViolation {
			return nil, fmt.Errorf("%s: %w", scope, repository.ErrUniqueViolation.SetError(pgError))
		}
		return nil, fmt.Errorf("%s: %w", scope, repository.ErrUnknown.SetError(pgError))
	}
	ur.logger.Info(
		"Created a user",
		slog.String("request_id", ctx.Value(util.RequestID).(string)),
		slog.String("scope", scope),
	)
	return user.ToModel(), nil
}

func (ur userRepository) DeleteByID(ctx context.Context, id int) (*model.User, error) {
	const scope = "userRepository#DeleteByID"
	user := &sqltype.User{}
	err := ur.db.QueryRow(
		`
			UPDATE "users_tab"
			SET "deleted_at" = $1
			WHERE "id" = $2 AND "deleted_at" IS NULL
			RETURNING "id", "email", "password", "first_name", "last_name", "created_at", "updated_at", "deleted_at";
		`,
		time.Now(),
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		ur.logger.Error(
			"Failed to soft delete a user by its ID",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", scope, repository.ErrDataNotFound.SetError(err))
		}
		pgError, ok := err.(*pgconn.PgError)
		if !ok {
			return nil, fmt.Errorf("%s: %w", scope, repository.ErrUnknown.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, repository.ErrUnknown.SetError(pgError))
	}
	ur.logger.Info(
		"Soft deleted a user by its ID",
		slog.String("request_id", ctx.Value(util.RequestID).(string)),
		slog.String("scope", scope),
	)
	return user.ToModel(), nil
}

func (ur userRepository) DeletePermanentlyByID(ctx context.Context, id int) (*model.User, error) {
	const scope = "userRepository#DeletePermanentlyByID"
	user := &sqltype.User{}
	err := ur.db.QueryRow(
		`
			DELETE FROM "users_tab"
			WHERE "id" = $1
			RETURNING "id", "email", "password", "first_name", "last_name", "created_at", "updated_at", "deleted_at";
		`,
		id,
	).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		ur.logger.Error(
			"Failed to delete a user permanently by its ID",
			err,
			slog.String("request_id", ctx.Value(util.RequestID).(string)),
			slog.String("scope", scope),
		)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", scope, repository.ErrDataNotFound.SetError(err))
		}
		pgError, ok := err.(*pgconn.PgError)
		if !ok {
			return nil, fmt.Errorf("%s: %w", scope, repository.ErrUnknown.SetError(err))
		}
		return nil, fmt.Errorf("%s: %w", scope, repository.ErrUnknown.SetError(pgError))
	}
	ur.logger.Info(
		"Deleted a user permanently by its ID",
		slog.String("request_id", ctx.Value(util.RequestID).(string)),
		slog.String("scope", scope),
	)
	return user.ToModel(), nil
}
