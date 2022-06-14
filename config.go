package whm

type Config struct {
	Host     string //with http:// or https://
	Port     string //whm port number
	ApiToken string //whm api token that you created
	User     string //whm user name : such as root
}

func NewConfig(host, port, user, apitoken string) *Config {
	return &Config{
		Host:     host,
		ApiToken: apitoken,
		Port:     port,
		User:     user,
	}
}
