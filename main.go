package main

import (
	"fmt"

	"github.com/chayutK/thaiid/thaiid"
)

func main() {
	id := "1234567890121"
	fmt.Println(thaiid.IsValid(id))

}
