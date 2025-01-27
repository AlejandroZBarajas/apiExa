package main

import (
	"ArquitecturaExagonal/src/infrastructure"
	"fmt"
)

func main() {
	fmt.Println("hola")
	infrastructure.ConnectDB()
}
