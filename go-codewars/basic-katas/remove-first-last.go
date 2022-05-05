package main

func removeBookends(str string) string {
	return str[1 : len(str)-1]
}
