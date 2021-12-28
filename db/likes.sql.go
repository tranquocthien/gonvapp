// Code generated by sqlc. DO NOT EDIT.
// source: likes.sql

package db

import (
	"context"
)

const checkAlreadyDislike = `-- name: CheckAlreadyDislike :one
SELECT exists(select 1 from book_group_likes where user_id = $1 and book_group_id = $2 and point < 0)
`

type CheckAlreadyDislikeParams struct {
	UserID      int32 `json:"userID"`
	BookGroupID int32 `json:"bookGroupID"`
}

func (q *Queries) CheckAlreadyDislike(ctx context.Context, arg CheckAlreadyDislikeParams) (bool, error) {
	row := q.db.QueryRow(ctx, checkAlreadyDislike, arg.UserID, arg.BookGroupID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkAlreadyLike = `-- name: CheckAlreadyLike :one
SELECT EXISTS(select 1 from book_group_likes where user_id = $1 and book_group_id = $2 and point > 0)
`

type CheckAlreadyLikeParams struct {
	UserID      int32 `json:"userID"`
	BookGroupID int32 `json:"bookGroupID"`
}

func (q *Queries) CheckAlreadyLike(ctx context.Context, arg CheckAlreadyLikeParams) (bool, error) {
	row := q.db.QueryRow(ctx, checkAlreadyLike, arg.UserID, arg.BookGroupID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkUnlike = `-- name: CheckUnlike :one
SELECT EXISTS(select 1 from book_group_likes where user_id = $1 and book_group_id = $2)
`

type CheckUnlikeParams struct {
	UserID      int32 `json:"userID"`
	BookGroupID int32 `json:"bookGroupID"`
}

func (q *Queries) CheckUnlike(ctx context.Context, arg CheckUnlikeParams) (bool, error) {
	row := q.db.QueryRow(ctx, checkUnlike, arg.UserID, arg.BookGroupID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const disLikes = `-- name: DisLikes :exec
INSERT INTO book_group_likes(user_id, book_group_id, point)
VALUES ($1, $2, -1)
`

type DisLikesParams struct {
	UserID      int32 `json:"userID"`
	BookGroupID int32 `json:"bookGroupID"`
}

func (q *Queries) DisLikes(ctx context.Context, arg DisLikesParams) error {
	_, err := q.db.Exec(ctx, disLikes, arg.UserID, arg.BookGroupID)
	return err
}

const getDislikes = `-- name: GetDislikes :one
SELECT coalesce(SUM(point), 0) as totalLikes FROM book_group_likes WHERE book_group_id = $1 AND point < 0
`

func (q *Queries) GetDislikes(ctx context.Context, bookGroupID int32) (interface{}, error) {
	row := q.db.QueryRow(ctx, getDislikes, bookGroupID)
	var totallikes interface{}
	err := row.Scan(&totallikes)
	return totallikes, err
}

const getLikes = `-- name: GetLikes :one
SELECT coalesce(SUM(point), 0) as totalLikes FROM book_group_likes WHERE book_group_id = $1 AND point > 0
`

func (q *Queries) GetLikes(ctx context.Context, bookGroupID int32) (interface{}, error) {
	row := q.db.QueryRow(ctx, getLikes, bookGroupID)
	var totallikes interface{}
	err := row.Scan(&totallikes)
	return totallikes, err
}

const likes = `-- name: Likes :exec
INSERT INTO book_group_likes(user_id, book_group_id, point)
VALUES ($1, $2, 1)
`

type LikesParams struct {
	UserID      int32 `json:"userID"`
	BookGroupID int32 `json:"bookGroupID"`
}

func (q *Queries) Likes(ctx context.Context, arg LikesParams) error {
	_, err := q.db.Exec(ctx, likes, arg.UserID, arg.BookGroupID)
	return err
}

const unlikes = `-- name: Unlikes :exec
DELETE FROM book_group_likes WHERE user_id = $1 AND book_group_id = $2
`

type UnlikesParams struct {
	UserID      int32 `json:"userID"`
	BookGroupID int32 `json:"bookGroupID"`
}

func (q *Queries) Unlikes(ctx context.Context, arg UnlikesParams) error {
	_, err := q.db.Exec(ctx, unlikes, arg.UserID, arg.BookGroupID)
	return err
}