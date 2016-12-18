# Class Scheduler #

The program reads a JSON file with course names and prerequisites and prints a course schedule that satisfies these conditions:

- All courses are taken
- No course is taken until all its prerequisites are satisfied.



## How to Use ##

On a Linux machine, on the command line, navigate to the folder that contains this readme, then navigate to the `dist` subdirectory.

> All paths in this documentation are relative to  [$GOPATH](https://golang.org/doc/code.html#Workspaces)/src/bitbucket.org/maxheiber/coding-challenge/` unless otherwsise noted.

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

2. Place the source files in your workspace so they [follow the conventions outlined in the Go docs](https://golang.org/doc/code.html#Workspaces). The path within your [$GOPATH](https://golang.org/doc/code.html#Workspaces) must be `src/bitbucket.org/maxheiber/coding-challenge/`.

3. You can run the `go test` command on the command line in the `src/schedule` to unit test the functionality for creating schedules. Running `go test` in `src/catalog` will test the program's ability to read in a JSON file with a list of courses, in addition to testing some of the abstractions used in the program. 

> testing the schedule module requires the github.com/mheiber/golang-utils/stringwriter lib.

4. Run `go test` in the `scheduler` directory to kick off a simple integration test.


> For information about benchmarking, see [Performance]("#performance")


### Design Explanation ###

#### How it Works ###

1. The `catalog` library parses the JSON file with information about courses. It creates a mapping from course names to courses and stores the prerequisites for each course.

2. The `schedule` library prints course names to stdout in an order that ensures no course is printed before any of its prerequisites. It follows an algorithm similar to that used in Facebook's [Flux](http://facebook.github.io/flux/) [Dispatcher](https://github.com/facebook/flux/blob/master/src/Dispatcher.js#L22). 

Start with any course
If all the course's prerequisites are satisfied, print the course name and mark the course as satisfied.
Otherwise, for each of the courses unsatisfied prerequisites, repeat the procedure for that prerequisite.
Repeat the procedure for the next course until no courses remain.

### Performance ###

**Terrible** (On^2). This is a topological sorting problem, for which there are [several simple linear time algorithms](https://en.wikipedia.org/wiki/Topological_sorting#Algorithms) that I should have used instead.

The worst-case inputs, performance wise, have the maximum number of prerequisites per course. This makes them "triangular":

{"course 4": "prerequisites": ["course 4", "course 3", "course 2", "course 1"]} 
{"course 3": "prerequisites": ["course 3", "course 2", "course 1"]}
{"course 2": "prerequisites": ["course 1"]}
{"course 1": "prerequisites": []}


Where the number of courses is `n`, the number of steps to produce a list of courses in order that respects prerequisites is roughly `(n^2 + n)/2`. That means that worst-case growth in running time is *(O(n^2))*. This gibes with what I saw in my benchmarks: number of calls to `schedule.Scheduler#ProcessCourseName` fits a quadratic regression line. See `documentation-assets/performance-plog.png`.

You can run the benchmarks from the `schedule` directory by running `go test -bench=.`

Also, if I were to do this again I'd use [streaming JSON parsing](https://golang.org/pkg/encoding/json/#example_Decoder_Decode_stream), which I'm not sure was a thing in the Go standard library when I wrote this.
