package pkg

import (
	"fmt"
	"module-test/summer"
)

func init() {
	fmt.Println("pkg init 1")
}
func init() {
	fmt.Println("pkg init 2", i)
}
func Pkg() {
	fmt.Println("pkg")
}

var i = summer.Io
