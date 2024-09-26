package constants

const DEFAULT_PORT string = "9090"

const DEFAULT_TOPIC_NAME string = "values"

var ENV_NAMES = EnvNames{
	BROKER_ADDRESS: "BROKER_ADDRESS",
	ENV_SOURCE:     "ENV_SOURCE",
	PORT:           "PORT",
}

var RESPONSE_INFO = ResponseInfo{
	BadRequest:          "BAD_REQUEST",
	InternalServerError: "INTERNAL_SERVER_ERROR",
	MissingData:         "MISSING_DATA",
	Ok:                  "OK",
}
