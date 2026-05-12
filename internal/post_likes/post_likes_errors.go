package plikes

import "errors"

var (
	ErrPostAlreadyLiked = errors.New("post already liked")
	ErrLikeNotFound     = errors.New("post like data not found")
)
