package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var catalog Catalog

type Course struct {
	Name          string
	Prerequisites []string
	isPending     bool
	isHandled     bool
}

func load(filename string) []Course {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("Error: can't open %v\n", e)
	}
	var courses []Course
	var unmarshallingErr = json.Unmarshal(file, &courses)

	if unmarshallingErr != nil {
		fmt.Printf("Error: invalid course list. Each course must have a name and list of prerequisites")
	}

	return courses
}

type Catalog map[string]*Course

func toCatalog(courses []Course) Catalog {
	catalog := make(map[string]*Course)

	for i, course := range courses {
		catalog[course.Name] = &courses[i]
	}

	return catalog
}

func (catalog Catalog) handleCourseName(courseName string) {
	fmt.Printf("Handling: %v\n", courseName)
}

func (catalog Catalog) processCourseName(courseName string) {

	course := catalog[courseName]

	if course.isHandled {
		return
	}

	if course.isPending {
		panic("CYCLICAL" + course.Name)
	}

	course.isPending = true
	for _, prerequisite := range course.Prerequisites {
		prerequisiteCourseName := catalog[prerequisite].Name
		catalog.processCourseName(prerequisiteCourseName)
	}
	course.isPending = false

	catalog.handleCourseName(courseName)
	course.isHandled = true

}

func (catalog Catalog) order() {
	for _, course := range catalog {
		//fmt.Println(i, course.Name)
		catalog.processCourseName(course.Name)
	}
}

func main() {

	jsonFile := os.Args[1]
	courses := load(jsonFile)
	//fmt.Printf("Results: %v\n", courses)
	catalog = toCatalog(courses)

	//fmt.Println(catalog)
	catalog.order()
	//processCoursesData(coursesData)
}
