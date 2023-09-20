//go:build integration
// +build integration

package db

import (
	"context"
	"testing"

	"github.com/mobamoh/commentmason/internal/comment"
	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "Test Slug",
			Body:   "Test body",
			Author: "Mo Bamoh",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, cmt.ID)

	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "Test Slug",
			Body:   "Test body",
			Author: "Mo Bamoh",
		})
		assert.NoError(t, err)
		assert.NotEmpty(t, cmt.ID)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})
}
