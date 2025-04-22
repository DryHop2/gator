package state

import (
	"context"

	"github.com/DryHop2/gator/internal/config"
	"github.com/DryHop2/gator/internal/database"
)

type State struct {
	Ctx context.Context
	DB  *database.Queries
	Cfg *config.Config
}
