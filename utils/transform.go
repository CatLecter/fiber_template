package utils

func IsBool(value string) bool {
	switch {
	case value == "" || value == "false":
		return false
	default:
		return true
	}
}
