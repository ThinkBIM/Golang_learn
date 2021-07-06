package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("%s", string(cmd))
}
