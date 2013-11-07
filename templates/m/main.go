package m

import (
	//"github.com/itang/gotang"
	"io/ioutil"
)

const s = `package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello")
}

`

func MakeWithMainTemplate() {
	ioutil.WriteFile("main.go", []byte(s), 0x666)
}
