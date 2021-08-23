package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Nonne46/Builds-List/internal/app/model"
	"github.com/Nonne46/Builds-List/internal/app/store"
)

type BuildRepository struct {
	store *Store
}

func (r *BuildRepository) FindById(id int) (*model.Build, error) {
	build := &model.Build{}

	if err := r.store.db.QueryRow(
		"SELECT * FROM build WHERE id = $1",
		id,
	).Scan(
		&build.Id,
		&build.Name,
		&build.Description,
		&build.NameAddr,
		&build.IsAlive,
		&build.Tags,
		&build.AuthorRepo,
		&build.ByondVersion,
		&build.Github,
		&build.BackupDate,
		&build.Thanks,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return build, nil
}

func (r *BuildRepository) FindBySearch(query string) []model.Build {
	builds := []model.Build{}
	searchedBuilds := []model.Build{}

	searchReq := strings.ToLower(query)

	rows, _ := r.store.db.Query("select * from build")
	defer rows.Close()

	for rows.Next() {
		p := model.Build{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.NameAddr, &p.IsAlive, &p.Tags, &p.AuthorRepo, &p.ByondVersion, &p.Github, &p.BackupDate, &p.Thanks)
		if err != nil {
			fmt.Println(err)
			continue
		}
		builds = append(builds, p)
	}

	for _, build := range builds {
		if !strings.Contains(strings.ToLower(build.Name), searchReq) &&
			!strings.Contains(strings.ToLower(build.Description), searchReq) &&
			!strings.Contains(strings.ToLower(build.Thanks), searchReq) &&
			!strings.Contains(strings.ToLower(build.Github), searchReq) &&
			!strings.Contains(strings.ToLower(build.AuthorRepo), searchReq) {
			continue
		}
		searchedBuilds = append(searchedBuilds, build)
	}
	return searchedBuilds
}

func (r *BuildRepository) GetBuilds() []model.Build {
	builds := []model.Build{}

	rows, _ := r.store.db.Query("select * from build")
	defer rows.Close()

	for rows.Next() {
		p := model.Build{}
		err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.NameAddr, &p.IsAlive, &p.Tags, &p.AuthorRepo, &p.ByondVersion, &p.Github, &p.BackupDate, &p.Thanks)
		if err != nil {
			fmt.Println(err)
			continue
		}
		builds = append(builds, p)
	}

	return builds
}
