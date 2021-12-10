package main

import (
	"flag"
	"fmt"
	"gosapher/cmd/cli"
)

func main() {
	operation := flag.String("op", "", "--op=encode or decode")
	str := flag.String("str", "", "--str=\"any secret value for encode or decode\"")

	flag.Parse()

	if *operation == ""{
		fmt.Println("Use --help")
		return
	}

	if *str == ""{
		fmt.Println("Use --help")
		return
	}

	if *operation == "encode"{
		fmt.Println(cli.Encode(*str))
	}

	if *operation == "decode"{
		fmt.Println(cli.Decode(*str))
	}

}
