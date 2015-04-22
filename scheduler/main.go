package main

import (
	"bitbucket.org/maxheiber/coding-challenge/catalog"
	"bitbucket.org/maxheiber/coding-challenge/schedule"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "Error: Expected the name of a JSON file as the first argument to scheduler\n")
		os.Exit(2)
	}
	jsonFile := os.Args[1]

	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Couldn't open file %v\n", err)
		os.Exit(2)
	}

	cat := new(catalog.Catalog)
	err = catalog.UnmarshalJSON(cat, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v\n", err)
		os.Exit(2)
	}

	//This is the most important part. See schedule.go
	err = schedule.Generate(os.Stdout, cat)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}

}
