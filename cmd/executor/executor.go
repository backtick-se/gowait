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
	ctx := context.Background()
	err = executor.Run(ctx)
	if err != nil {
		fmt.Println("execution failed:", err)
		os.Exit(1)
	}
}
