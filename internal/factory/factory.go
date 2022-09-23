package factory

import (
	"github.com/dakasakti/deploy-apps-hexagonal/config"
	"github.com/dakasakti/deploy-apps-hexagonal/database"
	"github.com/dakasakti/deploy-apps-hexagonal/internal/repository/user"
)

type Factory struct {
	UserRepository user.UserRepository
}

func NewFactory(config *config.AppConfig) *Factory {
	db, mc := database.InitConnection(config)

	return &Factory{
		user.NewUserRepository(db, mc),
	}
}
