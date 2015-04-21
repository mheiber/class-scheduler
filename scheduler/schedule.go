package main

import (
	"bitbucket.org/maxheiber/coding-challenge/catalog"
	// "errors"
	// "encoding/json"
	"fmt"
	// "os"
)

type scheduler struct {
	*catalog.Catalog
	isHandled map[string]bool
	isPending map[string]bool
	schedule  []string
}

func (s *scheduler) ProcessCourseName(courseName string) error {

	course := s.GetCourse(courseName)

	if s.isHandled[courseName] {
		return nil
	}

	if s.isPending[courseName] {
		err := fmt.Errorf("Error: Cyclical dependency. Taking course \"%v\" requires first taking course \"%v\"\n", courseName, courseName)
		return err
	}

	s.isPending[course.Name] = true
	for _, prerequisite := range course.Prerequisites {
		prerequisiteCourse := s.GetCourse(prerequisite)
		if prerequisiteCourse == nil {
			return fmt.Errorf("Error: \"%v\" is listed as a prerequisite of \"%v\" but is not in the list of courses\n", prerequisite, course.Name)
		}

		err := s.ProcessCourseName(prerequisiteCourse.Name)
		if err != nil {
			return err
		}
	}

	s.isPending[courseName] = false

	s.schedule = append(s.schedule, courseName)

	s.isHandled[courseName] = true

	return nil
}

func FromCatalog(cat *catalog.Catalog) (schedule []string, err error) {

	courseNames := cat.CourseNames()
	length := len(courseNames)

	s := &scheduler{
		cat,
		make(map[string]bool, length),
		make(map[string]bool, length),
		make([]string, 0, length),
	}

	for _, courseName := range courseNames {
		err := s.ProcessCourseName(courseName)
		if err != nil {
			return nil, err
		}

	}

	return s.schedule, nil
}
