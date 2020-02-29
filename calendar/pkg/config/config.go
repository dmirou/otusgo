package config

type Server struct {
	IP   string
	Port int
}

type Log struct {
	File  string
	Level string
}

type Config struct {
	Server Server
	Log    Log
}
