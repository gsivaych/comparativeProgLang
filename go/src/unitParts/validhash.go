package main

import (
  "fmt"
)

func main(){
  hash := "ae3h4bdy5ndg6j0000"
  difficulty := 3
  last_d_bytes := string(hash[len(hash)-difficulty:])
	count := 0;
	for i := 1 ; i <= len(last_d_bytes) ; i++ {
		if last_d_bytes[i-1] == 48 {
			count++
		}
		}
	if count == difficulty{
		fmt.Printf("true\n")
		} else {
			fmt.Printf("false\n")
		}
}
