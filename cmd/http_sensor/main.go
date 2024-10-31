package main

import (
	"event_hub/internal/config"
	"fmt"
)

func main() {
	foo := config.GetEnv("FOO", nil)
	fmt.Println(foo)
}
