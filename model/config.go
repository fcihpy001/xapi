package model

type Config struct {
	Mysql    Mysql
	Redis    Redis
	Version  string
	Dsn      string
	RedisURL string
	Server   Server
	Debug    bool
}

type Server struct {
	Host string
	Port string
}

type Mysql struct {
	Host     string
	Port     string
	Database string
	UserName string
	Password string
	Charset  string
	Loc      string
}
type Redis struct {
	Host string
	Port string
}
