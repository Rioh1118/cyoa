package main

import (
	"fmt"
	"log"
	"os"
	//"github.com/Rioh1118/cyoa"
)

func main() {
	file, err := os.Open("../templates/format.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}

}
