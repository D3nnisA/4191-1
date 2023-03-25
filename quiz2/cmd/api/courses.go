// Filename: cmd/api/courses.go

package main

import (
	"fmt"
	"net/http"

	"github.com/D3nnisA/4191-1/internal/data"
)

func (app *application) createCoursesHandler(w http.ResponseWriter, r *http.Request) {

	//Create a struct to hold a school that will be provided to us via the request
	var input struct {
		CourseCode   string `json:"Course Code"`
		CourseTitle  string `json:"Course Title"`
		CourseCredit int64  `json:"Course Credit"`
	}

	//decode our JSON request

	err := app.readJSON(w, r, &input)
	if err != nil {

		app.badRequestResponse(w, r, err)
		return
	}

	// Print the request

	fmt.Fprintf(w, "%+v\n", input)

}

func (app *application) showCoursesHandler(w http.ResponseWriter, r *http.Request) {
	version, err := app.readIDParams(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//fmt.Fprintf(w, "show details of courses %d\n", id)

	course := data.Courses{

		CourseCode:   "MATH2134",
		CourseTitle:  "Algebra",
		CourseCredit: "3",
		Version:      int32(version),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {

		app.notFoundResponse(w, r)

		return

	}
}
