package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

//represent one row of data in our courses table

type Courses struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"-"`
	CourseCode   string    `json:"Course Code"`
	CourseTitle  string    `json:"Course Title"`
	CourseCredit string    `json:"Course Credit"`
	Version      int32     `json:"version"`
}

// Define a courseModel which wraps a sql.DB connection pool
type CourseModel struct {
	DB *sql.DB
}

func (m CourseModel) Insert(Course *Courses) error {

	query := `
		INSERT INTO courses (CourseCode, CourseTitle, CourseCredit)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, version
	`
	// Collect the data fields into a slice
	args := []interface{}{

		Course.CourseCode,
		Course.CourseTitle,
		Course.CourseCredit,
	}
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&Course.ID, &Course.CreatedAt, &Course.Version)
}

// Get() allows us to retrieve a specific course
func (m CourseModel) Get(id int64) (*Courses, error) {

	// Ensure that there is a valid id
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	// Create the query
	query := `
		SELECT id, created_at, CourseCode, CourseTitle, CourseCredit, version
		FROM courses
		WHERE id = $1
	`
	// Declare a course variable to hold the returned data
	var course Courses
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute the query using QueryRow()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&course.ID,
		&course.CreatedAt,
		&course.CourseCode,
		&course.CourseTitle,
		&course.CourseCredit,
		&course.Version,
	)
	// Handle any errors
	if err != nil {
		// Check the type of error
		switch {

		case errors.Is(err, sql.ErrNoRows):

			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	// Success
	return &course, nil

}

// update
func (m CourseModel) Update(Course *Courses) error {
	// Create a query
	query := `
		UPDATE courses
		SET CourseCode = $1, CourseTitle = $2, CourseCredit = $3, version = version + 1
		WHERE id = $4
		AND version = $5
		RETURNING version
	`
	args := []interface{}{
		Course.CourseCode,
		Course.CourseTitle,
		Course.CourseCredit,
		Course.ID,
		Course.Version,
	}
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Check for edit conflicts
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&Course.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

// Delete() removes a specific course
func (m CourseModel) Delete(id int64) error {

	// Ensure that there is a valid id
	if id < 1 {
		return ErrRecordNotFound
	}
	// Create the delete query
	query := `
		DELETE FROM courses
		WHERE id = $1
	`
	// Create a context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// Cleanup to prevent memory leaks
	defer cancel()
	// Execute the query
	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	// Check how many rows were affected by the delete operation. We
	// call the RowsAffected() method on the result variable
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// Check if no rows were affected
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
