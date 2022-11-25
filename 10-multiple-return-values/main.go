package main

import (
	"strings"
	"fmt"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], ""
}

func main() {
	fn, ln := getInitials("Samuel Kelechukwu")

	fmt.Println(fn, ln, "Onyegbuna")
}