package schedule_test

import (
	"bitbucket.org/maxheiber/coding-challenge/catalog"
	"bitbucket.org/maxheiber/coding-challenge/schedule"
	"bytes"
	// "fmt"
	"github.com/mheiber/golang-utils/stringwriter"
	"io/ioutil"
	"strings"
	"testing"
)

type testCase struct {
	jsonFile string
	errMsg   string
}

func TestCyclicalDependency(t *testing.T) {
	tCase := &testCase{
		"./fixtures/cyclical.json",
		"Error: Cyclical dependency.",
	}
	tCase.run(t)
}

func TestMissingPrerequsite(t *testing.T) {
	tCase := &testCase{
		"./fixtures/missing.json",
		"is not in the list of courses",
	}
	tCase.run(t)
}

func TestResultsValid(t *testing.T) {

	fixtures := []string{"physics.json", "math.json", "physics2.json"}

	for _, fixture := range fixtures {
		//sorry Windows
		tCase := &testCase{"./fixtures/" + fixture, ""}
		tCase.run(t)
	}

}

func setup(jsonFile string) (*catalog.Catalog, error) {
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	cat := new(catalog.Catalog)
	err = catalog.UnmarshalJSON(cat, data)
	return cat, err
}

func (tCase *testCase) run(t *testing.T) {
	cat, err := setup(tCase.jsonFile)
	// fmt.Println(cat)
	//Test that appropriate err is returned if err expected
	if err != nil && err.Error() != tCase.errMsg {
		t.Errorf("Expected %v to equal %v", err, tCase.errMsg)
	}
	//Writing to a buf so we can get the course schedule as a []string
	sw := new(stringwriter.StringW)
	err = schedule.Generate(sw, cat)
	if err != nil && !strings.Contains(err.Error(), tCase.errMsg) {
		t.Errorf("Expected %v to contain %v", err, tCase.errMsg)
	}

	//list of course names from scheduler
	var results []string
	results = sw.Val()

	//See if the schedule is valid
	//Don't bother checking if we've already hit an err
	if err == nil {
		testValidLength(cat, results, t)
		testPrereqsSatisfied(cat, results, t)
	}
}

func testValidLength(cat *catalog.Catalog, results []string, t *testing.T) {
	if len(results) != len(cat.CourseNames()) {
		t.Errorf("Expected course schedule length to equal number of courses: \n%v%v\n%v%v", results, len(results), cat.CourseNames(), len(cat.CourseNames()))
	}
}

func testPrereqsSatisfied(cat *catalog.Catalog, results []string, t *testing.T) {

	isTaken := make(map[string]bool)

	//uncomment to test this test
	//results[0], results[len(results)-1] = results[len(results)-1], results[0]

	for _, cname := range results {
		isTaken[cname] = true
		prereqs := cat.GetCourse(cname).Prerequisites
		for _, prereqName := range prereqs {
			if !isTaken[prereqName] {
				t.Errorf("Took course \"%v\" before prereq \"%v\"", cname, prereqName)
			}
		}
	}
}

var testCat1000 = genTestCatalog(1000)

func BenchmarkSchedule1000(b *testing.B) {
	cat := testCat1000
	buf := new(bytes.Buffer)
	schedule.Generate(buf, cat)
}

var testCat2000 = genTestCatalog(2000)

func BenchmarkSchedule2000(b *testing.B) {
	cat := testCat2000
	buf := new(bytes.Buffer)
	schedule.Generate(buf, cat)
}

var testCat4000 = genTestCatalog(4000)

func BenchmarkSchedule4000(b *testing.B) {
	cat := testCat4000
	buf := new(bytes.Buffer)
	schedule.Generate(buf, cat)
}

var testCat8000 = genTestCatalog(8000)

func BenchmarkSchedule8000(b *testing.B) {
	cat := testCat8000
	buf := new(bytes.Buffer)
	schedule.Generate(buf, cat)
}

var testCat16000 = genTestCatalog(16000)

func BenchmarkSchedule16000(b *testing.B) {
	cat := testCat16000
	buf := new(bytes.Buffer)
	schedule.Generate(buf, cat)
}

func names(courses []catalog.Course) []string {
	names := make([]string, 0, len(courses))
	for _, crse := range courses {
		names = append(names, crse.Name)
	}
	return names
}

func genTestCatalog(courseCount int) *catalog.Catalog {
	courses := make([]catalog.Course, 0, courseCount)

	for i := courseCount - 1; i >= 0; i-- {
		crse := catalog.Course{Name: string(i), Prerequisites: names(courses)}
		courses = append(courses, crse)
	}

	return catalog.New(courses)
}
