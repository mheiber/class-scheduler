package main

import (
	"bitbucket.org/maxheiber/coding-challenge/catalog"
	//"errors"
	// "encoding/json"
	"fmt"
)

var isHandled = make(map[string]bool)

var isPending = make(map[string]bool)

func handleCourseName(courseName string) {
	fmt.Printf("Handling: %v\n", courseName)
}

func processCourseName(cat *catalog.Catalog, courseName string) {

	course := cat.GetCourse(courseName)

	if isHandled[courseName] {
		return
	}

	if isPending[courseName] {
		panic("CYCLICAL" + course.Name)
	}

	isPending[course.Name] = true
	for _, prerequisite := range course.Prerequisites {
		prerequisiteCourseName := cat.GetCourse(prerequisite).Name
		processCourseName(cat, prerequisiteCourseName)
	}
	isPending[courseName] = false

	handleCourseName(courseName)
	isHandled[courseName] = true

}

func Schedule(cat *catalog.Catalog) {
	for _, courseName := range cat.CourseNames() {
		processCourseName(cat, courseName)
	}
}
