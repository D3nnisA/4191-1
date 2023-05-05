package data

//represent one row of data in our courses table

type Courses struct {
	ID           int64  `json:"id"`
	CourseCode   string `json:"Course Code"`
	CourseTitle  string `json:"Course Title"`
	CourseCredit string `json:"Course Credit"`
	Version      int32  `json:"version"`
}
