package main

import "fmt"

var dataCache map[string]string

func main() {
	// data := [4]byte{65, 66, 67, 68}
	// str := string(data[:])
	// fmt.Print(str)

	for key, _ := range dataCache {
		fmt.Println(key)
	}

}
