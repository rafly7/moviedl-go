package configs

func ConfigServer() (string, string, string) {
	host := LoadEnv("APP_HOST")
	port := LoadEnv("APP_PORT")
	secretContext := LoadEnv("SECRET_CONTEXT")
	return host, port, secretContext
}
