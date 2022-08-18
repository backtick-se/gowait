package main

import (
	"cowait/core"
	"cowait/executor"

	"context"
	"fmt"
	"os"
)

// container client / process manager

func main() {
	// create executor
	envdef := os.Getenv(core.EnvTaskdef)
	executor, err := executor.NewFromEnv(envdef)
	if err != nil {
		fmt.Println("failed to create executor:", err)
		os.Exit(1)
	}

	// run task
	err = executor.Run(context.Background())
	if err != nil {
		fmt.Println("execution failed:", err)
		os.Exit(1)
	}
}
