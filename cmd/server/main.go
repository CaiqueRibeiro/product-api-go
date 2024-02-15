package main

import (
	"fmt"

	"github.com/CaiqueRibeiro/product-api/configs"
)

func main() {
	configs, err := configs.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	fmt.Println(configs)
}
