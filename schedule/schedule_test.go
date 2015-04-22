package schedule_test

import (
	"bitbucket.org/maxheiber/coding-challenge/catalog"
	"bitbucket.org/maxheiber/coding-challenge/schedule"
	"github.com/mheiber/golang-utils/stringwriter"
	// "fmt"
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
