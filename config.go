package whm

type Config struct {
	Host string
	Port string
	ApiToken  string
	Secret string
}

func NewConfig(host,apitoken, port string) *Config {
	return &Config{
		Host: host,
		ApiToken:  apitoken,
		Port: port,
	}
}