package utils

import (
	"fmt"

	"github.com/gorilla/securecookie"
)

func Test() {

	fmt.Println(securecookie.GenerateRandomKey(32))
	fmt.Println(securecookie.GenerateRandomKey(32))
}
