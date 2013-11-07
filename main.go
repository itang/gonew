package main

import (
	"flag"

  "github.com/itang/gonew/templates/m"
)

var (
	tmp = flag.String("t", "m", "The template")
)

func main() {
  flag.Parse()

  switch(*tmp) {
  case "m":
    m.MakeWithMainTemplate()
  }
}
