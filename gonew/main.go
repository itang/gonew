package main

import (
	"fmt"
	"os"

	"github.com/itang/gonew"
	_ "github.com/itang/gonew/templates/m"
)

func main() {
	cmd, args := getCmdFromArgs()
	fmt.Println(cmd, args)

	if cmd == "" {
		cmd = "main"
	}

	tmp := gonew.GetTemplateContainer().GetTemplate(cmd)

	if tmp != nil {
		fmt.Println(tmp.Name(), tmp.Description())
		tmp.Make(args)
	} else {
		fmt.Printf("cmd:%s 现在暂不支持.\n", cmd)
	}
}

func getCmdFromArgs() (cmd string, args []string) {
	as := os.Args[1:]

	if len(as) > 0 {
		cmd = as[0]
	}
	if len(as) > 1 {
		args = as[1:]
	}
	return
}
