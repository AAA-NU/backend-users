package storage

import (
	"context"

	"github.com/aaanu/backendusers/internal/domain/models"
)

func (s *Storage) User(ctx context.Context, tgID string) (*models.User, error) {
	var user models.User
	if err := s.db.WithContext(ctx).Where("telegram_id = ?", tgID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Storage) Users(ctx context.Context, role string) ([]models.User, error) {
	var users []models.User
	query := s.db.WithContext(ctx)
	if role != "" {
		query = query.Where("role = ?", role)
	}
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Storage) SaveUser(ctx context.Context, user *models.User) error {
	return s.db.WithContext(ctx).Save(user).Error
}

func (s *Storage) DeleteUser(ctx context.Context, tgID string) error {
	return s.db.WithContext(ctx).Where("telegram_id = ?", tgID).Delete(&models.User{}).Error
}

func (s *Storage) UpdateUser(ctx context.Context, tgID string, role string, language string) error {
	query := s.db.WithContext(ctx).Table("users").Where("telegram_id = ?", tgID)
	if role != "" {
		query = query.Update("role", role)
	}
	if language != "" {
		query = query.Update("language", language)
	}
	return query.Error
}
