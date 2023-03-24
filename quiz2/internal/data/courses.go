package data

//represent one row of data in our courses table

type Courses struct {
	CourseCode   int64  `json:"Course Code"`
	CourseTitle  string `json:"Course Title"`
	CourseCredit int64  `json:"Course Credit"`
	Version      int32  `json:"version"`
}
