package main

import (
	"flag"
	"fmt"

	"github.com/mei-rune/properties"
)

func main() {
	var updated string
	flag.StringVar(&updated, "variables", "", "")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		return
	}

	if updated == "" {
		fmt.Println("updated is empty")
		return
	}

	variables, err := properties.ReadProperties(updated)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, arg := range args {
		err := properties.UpdateProperties(arg, variables)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
