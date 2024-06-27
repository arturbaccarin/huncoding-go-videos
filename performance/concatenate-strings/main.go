package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

/*

fmt.Println("test" + "test")
fmt.Sprintf("%s %s", "test", "test")
strings.Join([]string{"test", "test"}, " ")
string.Builder

*/

func main() {
	s := [6]string{"one", "two", "3", "4", "five", "six"}

	str0 := joinStrDirect(s)
	log.Println(str0)

	str1 := joinStrWithSprintf(s)
	log.Println(str1)

	str2 := joinStrWithStringsJoin(s)
	log.Println(str2)

	str3 := joinStrWithBuilder(s)
	log.Println(str3)

}

func joinStrWithStringsJoin(s [6]string) (allStr string) {
	exectTime := time.Now()

	allStr = strings.Join(s[:], "")
	log.Printf("joinStrWithStringsJoin took: %s\n", time.Since(exectTime))

	return
}

func joinStrDirect(s [6]string) (allStr string) {
	exectTime := time.Now()

	for _, a := range s {
		allStr += a
	}

	log.Printf("joinStrDirect took: %s\n", time.Since(exectTime))

	return
}

func joinStrWithSprintf(s [6]string) (allStr string) {
	exectTime := time.Now()

	allStr = fmt.Sprintf("%s %s %s %s %s %s", s[0], s[1], s[2], s[3], s[4], s[5])

	log.Printf("joinStrWithSprintf took: %s\n", time.Since(exectTime))

	return
}
func joinStrWithBuilder(s [6]string) (allStr string) {
	exectTime := time.Now()

	var sb strings.Builder
	for _, a := range s {
		sb.WriteString(a)
	}

	log.Printf("joinStrWithBuilder took: %s\n", time.Since(exectTime))
	return sb.String()
}
