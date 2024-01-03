package note

import (
	"github.com/Pavel910862/microservices_course/week_3/internal/client/db"
	"github.com/Pavel910862/microservices_course/week_3/internal/repository"
	"github.com/Pavel910862/microservices_course/week_3/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(
	noteRepository repository.NoteRepository,
	txManager db.TxManager,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}
