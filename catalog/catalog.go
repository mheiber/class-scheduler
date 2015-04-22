package catalog

// A course consists of a name and a list of prerequisites.
// A catalog is a mapping from course names to courses.
// A catalog can be unmarshalled from a json array of courses

import (
	"encoding/json"
	"errors"
)

type Course struct {
	Name          string
	Prerequisites []string
}

type Catalog struct {
	m           map[string]*Course
	courseNames []string
}

func (cat *Catalog) GetCourse(name string) *Course {
	return cat.m[name]
}

func (cat *Catalog) CourseNames() []string {
	return cat.courseNames
}

func New(courses []Course) *Catalog {
	cat := new(Catalog)
	cat.m = make(map[string]*Course)
	cat.courseNames = make([]string, 0, len(courses))
	for i, course := range courses {
		cat.m[course.Name] = &courses[i]
		cat.courseNames = append(cat.courseNames, course.Name)
	}

	return cat
}

func UnmarshalJSON(cat *Catalog, data []byte) error {
	courses := make([]Course, 1)
	err := json.Unmarshal(data, &courses)
	if err != nil {
		return errors.New("Error: Invalid course list. Each course must have a name and an array of prerequisites")
	}
	//pointer that is passed in will point to the new Catalog
	*cat = *New(courses)

	return nil
}
