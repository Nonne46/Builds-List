package sqlstore

import (
	"database/sql"

	"github.com/Nonne46/Builds-List/internal/app/model"
	"github.com/Nonne46/Builds-List/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) CreateUser(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"INSERT INTO users (name, email, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		u.Name,
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) FindById(id int) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
