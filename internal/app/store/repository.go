package store

import "github.com/Nonne46/Builds-List/internal/app/model"

type BuildRepository interface {
	//	Create(*model.Build) error
	FindById(int) (*model.Build, error)
	FindBySearch(string) []model.Build
	GetBuilds() []model.Build
}

type CommentRepository interface {
	FindByBuildId(int) []model.Comment
	CountByBuildId(int) int
}
