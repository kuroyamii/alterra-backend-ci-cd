package config

import "os"

func GetEnvVariables() map[string]string {
	envVariables := make(map[string]string)

	envVariables["SERVER_ADDRESS"] = os.Getenv("SERVER_ADDRESS")
	envVariables["DB_ADDRESS"] = os.Getenv("DB_ADDRESS")
	envVariables["DB_USERNAME"] = os.Getenv("DB_USERNAME")
	envVariables["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVariables["DB_NAME"] = os.Getenv("DB_NAME")
	return envVariables
}
