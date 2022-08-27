package main

import (
	"context"
	"fmt"
	"rss/drivers"
)

func main() {
	ctx := context.Background()
	newsDriver, err := drivers.InitializeNewsDriver(ctx)
	if err != nil {
		fmt.Printf("failed to create NewsDriver: %s\n", err)
	}

	newsDriver.Run(ctx)
}
