package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.Debug()
	Print()
}

func Print() {
	fmt.Println("Hi hello")
}
