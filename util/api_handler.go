package util

// ResultHandler handle api result and reason
func ResultHandler(result string, reason string) string {
	return "{\"result\": \"" + result + "\", \"reason\": \"" + reason + "\"}"
}
