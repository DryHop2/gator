package state

import (
	"github.com/DryHop2/gator/internal/config"
	"github.com/DryHop2/gator/internal/database"
)

type State struct {
	DB  *database.Queries
	Cfg *config.Config
}
