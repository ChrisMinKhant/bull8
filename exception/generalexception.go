package exception

import (
	"github.com/sirupsen/logrus"
)

type generalException struct{}

var generalExceptionInstance *generalException

func GetGeneralExceptionInstance() *generalException {

	if generalExceptionInstance != nil {
		return generalExceptionInstance
	}

	generalExceptionInstance = &generalException{}

	return generalExceptionInstance
}

func (generalException *generalException) RecoverPanic() {
	if recoveryStatus := recover(); recoveryStatus != nil {

		logrus.Infof("System has been recovered from exception ::: [ %v ]\n", recoveryStatus)
		return

	}
}
