package main

import (
	"fmt"
	"strings"
)

func reverse(str string) {

	st := strings.Split(str, " ")

	for i, j := 0, len(st)-1; i < j; i, j = i+1, i-j {
		st[i], st[j] = st[j], st[i]
	}

	str = strings.Join(st, " ")

	fmt.Println(str)

}

func main() {
	reverse("abc")
}
