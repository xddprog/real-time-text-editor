package setup

import (
	"fmt"
	"golang/internal/core/repositories"
	"golang/internal/core/services"
	"golang/internal/handlers/v1"
	"golang/internal/infrastructure/config"
	"golang/internal/infrastructure/types"

	"github.com/jackc/pgx/v5/pgxpool"
)


func InitNewHandler[T types.HandlerInterface](emptyHandler T, db *pgxpool.Pool) (T, error) {
	switch h := any(emptyHandler).(type) {
	case *handlers.UserHandler:
		repository := &repositories.UserRepository{DB: db}
		service := &services.UserService{Repository: repository}
		*h = handlers.UserHandler{Service: service}
		return any(h).(T), nil
		
	case *handlers.AuthHandler:		
		cfg, err := config.LoadJwtConfig()
		if err != nil {
			panic(err)
		}
		
		repository := &repositories.UserRepository{DB: db}
		service := &services.AuthService{Repository: repository, Config: cfg}
		*h = handlers.AuthHandler{Service: service}
		return any(h).(T), nil

	case *handlers.DocumentHandler:
		repository := &repositories.DocumentRepository{DB: db}
		service := &services.DocumentService{Repository: repository}
		*h = handlers.DocumentHandler{Service: service}
		return any(h).(T), nil
		
	default:
		return emptyHandler, fmt.Errorf("undefined handler type: %T", emptyHandler)
	}
}