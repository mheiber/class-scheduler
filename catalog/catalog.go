package catalog

import (
	"encoding/json"
	"errors"
)

type Course struct {
	Name          string
	Prerequisites []string
}

type Catalog map[string]*Course

//Doesn't return a pointer since maps are reference types
func New(courses []Course) *Catalog {
	cat := make(Catalog)

	for i, course := range courses {
		cat[course.Name] = &courses[i]
	}

	return &cat
}

func UnmarshalJSON(cat *Catalog, data []byte) error {

	var courses []Course
	err := json.Unmarshal(data, &courses)
	if err != nil {
		return errors.New("Error: Invalid course list. Each course must have a name and list of prerequisites")
	}

	cat = New(courses)

	return nil
}
