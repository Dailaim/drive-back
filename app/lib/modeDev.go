package lib

var Mode = CheckEnv("MODE", "prod")

func IsProduction() bool {
	return CheckEnv("MODE", "prod") == "prod"
}
