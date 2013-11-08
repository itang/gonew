package m

import (
	"io/ioutil"

	//"github.com/itang/gotang"
	"github.com/itang/gonew"
)

const s = `package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello")
}

`

func init() {
	gonew.GetTemplateContainer().RegistTemplate("main",&MainTemplate{"main", "main template"})
}

type MainTemplate struct {
	name        string
	description string
}

func (this MainTemplate) Make(args []string) error {
	var name = "main.go"
	if len(args) > 0 {
		name = args[0] + ".go"
	}
	ioutil.WriteFile(name, []byte(s), 0666)
  return nil
}

func (this MainTemplate) Name() string{
  return this.name
}

func (this MainTemplate) Description() string{
  return this.description
}

