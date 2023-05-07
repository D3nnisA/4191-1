// Filename: cmd/api/courses.go

package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/D3nnisA/4191-1/internal/data"
)

func (app *application) createCoursesHandler(w http.ResponseWriter, r *http.Request) {

	//Create a struct to hold a course that will be provided to us via the request
	var input struct {
		//ID           int64  `json:"id"`
		CourseCode   string `json:"Course Code"`
		CourseTitle  string `json:"Course Title"`
		CourseCredit string `json:"Course Credit"`
		//Version      int32  `json:"version"`
	}

	//decode our JSON request

	err := app.readJSON(w, r, &input)
	if err != nil {

		app.badRequestResponse(w, r, err)
		return
	}

	//Copy the values from the input struct to a new School struct
	course := &data.Courses{
		//ID:           input.ID,
		CourseCode:   input.CourseCode,
		CourseTitle:  input.CourseTitle,
		CourseCredit: input.CourseCredit,
	}

	// //write our validated school to database

	// err = app.models.Courses.Insert(course)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }

	// //set the creation header
	// headers := make(http.Header)
	// headers.Set("Location", fmt.Sprintf("/v1/courses/%d", course.ID))

	//write the response

	// err = app.writeJSON(w, http.StatusCreated, envelope{"course": course}, headers)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// }

	// Create a School
	err = app.models.Courses.Insert(course)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// Create a Location header for the newly created resource/course
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/courses/%d", course.ID))

	// Write the JSON response with 201 - Created status code with the body
	// being the School data and the header being the headers map
	err = app.writeJSON(w, http.StatusCreated, envelope{"course": course}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showCoursesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// Fetch the specific school
	course, err := app.models.Courses.Get(id)
	// Handle errors
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Write the data returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {

		app.serverErrorResponse(w, r, err)

		return

	}
}
