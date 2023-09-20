package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/mobamoh/commentmason/internal/comment"
)

func commentRowToComment(cmtRow CommentRow) comment.Comment {
	return comment.Comment{
		ID:     cmtRow.ID,
		Slug:   cmtRow.Slug.String,
		Body:   cmtRow.Body.String,
		Author: cmtRow.Author.String,
	}
}

type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func (db *Database) GetComment(ctx context.Context, id string) (comment.Comment, error) {
	// _, err := db.Client.ExecContext(ctx, SELECT pg_sleep(16))
	var cmtRow CommentRow
	row := db.Client.QueryRowContext(ctx,
		`SELECT id, slug, body, author
		FROM comments
		WHERE id = $1`, id,
	)

	if err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author); err != nil {
		return comment.Comment{}, fmt.Errorf("error fetching comment with id:%s", id)
	}
	return commentRowToComment(cmtRow), nil
}

func (db *Database) CreateComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {
	cmt.ID = uuid.NewString()
	cmtRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	rows, err := db.Client.NamedQueryContext(ctx,
		`INSERT INTO comments (id, slug, body, author)
	VALUES (:id, :slug, :body, :author)`, cmtRow)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to insert comment: %w", err)
	}
	if err = rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("faled to close rows: %w", err)
	}
	return cmt, nil
}

func (db *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := db.Client.ExecContext(ctx,
		`DELETE FROM comments WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}
	return nil
}

func (db *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {

	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
	}
	rows, err := db.Client.NamedQueryContext(ctx,
		`UPDATE comments
		SET slug = :slug, body = :body, author = :author
		WHERE id = :id`, cmtRow)
	if err != nil {
		return comment.Comment{}, fmt.Errorf("failed to update comment: %w", err)
	}
	if err = rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("faled to close rows: %w", err)
	}
	return commentRowToComment(cmtRow), nil
}
