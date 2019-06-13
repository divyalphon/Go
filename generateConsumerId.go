package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func generateRequestId() string {
	consumerRequestId, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	replaceByte := Bytereplace(consumerRequestId)
	aa := BytesToString(replaceByte)
	bb := strings.ToLower(aa)
	return bb

}

func Bytereplace(data []byte) []byte {
	return bytes.Replace(data, []byte("-"), []byte(""), -1)
}
func BytesToString(data []byte) string {
	return string(data[:])
}

func main() {
	request_id := generateRequestId()
	fmt.Println(request_id)
}
