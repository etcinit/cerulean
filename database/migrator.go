package database

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/etcinit/cerulean/database/models"
	"github.com/etcinit/ohmygorm"
)

// MigratorService provides the migration command
type MigratorService struct {
	Migrations  *ohmygorm.MigrationsService  `inject:""`
	Connections *ohmygorm.ConnectionsService `inject:""`
}

// Run runs all the migrations for Cerulean
func (m *MigratorService) Run(c *cli.Context) {
	fmt.Println("Running migrations...")

	db, _ := m.Connections.Make()
	db.LogMode(true)

	m.Migrations.Run([]interface{}{
		&models.Article{},
		&models.User{},
	})

	fmt.Println("✔︎ Done!")
}
