package builtin

import (
	"errors"
	"fmt"
	"io/ioutil"

	//"github.com/itang/gotang"
	"github.com/itang/gonew"
)

func init() {
	gonew.GetTemplateContainer().RegistTemplate("main", &MainTemplate{BuiltinTemplate{"main", "Main template", "e.g. gonew main some"}})
	gonew.GetTemplateContainer().RegistTemplate("package", &PackageTemplate{BuiltinTemplate{"package", "New package", "e.g. gonew package gotang"}})
	gonew.GetTemplateContainer().RegistTemplate("test", &TestTemplate{BuiltinTemplate{"test", "Test template", "e.g. gonew test gotang"}})
}

type BuiltinTemplate struct {
	name        string
	description string
	usage       string
}

func (this BuiltinTemplate) Name() string {
	return this.name
}

func (this BuiltinTemplate) Description() string {
	return this.description
}

func (this BuiltinTemplate) Usage() string {
	return this.usage
}

type MainTemplate struct {
	BuiltinTemplate
}

const mainSource = `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
}

`

func (this MainTemplate) Make(args []string) error {
	var file = "main.go"
	if len(args) > 0 {
		file = args[0] + ".go"
	}

	return ioutil.WriteFile(file, []byte(mainSource), 0666)
}

type PackageTemplate struct {
	BuiltinTemplate
}

const packageSource = `package %s

import (
	//"fmt"
)

//type SomeStruct struct {

//}

func Xxx() {

}

`

func (this PackageTemplate) Make(args []string) error {
	if len(args) == 0 {
		return errors.New("Package name can't be empty! " + this.Usage())
	}

	var (
		name = args[0]
		file = name + ".go"
	)

	return ioutil.WriteFile(file, []byte(fmt.Sprintf(packageSource, name)), 0666)
}

type TestTemplate struct {
	BuiltinTemplate
}

const testSource = `package %s

import (
	"testing"
)

func TestXXX(t *testing.T) {

}

`

func (this TestTemplate) Make(args []string) error {
	if len(args) == 0 {
		return errors.New("Test package name can't be empty! " + this.Usage())
	}

	var (
		name = args[0]
		file = name + "_test.go"
	)

	return ioutil.WriteFile(file, []byte(fmt.Sprintf(testSource, name)), 0666)
}
