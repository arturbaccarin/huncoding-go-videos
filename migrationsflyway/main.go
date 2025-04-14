package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Starting migrations")

	cmd := exec.Command("flyway", "migrate", "-configFile=flyway.conf")
	resp, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	fmt.Println(string(resp))
	fmt.Println("Finished migrations")

}
