package database

import (
	"github.com/ChrisMinKhant/megoyougo_framework/entity"
	"github.com/ChrisMinKhant/megoyougo_framework/entity/car"
	"github.com/ChrisMinKhant/megoyougo_framework/util"
	"gorm.io/gorm"
)

type postgres struct {
	dbConnection *gorm.DB
	entities     []entity.Entity
}

func NewPostgres() *postgres {
	return &postgres{
		dbConnection: util.GetDbHelperInstance().Connect(),
	}
}

/*
 * All the entities must be register here.
 * Then, those will be auto migrated and db tables
 * will be produced.
 */

func (postgres *postgres) register() {

	pushedEntities := [...]entity.Entity{
		car.New(),
	}

	postgres.entities = pushedEntities[:]
}

func (postgres *postgres) migrate() {

	// Auto mirating each entities.
	for _, entity := range postgres.entities {
		entity.Initialize(postgres.dbConnection)
	}
}

func (postgres *postgres) Initialize() {
	postgres.register()
	postgres.migrate()
}
