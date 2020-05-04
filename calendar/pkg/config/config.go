package config

type Info struct {
	IP   string
	Port int
}

type Server struct {
	IP   string
	Port int
}

type Log struct {
	File  string
	Level string
}

type Config struct {
	Info   Info
	Log    Log
	Server Server
}
