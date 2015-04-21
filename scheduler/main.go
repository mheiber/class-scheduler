package main

import (
	"bitbucket.org/maxheiber/coding-challenge/catalog"
	//"errors"
	// "encoding/json"
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

	data, _ := ioutil.ReadFile(jsonFile)

	cat := new(catalog.Catalog)
	_ = catalog.UnmarshalJSON(cat, data)

	schedule(cat)

}
