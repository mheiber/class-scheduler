package schedule

import (
	"bitbucket.org/maxheiber/coding-challenge/course"
	"fmt"
	"io"
)

// schedule.Generate takes an array of courses and writes
// a list of courses in order so that all of the
// prerequistes will be satisfied

type scheduler struct {
	w             io.Writer
	coursesByName map[string]*course.Course
	isHandled     map[string]bool
	isPending     map[string]bool
}

func (s *scheduler) writeln(str string) {
	line := []byte(str + "\n")
	s.w.Write(line)
}

func (s *scheduler) ProcessCourse(course *course.Course) error {
	if s.isHandled[course.Name] {
		return nil
	}

	if s.isPending[course.Name] {
		err := fmt.Errorf("Error: Cyclical dependency. Taking course \"%v\" requires first taking course \"%v\"\n", course.Name, course.Name)
		return err
	}
	s.isPending[course.Name] = true
	for _, prerequisite := range course.Prerequisites {
		prerequisiteCourse := s.coursesByName[prerequisite]
		if prerequisiteCourse == nil {
			return fmt.Errorf("Error: \"%v\" is listed as a prerequisite of \"%v\" but is not in the list of courses\n", prerequisite, course.Name)
		}
		err := s.ProcessCourse(prerequisiteCourse)
		if err != nil {
			return err
		}
	}

	s.isPending[course.Name] = false

	s.writeln(course.Name)

	s.isHandled[course.Name] = true

	return nil
}

func Generate(w io.Writer, courses []course.Course) error {
	length := 0

	coursesByName := make(map[string]*course.Course)

	for index, course := range courses {
		length = index
		coursesByName[course.Name] = &courses[index]
	}

	s := &scheduler{
		w,
		coursesByName,
		make(map[string]bool, length), //record of which courses have been taken
		make(map[string]bool, length), //record of current course for which we're satisfying prereqs
	}

	for _, course := range courses {
		err := s.ProcessCourse(&course)
		if err != nil {
			return err
		}

	}

	return nil
}
