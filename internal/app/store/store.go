package store

// Store ...
type Store interface {
	Build() BuildRepository
	Comment() CommentRepository
}
