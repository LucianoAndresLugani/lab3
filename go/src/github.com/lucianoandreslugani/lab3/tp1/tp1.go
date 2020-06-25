package main

import (
	"fmt"

	"github.com/lucianoandreslugani/lab3/tp2"
)

var desv float32

func main() {

	fmt.Println("ingrese desviacion deseada")
	fmt.Scan(&desv)
	tp2.Inicio(desv)
}
