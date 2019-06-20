package config

//Config will hold structure from config/default.toml
type Config struct {
	Data JsonFile
}

//JsonFile will hold json files inside [Data]
type JsonFile struct {
	Organization string
	Ticket       string
	User         string
}
