package utils

var Consts = struct {
	//HTTP
	RequiredParams string

	//SQL
	SqlNotFound string
}{
	RequiredParams: "CreateUser handler: One or more required params are missing or empty value(s).",
	SqlNotFound:    "The item was not found",
}
