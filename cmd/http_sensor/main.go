package main

import (
	"fmt"
	"github.com/tiko-star/eh/internal/config"
)

func main() {
	foo := config.GetEnv("FOO", nil)
	fmt.Println(foo)
}
