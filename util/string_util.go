package util

import "strings"

func GetEtcdErrorType(err string) []string {
	strs := strings.Split(err, ":")
	for i := 0; i < len(strs); i++ {
		strs[i] = strings.TrimSpace(strs[i])
	}
	return strs
}
