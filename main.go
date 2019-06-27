package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func randomPassword(length int) (pwd string, err error) {
	numBytes := length * 3 / 4
	if numBytes == 0 {
		numBytes = 1
	}
	buffer := make([]byte, numBytes)
	_, err = io.ReadFull(rand.Reader, buffer)
	if err != nil {
		fmt.Println("problem geberating password")
		return "", err
	}
	pwd = base64.StdEncoding.EncodeToString(buffer)
	return pwd, nil

}

func main() {
	randPwd, err := randomPassword(24)
	if err == nil {
		fmt.Println(randPwd)
	}
}
