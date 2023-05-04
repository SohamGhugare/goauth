package main

import (
	"fmt"
	"goauth/initializers"
)

func init() {
	initializers.LoadEnvVars()
}

func main() {
	fmt.Println("Hello World")

}
