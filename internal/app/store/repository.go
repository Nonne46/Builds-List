package store

import "github.com/Nonne46/Builds-List/internal/app/model"

type BuildRepository interface {
	//	Create(*model.Build) error
	FindById(int) (*model.Build, error)
	FindBySearch(string) []model.Build
	GetBuilds() []model.Build
}

type CommentRepository interface {
	AddComment(*model.Comment) error
	FindByBuildId(int) []model.Comment
	CountByBuildId(int) int
}

type UserRepository interface {
	CreateUser(*model.User) error
	FindById(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
