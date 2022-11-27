package main

import (
	"fmt"

	"github.com/govnokode/Testers/unit"
)

func main() {
	t := unit.TestStr{"shuhrat"}
	fmt.Println("Hello, ", t.Name)
}
