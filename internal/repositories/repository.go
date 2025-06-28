// Package repositories содержит репозитории для работы с данными.
package repositories

// Repository содержит все репозитории приложения.
type Repository struct{ UserRepositoryInterface }

// NewRepository создает новый экземпляр репозитория.
func NewRepository() *Repository {
	return &Repository{
		UserRepositoryInterface: NewUserRepository(),
	}
}
