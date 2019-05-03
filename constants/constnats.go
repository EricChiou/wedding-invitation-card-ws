package cons

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
