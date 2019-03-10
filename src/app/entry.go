package main

import (
	"fmt"
	"greet"
	_ "greet"
	"greet/de"
	child "greet/greet"
	// "github.com/jinzhu/gorm"
)

var (
	a int = b
	b int = f()
	c int = 1
)

func init() {
	fmt.Println("app/entry.go ==> init() [1]")
}

func init() {
	fmt.Println("app/entry.go ==> init() [2]")
}

func main() {
	fmt.Println("app/entry.go ==> main()")
	fmt.Println(greet.Morning)
	fmt.Println(de.Morning)
	fmt.Println("version ===> ", version)
	fmt.Println(a, b, c)

	fmt.Println(child.Message)

	// fmt.Println(gorm.SomeExportedMember)
}
