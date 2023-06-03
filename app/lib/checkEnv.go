package lib

import "os"


func CheckEnv(env string, def string) string {
	variable := os.Getenv(env)
	if len(variable) == 0 {
		return def
	}
	return variable
}
