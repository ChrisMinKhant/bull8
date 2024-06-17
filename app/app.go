package app

import (
	"net/http"

	"github.com/ChrisMinKhant/megoyougo_framework/database"
	"github.com/ChrisMinKhant/megoyougo_framework/provider"
	"github.com/ChrisMinKhant/megoyougo_framework/provider/handlerprovider"
	"github.com/ChrisMinKhant/megoyougo_framework/util"
	"github.com/sirupsen/logrus"
)

/*
 * The app package prepares all the necessary
 * packages to be booted up just by one invoking
 * from cmd.
 */

type app struct {
	handlerprovider provider.Provider
	database        database.Database
}

func NewApp() *app {
	return &app{
		handlerprovider: handlerprovider.NewHandlerProvider(),
		database:        database.NewPostgres(),
	}
}

func (app *app) Start() {

	servedPort := util.NewEnvHelper().Get("server.port")

	/*
	 * All the handlers are booted up here.
	 */

	app.handlerprovider.Register()

	/*
	 * All entities of the database are booted up here.
	 */

	app.database.Initialize()

	logrus.Infof("The server is started and serving at port ::: %v\n", servedPort)
	/*
	 * Then, the application is serve at defined port.
	 */

	http.ListenAndServe(":"+servedPort, NewGateWay())
}
