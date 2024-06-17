package util

import (
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type envHelper struct {
	env map[any]any
}

var envHelperInstance *envHelper

func NewEnvHelper() *envHelper {

	if envHelperInstance != nil {
		return envHelperInstance
	}

	envHelperInstance = &envHelper{}

	return envHelperInstance
}

func (envHelper *envHelper) Get(key string) string {

	if envHelper.env == nil {

		log.Println("env value does not exist.")

		readBytes, error := os.ReadFile("env.yaml")

		if error != nil {
			log.Panicf("Error occurred at opening env.yaml file ::: %v\n", error.Error())
		}

		error = yaml.Unmarshal(readBytes, &envHelper.env)

		if error != nil {
			log.Panicf("Error occurred at unmarshalling read bytes from env.yaml file ::: %v\n", error.Error())
		}
	}

	if strings.Contains(key, ".") {
		nestedKey := strings.Split(key, ".")

		rootValue := envHelper.env[nestedKey[0]].(map[string]any)
		return rootValue[nestedKey[1]].(string)
	}

	return envHelper.env[key].(string)

}
