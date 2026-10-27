package main

import (
	"fmt"

	"github.com/Kipow/myproject/mathutils"
)

func main() {
	a := 20
	b := 10

	fmt.Println("Қосу:", mathutils.Add(a, b))
	fmt.Println("Азайту:", mathutils.Subtract(a, b))
}
