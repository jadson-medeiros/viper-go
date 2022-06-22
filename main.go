package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// This is usually set outside the program
	os.Setenv("PREFIX_VAR", "hello")
	os.Setenv("PREFIX_VAR2", "hello2")

	viper.SetEnvPrefix("prefix")
	viper.AutomaticEnv()

	fmt.Println(viper.Get("var"))
	fmt.Println(viper.Get("var2"))

	if !viper.IsSet("home") {
		fmt.Println("The key 'home' is not set")
	}
}
