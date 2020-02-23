package usecase

import (
	"context"

	"github.com/dmirou/otusgo/calendar/pkg/event"
)

type UseCase struct {
	repo event.Repository
}

func New(repo event.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) CreateEvent(ctx context.Context, e *event.Event) error {
	return uc.repo.Create(ctx, e)
}

func (uc *UseCase) GetEventByID(ctx context.Context, id event.ID) (*event.Event, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *UseCase) UpdateEvent(ctx context.Context, e *event.Event) error {
	return uc.repo.Update(ctx, e)
}

func (uc *UseCase) DeleteEvent(ctx context.Context, id event.ID) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *UseCase) ListEventsByDate(ctx context.Context, year, month, day int) ([]*event.Event, error) {
	return uc.repo.FindByDate(ctx, year, month, day)
}
