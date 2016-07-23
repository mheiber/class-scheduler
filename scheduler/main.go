package main

import (
	"bitbucket.org/maxheiber/coding-challenge/course"
	"bitbucket.org/maxheiber/coding-challenge/schedule"
	"encoding/json"
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
	courses := getCourses(jsonData)

	//Prints courses to Stdout, not taking a course
	//until prerequisites are satisfied
	err := schedule.Generate(os.Stdout, courses)
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

func getCourses(jsonData []byte) (courses []course.Course) {
	courses = make([]course.Course, 1)
	err := json.Unmarshal(jsonData, &courses)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: Invalid course list. Each course must have a name and an array of prerequisites")
	}

	return
}
