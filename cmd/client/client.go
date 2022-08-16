package main

import (
	"cowait/client"
	"cowait/core"

	"fmt"
	"os"
)

// container client / process manager

func main() {
	// create executor
	envdef := os.Getenv(core.EnvTaskdef)
	executor, err := client.ExecutorFromEnv(envdef)
	if err != nil {
		fmt.Println("failed to create executor:", err)
		os.Exit(1)
	}

	// run task
	err = executor.Run()
	if err != nil {
		fmt.Println("execution failed:", err)
		os.Exit(1)
	}
}
