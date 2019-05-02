package util

type dbConnect struct {
	Account  string
	Password string
	DBName   string
}

type apiResult struct {
	Success     string
	FormatError string
	DBError     string
}

// Result api result
var Result = apiResult{
	Success:     "success",
	FormatError: "formatError",
	DBError:     "dbError"}

// DBConnect db config
var DBConnect = dbConnect{
	Account:  "glory",
	Password: "aA!29621097aA!",
	DBName:   "wedding"}
