package main

import (
	"event_hub/internal/config"
	"fmt"
)

func main() {
	bar := config.GetEnv("BAR", nil)
	fmt.Println(bar)
}
