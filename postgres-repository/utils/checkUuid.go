package utils

import "regexp"

var uuidRegExp = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

func CheckUuid(value string) bool {
	return uuidRegExp.MatchString(value)
}
