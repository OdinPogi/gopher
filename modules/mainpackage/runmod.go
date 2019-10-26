package main

import (
	"fmt"

	"github.com/OdinPogi/gopher/modules/libpackage"
)

const Name = libpackage.Name

func main() {
	fmt.Println("")
	fmt.Println("  Congratulations. You started successfully !")
	fmt.Println("  You'll see, Go is simple and really slick.")
	fmt.Println("")
	fmt.Println("  Oh no! a typo -> TEST22")
	fmt.Println("")
	fmt.Println("  Correct it and re-run this program.")
	fmt.Println("")
	fmt.Println("  You can modify this program by tweaking `main.go`")
	fmt.Println("")
	fmt.Println("Name value is %s", Name)
	fmt.Println("")
}
