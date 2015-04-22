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

	jsonData := readFile(jsonFile)

	//a catalog maps course names to courses
	cat := getCatalog(jsonData)

	//Prints courses to Stdout, not taking a course
	//until prerequisites are satisfied
	err := schedule.Generate(os.Stdout, cat)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(2)
	}
}

func readFile(jsonFile string) []byte {
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Couldn't open file %v\n", err)
		os.Exit(2)
	}
	return data
}

func getCatalog(jsonData []byte) *catalog.Catalog {
	cat := new(catalog.Catalog)
	err := catalog.UnmarshalJSON(cat, jsonData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%v\n", err)
		os.Exit(2)
	}

	return cat
}
