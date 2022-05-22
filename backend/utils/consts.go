package utils

var Consts = struct {
	//HTTP
	RequiredParams string
	InvalidPayload string

	//SQL
	SqlNotFound string

	//CACHE
	KeyNotExist string
}{
	RequiredParams: "CreateUser handler: One or more required params are missing or empty value(s).",
	InvalidPayload: "Invalid request payload.",
	SqlNotFound:    "The item was not found.",
	KeyNotExist: "does not exist.",
}
