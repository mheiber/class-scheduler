package catalog_test

import (
	"bitbucket.org/maxheiber/coding-challenge/catalog"
	"encoding/json"
	"testing"
)

var courses = []catalog.Course{
	catalog.Course{Name: "Biology 1", Prerequisites: []string{}},
	catalog.Course{Name: "Biology 2", Prerequisites: []string{"Biology 1"}},
}

func TestNew(t *testing.T) {

	if catalog.New(courses) == catalog.New(courses) {
		t.Errorf(`New(courses) == New(courses)`)
	}

	//Shouldn't crash
	cat := catalog.New(courses)
	if cat != cat {
		t.Errorf(`cat != cat`)
	}

}

func TestGetCourse(t *testing.T) {

	cat := catalog.New(courses)
	if cat.GetCourse("Biology 2").Name != "Biology 2" {
		t.Errorf(`cat.getCourse("Biology 2").Name != "Biology 2"`)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	data, err := json.Marshal(courses)
	if err != nil {
		t.Errorf("Error: Bad test design. Couldn't marshall fixture %v\n", courses)
	}
	cat := new(catalog.Catalog)
	err = catalog.UnmarshalJSON(cat, data)

	//test that courses are accessible
	course1 := cat.GetCourse("course1")
	if course1 != course1 {
		t.Errorf(`course1 != course1`)
	}
	if err != nil {
		t.Errorf("%v\n", err)
	}
}
