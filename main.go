package main

import (
	"fmt"

	"github.com/andresterba/config-example/config"
)

func main() {
	config := config.GetConfig()

	fmt.Printf("%+v\n", config.DB)
	fmt.Printf("%+v\n", config.Web)
}
