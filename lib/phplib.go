package lib

func End(str string) string {
	length := len(str)
	index := length - 1
	return str[index:]
}
