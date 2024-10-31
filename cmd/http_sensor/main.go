package main

import (
	"fmt"
	"github.com/tiko-star/event_hub/internal/config"
)

func main() {
	foo := config.GetEnv("FOO", nil)
	fmt.Println(foo)
}
