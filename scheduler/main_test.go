package main

import (
	"os"
)

func ExampleMain() {
	//testing that the scheduler will run correctly given a
	//file with only one course. See the schedule and catalog
	//libs for comprehensive test suites
	os.Args = append(os.Args, "./fixtures/one-course.json")
	main()
	// Output:
	// English
}
