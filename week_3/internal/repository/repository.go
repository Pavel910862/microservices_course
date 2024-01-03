package repository

import (
	"context"

	"github.com/Pavel910862/microservices_course/week_3/internal/model"
)

type NoteRepository interface {
	Create(ctx context.Context, info *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
}
