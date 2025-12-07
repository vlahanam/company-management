package repositories

import "context"

func (s *mysqlStorage) GetUserRoleNames(ctx context.Context, userID uint64) ([]string, error) {
	var roleNames []string

	err := s.db.WithContext(ctx).
		Table("roles").
		Select("roles.name").
		Joins("INNER JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Pluck("name", &roleNames).Error

	if err != nil {
		return nil, err
	}

	return roleNames, nil
}
