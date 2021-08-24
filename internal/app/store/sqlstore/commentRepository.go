package sqlstore

import (
	"fmt"

	"github.com/Nonne46/Builds-List/internal/app/model"
)

type CommentRepository struct {
	store *Store
}

func (r *CommentRepository) AddComment(c *model.Comment) error {
	c.CheckUsername()

	if err := c.CheckMessage(); err != nil {
		return err
	}

	_, err := r.store.db.Exec(
		"INSERT INTO comments (idPage, userName, commentText, time) VALUES ($1, $2, $3, $4)",
		c.IdPage,
		c.Username,
		c.Comment,
		c.Time,
	)
	return err
}

func (r *CommentRepository) FindByBuildId(id int) []model.Comment {
	comments := []model.Comment{}

	rows, _ := r.store.db.Query("SELECT * FROM comments WHERE idPage LIKE ? ORDER BY id DESC", id)
	defer rows.Close()

	for rows.Next() {
		p := model.Comment{}
		err := rows.Scan(&p.Id, &p.IdPage, &p.Username, &p.Comment, &p.Time)
		if err != nil {
			fmt.Println(err)
			continue
		}
		comments = append(comments, p)
	}

	return comments
}

func (r *CommentRepository) CountByBuildId(id int) int {
	var count int

	row := r.store.db.QueryRow("SELECT COUNT(*) FROM comments WHERE idPage = $1", id)
	row.Scan(&count)

	return count
}
