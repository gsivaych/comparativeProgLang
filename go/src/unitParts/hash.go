package main

import (
  "fmt"
  "strconv"
  "strings"
  "encoding/hex"
  "crypto/sha256"
)

func makeHash() []byte {
  var x []byte = []byte{0x00}
	HashString := strings.Join([]string{
		hex.EncodeToString(x),
		strconv.Itoa(0),
		strconv.Itoa(2),
		"",
		"242278",
		}, ":")
    fmt.Println(HashString)
    h := sha256.New()
    h.Write([]byte(HashString))
    return h.Sum(nil)
}

func main(){
  fmt.Printf(hex.EncodeToString(makeHash()))
}
