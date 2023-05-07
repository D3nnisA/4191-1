package data

import (
	"context"
	"database/sql"
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

// Define a SchoolModel which wraps a sql.DB connection pool
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

// Get() allows us to retrieve a specific School
func (m CourseModel) Get(id int64) (*Courses, error) {
	return nil, nil
}

// update
func (m CourseModel) Update(Course *Courses) error {
	return nil
}

// Delete() removes a specific School
func (m CourseModel) Delete(id int64) error {
	return nil
}
