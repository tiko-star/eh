package main

import (
	"fmt"
	"github.com/tiko-star/event_hub/internal/config"
)

func main() {
	bar := config.GetEnv("BAR", nil)
	fmt.Println(bar)
}
