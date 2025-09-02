package service

import (
	"context"
	"log/slog"

	"github.com/aaanu/backendusers/internal/customerrors"
	"github.com/aaanu/backendusers/internal/domain/models"
	"github.com/aaanu/backendusers/internal/domain/requests"
)

type UsersStorage interface {
	User(ctx context.Context, tgID string) (*models.User, error)
	Users(ctx context.Context, role string) ([]models.User, error)
	SaveUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, tgID string, role string, language string) error
	DeleteUser(ctx context.Context, tgID string) error
}

type UsersService struct {
	log     *slog.Logger
	storage UsersStorage
}

func New(log *slog.Logger, storage UsersStorage) *UsersService {
	log = log.With("service", "users")
	return &UsersService{
		log:     log,
		storage: storage,
	}
}

func (s *UsersService) User(ctx context.Context, tgID string) (*models.User, error) {
	const op = "service.User"

	log := s.log.With("op", op)
	log.Info("get user", "telegram_id", tgID)

	user, err := s.storage.User(ctx, tgID)
	if err != nil {
		err = customerrors.New(err.Error(), customerrors.ErrNotFound)
		log.Error("failed to get user", "error", err)
		return nil, err
	}

	log.Info("got user", "user", user)
	return user, nil
}

func (s *UsersService) Users(ctx context.Context, role string) ([]models.User, error) {
	const op = "service.Users"

	log := s.log.With("op", op)
	log.Info("get users", "role", role)

	users, err := s.storage.Users(ctx, role)
	if err != nil {
		err = customerrors.New(err.Error(), customerrors.ErrBadRequest)
		log.Error("failed to get users", "error", err)
		return nil, err
	}

	log.Info("got users", "users", users)
	return users, nil
}

func (s *UsersService) SaveUser(ctx context.Context, userRequest *requests.SaveUserRequest) error {
	const op = "service.SaveUser"

	log := s.log.With("op", op)
	log.Info("save user", "user", userRequest)

	user := &models.User{
		TelegramID: userRequest.TelegramID,
		Role:       "student",
		Language:   userRequest.Language,
	}

	if err := s.storage.SaveUser(ctx, user); err != nil {
		err = customerrors.New(err.Error(), customerrors.ErrBadRequest)
		log.Error("failed to save user", "error", err)
		return err
	}

	log.Info("saved user", "user", user)
	return nil
}

func (s *UsersService) UpdateUser(ctx context.Context, tgID string, role string, language string) error {
	const op = "service.UpdateUser"

	log := s.log.With("op", op)
	log.Info("update user", "telegram_id", tgID, "role", role, "language", language)

	if err := s.storage.UpdateUser(ctx, tgID, role, language); err != nil {
		err = customerrors.New(err.Error(), customerrors.ErrBadRequest)
		log.Error("failed to update user", "error", err)
		return err
	}

	log.Info("updated user", "telegram_id", tgID, "role", role, "language", language)
	return nil
}

func (s *UsersService) DeleteUser(ctx context.Context, tgID string, fromUserID string) error {
	const op = "service.DeleteUser"

	log := s.log.With("op", op)
	log.Info("delete user", "telegram_id", tgID)

	user, err := s.storage.User(ctx, fromUserID)
	if err != nil {
		err = customerrors.New(err.Error(), customerrors.ErrForbidden)
		log.Error("failed to get user", "error", err)
		return err
	}

	if user.Role != "admin" {
		return nil
	}

	if err := s.storage.DeleteUser(ctx, tgID); err != nil {
		err = customerrors.New(err.Error(), customerrors.ErrBadRequest)
		log.Error("failed to delete user", "error", err)
		return err
	}

	log.Info("deleted user", "telegram_id", tgID)
	return nil
}
