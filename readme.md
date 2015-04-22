# Solution to Clever Coding Challenge - Class Scheduler #

The program reads a JSON file with course names and prerequisites and prints a course schedule that satisfies these conditions:

- All courses are taken
- No course is taken until all its prerequisites are satisfied.



## How to Use ##

On the command line, navigate to the folder that contains this readme, then navigate to the `dist` subdirectory.

Run scheduler on a json file, such as the physics.json example file in the `examples` folder:

`./scheduler ../examples/physics.json` 

A course schedule will print to stdout. For example running the command above will print something close to the following:

```
Calculus
Scientific Thinking
Differential Equations
Intro to Physics
Relativity
```

> Note: There can be multiple valid schedules for a given list of courses.

The format of the JSON files must be:


- An array of objects
- Each object has at least these keys: 
    - name (type string) 
    - prerequisites (array of strings)


### Testing ###

1. Make sure you have [Go installed](https://golang.org/doc/install)

2. Place the source files in your workspace so they [follow the conventions outlined in the Go docs](https://golang.org/doc/code.html#Workspaces).

3. You can run the `go test` command on the command line in the `src/schedule` to unit test the functionality for creating schedules. Running `go test` in `src/catalog` will test the program's ability to read in a JSON file with a list of courses, in addition to testing some of the abstractions used in the program. 

4. Run `go test` in `src/scheduler` to kick off a simple integration test.


> For information about benchmarking, see __


### Design Explanation ###

#### How it Works ###

1. The `catalog` library parses the JSON file with information about courses. It creates a mapping from course names to courses and stores the prerequisites for each course.

2. The `schedule` library prints course names to stdout in an order that ensures no course is printed before any of its prerequisites. It follows an algorithm similar to that used in Facebook's [Flux Dispatcher]( ... ). 

Start with any course
If all the course's prerequisites are satisfied, print the course name and mark the course as satisfied.
Otherwise, for each of the courses unsatisfied prerequisites, repeat the procedure for that prerequisite.
Repeat the procedure for the next course until no courses remain.

#### Choices ###

Since the code will be assessed partially based on correctness and speed, I needed to use a fast language. I also needed a language conducive to good code style and architecture. For those reasons, [Go](golang.org) is a good fit. This required me to learn how to write correct, clear, and idiomatic Go.

A challenge was determining the right level of abstraction to use. I could have skipped creating the `catalog.Catalog` struct type and instead just used a map from strings to `catalog.Course`s. This would have spared the need to implement the `catalog#CourseNames` and `catalog#GetCourse` getters. It was a close judgment call, but I used a custom struct type because it enabled me to implement the `json.Unmarshaler` interface and have clear and idiomatic code for reading in JSON files (see `src/scheduler/main.go`).

### Performance ###

The worst-case inputs, performance wise, have the maximum number of prerequisites per course. This makes them "triangular":

Course{name: course1, }
