package main

import (
	"fmt"
	"github.com/tiko-star/eh/internal/config"
)

func main() {
	bar := config.GetEnv("BAR", nil)
	fmt.Println(bar)
}
