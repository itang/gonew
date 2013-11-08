package main

import (
	"fmt"
	"os"

	"github.com/itang/gonew"
	_ "github.com/itang/gonew/templates/builtin"
)

func main() {
	tname, args := getTemplateNameFromArgs()
	tmp := gonew.GetTemplateContainer().GetTemplate(tname)

	if tmp == nil {
		fmt.Println("gonew: Generate scaffolding for a new project based on a template\n")
		if tname != "" {
			fmt.Printf("ERROR: template:'%s' 现在暂不支持.\n\n", tname)
		}
		printTemplateList()
		return
	}

	fmt.Printf("Generating a scaffolding based on the '%v' template.\n", tmp.Name())
	err := tmp.Make(args)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
}

func printTemplateList() {
	fmt.Println("Several templates are available:")
	for i, v := range gonew.GetTemplateContainer().GetTemplates() {
		fmt.Printf("%v\t%s\t%s\n", i, v.Name(), v.Description())
	}
	fmt.Println("\ne.g. gonew main xxx\n")
}

func getTemplateNameFromArgs() (cmd string, args []string) {
	as := os.Args[1:]

	if len(as) > 0 {
		cmd = as[0]
	}
	if len(as) > 1 {
		args = as[1:]
	}

	return
}
