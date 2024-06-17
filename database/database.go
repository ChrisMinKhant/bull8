package database

type Database interface {
	register()
	migrate()
	Initialize()
}
