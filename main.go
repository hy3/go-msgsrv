package main

import (
	"flag"
	"fmt"
	"net/http"
)

const defaultPort = "80"

type arguments struct {
	port string
}

func main() {
	args := fetchArgs()
	handler := SetupHandler()
	fmt.Println("Start to listen port ", args.port, "...")
	if err := http.ListenAndServe(":"+args.port, handler); err != nil {
		fmt.Println(err)
	}
}

func fetchArgs() *arguments {
	args := new(arguments)
	flag.StringVar(&args.port, "p", defaultPort, "port for http listen.")
	flag.StringVar(&args.port, "port", defaultPort, "port for http listen.")
	flag.Parse()
	return args
}
