package catalog

import (
	"encoding/json"
	"errors"
)

type Course struct {
	Name          string
	Prerequisites []string
}

type Catalog struct {
	m map[string]*Course
}

func (cat *Catalog) GetCourse(name string) *Course {
	return cat.m[name]
}

func (cat *Catalog) CourseNames() []string {
	keys := make([]string, len(cat.m))
	i := 0
	for key := range cat.m {
		keys[i] = key
		i += 1
	}
	return keys
}

func New(courses []Course) *Catalog {
	cat := new(Catalog)
	cat.m = make(map[string]*Course)

	for i, course := range courses {
		cat.m[course.Name] = &courses[i]
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
