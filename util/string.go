package util

import (
	"fmt"
	"strings"
)

func GetSayHi(name string) string {
	return fmt.Sprintf("Hi, %s !", name)
}

func Trim(str string) string {
	return strings.TrimSpace(str)
}
