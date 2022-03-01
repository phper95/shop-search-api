package main

import (
	"fmt"
	"shop-search-api/config"
)

func init() {
	config.LoadConfig()
}
func main() {

	fmt.Println(config.Cfg)

}
