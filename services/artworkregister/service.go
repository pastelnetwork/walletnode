package artworkregister

import (
	"context"

	"github.com/pastelnetwork/go-commons/errors"
	"github.com/pastelnetwork/go-pastel"
	"github.com/pastelnetwork/walletnode/storage"
	"golang.org/x/sync/errgroup"
)

// const logPrefix = "[artwork]"

// Service represent artwork service.
type Service struct {
	db     storage.KeyValue
	pastel pastel.Client
	worker *Worker
	tasks  []*Task
}

// Run starts worker
func (service *Service) Run(ctx context.Context) error {
	//return service.worker.Run(ctx)

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() (err error) {
		defer errors.Recover(func(rec error) { err = rec })
		return service.worker.Run(ctx)
	})
	task := NewTask(service, &Ticket{})
	service.worker.AddTask(ctx, task)

	return group.Wait()
}

// Tasks returns all tasks.
func (service *Service) Tasks() []*Task {
	return service.tasks
}

// Task returns the task of the registration artwork.
func (service *Service) Task(taskID int) *Task {
	for _, task := range service.tasks {
		if task.ID == taskID {
			return task
		}
	}
	return nil
}

// Register runs a new task of the registration artwork and returns its taskID.
func (service *Service) Register(ctx context.Context, ticket *Ticket) (int, error) {
	// NOTE: for testing
	task := NewTask(service, ticket)
	service.tasks = append(service.tasks, task)
	service.worker.AddTask(ctx, task)

	return task.ID, nil
}

// NewService returns a new Service instance.
func NewService(db storage.KeyValue, pastel pastel.Client) *Service {
	return &Service{
		db:     db,
		pastel: pastel,
		worker: NewWorker(),
	}
}
