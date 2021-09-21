package db

import (
	"database/sql"

	"baltard/internal/domain/model"
	repo "baltard/internal/domain/repository"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryImpl struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repo.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (u UserRepositoryImpl) Create(uid, secret string) (*model.User, error) {
	rows, err := u.DB.Exec(`
		INSERT INTO
			users (
				uid,
				generated_secret
			)
		VALUES (?, ?)
	`,
		uid,
		secret,
	)
	if err != nil {
		return nil, err
	}

	insertedId, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	eu := model.User{
		Id:     int(insertedId),
		Uid:    uid,
		Secret: secret,
	}
	return &eu, nil
}

func (u UserRepositoryImpl) FindById(userId int) (*model.User, error) {
	user := model.User{}
	row := u.DB.QueryRowx(`
		SELECT
			id,
			uid,
			generated_secret
		FROM
			users
		WHERE
			id = ?
	`, userId)
	if err := row.StructScan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepositoryImpl) FindByUid(uid string) (*model.User, error) {
	user := model.User{}
	row := u.DB.QueryRowx(`
		SELECT
			id,
			uid,
			generated_secret
		FROM
			users
		WHERE
			uid = ?
	`, uid)
	if err := row.StructScan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u UserRepositoryImpl) AddCompletionCode(userId, code int) error {
	_, err := u.DB.Exec(`
		INSERT INTO 
			completion_codes (
				user_id, 
				completion_code
			)
		VALUES (?, ?)`,
		userId,
		code,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepositoryImpl) GetCompletionCodeById(userId int) (int, error) {
	var code sql.NullInt64
	row := u.DB.QueryRow(`
		SELECT
			completion_code
		FROM
			completion_codes
		WHERE
			user_id = ?
	`, userId)

	if err := row.Scan(&code); err != nil {
		return 0, err
	}

	if code.Valid {
		return int(code.Int64), nil
	} else {
		return 42, sql.ErrNoRows
	}
}