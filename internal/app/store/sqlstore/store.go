package sqlstore

import (
	"database/sql"

	"github.com/Nonne46/Builds-List/internal/app/store"
)

type Store struct {
	db                *sql.DB
	buildRepository   *BuildRepository
	commentRepository *CommentRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Build() store.BuildRepository {
	if s.buildRepository != nil {
		return s.buildRepository
	}

	s.buildRepository = &BuildRepository{
		store: s,
	}

	return s.buildRepository
}

func (s *Store) Comment() store.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}

	s.commentRepository = &CommentRepository{
		store: s,
	}

	return s.commentRepository
}
