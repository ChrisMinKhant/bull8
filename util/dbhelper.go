package util

import (
	"log"

	"github.com/ChrisMinKhant/megoyougo_framework/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
 * Db helper struct is for providing
 * extra-necessary functionalities for
 * application's database function.
 */

type dbHelper struct {
	dbConnection *gorm.DB
	envHelper    *envHelper
	exception    exception.Exception
}

var dbHelperInstance *dbHelper

/*
 * This function make sure and provide the
 * only singleton instance of dbHelper. Which
 * can reduce from establishing database connection
 * and its instances multiple times.
 */

func GetDbHelperInstance() *dbHelper {

	if dbHelperInstance != nil {
		return dbHelperInstance
	}

	dbHelperInstance = &dbHelper{
		envHelper: NewEnvHelper(),
		exception: exception.GetGeneralExceptionInstance(),
	}

	return dbHelperInstance
}

/*
 * Creating database connection instance in
 * singleton manner.
 */
func (dbHelper *dbHelper) Connect() *gorm.DB {

	defer dbHelper.exception.RecoverPanic()

	dbStream := dbHelper.getDatabaseStream()

	/*
	 * " gorm.DB " pointer is if and only if
	 * there is no such pointer is stored in
	 * " dbHelper.dbConnection ".
	 */

	if dbHelper.dbConnection != nil {
		return dbHelper.dbConnection
	}

	dbConnection, error := gorm.Open(postgres.Open(dbStream), &gorm.Config{})

	if error != nil {
		log.Panicf("Error occured at connecting to database with stream ::: %v", dbStream)
	}

	dbHelper.dbConnection = dbConnection

	return dbHelper.dbConnection
}

func (dbHelper *dbHelper) getDatabaseStream() string {
	return "postgres://" + dbHelper.envHelper.Get("database.username") +
		":" + dbHelper.envHelper.Get("database.password") +
		"@" + dbHelper.envHelper.Get("database.host") +
		":" + dbHelper.envHelper.Get("database.port") +
		"/" + dbHelper.envHelper.Get("database.databaseName")
}
