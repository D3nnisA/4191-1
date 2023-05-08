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
		CourseCode   string `json:"Course Code"`
		CourseTitle  string `json:"Course Title"`
		CourseCredit string `json:"Course Credit"`
	}

	//decode our JSON request

	err := app.readJSON(w, r, &input)
	if err != nil {

		app.badRequestResponse(w, r, err)
		return
	}

	//Copy the values from the input struct to a new course struct
	course := &data.Courses{
		//ID:           input.ID,
		CourseCode:   input.CourseCode,
		CourseTitle:  input.CourseTitle,
		CourseCredit: input.CourseCredit,
	}

	// Create a course
	err = app.models.Courses.Insert(course)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// Create a Location header for the newly created resource/course
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/courses/%d", course.ID))

	// Write the JSON response with 201 - Created status code with the body
	// being the course data and the header being the headers map
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

	// Fetch the specific course
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

func (app *application) updateCoursesHandler(w http.ResponseWriter, r *http.Request) {
	// This method does a partial replacement
	// Get the id for the course that needs updating
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Fetch the orginal record from the database
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
	// Create an input struct to hold data read in from the client
	// We update input struct to use pointers because pointers have a
	// default value of nil
	// If a field remains nil then we know that the client did not update it
	var input struct {
		CourseCode   *string `json:"Course Code"`
		CourseTitle  *string `json:"Course Title"`
		CourseCredit *string `json:"Course Credit"`
	}

	// Initialize a new json.Decoder instance
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Check for updates
	if input.CourseCode != nil {
		course.CourseCode = *input.CourseCode
	}
	if input.CourseTitle != nil {
		course.CourseTitle = *input.CourseTitle
	}
	if input.CourseCredit != nil {
		course.CourseCredit = *input.CourseCredit
	}

	// Pass the updated course record to the Update() method
	err = app.models.Courses.Update(course)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Write the data returned by Get()
	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteCoursesHandler(w http.ResponseWriter, r *http.Request) {
	// Get the id for the Course that needs updating
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Delete the Course from the database. Send a 404 Not Found status code to the
	// client if there is no matching record
	err = app.models.Courses.Delete(id)
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
	// Return 200 Status OK to the client with a success message
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "course successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
