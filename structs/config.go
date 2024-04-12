package structs

type Config struct {
	ServerPort       string
	ApiKey           string
	AdminKey         string
	TransferProtocol string
	AllowAdmin       bool
	AllowHosting     bool
	NeedKey          bool
	DBName           string
	MongoURI         string
}
