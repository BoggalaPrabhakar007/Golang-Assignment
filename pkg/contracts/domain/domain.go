package domain

//Config holds the all the configuration details
type Config struct {
	Server   Server
	Database Database
	FilePath FilePath
}

//Server holds the all the server details
type Server struct {
	Port string
}

//Database holds the all the database details
type Database struct {
	Username string
	Password string
	Port     string
	ConnStr  string
}

//FilePath holds the all the file path details
type FilePath struct {
	FilePath string
}
