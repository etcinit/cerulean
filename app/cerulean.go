package app

import (
	"github.com/etcinit/cerulean/database"
	"github.com/etcinit/ohmygorm"
	"github.com/jacobstr/confer"
)

// Cerulean is the root node of the DI graph
type Cerulean struct {
	Config      *confer.Config               `inject:""`
	Connections *ohmygorm.ConnectionsService `inject:""`
	Engine      *EngineService               `inject:""`
	Serve       *ServeService                `inject:""`
	Migrator    *database.MigratorService    `inject:""`
}
