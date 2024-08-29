package constants

type EnvNames struct {
	BrokerAddress string
	Port          string
}

type ResponseInfo struct {
	BadRequest          string
	InternalServerError string
	MissingData         string
	Ok                  string
}
