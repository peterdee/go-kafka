package constants

type EnvNames struct {
	BROKER_ADDRESS string
	ENV_SOURCE     string
	PORT           string
}

type ResponseInfo struct {
	BadRequest          string
	InternalServerError string
	MissingData         string
	Ok                  string
}
